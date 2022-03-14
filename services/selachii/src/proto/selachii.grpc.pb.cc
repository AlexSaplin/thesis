// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: selachii.proto

#include "selachii.pb.h"
#include "selachii.grpc.pb.h"

#include <functional>
#include <grpcpp/impl/codegen/async_stream.h>
#include <grpcpp/impl/codegen/async_unary_call.h>
#include <grpcpp/impl/codegen/channel_interface.h>
#include <grpcpp/impl/codegen/client_unary_call.h>
#include <grpcpp/impl/codegen/client_callback.h>
#include <grpcpp/impl/codegen/message_allocator.h>
#include <grpcpp/impl/codegen/method_handler.h>
#include <grpcpp/impl/codegen/rpc_service_method.h>
#include <grpcpp/impl/codegen/server_callback.h>
#include <grpcpp/impl/codegen/server_callback_handlers.h>
#include <grpcpp/impl/codegen/server_context.h>
#include <grpcpp/impl/codegen/service_type.h>
#include <grpcpp/impl/codegen/sync_stream.h>
namespace selachii {

static const char* Selachii_method_names[] = {
  "/selachii.Selachii/LoadModel",
  "/selachii.Selachii/UnloadModel",
  "/selachii.Selachii/Run",
};

std::unique_ptr< Selachii::Stub> Selachii::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  (void)options;
  std::unique_ptr< Selachii::Stub> stub(new Selachii::Stub(channel));
  return stub;
}

Selachii::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel)
  : channel_(channel), rpcmethod_LoadModel_(Selachii_method_names[0], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_UnloadModel_(Selachii_method_names[1], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_Run_(Selachii_method_names[2], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::Status Selachii::Stub::LoadModel(::grpc::ClientContext* context, const ::selachii::LoadRequest& request, ::selachii::LoadResponse* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_LoadModel_, context, request, response);
}

void Selachii::Stub::experimental_async::LoadModel(::grpc::ClientContext* context, const ::selachii::LoadRequest* request, ::selachii::LoadResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc_impl::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_LoadModel_, context, request, response, std::move(f));
}

void Selachii::Stub::experimental_async::LoadModel(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::selachii::LoadResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc_impl::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_LoadModel_, context, request, response, std::move(f));
}

void Selachii::Stub::experimental_async::LoadModel(::grpc::ClientContext* context, const ::selachii::LoadRequest* request, ::selachii::LoadResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc_impl::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_LoadModel_, context, request, response, reactor);
}

void Selachii::Stub::experimental_async::LoadModel(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::selachii::LoadResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc_impl::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_LoadModel_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::selachii::LoadResponse>* Selachii::Stub::AsyncLoadModelRaw(::grpc::ClientContext* context, const ::selachii::LoadRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc_impl::internal::ClientAsyncResponseReaderFactory< ::selachii::LoadResponse>::Create(channel_.get(), cq, rpcmethod_LoadModel_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::selachii::LoadResponse>* Selachii::Stub::PrepareAsyncLoadModelRaw(::grpc::ClientContext* context, const ::selachii::LoadRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc_impl::internal::ClientAsyncResponseReaderFactory< ::selachii::LoadResponse>::Create(channel_.get(), cq, rpcmethod_LoadModel_, context, request, false);
}

::grpc::Status Selachii::Stub::UnloadModel(::grpc::ClientContext* context, const ::selachii::UnloadRequest& request, ::selachii::UnloadResponse* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_UnloadModel_, context, request, response);
}

void Selachii::Stub::experimental_async::UnloadModel(::grpc::ClientContext* context, const ::selachii::UnloadRequest* request, ::selachii::UnloadResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc_impl::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_UnloadModel_, context, request, response, std::move(f));
}

void Selachii::Stub::experimental_async::UnloadModel(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::selachii::UnloadResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc_impl::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_UnloadModel_, context, request, response, std::move(f));
}

