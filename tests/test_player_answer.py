import random
import uuid

from django.urls import reverse
from rest_framework import status

from qa_server.models import Player, Quiz
from tests.common import GameFlowTestCaseBase


class PlayerAnswerTestCase(GameFlowTestCaseBase):
    def setUp(self):
        super().setUp()

    def test_player_answer_quiz_unknown_player(self):
        feed_url = reverse("feeds", kwargs={"player_uuid": self.test_player_uuid})
        feed_response = self.client.get(feed_url)
        self.assertEqual(feed_response.status_code, status.HTTP_200_OK)
        self.assertIsNotNone(feed_response.data["quiz_uuid"])
        self.assertEqual(len(feed_response.data["options"]), 4)

        feed_quiz_uuid = feed_response.data["quiz_uuid"]
        feed_quiz_options = feed_response.data["options"]

        answer_url = reverse("answers")
        answer_paylaod = {
            "player_uuid": uuid.uuid4(),
            "quiz_uuid": feed_quiz_uuid,
            "answer": [x for x in feed_quiz_options if x[-1] == "c"][0],
        }
        answer_response = self.client.post(answer_url, answer_paylaod, format="json")
        self.assertEqual(answer_response.status_code, status.HTTP_404_NOT_FOUND)
        self.assertIsNotNone(answer_response.data["error_message"])

        check_url = f"{answer_url}?player_uuid={self.test_player_uuid}&quiz_uuid={feed_quiz_uuid}"
        check_response = self.client.get(check_url)
        self.assertEqual(len(check_response.data), 0)

    def test_player_answer_quiz_unknown_quiz(self):
        feed_url = reverse("feeds", kwargs={"player_uuid": self.test_player_uuid})
        feed_response = self.client.get(feed_url)
        self.assertEqual(feed_response.status_code, status.HTTP_200_OK)
        self.assertIsNotNone(feed_response.data["quiz_uuid"])
        self.assertEqual(len(feed_response.data["options"]), 4)

        feed_quiz_uuid = feed_response.data["quiz_uuid"]
        feed_quiz_options = feed_response.data["options"]

        answer_url = reverse("answers")
        answer_paylaod = {
            "player_uuid": self.test_player_uuid,
            "quiz_uuid": uuid.uuid4(),
            "answer": [x for x in feed_quiz_options if x[-1] == "c"][0],
        }
        answer_response = self.client.post(answer_url, answer_paylaod, format="json")
        self.assertEqual(answer_response.status_code, status.HTTP_404_NOT_FOUND)
        self.assertIsNotNone(answer_response.data["error_message"])

        check_url = f"{answer_url}?player_uuid={self.test_player_uuid}&quiz_uuid={feed_quiz_uuid}"
        check_response = self.client.get(check_url)
        self.assertEqual(len(check_response.data), 0)

    def test_player_answer_quiz_correct(self):
        feed_url = reverse("feeds", kwargs={"player_uuid": self.test_player_uuid})
        feed_response = self.client.get(feed_url)
        self.assertEqual(feed_response.status_code, status.HTTP_200_OK)
        self.assertIsNotNone(feed_response.data["quiz_uuid"])
        self.assertEqual(len(feed_response.data["options"]), 4)

        feed_quiz_uuid = feed_response.data["quiz_uuid"]
        feed_quiz_options = feed_response.data["options"]

        answer_url = reverse("answers")
        answer_paylaod = {
            "player_uuid": self.test_player_uuid,
            "quiz_uuid": feed_quiz_uuid,
            "answer": [x for x in feed_quiz_options if x[-1] == "c"][0],
        }
        answer_response = self.client.post(answer_url, answer_paylaod, format="json")
        self.assertEqual(answer_response.status_code, status.HTTP_201_CREATED)
        self.assertTrue(answer_response.data["correct"])

        check_url = f"{answer_url}?player_uuid={self.test_player_uuid}&quiz_uuid={feed_quiz_uuid}"
        check_response = self.client.get(check_url)
        self.assertEqual(check_response.status_code, status.HTTP_200_OK)
        self.assertEqual(len(check_response.data), 1)
        self.assertTrue(check_response.data[0]["correct"])

    def test_player_answer_quiz_incorrect(self):
        feed_url = reverse("feeds", kwargs={"player_uuid": self.test_player_uuid})
        feed_response = self.client.get(feed_url)
        self.assertEqual(feed_response.status_code, status.HTTP_200_OK)
        self.assertIsNotNone(feed_response.data["quiz_uuid"])
        self.assertEqual(len(feed_response.data["options"]), 4)

        feed_quiz_uuid = feed_response.data["quiz_uuid"]
        feed_quiz_options = feed_response.data["options"]

        answer_url = reverse("answers")
        answer_paylaod = {
            "player_uuid": self.test_player_uuid,
            "quiz_uuid": feed_quiz_uuid,
            "answer": [x for x in feed_quiz_options if x[-1] != "c"][0],
        }
        answer_response = self.client.post(answer_url, answer_paylaod, format="json")
        self.assertEqual(answer_response.status_code, status.HTTP_201_CREATED)
        self.assertFalse(answer_response.data["correct"])

        check_url = f"{answer_url}?player_uuid={self.test_player_uuid}&quiz_uuid={feed_quiz_uuid}"
        check_response = self.client.get(check_url)
        self.assertEqual(check_response.status_code, status.HTTP_200_OK)
        self.assertEqual(len(check_response.data), 1)
        self.assertFalse(check_response.data[0]["correct"])

    def test_player_answer_quiz_conflict(self):
        feed_url = reverse("feeds", kwargs={"player_uuid": self.test_player_uuid})
        feed_response = self.client.get(feed_url)
        self.assertIsNotNone(feed_response.data["quiz_uuid"])
        self.assertEqual(len(feed_response.data["options"]), 4)

        feed_quiz_uuid = feed_response.data["quiz_uuid"]
        feed_quiz_options = feed_response.data["options"]

        answer_url = reverse("answers")
        answer_paylaod = {
            "player_uuid": self.test_player_uuid,
            "quiz_uuid": feed_quiz_uuid,
            "answer": [x for x in feed_quiz_options if x[-1] == "c"][0],
        }
        answer_response = self.client.post(answer_url, answer_paylaod, format="json")
        self.assertEqual(answer_response.status_code, status.HTTP_201_CREATED)
        self.assertTrue(answer_response.data["correct"])

        conflict_answer_response = self.client.post(
            answer_url, answer_paylaod, format="json"
        )
        self.assertEqual(conflict_answer_response.status_code, status.HTTP_409_CONFLICT)
        self.assertIsNotNone(conflict_answer_response.data["error_message"])

    def test_player_answer_quizzes_completely(self):
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

        check_url = f"{answer_url}?player_uuid={self.test_player_uuid}"
        check_response = self.client.get(check_url)
        self.assertEqual(len(check_response.data), self.test_quizzes_amount)
