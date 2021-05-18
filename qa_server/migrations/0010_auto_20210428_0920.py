# Generated by Django 3.1.7 on 2021-04-28 09:20

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('qa_server', '0009_auto_20210416_1134'),
    ]

    operations = [
        migrations.RenameField(
            model_name='quiz',
            old_name='comment',
            new_name='photo_name',
        ),
        migrations.AlterField(
            model_name='player',
            name='platform',
            field=models.CharField(choices=[('Messenger', 'Messenger'), ('Telegram', 'Telegram'), ('Discord', 'Discord'), ('Netcat', 'Netcat'), ('Line', 'Line'), ('Mewe', 'Mewe')], default='Discord', max_length=16),
        ),
    ]
