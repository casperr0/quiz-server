import random
import uuid

from django.urls import reverse
from rest_framework import status

from qa_server.models import Player, Quiz
from tests.common import GameFlowTestCaseBase


class RandomQuizTestCase(GameFlowTestCaseBase):
    def setUp(self):
        super().setUp()

    def test_ask_for_random_quiz(self):
        rand_url = reverse("rand")
        for i in range(self.test_quizzes_amount + 1):
            response = self.client.get(rand_url)
            self.assertEqual(response.status_code, status.HTTP_200_OK)
            self.assertIsNotNone(response.data["quiz_uuid"])
            self.assertEqual(len(response.data["options"]), 4)
