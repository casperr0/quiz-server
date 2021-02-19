import random
import uuid

from django.urls import reverse
from rest_framework import status

from qa_server.models import Player, Quiz
from tests.common import GameFlowTestCaseBase


class PlayerFeedTestCase(GameFlowTestCaseBase):
    def setUp(self):
        super().setUp()

    def test_feed_player_quizzes_unknown_player(self):
        feed_url = reverse("feeds", kwargs={"player_uuid": uuid.uuid4()})
        response = self.client.get(feed_url)
        self.assertEqual(response.status_code, status.HTTP_404_NOT_FOUND)
        self.assertIsNotNone(response.data["error_message"])

    def test_feed_player_quizzes(self):
        feed_url = reverse("feeds", kwargs={"player_uuid": self.test_player_uuid})
        response = self.client.get(feed_url)
        self.assertEqual(response.status_code, status.HTTP_200_OK)
        self.assertIsNotNone(response.data["quiz_uuid"])
        self.assertEqual(len(response.data["options"]), 4)

    def test_feed_player_quizzes_no_left(self):
        feed_url = reverse("feeds", kwargs={"player_uuid": self.test_player_uuid})
        for i in range(self.test_quizzes_amount):
            print(i)
            feed_response = self.client.get(feed_url)
            feed_quiz_uuid = feed_response.data["quiz_uuid"]
            feed_quiz_options = feed_response.data["options"]

            answer_url = reverse("answers")
            answer_paylaod = {
                "player_uuid": self.test_player_uuid,
                "quiz_uuid": feed_quiz_uuid,
                "answer": [x for x in feed_quiz_options if x[-1] == "c"][0],
            }
            answer_response = self.client.post(
                answer_url, answer_paylaod, format="json"
            )
            self.assertEqual(answer_response.status_code, status.HTTP_201_CREATED)

        feed_response = self.client.get(feed_url)
        self.assertEqual(feed_response.status_code, status.HTTP_204_NO_CONTENT)
