# quiz-server
Backend server for quiz contest.

# Prerequisites
- go 1.15
- docker 19.03.12
- docker-compose 1.26.2

# Getting Started

## Bear Running
### Run Service
Remember to run a postgreSQL databse on localhost.
```zsh
go run cmd/main.go
```

### Reset Before
Reset all records in the current database.
```zsh
go run cmd/main.go -reset
```

### Load Dev Data
Load dev data from external static file within directory `example_data/`.
```zsh
go run cmd/main.go -loaddev
```

### Load Prod Data
Load prod data from external static file within directory `data/`.
```zsh
go run cmd/main.go -load
```

### Run Functional Verification Test
Remember to run the service before execute fvt.
```zsh
go run cmd/main.go -fvt
```

## Containerization
### Build the docker image
```zsh
docker build -t rainrainwu/quiz-server:latest .
```

### Start services
Server will listen on the port 8080, so make sure it was available.
```zsh
docker-compose up
```

### Stop services
```zsh
docker-compose down
```

# Usage
The service provides an interacting interface via RESTful API, below is the list of all endpoints:
```
GET     /v1/players
POST    /v1/players
DELETE  /v1/players/:player_name
GET     /v1/players/:player_name/feed
GET     /v1/players/:player_name/rand

GET     /v1/quizzes?tag=<tag_name>
POST    /v1/quizzes
GET     /v1/quizzes/:quiz_number
DELETE  /v1/quizzes/:quiz_number

GET     /v1/quizzes/:quiz_number/tags
POST    /v1/quizzes/:quiz_number/tags
DELETE  /v1/quizzes/:quiz_number/tags/:tag_name

GET     /v1/tags
POST    /v1/tags

GET     /v1/answers?player=<player_name>&quiz=<quiz_number>
POST    /v1/answers

GET     /v1/provokes?correct=<correctness>
POST    /v1/provokes
```

## Players
You can create new players, list all players by score, and feed a quiz for a user through these api endpoints.

## Quizzes
You can create new quizzes, access current quizzes, and update tags of a quiz through these api endpoints.

## Tags
You can create or delete tags through these api endpoints.

## Answers
You can upload or query answers from players to quizzes through these api endpoints.

## Provokes
You can uplaod new provoke messages or query messages by correctness through these api endpoints.

## More Examples and References
Please refer to `out.md` for more details of requests and responses.

# Contributors
[RainrainWu](https://github.com/RainrainWu)