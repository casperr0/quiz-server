import random
import uuid

from django.conf import settings
from django.core import exceptions
from django.db.models import Count, Q
from rest_framework import status
from rest_framework.response import Response
from rest_framework.views import APIView

from config import ADMIN_PASSWORD
from qa_server.models import Answer, Player, Provoke, Quiz


class QuizzesView(APIView):
    def get(self, request, format=None):
        quizzes = [quiz.get_json() for quiz in Quiz.objects.all()]
        return Response(quizzes)


class QuizView(APIView):
    def get(self, request, *args, **kwargs):
        try:
            quiz = Quiz.objects.get(quiz_uuid=kwargs["quiz_uuid"])
            return Response(
                quiz.get_json(),
                status=status.HTTP_200_OK,
            )
        except exceptions.ValidationError:
            return Response(
                {
                    "error_message": f"payload validation failed! your uuid format may be incorrect"
                },
                status=status.HTTP_400_BAD_REQUEST,
            )
        except Quiz.DoesNotExist:
            quiz_uuid = kwargs["quiz_uuid"]
            return Response(
                {"error_message": f"quiz with quiz_uuid {quiz_uuid} does not exist."},
                status=status.HTTP_404_NOT_FOUND,
            )

    def delete(self, request, *args, **kwargs):
        try:
            quiz = Quiz.objects.get(quiz_uuid=kwargs["quiz_uuid"])
            quiz_uuid = quiz.quiz_uuid
            if (
                "password" not in request.data
                or request.data["password"] != ADMIN_PASSWORD
            ):
                return Response(
                    {"error_message": "admin password is incorrect or not provided"},
                    status=status.HTTP_403_FORBIDDEN,
                )

            quiz.delete()
            return Response(
                {"quiz_uuid": quiz_uuid},
                status=status.HTTP_200_OK,
            )
        except exceptions.ValidationError:
            return Response(
                {
                    "error_message": f"payload validation failed! your uuid format may be incorrect"
                },
                status=status.HTTP_400_BAD_REQUEST,
            )
        except Quiz.DoesNotExist:
            quiz_uuid = kwargs["quiz_uuid"]
            return Response(
                {"error_message": f"quiz with quiz_uuid {quiz_uuid} does not exist."},
                status=status.HTTP_404_NOT_FOUND,
            )


class PlayersView(APIView):
    def get(self, request, *args, **kwargs):
        players = [player.get_json() for player in Player.objects.all()]
        return Response(players, status=status.HTTP_200_OK)

    def post(self, request, *args, **kwargs):
        try:
            new_player = Player.objects.create(
                name=request.data["name"],
                platform=Player.parse_platform(request.data["platform"]),
                platform_userid=request.data["platform_userid"],
                company=request.data["company"],
            )
            return Response(new_player.get_json(), status=status.HTTP_201_CREATED)
        except KeyError:
            return Response(
                {
                    "error_message": "Invalid payload, some required fields are not provided"
                },
                status=status.HTTP_400_BAD_REQUEST,
            )
        except ValueError as err:
            return Response(
                {"error_message": str(err)},
                status=status.HTTP_400_BAD_REQUEST,
            )

    def delete(self, request, *args, **kwargs):
        if not settings.DEBUG:
            return Response(
                {
                    "error_message": "Deleting all of players are not allowed in the production mode"
                },
                status=status.HTTP_405_METHOD_NOT_ALLOWED,
            )
        Player.objects.all().delete()
        return Response(
            {"message": "all players were deleted."},
            status=status.HTTP_200_OK,
        )


class PlayerView(APIView):
    def get(self, request, *args, **kwargs):
        try:
            player = Player.objects.get(player_uuid=kwargs["player_uuid"])
            return Response(
                player.get_json(),
                status=status.HTTP_200_OK,
            )
        except exceptions.ValidationError:
            return Response(
                {
                    "error_message": f"payload validation failed! your uuid format may be incorrect"
                },
                status=status.HTTP_400_BAD_REQUEST,
            )
        except Player.DoesNotExist:
            player_uuid = kwargs["player_uuid"]
            return Response(
                {
                    "error_message": f"player with player_uuid {player_uuid} does not exist."
                },
                status=status.HTTP_404_NOT_FOUND,
            )


class MappingView(APIView):
    def get(self, request, *args, **kwargs):
        try:
            player = Player.objects.get(platform_userid=kwargs["platform_userid"])
            return Response(
                player.get_json(),
                status=status.HTTP_200_OK,
            )
        except exceptions.ValidationError:
            return Response(
                {
                    "error_message": f"payload validation failed! your uuid format may be incorrect"
                },
                status=status.HTTP_400_BAD_REQUEST,
            )
        except Player.DoesNotExist:
            platform_userid = kwargs["platform_userid"]
            return Response(
                {
                    "error_message": f"player with platform_userid {platform_userid} does not exist."
                },
                status=status.HTTP_404_NOT_FOUND,
            )


