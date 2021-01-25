import random
import uuid

from django.contrib.postgres.fields import ArrayField
from django.db import models


class Quiz(models.Model):
    class Domain(models.TextChoices):
        AI = "Artificial Intelligence"
        SOFTWARE_ENGINEERING = "Software Engineering"
        NETWORKING = "Networking"
        OPERATING_SYSTEM = "Operating System"
        PROGRAMMING_LANGUAGE = "Programming Language"
        WEB_AND_BROWSER = "Web and Browser"
        ALGORITHM = "Algorithm"
        HARDWARE = "Hardware"
        CAMPUS = "Campus"
        GAME = "Game"
        ANIMATION = "Animation"
        COMMON_SENSE = "Common Sense"
        CCNS = "CCNS"

    class Level(models.TextChoices):
        EAZY = "Eazy"
        MEDIUM = "Medium"
        HARD = "Hard"

    quiz_uuid = models.UUIDField(db_index=True, default=uuid.uuid4)
    author = models.CharField(max_length=64)
    domain = models.CharField(
        max_length=64, choices=Domain.choices, default=Domain.COMMON_SENSE
    )
    description = models.CharField(max_length=255)
    level = models.CharField(max_length=16, choices=Level.choices, default=Level.MEDIUM)
    correct_answer = models.CharField(max_length=255)
    wrong_answers = ArrayField(models.CharField(max_length=255), size=3)
    comment = models.TextField()

    def get_json(self):
        options = [self.correct_answer] + self.wrong_answers
        random.shuffle(options)
        return {
            "quiz_uuid": str(self.quiz_uuid),
            "author": self.author,
            "domain": str(self.domain),
            "description": self.description,
            "level": str(self.level),
            "options": options,
            "comment": self.comment,
        }


class Player(models.Model):
    class Platform(models.Choices):
        MESSENGER = "Messenger"
        TELEGRAM = "TELEGRAM"
        DISCORD = "Discord"
        NETCAT = "Netcat"
        LINE = "Line"
        MEWE = "Mewe"

    player_uuid = models.UUIDField(db_index=True, default=uuid.uuid4)
    name = models.CharField(max_length=32)
    platform = models.CharField(
        max_length=16, choices=Platform.choices, default=Platform.DISCORD
    )

    @classmethod
    def parse_platform(cls, digest):
        digest = digest.lower()
        if digest == "messenger":
            return cls.Platform.MESSENGER
        elif digest == "telegram":
            return cls.Platform.TELEGRAM
        elif digest == "discord":
            return cls.Platform.DISCORD
        elif digest == "netcat":
            return cls.Platform.NETCAT
        elif digest == "line":
            return cls.Platform.LINE
        elif digest == "mewe":
            return cls.Platform.MEWE
        else:
            raise ValueError(f"Unknown platform type {digest}")

    def get_json(self):
        correct_count = Answer.objects.filter(
            player__player_uuid=self.player_uuid, correct=True
        ).count()
        incorrect_count = Answer.objects.filter(
            player__player_uuid=self.player_uuid, correct=False
        ).count()
        no_answer_count = Quiz.objects.all().count() - correct_count - incorrect_count
        return {
            "player_uuid": str(self.player_uuid),
            "name": self.name,
            "platform": str(self.platform),
            "correct_count": correct_count,
            "incorrect_count": incorrect_count,
            "no_answer_count": no_answer_count,
        }


class Answer(models.Model):

    player = models.ForeignKey(Player, on_delete=models.CASCADE)
    quiz = models.ForeignKey(Quiz, on_delete=models.CASCADE)
    correct = models.BooleanField(default=False)

    def get_json(self):
        return {
            "player_uuid": str(self.player.player_uuid),
            "quiz_uuid": str(self.quiz.quiz_uuid),
            "correct": self.correct,
        }


class Provoke(models.Model):

    message = models.CharField(max_length=255)
    correct = models.BooleanField(default=False)

    def get_json(self):
        return {
            "message": self.message,
            "correct": self.correct,
        }
