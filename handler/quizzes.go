package handler

import (
	"fmt"
	"strconv"

	"github.com/ccns/quiz-server/db"
	"github.com/gin-gonic/gin"
)

// GetQuizzesHandler handles get requests on route /quizzes.
func GetQuizzesHandler(ctx *gin.Context) {

	data, err := db.ListQuizzes()
	if err != nil {
		status := Status{500, err.Error()}
		resp, _ := BuildHATEOAS(nil, status, nil, nil)
		ctx.String(500, resp)
		return
	}

	status := Status{200, "all quizzes listed successfully."}
	resp, _ := BuildHATEOAS(nil, status, data, nil)
	ctx.String(200, resp)
}

// PostQuizzesHandler handles post requests on route /quizzes.
func PostQuizzesHandler(ctx *gin.Context) {

	var quiz db.Quiz
	err := ctx.BindJSON(&quiz)
	if err != nil {
		status := Status{400, err.Error()}
		resp, _ := BuildHATEOAS(nil, status, nil, nil)
		ctx.String(400, resp)
		return
	}

	err = db.CreateQuiz(quiz)
	if err != nil {
		status := Status{500, err.Error()}
		resp, _ := BuildHATEOAS(nil, status, quiz, nil)
		ctx.String(500, resp)
		return
	}

	status := Status{201, fmt.Sprintf("quiz number %d created successfully.", quiz.Number)}
	resp, _ := BuildHATEOAS(nil, status, quiz, nil)
	ctx.String(201, resp)
}

// UpdateQuizHandler handles update requests on route /quizzes/<quiz_number>.
func UpdateQuizHandler(ctx *gin.Context) {}

// DeleteQuizHandler handles delete requests on route /quizzes/<quiz_number>.
func DeleteQuizHandler(ctx *gin.Context) {

	number, err := strconv.Atoi(ctx.Param("quiz_number"))
	if err != nil {
		status := Status{400, err.Error()}
		resp, _ := BuildHATEOAS(nil, status, nil, nil)
		ctx.String(400, resp)
		return
	}

	data, err := db.GetQuiz(number)
	if err != nil {
		status := Status{400, err.Error()}
		resp, _ := BuildHATEOAS(nil, status, nil, nil)
		ctx.String(400, resp)
		return
	}
	db.DeleteQuiz(number)
	status := Status{200, fmt.Sprintf("quiz number %d deleted successfully.", number)}
	resp, _ := BuildHATEOAS(nil, status, data, nil)
	ctx.String(200, resp)
}

// GetQuizTagsHandler handles get requests on route /quizzes/<quiz_number>/tags.
func GetQuizTagsHandler(ctx *gin.Context) {

	number, err := strconv.Atoi(ctx.Param("quiz_number"))
	if err != nil {
		status := Status{400, err.Error()}
		resp, _ := BuildHATEOAS(nil, status, nil, nil)
		ctx.String(400, resp)
		return
	}

	quizFound, err := db.GetQuiz(number)
	if err != nil {
		status := Status{400, err.Error()}
		resp, _ := BuildHATEOAS(nil, status, nil, nil)
		ctx.String(400, resp)
		return
	}

	var data []db.Tag
	data, err = db.QueryTags(quizFound.Number)
	if err != nil {
		status := Status{500, err.Error()}
		resp, _ := BuildHATEOAS(nil, status, nil, nil)
		ctx.String(500, resp)
		return
	}

	status := Status{200, fmt.Sprintf("tags of quiz number %d listed successfully.", quizFound.Number)}
	resp, _ := BuildHATEOAS(nil, status, data, nil)
	ctx.String(200, resp)
}

// PostQuizTagsHandler handles post requests on route /quizzes/<quiz_number>/tags.
func PostQuizTagsHandler(ctx *gin.Context) {

	number, err := strconv.Atoi(ctx.Param("quiz_number"))
	if err != nil {
		status := Status{400, err.Error()}
		resp, _ := BuildHATEOAS(nil, status, nil, nil)
		ctx.String(400, resp)
		return
	}

	quizFound, err := db.GetQuiz(number)
	if err != nil {
		status := Status{400, err.Error()}
		resp, _ := BuildHATEOAS(nil, status, nil, nil)
		ctx.String(400, resp)
		return
	}

	var tag db.Tag
	err = ctx.BindJSON(&tag)
	if err != nil {
		status := Status{400, err.Error()}
		resp, _ := BuildHATEOAS(nil, status, nil, nil)
		ctx.String(400, resp)
		return
	}

	tagFound, err := db.GetTag(tag.Name)
	if err != nil {
		status := Status{400, err.Error()}
		resp, _ := BuildHATEOAS(nil, status, nil, nil)
		ctx.String(400, resp)
		return
	}

	db.RegisterTag(quizFound.Number, tagFound.Name)
	message := fmt.Sprintf("quiz number %d registered with tag %s successfully.", quizFound.Number, tagFound.Name)
	status := Status{200, message}
	resp, _ := BuildHATEOAS(nil, status, tag, nil)
	ctx.String(200, resp)
}

// DeleteQuizTagHandler handles delete requests on route /quizzes/<quiz_number>/tags/<tag_name>.
func DeleteQuizTagHandler(ctx *gin.Context) {

	number, err := strconv.Atoi(ctx.Param("quiz_number"))
	if err != nil {
		status := Status{400, err.Error()}
		resp, _ := BuildHATEOAS(nil, status, nil, nil)
		ctx.String(400, resp)
		return
	}

	quizFound, err := db.GetQuiz(number)
	if err != nil {
		status := Status{400, err.Error()}
		resp, _ := BuildHATEOAS(nil, status, nil, nil)
		ctx.String(400, resp)
		return
	}

	tagName := ctx.Param("tag_name")
	if tagName == "" {
		status := Status{400, "tag name not specified."}
		resp, _ := BuildHATEOAS(nil, status, nil, nil)
		ctx.String(400, resp)
		return
	}

	tagFound, err := db.GetTag(tagName)
	if err != nil {
		status := Status{400, err.Error()}
		resp, _ := BuildHATEOAS(nil, status, nil, nil)
		ctx.String(400, resp)
		return
	}

	db.DeregisterTag(quizFound.Number, tagFound.Name)
	message := fmt.Sprintf("quiz number %d deregistered with tag %s successfully.", quizFound.Number, tagFound.Name)
	status := Status{200, message}
	resp, _ := BuildHATEOAS(nil, status, tagFound.Name, nil)
	ctx.String(200, resp)
}