class AnswersView(APIView):
    def get(self, request, *args, **kwargs):
        answers = Answer.objects.all()
        if request.GET.get("player_uuid"):
            target_player = Player.objects.get(
                player_uuid=request.GET.get("player_uuid"),
            )
            answers = answers.filter(player=target_player)
        if request.GET.get("quiz_uuid"):
            target_quiz = Quiz.objects.get(
                quiz_uuid=request.GET.get("quiz_uuid"),
            )
            answers = answers.filter(quiz=target_quiz)
        return Response([answer.get_json() for answer in answers])

    def post(self, request, *args, **kwargs):

        try:
            target_player_uuid = request.data["player_uuid"]
            target_player = Player.objects.get(
                player_uuid=target_player_uuid,
            )
            target_quiz_uuid = request.data["quiz_uuid"]
            target_quiz = Quiz.objects.get(
                quiz_uuid=target_quiz_uuid,
            )
            exist_answer = Answer.objects.filter(player=target_player, quiz=target_quiz)
            if exist_answer:
                return Response(
                    {
                        "error_message": f"answer of player {target_player_uuid} to quiz {target_quiz_uuid} already exist.",
                        "correct": bool(
                            request.data["answer"] == target_quiz.correct_answer
                        ),
                    },
                    status=status.HTTP_409_CONFLICT,
                )

            new_answer = Answer.objects.create(
                player=target_player,
                quiz=target_quiz,
                correct=bool(request.data["answer"] == target_quiz.correct_answer),
            )
            return Response(new_answer.get_json(), status=status.HTTP_201_CREATED)
        except exceptions.ValidationError:
            return Response(
                {
                    "error_message": f"payload validation failed! your uuid format may be incorrect"
                },
                status=status.HTTP_400_BAD_REQUEST,
            )
        except Player.DoesNotExist:
            player_uuid = request.data["player_uuid"]
            return Response(
                {"error_message": f"player with uuid {player_uuid} does not exist."},
                status=status.HTTP_404_NOT_FOUND,
            )
        except Quiz.DoesNotExist:
            quiz_uuid = request.data["quiz_uuid"]
            return Response(
                {"error_message": f"quiz with uuid {quiz_uuid} does not exist."},
                status=status.HTTP_404_NOT_FOUND,
            )

    def delete(self, request, *args, **kwargs):
        if not settings.DEBUG:
            return Response(
                {
                    "error_message": "Deleting all of answers are not allowed in the production mode"
                },
                status=status.HTTP_405_METHOD_NOT_ALLOWED,
            )
        Answer.objects.all().delete()
        return Response(
            {"message": "all answers were deleted."},
            status=status.HTTP_204_NO_CONTENT,
        )


class FeedsView(APIView):
    def get(self, request, *args, **kwargs):
        try:
            target_player_uuid = kwargs["player_uuid"]
            target_player = Player.objects.get(player_uuid=target_player_uuid)

            all_quizzes = Quiz.objects.all()
            all_quizzes_uuids = [quiz.get_json()["quiz_uuid"] for quiz in all_quizzes]

            answered_quizzes = Answer.objects.all().filter(
                player__player_uuid=target_player_uuid
            )
            answered_quizzes_uuids = [
                answer.get_json()["quiz_uuid"] for answer in answered_quizzes
            ]

            candidates_quizzes_uuids = [
                quiz_uuid
                for quiz_uuid in all_quizzes_uuids
                if quiz_uuid not in answered_quizzes_uuids
            ]
            if not candidates_quizzes_uuids:
                return Response(
                    {},
                    status=status.HTTP_204_NO_CONTENT,
                )

            feed_quiz_uuid = random.choice(candidates_quizzes_uuids)
            feed_quiz = Quiz.objects.get(quiz_uuid=feed_quiz_uuid)
            return Response(feed_quiz.get_json(), status=status.HTTP_200_OK)
        except exceptions.ValidationError:
            return Response(
                {
                    "error_message": f"payload validation failed! your uuid format may be incorrect"
                },
                status=status.HTTP_400_BAD_REQUEST,
            )
        except Player.DoesNotExist:
            player_uuid = kwargs["player_uuid"]
            return Response(
                {"error_message": f"player with uuid {player_uuid} does not exist."},
                status=status.HTTP_404_NOT_FOUND,
            )


class RandView(APIView):
    def get(self, request, *args, **kwargs):
        quizzes = Quiz.objects.all()
        if not quizzes:
            return Response(
                {},
                status=status.HTTP_204_NO_CONTENT,
            )
        return Response(random.choice(quizzes).get_json(), status=status.HTTP_200_OK)


class ProvokesView(APIView):
    def get(self, request, *args, **kwargs):
        provokes = Provoke.objects.all()
        correctness = request.GET.get("correct")
        if correctness is not None:
            provokes = provokes.filter(correct=correctness.lower() == "true")
        return Response(
            [provoke.get_json() for provoke in provokes], status=status.HTTP_200_OK
        )


class LeaderboardView(APIView):
    def get(self, request, *args, **kwargs):
        players = Player.objects.annotate(
            score=Count("answer", filter=Q(answer__correct=True))
        ).order_by("-score")
        leaderboard = []
        for player in players:
            player_json = {
                "name": player.name,
                "platform": str(player.platform),
                "score": player.score,
            }
            leaderboard.append(player_json)
        return Response(leaderboard, status=status.HTTP_200_OK)
