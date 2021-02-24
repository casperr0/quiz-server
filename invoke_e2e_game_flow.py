import requests
from loguru import logger

PLAYERS_ENDPOINT = "http://127.0.0.1:8000/players/"
FEEDS_ENDPOINT = "http://127.0.0.1:8000/feeds/"
ANSWERS_ENDPOINT = "http://127.0.0.1:8000/answers/"
PLAYER_NAME = "E2E_PLAYER"
PLAYER_PLATFORM = "Messenger"
PLAYER_PLATFORM_USERID = "LineUserz@5726"
PLAYER_ANSWERS_AMOUNT = 15

def run_game_flow(answer_amount):
    player_payload = {
        "name": PLAYER_NAME,
        "platform": PLAYER_PLATFORM,
        "platform_userid": PLAYER_ANSWERS_AMOUNT,
    }
    player_response = requests.post(
        PLAYERS_ENDPOINT,
        data=player_payload
    )
    player_uuid = player_response.json()["player_uuid"]
    logger.info(f"Player {player_uuid} register successfully!")

    feed_url = f"{FEEDS_ENDPOINT}{player_uuid}/"
    for i in range(answer_amount):
        feed_response = requests.get(feed_url)
        quiz = feed_response.json()
        logger.info(f"Get quiz {i} {quiz}")

        answer_paylaod = {
            "player_uuid": player_uuid,
            "quiz_uuid": quiz["quiz_uuid"],
            "answer": quiz["options"][0],
        }
        answer_response = requests.post(
            ANSWERS_ENDPOINT,
            data=answer_paylaod
        )
        result = answer_response.json()
        logger.info(f"Get result {i} {result}")

    check_url = f"{PLAYERS_ENDPOINT}{player_uuid}"
    check_response = requests.get(check_url)
    profile = check_response.json()
    logger.info(f"Check profile {profile}")

    # delete method are not allowed in the production mode.
    clean_response = requests.delete(PLAYERS_ENDPOINT)
    message = clean_response.json()
    logger.info(f"Clean player {message}")

run_game_flow(PLAYER_ANSWERS_AMOUNT)
