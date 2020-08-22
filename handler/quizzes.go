package handler

import "github.com/gin-gonic/gin"

// GetQuizzesHandler handles get requests on route /quizzes.
func GetQuizzesHandler(ctx *gin.Context) {}

// PostQuizzesHandler handles post requests on route /quizzes.
func PostQuizzesHandler(ctx *gin.Context) {}

// UpdateQuizHandler handles update requests on route /quizzes/<quiz_id>.
func UpdateQuizHandler(ctx *gin.Context) {}

// DeleteQuizHandler handles delete requests on route /quizzes/<quiz_id>.
func DeleteQuizHandler(ctx *gin.Context) {}

// GetQuizTagsHandler handles get requests on route /quizzes/<quiz_id>/tags.
func GetQuizTagsHandler(ctx *gin.Context) {}

// PostQuizTagsHandler handles post requests on route /quizzes/<quiz_id>/tags.
func PostQuizTagsHandler(ctx *gin.Context) {}

// DeleteQuizTagHandler handles delete requests on route /quizzes/<quiz_id>/tags/<tag_name>.
func DeleteQuizTagHandler(ctx *gin.Context) {}
