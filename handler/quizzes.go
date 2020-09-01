package handler

import (
	"fmt"
	"strconv"

	"github.com/ccns/quiz-server/db"
	"github.com/gin-gonic/gin"
)

// GetQuizzesHandler handles GET requests on route /quizzes.
func GetQuizzesHandler(ctx *gin.Context) {

	links := map[string]LinkDetail{
		"self": LinkDetail{"/v1/quizzes"},
	}

	tagName := ctx.DefaultQuery("tag", "")

	if tagName != "" {

		data, err := db.QueryQuizzes(tagName)
		if err != nil {
			resp, _ := BuildHATEOAS(links, Status{500, err.Error()}, nil, nil)
			ctx.String(500, resp)
			return
		}

		resp, _ := BuildHATEOAS(links, Status{200, "quizzes listed successfully."}, data, nil)
		ctx.String(200, resp)
		return
	}

	data, err := db.ListQuizzes()
	if err != nil {
		resp, _ := BuildHATEOAS(links, Status{500, err.Error()}, nil, nil)
		ctx.String(500, resp)
		return
	}

	resp, _ := BuildHATEOAS(links, Status{200, "quizzes listed successfully."}, data, nil)
	ctx.String(200, resp)
}

// PostQuizzesHandler handles POST requests on route /quizzes.
func PostQuizzesHandler(ctx *gin.Context) {

	links := map[string]LinkDetail{
		"self": LinkDetail{"/v1/quizzes"},
	}

	var quiz db.Quiz
	err := ctx.BindJSON(&quiz)
	if err != nil {
		resp, _ := BuildHATEOAS(links, Status{400, err.Error()}, nil, nil)
		ctx.String(400, resp)
		return
	}

	err = db.CreateQuiz(quiz)
	if err != nil {
		if err.Error() == fmt.Sprintf("quiz number %d already existed", quiz.Number) {
			resp, _ := BuildHATEOAS(links, Status{409, err.Error()}, quiz, nil)
			ctx.String(409, resp)
			return
		}
		resp, _ := BuildHATEOAS(links, Status{500, err.Error()}, quiz, nil)
		ctx.String(500, resp)
		return
	}
	links["quiz"] = LinkDetail{fmt.Sprintf("/v1/quizzes/%d", quiz.Number)}
	links["tags"] = LinkDetail{fmt.Sprintf("/v1/quizzes/%d/tags", quiz.Number)}

	status := Status{201, fmt.Sprintf("quiz number %d created successfully.", quiz.Number)}
	resp, _ := BuildHATEOAS(links, status, quiz, nil)
	ctx.String(201, resp)
}

// GetQuizHandler handles get requests on route /quizzes/:quiz_number.
func GetQuizHandler(ctx *gin.Context) {

	links := map[string]LinkDetail{
		"list": LinkDetail{"/v1/quizzes"},
	}

	quizNumber, err := strconv.Atoi(ctx.Param("quiz_number"))
	if err != nil {
		resp, _ := BuildHATEOAS(links, Status{400, err.Error()}, nil, nil)
		ctx.String(400, resp)
		return
	}

	data, err := db.GetQuiz(quizNumber)
	if err != nil {
		resp, _ := BuildHATEOAS(nil, Status{400, err.Error()}, nil, nil)
		ctx.String(400, resp)
		return
	}

	var tags []db.Tag
	tags, err = db.QueryTags(quizNumber)
	if err != nil {
		resp, _ := BuildHATEOAS(links, Status{500, err.Error()}, data, nil)
		ctx.String(500, resp)
		return
	}
	embedded := []HATEOAS{
		HATEOAS{
			LinksSection:    nil,
			StatusSection:   Status{200, fmt.Sprintf("tags listed successfully.")},
			DataSection:     tags,
			EmbeddedSection: nil,
		},
	}
	links["self"] = LinkDetail{fmt.Sprintf("/v1/quizzes/%d", quizNumber)}
	links["tags"] = LinkDetail{fmt.Sprintf("/v1/quizzes/%d/tags", quizNumber)}
	links["answers"] = LinkDetail{fmt.Sprintf("/v1/answers?quiz=%d", quizNumber)}

	status := Status{200, fmt.Sprintf("quiz number %d accessed successfully.", quizNumber)}
	resp, _ := BuildHATEOAS(links, status, data, embedded)
	ctx.String(200, resp)
}

// DeleteQuizHandler handles delete requests on route /quizzes/:quiz_number.
func DeleteQuizHandler(ctx *gin.Context) {

	links := map[string]LinkDetail{
		"list": LinkDetail{"/v1/quizzes"},
	}

	quizNumber, err := strconv.Atoi(ctx.Param("quiz_number"))
	if err != nil {
		resp, _ := BuildHATEOAS(links, Status{400, err.Error()}, nil, nil)
		ctx.String(400, resp)
		return
	}

	data, err := db.GetQuiz(quizNumber)
	if err != nil {
		resp, _ := BuildHATEOAS(links, Status{400, err.Error()}, nil, nil)
		ctx.String(400, resp)
		return
	}
	db.DeleteQuiz(quizNumber)
	status := Status{200, fmt.Sprintf("quiz number %d deleted successfully.", quizNumber)}
	resp, _ := BuildHATEOAS(links, status, data, nil)
	ctx.String(200, resp)
}