void Selachii::Stub::experimental_async::UnloadModel(::grpc::ClientContext* context, const ::selachii::UnloadRequest* request, ::selachii::UnloadResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc_impl::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_UnloadModel_, context, request, response, reactor);
}

void Selachii::Stub::experimental_async::UnloadModel(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::selachii::UnloadResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc_impl::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_UnloadModel_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::selachii::UnloadResponse>* Selachii::Stub::AsyncUnloadModelRaw(::grpc::ClientContext* context, const ::selachii::UnloadRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc_impl::internal::ClientAsyncResponseReaderFactory< ::selachii::UnloadResponse>::Create(channel_.get(), cq, rpcmethod_UnloadModel_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::selachii::UnloadResponse>* Selachii::Stub::PrepareAsyncUnloadModelRaw(::grpc::ClientContext* context, const ::selachii::UnloadRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc_impl::internal::ClientAsyncResponseReaderFactory< ::selachii::UnloadResponse>::Create(channel_.get(), cq, rpcmethod_UnloadModel_, context, request, false);
}

::grpc::Status Selachii::Stub::Run(::grpc::ClientContext* context, const ::selachii::RunRequest& request, ::selachii::RunResponse* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_Run_, context, request, response);
}

void Selachii::Stub::experimental_async::Run(::grpc::ClientContext* context, const ::selachii::RunRequest* request, ::selachii::RunResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc_impl::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_Run_, context, request, response, std::move(f));
}

void Selachii::Stub::experimental_async::Run(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::selachii::RunResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc_impl::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_Run_, context, request, response, std::move(f));
}

void Selachii::Stub::experimental_async::Run(::grpc::ClientContext* context, const ::selachii::RunRequest* request, ::selachii::RunResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc_impl::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_Run_, context, request, response, reactor);
}

void Selachii::Stub::experimental_async::Run(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::selachii::RunResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc_impl::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_Run_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::selachii::RunResponse>* Selachii::Stub::AsyncRunRaw(::grpc::ClientContext* context, const ::selachii::RunRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc_impl::internal::ClientAsyncResponseReaderFactory< ::selachii::RunResponse>::Create(channel_.get(), cq, rpcmethod_Run_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::selachii::RunResponse>* Selachii::Stub::PrepareAsyncRunRaw(::grpc::ClientContext* context, const ::selachii::RunRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc_impl::internal::ClientAsyncResponseReaderFactory< ::selachii::RunResponse>::Create(channel_.get(), cq, rpcmethod_Run_, context, request, false);
}

Selachii::Service::Service() {
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      Selachii_method_names[0],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< Selachii::Service, ::selachii::LoadRequest, ::selachii::LoadResponse>(
          [](Selachii::Service* service,
             ::grpc_impl::ServerContext* ctx,
             const ::selachii::LoadRequest* req,
             ::selachii::LoadResponse* resp) {
               return service->LoadModel(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      Selachii_method_names[1],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< Selachii::Service, ::selachii::UnloadRequest, ::selachii::UnloadResponse>(
          [](Selachii::Service* service,
             ::grpc_impl::ServerContext* ctx,
             const ::selachii::UnloadRequest* req,
             ::selachii::UnloadResponse* resp) {
               return service->UnloadModel(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      Selachii_method_names[2],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< Selachii::Service, ::selachii::RunRequest, ::selachii::RunResponse>(
          [](Selachii::Service* service,
             ::grpc_impl::ServerContext* ctx,
             const ::selachii::RunRequest* req,
             ::selachii::RunResponse* resp) {
               return service->Run(ctx, req, resp);
             }, this)));
}

Selachii::Service::~Service() {
}

::grpc::Status Selachii::Service::LoadModel(::grpc::ServerContext* context, const ::selachii::LoadRequest* request, ::selachii::LoadResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status Selachii::Service::UnloadModel(::grpc::ServerContext* context, const ::selachii::UnloadRequest* request, ::selachii::UnloadResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status Selachii::Service::Run(::grpc::ServerContext* context, const ::selachii::RunRequest* request, ::selachii::RunResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


}  // namespace selachii
