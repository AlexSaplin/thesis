# Generated by Django 3.0 on 2020-06-19 21:22

from django.db import migrations


class Migration(migrations.Migration):

    dependencies = [
        ('billing', '0001_initial'),
    ]

    operations = [
        migrations.RenameField(
            model_name='paymentoperation',
            old_name='client_secret',
            new_name='payment_id',
        ),
        migrations.RemoveField(
            model_name='paymentoperation',
            name='payment_intent_id',
        ),
    ]
