from django.urls import re_path

from qa_server.views import (
    AnswersView,
    FeedsView,
    LeaderboardView,
    MappingView,
    PlayersView,
    PlayerView,
    ProvokesView,
    QuizzesView,
    RandView,
)

urlpatterns = [
    re_path(r"^leaderboard/$", LeaderboardView.as_view(), name="leaderboard"),
    re_path(r"^quizzes/$", QuizzesView.as_view(), name="quizzes"),
    re_path(r"^players/$", PlayersView.as_view(), name="players"),
    re_path(
        r"^players/(?P<player_uuid>[0-9a-f-]+)/$", PlayerView.as_view(), name="player"
    ),
    re_path(
        r"^mappings/(?P<platform_userid>.+)/$", MappingView.as_view(), name="mapping"
    ),
    re_path(r"^answers/$", AnswersView.as_view(), name="answers"),
    re_path(r"^feeds/(?P<player_uuid>[0-9a-f-]+)/$", FeedsView.as_view(), name="feeds"),
    re_path(r"^rand/$", RandView.as_view(), name="rand"),
    re_path(r"^provokes/$", ProvokesView.as_view(), name="provokes"),
]
