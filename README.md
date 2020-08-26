# quiz-server
Backend server for quiz contest.

# Prerequisites
- go 1.15
- docker 19.03.12
- docker-compose 1.26.2

# RESTful API endpoints
```
GET     /v1/players
POST    /v1/players
DELETE  /v1/players/:player_name
GET     /v1/players/:player_name/feed

GET     /v1/quizzes?tag=<tag_name>
POST    /v1/quizzes
GET     /v1/quizzes/:quiz_number
DELETE  /v1/quizzes/:quiz_number

GET     /v1/quizzes/:quiz_number/tags
POST    /v1/quizzes/:quiz_number/tags
DELETE  /v1/quizzes/:quiz_number/tags/:tag_name

GET     /v1/answers?player=<player_name>&quiz=<quiz_number>
POST    /v1/answers

GET     /v1/provokes?correct=<correctness>
POST    /v1/provokes
```

# Getting Started
## Build the docker image
```zsh
docker build -t ccns/quiz-server:latest .
```

## Start services
Server will listen on the port 8080, so make sure it was available.
```zsh
docker-compose up
```

## Stop services
```zsh
docker-compose down
```

# Contributors
[RainrainWu](https://github.com/RainrainWu)