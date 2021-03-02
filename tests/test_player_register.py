import uuid

from django.urls import reverse
from rest_framework import status

from qa_server.models import Player
from tests.common import GameFlowTestCaseBase


class PlayerRegisterTestCase(GameFlowTestCaseBase):
    def setUp(self):
        super().setUp()

    def test_player_register_without_name(self):
        register_url = reverse("players")
        payload = {
            "platform": "Line",
            "platform_userid": "LineUser#3682",
        }
        register_response = self.client.post(register_url, payload, format="json")
        self.assertEqual(register_response.status_code, status.HTTP_400_BAD_REQUEST)
        self.assertIn("error_message", register_response.data)

    def test_player_register_unknow_platform(self):
        register_url = reverse("players")
        payload = {
            "name": "sample_user",
            "platform": "unknown_platform",
            "platform_userid": "User#4771",
        }
        register_response = self.client.post(register_url, payload, format="json")
        self.assertEqual(register_response.status_code, status.HTTP_400_BAD_REQUEST)
        self.assertIn("error_message", register_response.data)

    def test_player_register_success(self):
        register_url = reverse("players")
        payload = {
            "name": "sample_user",
            "platform": "Telegram",
            "platform_userid": "TelegramUser#4771",
        }
        register_response = self.client.post(register_url, payload, format="json")
        self.assertEqual(register_response.status_code, status.HTTP_201_CREATED)
        self.assertEqual(register_response.data["name"], "sample_user")
        self.assertEqual(register_response.data["platform"].upper(), Player.Platform.TELEGRAM.name)
        self.assertEqual(register_response.data["platform_userid"], "TelegramUser#4771")
