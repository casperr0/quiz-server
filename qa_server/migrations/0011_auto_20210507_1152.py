# Generated by Django 3.1.7 on 2021-05-07 11:52

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('qa_server', '0010_auto_20210428_0920'),
    ]

    operations = [
        migrations.AlterField(
            model_name='player',
            name='platform',
            field=models.CharField(choices=[('Messenger', 'Messenger'), ('Telegram', 'Telegram'), ('Discord', 'Discord'), ('Netcat', 'Netcat'), ('Line', 'Line'), ('Mewe', 'Mewe')], default='Discord', max_length=16),
        ),
    ]
