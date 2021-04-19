import math

import pandas as pd

from qa_server.models import Provoke, Quiz

SAMPLE_QUIZZES_FILE = "quizzes.csv"
SAMPLE_PROVOKES_FILE = "provokes.csv"


def parse_domain(digest):
    if digest == "AI":
        return Quiz.Domain.AI
    elif digest == "軟體工程":
        return Quiz.Domain.SOFTWARE_ENGINEERING
    elif digest == "網路":
        return Quiz.Domain.NETWORKING
    elif digest == "作業系統":
        return Quiz.Domain.OPERATING_SYSTEM
    elif digest == "程式語言":
        return Quiz.Domain.PROGRAMMING_LANGUAGE
    elif digest == "網頁 / 瀏覽器":
        return Quiz.Domain.WEB_AND_BROWSER
    elif digest == "演算法":
        return Quiz.Domain.ALGORITHM
    elif digest == "硬體":
        return Quiz.Domain.HARDWARE
    elif digest == "校園":
        return Quiz.Domain.CAMPUS
    elif digest == "遊戲":
        return Quiz.Domain.GAME
    elif digest == "動漫":
        return Quiz.Domain.ANIMATION
    elif digest == "常識":
        return Quiz.Domain.COMMON_SENSE
    elif digest == "CCNS":
        return Quiz.Domain.CCNS
    elif digest == "Vtuber":
        return Quiz.Domain.VTUBER
    else:
        raise ValueError(f"Unknown domain tag {digest}")


def parse_level(digest):

    if (type(digest) == float or type(digest) == int):

        return Quiz.Level.MEDIUM
    elif digest[:2] == "簡單":
        return Quiz.Level.EAZY
    elif digest[:2] == "中等":
        return Quiz.Level.MEDIUM
    elif digest[:2] == "困難":
        return Quiz.Level.HARD
    else:
        raise ValueError(f"Unknown level option {digest}")


def parse_correctness(digest):

    return digest == "correct"



def run():

    Quiz.objects.all().delete()

    df = pd.read_csv(SAMPLE_QUIZZES_FILE)
    quizzes = []
    for _, r in df.iterrows():
        row = r.to_list()
        quizzes.append(
            Quiz(
                author=row[1],
                description=row[2],
                level=parse_level(row[3]),
                correct_answer=row[4],
                wrong_answers=row[5:8],
                comment=row[8],

            )
        )
    Quiz.objects.bulk_create(quizzes)

    Provoke.objects.all().delete()

    df = pd.read_csv(SAMPLE_PROVOKES_FILE)
    provokes = []
    for _, r in df.iterrows():
        row = r.to_list()
        provokes.append(
            Provoke(
                correct=parse_correctness(row[1]),
                message=row[2],
            )
        )
    Provoke.objects.bulk_create(provokes)
