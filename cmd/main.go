package main

import (
	"flag"

	"github.com/ccns/quiz-server/utils"

	"github.com/gin-gonic/gin"

	_ "github.com/ccns/quiz-server/config"
	"github.com/ccns/quiz-server/db"
	"github.com/ccns/quiz-server/fvt"
	"github.com/ccns/quiz-server/handler"
)

func main() {

	fvtFlag := flag.Bool("fvt", false, "run functional verification test")
	loadFlag := flag.Bool("load", false, "load prod data from external file.")
	loaddevFlag := flag.Bool("loaddev", false, "load dev data from external file.")
	resetFlag := flag.Bool("reset", false, "reset records in current database before run service.")
	flag.Parse()

	if *fvtFlag {
		runFVT()
	} else if *loadFlag {
		utils.LoadAll(true)
	} else if *loaddevFlag {
		utils.LoadAll(false)
	} else if *resetFlag {
		db.ConnectDatabase(true)
	} else {
		db.ConnectDatabase(false)
		runService()
	}
}

func runFVT() {

	fvt.VerifyPostProvokes()
	fvt.VerifyGetProvokes()

	fvt.VerifyPostTags()
	fvt.VerifyGetTags()

	fvt.VerifyPostPlayers()
	fvt.VerifyGetPlayers()
	fvt.VerifyGetPlayerFeed()

	fvt.VerifyPostQuizzes()
	fvt.VerifyGetQuizzes()
	fvt.VerifyGetQuiz()

	fvt.VerifyPostQuizTags()
	fvt.VerifyGetQuizTags()

	fvt.VerifyPostAnswers()
	fvt.VerifyGetAnswers()

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
		v1.GET("/players/:player_name/feed", handler.GetPlayerFeedHandler)

		v1.GET("/quizzes", handler.GetQuizzesHandler)
		v1.POST("/quizzes", handler.PostQuizzesHandler)
		v1.GET("/quizzes/:quiz_number", handler.GetQuizHandler)
		v1.DELETE("/quizzes/:quiz_number", handler.DeleteQuizHandler)

		v1.GET("/quizzes/:quiz_number/tags", handler.GetQuizTagsHandler)
		v1.POST("/quizzes/:quiz_number/tags", handler.PostQuizTagsHandler)
		v1.DELETE("/quizzes/:quiz_number/tags/:tag_name", handler.DeleteQuizTagHandler)

		v1.GET("/tags", handler.GetTagsHandler)
		v1.POST("/tags", handler.PostTagsHandler)

		v1.GET("/answers", handler.GetAnswersHandler)
		v1.POST("/answers", handler.PostAnswersHandler)

		v1.GET("/provokes", handler.GetProvokesHandler)
		v1.POST("/provokes", handler.PostProvokesHandler)
	}
	router.Run()
}
