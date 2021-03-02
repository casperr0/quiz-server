import uuid

from django.urls import reverse
from rest_framework import status

from qa_server.models import Answer, Player, Quiz
from tests.common import GameFlowTestCaseBase


class PlayerProfileCase(GameFlowTestCaseBase):
    def setUp(self):
        super().setUp()

    def test_fetch_player_profile_wrong_format(self):
        profile_url = "/players/0000-0000/"
        profile_response = self.client.get(profile_url)
        self.assertEqual(profile_response.status_code, status.HTTP_400_BAD_REQUEST)

    def test_fetch_player_profile_not_found(self):
        profile_url = reverse("player", kwargs={"player_uuid": str(uuid.uuid4())})
        profile_response = self.client.get(profile_url)
        self.assertEqual(profile_response.status_code, status.HTTP_404_NOT_FOUND)

    def test_fetch_player_profile_succedd(self):
        profile_url = reverse("player", kwargs={"player_uuid": self.test_player_uuid})
        profile_response = self.client.get(profile_url)
        self.assertEqual(profile_response.status_code, status.HTTP_200_OK)
        self.assertEqual(profile_response.data["name"], self.test_player_name)
        self.assertEqual(profile_response.data["platform"], self.test_player_platform)
        self.assertEqual(profile_response.data["platform_userid"], self.test_player_platform_userid)

    def test_platform_userid_mapping_not_found(self):
        mapping_url = reverse("mapping", kwargs={"platform_userid": "unknown-platform-userid"})
        mapping_response = self.client.get(mapping_url)
        self.assertEqual(mapping_response.status_code, status.HTTP_404_NOT_FOUND)

    def test_platform_userid_mapping_success(self):
        mapping_url = reverse("mapping", kwargs={"platform_userid": self.test_player_platform_userid})
        mapping_response = self.client.get(mapping_url)
        self.assertEqual(mapping_response.status_code, status.HTTP_200_OK)
        self.assertEqual(mapping_response.data["name"], self.test_player_name)
        self.assertEqual(mapping_response.data["platform"], self.test_player_platform)
