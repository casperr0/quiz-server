import random
import uuid

from django.urls import reverse
from rest_framework import status

from qa_server.models import Player, Quiz
from tests.common import GameFlowTestCaseBase


class FetchProvokeMessagesTestCase(GameFlowTestCaseBase):
    def setUp(self):
        super().setUp()

    def test_fetch_all_provoke_messages(self):
        provoke_url = reverse("provokes")
        response = self.client.get(provoke_url)
        self.assertEqual(response.status_code, status.HTTP_200_OK)
        self.assertEqual(
            len(response.data),
            self.test_provokes_correct_amount + self.test_provokes_incorrect_amount,
        )

    def test_fetch_correct_provoke_messages(self):
        provoke_url = reverse("provokes")
        correct_provoke_url = f"{provoke_url}?correct=True"
        response = self.client.get(correct_provoke_url)
        self.assertEqual(response.status_code, status.HTTP_200_OK)
        self.assertEqual(len(response.data), self.test_provokes_correct_amount)

    def test_fetch_incorrect_provoke_messages(self):
        provoke_url = reverse("provokes")
        correct_provoke_url = f"{provoke_url}?correct=False"
        response = self.client.get(correct_provoke_url)
        self.assertEqual(response.status_code, status.HTTP_200_OK)
        self.assertEqual(len(response.data), self.test_provokes_incorrect_amount)
