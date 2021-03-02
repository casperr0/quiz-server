import uuid

from django.urls import reverse
from rest_framework import status

from config import ADMIN_PASSWORD
from qa_server.models import Quiz
from tests.common import GameFlowTestCaseBase


class AdminDeleteQuizTestCase(GameFlowTestCaseBase):
    def setUp(self):
        super().setUp()

    def test_delete_quiz_not_found(self):
        delete_url = reverse("quiz", kwargs={"quiz_uuid": uuid.uuid4()})
        delete_response = self.client.delete(delete_url)
        self.assertEqual(delete_response.status_code, status.HTTP_404_NOT_FOUND)
        self.assertIn("error_message", delete_response.data)

    def test_delete_quiz_forbidden_no_password(self):
        delete_url = reverse("quiz", kwargs={"quiz_uuid": Quiz.objects.first().quiz_uuid})
        delete_response = self.client.delete(delete_url)
        self.assertEqual(delete_response.status_code, status.HTTP_403_FORBIDDEN)
        self.assertIn("error_message", delete_response.data)

    def test_delete_quiz_forbidden_wrong_password(self):
        delete_url = reverse("quiz", kwargs={"quiz_uuid": Quiz.objects.first().quiz_uuid})
        delete_response = self.client.delete(delete_url)
        self.assertEqual(delete_response.status_code, status.HTTP_403_FORBIDDEN)
        self.assertIn("error_message", delete_response.data)

    def test_delete_quiz_success(self):
        quiz_uuid = Quiz.objects.first().quiz_uuid
        delete_url = reverse("quiz", kwargs={"quiz_uuid": quiz_uuid})
        payload = {
            "password": ADMIN_PASSWORD
        }
        delete_response = self.client.delete(delete_url, payload, format="json")
        self.assertEqual(delete_response.status_code, status.HTTP_200_OK)
        self.assertEqual(delete_response.data["quiz_uuid"], quiz_uuid)

        get_url = reverse("quiz", kwargs={"quiz_uuid": quiz_uuid})
        get_response = self.client.get(get_url)
        self.assertEqual(get_response.status_code, status.HTTP_404_NOT_FOUND)
        self.assertIn("quiz_uuid", delete_response.data)
