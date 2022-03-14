"""Model views are defined here"""
from drf_yasg.utils import swagger_auto_schema
from rest_framework import generics, permissions
from rest_framework.decorators import permission_classes
from rest_framework.response import Response
from rest_framework.views import APIView
from .models import PaymentOperation
from django.db import transaction
from django.conf import settings

import datetime
import typing
from gorilla_pb.gorilla_pb2 import AddDeltasRequest, AddDeltasResponse, Delta, \
                                   GetBalanceRequest, GetBalanceResponse, \
                                   GetDeltasRequest, GetDeltasResponse

from gorilla_pb.gorilla_pb2_grpc import GorillaStub
from google.protobuf.json_format import MessageToDict

import stripe


def get_customer_id_by_user(user):
    stripe.api_key = settings.STRIPE_SECRET
    if not user.stripe_customer_id:
        customer_data = stripe.Customer.list(email=user.email).data
        if len(customer_data) == 0:
            # Creating new customer
            customer = stripe.Customer.create(name=user.name, email=user.email)
        else:
            customer = customer_data[0]
        user.stripe_customer_id = customer['id']
        user.save()
    return user.stripe_customer_id


@permission_classes((permissions.IsAuthenticated, ))
class InitiatePaymentView(APIView):
    @swagger_auto_schema()
    def post(self, request, *args, **kwargs):
        stripe.api_key = settings.STRIPE_SECRET
        user = self.request.user
        amount = int(float(self.kwargs['amount']) * 100)
        # redirect_url = self.kwargs['redirect_url']

        customer_id = get_customer_id_by_user(user)

        stripe.InvoiceItem.create(
            customer=customer_id,
            price=settings.STRIPE_PRICE_ID,
            quantity=amount,
        )

        invoice = stripe.Invoice.create(
            customer=customer_id,
            auto_advance=True,
        )

        invoice_finalized = stripe.Invoice.finalize_invoice(invoice['id'])

        payment_operation = PaymentOperation(user=user, amount=amount / 100, payment_id=invoice_finalized.id)
        payment_operation.save()

        return Response(data={'redirect_url': invoice_finalized.hosted_invoice_url}, status=200)


def make_delta_for_payment_operation(operation: PaymentOperation):
    creation: datetime.datetime = operation.creation_time
    creation = creation.replace(hour=0, minute=0, second=0, microsecond=0)
    return Delta(
        Date=int(creation.timestamp()),
        Category="PAYMENT",
        Balance=float(operation.amount),
        ObjectID='00000000-0000-0000-0000-000000000000',
        ObjectType='UNKNOWN',
        OwnerID=str(operation.user.uuid),
    )


def send_bill_to_gorilla(gorilla: GorillaStub, deltas: typing.List[Delta]):
    _: AddDeltasResponse = gorilla.AddDeltas(AddDeltasRequest(Deltas=deltas))


def get_user_balance_from_gorilla(gorilla: GorillaStub, user):
    response: GetBalanceResponse = gorilla.GetBalance(GetBalanceRequest(OwnerID=str(user.uuid)))
    return response.Balance

# message GetDeltasRequest {
#     string OwnerID = 1;
#     string ModelID = 2; // Empty means all models
#     int64  FirstDate = 3; // Unix timestamp of date in UTC
#     int64  LastDate = 4; // Unix timestamp of date in UTC
#     bool   UseCategories = 5; // If true split by categories
# }
def get_user_transactions_from_gorilla(gorilla: GorillaStub, user, begin_timestamp: int, end_timestamp: int, use_categories: bool):
    response: GetDeltasResponse = gorilla.GetDeltas(GetDeltasRequest(
        OwnerID=str(user.uuid),
        FirstDate=begin_timestamp,
        LastDate=end_timestamp,
        UseCategories=use_categories,
    ))
    return response


@transaction.atomic
def process_all_user_pending_payments(user):
    stripe.api_key = settings.STRIPE_SECRET
    queryset = PaymentOperation.objects.all().filter(user=user)
    deltas_evaluated = []

    for payment_operation in queryset:
        if payment_operation.status == 'paid' or payment_operation.status == 'void' \
        or payment_operation.status[0:5] == 'ERROR':
            continue

        try:
            payment = stripe.Invoice.retrieve(payment_operation.payment_id)
        except Exception as e:
            payment_operation.status = ('ERROR: ' + str(e))[:512]
            payment_operation.save()
            continue

        if payment.status == 'paid' and payment_operation.status != 'paid':
            deltas_evaluated.append(make_delta_for_payment_operation(payment_operation))
            payment_operation.status = payment.status
        else:
            payment_operation.status = payment.status
        payment_operation.save()
    if len(deltas_evaluated) > 0:
        send_bill_to_gorilla(gorilla=settings.GORILLA_STUB, deltas=deltas_evaluated)


@permission_classes((permissions.IsAuthenticated, ))
class GetCurrentBalanceAPIView(APIView):
    @swagger_auto_schema()
    def get(self, request, *args, **kwargs):
        process_all_user_pending_payments(user=self.request.user)
        balance = get_user_balance_from_gorilla(gorilla=settings.GORILLA_STUB, user=self.request.user)
        return Response(data={'money': balance})


@permission_classes((permissions.IsAuthenticated, ))
class GetTransactionsListAPIView(APIView):
    @swagger_auto_schema()
    def get(self, request, *args, **kwargs):
        user = self.request.user
        timestamp_begin = int(self.kwargs['timestamp_begin'])
        timestamp_end = int(self.kwargs['timestamp_end'])
        split_by_categories = True

        transactions = get_user_transactions_from_gorilla(settings.GORILLA_STUB, user,
                                                          timestamp_begin, timestamp_end,
                                                          split_by_categories)
        return Response(data=MessageToDict(transactions))

