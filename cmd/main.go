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
	r.GET("/quizzes", handler.GetQuizzesHandler)
	r.POST("/quizzes", handler.PostQuizzesHandler)
	r.DELETE("/quizzes/:quiz_number", handler.DeleteQuizHandler)
	r.GET("/quizzes/:quiz_number/tags", handler.GetQuizTagsHandler)
	r.POST("/quizzes/:quiz_number/tags", handler.PostQuizTagsHandler)
	r.DELETE("/quizzes/:quiz_number/tags/:tag_name", handler.DeleteQuizTagHandler)
	r.Run()
}
