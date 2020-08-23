package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"

	_ "github.com/ccns/quiz-server/db"
	"github.com/ccns/quiz-server/handler"
	_ "github.com/ccns/quiz-server/utils"
	// "github.com/ccns/quiz-server/config"
)

func main() {

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/players", handler.GetPlayersHandler)
		v1.POST("/players", handler.PostPlayersHandler)
		v1.DELETE("/players/:player_name", handler.DeletePlayerHandler)

		v1.GET("/quizzes", handler.GetQuizzesHandler)
		v1.POST("/quizzes", handler.PostQuizzesHandler)
		v1.DELETE("/quizzes/:quiz_number", handler.DeleteQuizHandler)

		v1.GET("/quizzes/:quiz_number/tags", handler.GetQuizTagsHandler)
		v1.POST("/quizzes/:quiz_number/tags", handler.PostQuizTagsHandler)
		v1.DELETE("/quizzes/:quiz_number/tags/:tag_name", handler.DeleteQuizTagHandler)

		v1.GET("/answers", handler.GetAnswersHandler)
		v1.POST("/answers", handler.PostAnswersHandler)

		v1.GET("/provokes", handler.GetProvokesHandler)
		v1.POST("/provokes", handler.PostAnswersHandler)
	}
	r.Run()
}
