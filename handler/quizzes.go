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
		resp, _ := BuildHATEOAS(nil, Status{500, err.Error()}, nil, nil)
		ctx.String(500, resp)
		return
	}

	status := Status{200, "quizzes listed successfully."}
	resp, _ := BuildHATEOAS(nil, status, data, nil)
	ctx.String(200, resp)
}

// PostQuizzesHandler handles post requests on route /quizzes.
func PostQuizzesHandler(ctx *gin.Context) {

	var quiz db.Quiz
	err := ctx.BindJSON(&quiz)
	if err != nil {
		resp, _ := BuildHATEOAS(nil, Status{400, err.Error()}, nil, nil)
		ctx.String(400, resp)
		return
	}

	err = db.CreateQuiz(quiz)
	if err != nil {
		resp, _ := BuildHATEOAS(nil, Status{500, err.Error()}, quiz, nil)
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

	quizNumber, err := strconv.Atoi(ctx.Param("quiz_number"))
	if err != nil {
		resp, _ := BuildHATEOAS(nil, Status{400, err.Error()}, nil, nil)
		ctx.String(400, resp)
		return
	}

	data, err := db.GetQuiz(quizNumber)
	if err != nil {
		resp, _ := BuildHATEOAS(nil, Status{400, err.Error()}, nil, nil)
		ctx.String(400, resp)
		return
	}
	db.DeleteQuiz(quizNumber)
	status := Status{200, fmt.Sprintf("quiz number %d deleted successfully.", quizNumber)}
	resp, _ := BuildHATEOAS(nil, status, data, nil)
	ctx.String(200, resp)
}

// GetQuizTagsHandler handles get requests on route /quizzes/<quiz_number>/tags.
func GetQuizTagsHandler(ctx *gin.Context) {

	quizNumber, err := strconv.Atoi(ctx.Param("quiz_number"))
	if err != nil {
		resp, _ := BuildHATEOAS(nil, Status{400, err.Error()}, nil, nil)
		ctx.String(400, resp)
		return
	}

	var data []db.Tag
	data, err = db.QueryTags(quizNumber)
	if err != nil {
		resp, _ := BuildHATEOAS(nil, Status{500, err.Error()}, nil, nil)
		ctx.String(500, resp)
		return
	}

	status := Status{200, fmt.Sprintf("tags of quiz number %d listed successfully.", quizNumber)}
	resp, _ := BuildHATEOAS(nil, status, data, nil)
	ctx.String(200, resp)
}

// PostQuizTagsHandler handles post requests on route /quizzes/<quiz_number>/tags.
func PostQuizTagsHandler(ctx *gin.Context) {

	quizNumber, err := strconv.Atoi(ctx.Param("quiz_number"))
	if err != nil {
		resp, _ := BuildHATEOAS(nil, Status{400, err.Error()}, nil, nil)
		ctx.String(400, resp)
		return
	}

	var tag db.Tag
	err = ctx.BindJSON(&tag)
	if err != nil {
		resp, _ := BuildHATEOAS(nil, Status{400, err.Error()}, nil, nil)
		ctx.String(400, resp)
		return
	}

	db.RegisterTag(quizNumber, tag.Name)
	message := fmt.Sprintf("quiz number %d registered with tag %s successfully.", quizNumber, tag.Name)
	status := Status{200, message}
	resp, _ := BuildHATEOAS(nil, status, tag, nil)
	ctx.String(200, resp)
}

// DeleteQuizTagHandler handles delete requests on route /quizzes/<quiz_number>/tags/<tag_name>.
func DeleteQuizTagHandler(ctx *gin.Context) {

	quizNumber, err := strconv.Atoi(ctx.Param("quiz_number"))
	if err != nil {
		resp, _ := BuildHATEOAS(nil, Status{400, err.Error()}, nil, nil)
		ctx.String(400, resp)
		return
	}

	tagName := ctx.Param("tag_name")
	if tagName == "" {
		resp, _ := BuildHATEOAS(nil, Status{400, "tag name not specified."}, nil, nil)
		ctx.String(400, resp)
		return
	}

	db.DeregisterTag(quizNumber, tagName)
	message := fmt.Sprintf("quiz number %d deregistered with tag %s successfully.", quizNumber, tagName)
	status := Status{200, message}
	resp, _ := BuildHATEOAS(nil, status, tagName, nil)
	ctx.String(200, resp)
}
