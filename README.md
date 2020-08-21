# quiz-server
Backend server for quiz contest.

# Prerequisites
- go 1.15
- docker 19.03.12
- docker-compose 1.26.2

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