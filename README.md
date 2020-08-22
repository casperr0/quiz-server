# quiz-server
Backend server for quiz contest.

# Prerequisites
- go 1.15
- docker 19.03.12
- docker-compose 1.26.2

# RESTful API endpoints
```
GET     /officers
POST    /officers
UPDATE  /officers/<officer_name>
DELETE  /officers/<officer_name>

GET     /officers/<officer_name>/roles
POST    /officers/<officer_name>/roles
DELETE  /officers/<officer_name>/roles/<role_name>

GET     /players
POST    /players
DELETE  /players/<player_name>

GET     /quizzes
POST    /quizzes
UPDATE  /quizzes/<quiz_id>
DELETE  /quizzes/<quiz_id>

GET     /quizzes/<quiz_id>/tags
POST    /quizzes/<quiz_id>/tags
DELETE  /quizzes/<quiz_id>/tags/<tag_name>
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