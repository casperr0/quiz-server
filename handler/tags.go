package handler

import (
	"fmt"

	"github.com/ccns/quiz-server/db"
	"github.com/gin-gonic/gin"
)

// GetTagsHandler handlers GET requests on route /tags.
func GetTagsHandler(ctx *gin.Context) {

	links := map[string]LinkDetail{
		"self": LinkDetail{"/v1/tags"},
	}

	data, err := db.ListTags()
	if err != nil {
		resp, _ := BuildHATEOAS(links, Status{500, err.Error()}, nil, nil)
		ctx.String(500, resp)
		return
	}

	status := Status{200, "tags listed successfully."}
	resp, _ := BuildHATEOAS(links, status, data, nil)
	ctx.String(200, resp)
}

// PostTagsHandler handlers POST requests on route /tags.
func PostTagsHandler(ctx *gin.Context) {

	links := map[string]LinkDetail{
		"self": LinkDetail{"/v1/tags"},
	}

	var tag db.Tag
	err := ctx.BindJSON(&tag)
	if err != nil {
		resp, _ := BuildHATEOAS(links, Status{400, err.Error()}, nil, nil)
		ctx.String(400, resp)
		return
	}

	err = db.CreateTag(tag.Name)
	if err != nil {
		if err.Error() == fmt.Sprintf("tag %s already existed", tag.Name) {
			resp, _ := BuildHATEOAS(links, Status{409, err.Error()}, nil, nil)
			ctx.String(409, resp)
			return
		}
		resp, _ := BuildHATEOAS(links, Status{500, err.Error()}, nil, nil)
		ctx.String(500, resp)
		return
	}

	status := Status{201, fmt.Sprintf("tag %s created successfully.", tag.Name)}
	resp, _ := BuildHATEOAS(links, status, tag, nil)
	ctx.String(201, resp)
}
