# Generated by Django 3.0 on 2020-06-19 21:22

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('accounts', '0003_auto_20200619_1808'),
    ]

    operations = [
        migrations.AddField(
            model_name='user',
            name='money',
            field=models.IntegerField(default=0),
            preserve_default=False,
        ),
    ]
