package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"

	_ "github.com/ccns/quiz-server/config"
	_ "github.com/ccns/quiz-server/db"
	"github.com/ccns/quiz-server/fvt"
	"github.com/ccns/quiz-server/handler"
	_ "github.com/ccns/quiz-server/utils"
)

func main() {

	fvtFlag := flag.Bool("fvt", false, "run functional verification test")
	flag.Parse()

	if *fvtFlag {
		runFVT()
	} else {
		runService()
	}
}

func runFVT() {

	fvt.VerifyGetProvokes()
	fvt.VerifyPostProvokes()

	fvt.VerifyGetPlayers()
	fvt.VerifyPostPlayers()

	fvt.VerifyGetQuizzes()
	fvt.VerifyPostQuizzes()
	fvt.VerifyGetQuiz()

	fvt.VerifyGetAnswers()
	fvt.VerifyPostAnswers()

	fvt.VerifyDeleteQuiz()
	fvt.VerifyDeletePlayer()
}

func runService() {

	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/players", handler.GetPlayersHandler)
		v1.POST("/players", handler.PostPlayersHandler)
		v1.DELETE("/players/:player_name", handler.DeletePlayerHandler)

		v1.GET("/quizzes", handler.GetQuizzesHandler)
		v1.POST("/quizzes", handler.PostQuizzesHandler)
		v1.GET("/quizzes/:quiz_number", handler.GetQuizHandler)
		v1.DELETE("/quizzes/:quiz_number", handler.DeleteQuizHandler)

		v1.GET("/quizzes/:quiz_number/tags", handler.GetQuizTagsHandler)
		v1.POST("/quizzes/:quiz_number/tags", handler.PostQuizTagsHandler)
		v1.DELETE("/quizzes/:quiz_number/tags/:tag_name", handler.DeleteQuizTagHandler)

		v1.GET("/answers", handler.GetAnswersHandler)
		v1.POST("/answers", handler.PostAnswersHandler)

		v1.GET("/provokes", handler.GetProvokesHandler)
		v1.POST("/provokes", handler.PostProvokesHandler)
	}
	router.Run()
}
