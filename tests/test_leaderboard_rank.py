import random
import uuid

from django.urls import reverse
from rest_framework import status

from qa_server.models import Answer, Player, Quiz
from tests.common import GameFlowTestCaseBase


class LeaderboardTestCase(GameFlowTestCaseBase):
    def setUp(self):
        super().setUp()

    def __mock_player(self, name, platform, platform_userid, score):

        if score > Quiz.objects.all().count():
            raise ValueError(f"Score exceeds the total amount of quizzess.")

        url = reverse("players")
        payload = {
            "name": name,
            "platform": platform,
            "platform_userid": platform_userid,
        }
        response = self.client.post(url, payload, format="json")

        mock_player_uuid = response.data["player_uuid"]

        feed_url = reverse("feeds", kwargs={"player_uuid": mock_player_uuid})
        for i in range(score):
            profile_url = reverse("player", kwargs={"player_uuid": mock_player_uuid})
            profile_response = self.client.get(profile_url)

            feed_url = reverse("feeds", kwargs={"player_uuid": mock_player_uuid})
            feed_response = self.client.get(feed_url)
            self.assertEqual(feed_response.status_code, status.HTTP_200_OK)

            feed_quiz_uuid = feed_response.data["quiz_uuid"]
            feed_quiz_options = feed_response.data["options"]

            answer_url = reverse("answers")
            answer_paylaod = {
                "player_uuid": mock_player_uuid,
                "quiz_uuid": feed_quiz_uuid,
                "answer": [x for x in feed_quiz_options if x[-1] == "c"][0],
            }
            answer_response = self.client.post(answer_url, answer_paylaod, format="json")
            self.assertEqual(answer_response.status_code, status.HTTP_201_CREATED)

        check_url = f"{answer_url}?player_uuid={mock_player_uuid}"
        check_response = self.client.get(check_url)
        self.assertEqual(len(check_response.data), score)

    def test_rank_players_by_score(self):

        self.__mock_player("rank_one", "Discord", "rank_one_userid", self.test_quizzes_amount)
        self.__mock_player("rank_three", "Line", "rank_three_userid", self.test_quizzes_amount-2)
        self.__mock_player("rank_two", "mewe", "rank_two_userid", self.test_quizzes_amount-1)

        leaderboard_url = reverse("leaderboard")
        response = self.client.get(leaderboard_url)
        self.assertEqual(len(response.data), 4)
        self.assertEqual(response.data[0]["name"], "rank_one")
        self.assertEqual(response.data[1]["name"], "rank_two")
        self.assertEqual(response.data[2]["name"], "rank_three")