// GetQuizTagsHandler handles GET requests on route /quizzes/:quiz_number/tags.
func GetQuizTagsHandler(ctx *gin.Context) {

	links := map[string]LinkDetail{}

	quizNumber, err := strconv.Atoi(ctx.Param("quiz_number"))
	if err != nil {
		resp, _ := BuildHATEOAS(links, Status{400, err.Error()}, nil, nil)
		ctx.String(400, resp)
		return
	}

	var data []db.Tag
	data, err = db.QueryTags(quizNumber)
	if err != nil {
		resp, _ := BuildHATEOAS(links, Status{500, err.Error()}, nil, nil)
		ctx.String(500, resp)
		return
	}
	links["quiz"] = LinkDetail{fmt.Sprintf("/v1/quizzes/%d", quizNumber)}
	links["self"] = LinkDetail{fmt.Sprintf("/v1/quizzes/%d/tags", quizNumber)}
	links["answers"] = LinkDetail{fmt.Sprintf("/v1/answers?quiz=%d", quizNumber)}

	status := Status{200, fmt.Sprintf("tags of quiz number %d listed successfully.", quizNumber)}
	resp, _ := BuildHATEOAS(links, status, data, nil)
	ctx.String(200, resp)
}

// PostQuizTagsHandler handles POST requests on route /quizzes/:quiz_number/tags.
func PostQuizTagsHandler(ctx *gin.Context) {

	links := map[string]LinkDetail{}

	quizNumber, err := strconv.Atoi(ctx.Param("quiz_number"))
	if err != nil {
		resp, _ := BuildHATEOAS(links, Status{400, err.Error()}, nil, nil)
		ctx.String(400, resp)
		return
	}

	var tag db.Tag
	err = ctx.BindJSON(&tag)
	if err != nil {
		resp, _ := BuildHATEOAS(links, Status{400, err.Error()}, nil, nil)
		ctx.String(400, resp)
		return
	}

	_, err = db.GetTag(tag.Name)
	if err != nil {
		resp, _ := BuildHATEOAS(links, Status{400, err.Error()}, nil, nil)
		ctx.String(400, resp)
		return
	}

	err = db.RegisterTag(quizNumber, tag.Name)
	if err != nil {
		resp, _ := BuildHATEOAS(links, Status{500, err.Error()}, nil, nil)
		ctx.String(500, resp)
		return
	}
	links["quiz"] = LinkDetail{fmt.Sprintf("/v1/quizzes/%d", quizNumber)}
	links["self"] = LinkDetail{fmt.Sprintf("/v1/quizzes/%d/tags", quizNumber)}
	links["tag"] = LinkDetail{fmt.Sprintf("/v1/quizzes/%d/tags/%s", quizNumber, tag.Name)}
	links["related_quizzes"] = LinkDetail{fmt.Sprintf("/v1/quizzes?tag=%s", tag.Name)}
	links["answers"] = LinkDetail{fmt.Sprintf("/v1/answers?quiz=%d", quizNumber)}

	message := fmt.Sprintf("quiz number %d registered with tag %s successfully.", quizNumber, tag.Name)
	status := Status{201, message}
	resp, _ := BuildHATEOAS(nil, status, tag, nil)
	ctx.String(201, resp)
}

// DeleteQuizTagHandler handles delete requests on route /quizzes/:quiz_number/tags/:tag_name.
func DeleteQuizTagHandler(ctx *gin.Context) {

	links := map[string]LinkDetail{}

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

	err = db.DeregisterTag(quizNumber, tagName)
	if err != nil {
		resp, _ := BuildHATEOAS(links, Status{500, err.Error()}, nil, nil)
		ctx.String(500, resp)
		return
	}
	links["quiz"] = LinkDetail{fmt.Sprintf("/v1/quizzes/%d", quizNumber)}
	links["self"] = LinkDetail{fmt.Sprintf("/v1/quizzes/%d/tags", quizNumber)}
	links["tag"] = LinkDetail{fmt.Sprintf("/v1/quizzes/%d/tags/%s", quizNumber, tagName)}
	links["related_quizzes"] = LinkDetail{fmt.Sprintf("/v1/quizzes?tag=%s", tagName)}
	links["answers"] = LinkDetail{fmt.Sprintf("/v1/answers?quiz=%d", quizNumber)}

	message := fmt.Sprintf("quiz number %d deregistered with tag %s successfully.", quizNumber, tagName)
	status := Status{200, message}
	resp, _ := BuildHATEOAS(nil, status, tagName, nil)
	ctx.String(200, resp)
}
