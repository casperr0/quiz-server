package handler

import (
	"strconv"

	"github.com/ccns/quiz-server/db"
	"github.com/gin-gonic/gin"
)

// Answer describe the structue of an answer from player to quiz.
type Answer struct {
	PlayerName string `json:"player_name"`
	QuizNumber int    `json:"quiz_number"`
	Correct    bool   `json:"correct"`
}

// GetAnswersHandler handlers get requests on route /answers.
func GetAnswersHandler(ctx *gin.Context) {

	playerName := ctx.DefaultQuery("player", "")
	quizNumber := ctx.DefaultQuery("quiz", "")

	if playerName != "" {

		data, err := db.QueryQuizAnswersByPlayer(playerName)
		if err != nil {
			status := Status{500, err.Error()}
			resp, _ := BuildHATEOAS(nil, status, nil, nil)
			ctx.String(500, resp)
			return
		}

		status := Status{200, "answers listed successfully."}
		resp, _ := BuildHATEOAS(nil, status, data, nil)
		ctx.String(200, resp)

	} else if quizNumber != "" {

		quizNumber, err := strconv.Atoi(ctx.Param("quiz_number"))
		if err != nil {
			resp, _ := BuildHATEOAS(nil, Status{400, err.Error()}, nil, nil)
			ctx.String(400, resp)
			return
		}

		data, err := db.QueryPlayerAnswersByQuiz(quizNumber)
		if err != nil {
			resp, _ := BuildHATEOAS(nil, Status{500, err.Error()}, nil, nil)
			ctx.String(500, resp)
			return
		}

		status := Status{200, "answers listed successfully."}
		resp, _ := BuildHATEOAS(nil, status, data, nil)
		ctx.String(200, resp)
	}

	data, err := db.ListAnswers()
	if err != nil {
		resp, _ := BuildHATEOAS(nil, Status{500, err.Error()}, nil, nil)
		ctx.String(500, resp)
		return
	}

	status := Status{200, "answers listed successfully."}
	resp, _ := BuildHATEOAS(nil, status, data, nil)
	ctx.String(200, resp)
}

// PostAnswersHandler handlers post requests on route /answers.
func PostAnswersHandler(ctx *gin.Context) {

	var answer Answer
	err := ctx.BindJSON(&answer)
	if err != nil {
		resp, _ := BuildHATEOAS(nil, Status{400, err.Error()}, nil, nil)
		ctx.String(400, resp)
		return
	}

	err = db.RegisterAnswer(answer.PlayerName, answer.QuizNumber, answer.Correct)
	if err != nil {
		resp, _ := BuildHATEOAS(nil, Status{400, err.Error()}, nil, nil)
		ctx.String(400, resp)
		return
	}

	status := Status{201, "answer created successfully."}
	resp, _ := BuildHATEOAS(nil, status, answer, nil)
	ctx.String(201, resp)
}
