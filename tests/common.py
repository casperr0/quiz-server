import random

from django.urls import reverse
from rest_framework import status
from rest_framework.test import APITestCase

from qa_server.models import Player, Provoke, Quiz


class GameFlowTestCaseBase(APITestCase):
    def setUp(self):
        super().setUp()
        self.test_quizzes_amount = 10
        self.test_player_name = "test_player"
        self.test_player_platform = "Discord"
        self.test_provokes_correct_amount = 6
        self.test_provokes_incorrect_amount = 4
        self.__setup_quizzes()
        self.__setup_player()
        self.__setup_provoke_messages()

    def __setup_quizzes(self):
        quizzes_list = []
        for i in range(self.test_quizzes_amount):
            quizzes_list.append(
                Quiz(
                    author="unit_tester",
                    domain=random.choice(Quiz.Domain.choices)[0],
                    description=f"Quiz {i}",
                    level=random.choice(Quiz.Level.choices)[0],
                    correct_answer=f"{i}_c",
                    wrong_answers=[f"{i}_w_{x}" for x in range(1, 4)],
                    comment=f"comment {i}",
                )
            )
        Quiz.objects.bulk_create(quizzes_list)

    def __setup_player(self):
        url = reverse("players")
        payload = {
            "name": self.test_player_name,
            "platform": self.test_player_platform,
        }
        response = self.client.post(url, payload, format="json")
        self.assertEqual(response.status_code, status.HTTP_201_CREATED)
        self.assertEqual(Player.objects.count(), 1)
        self.assertEqual(Player.objects.get().name, self.test_player_name)
        self.test_player_uuid = response.data["player_uuid"]

    def __setup_provoke_messages(self):
        provokes_list = []
        for i in range(self.test_provokes_correct_amount):
            provokes_list.append(
                Provoke(
                    message=f"Correct provoke {i}",
                    correct=True,
                )
            )
        for i in range(self.test_provokes_incorrect_amount):
            provokes_list.append(
                Provoke(
                    message=f"Incorrect provoke {i}",
                    correct=False,
                )
            )
        Provoke.objects.bulk_create(provokes_list)
