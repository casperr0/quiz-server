package handler

import (
	"github.com/ccns/quiz-server/db"
	"github.com/gin-gonic/gin"
)

// GetProvokesHandler handlers GET requests on route /provokes.
func GetProvokesHandler(ctx *gin.Context) {

	links := map[string]LinkDetail{
		"self": LinkDetail{"/v1/provokes"},
	}

	correctness := ctx.DefaultQuery("correct", "")

	if correctness != "" {

		if correctness == "true" {

			data, err := db.QueryProvokes(true)
			if err != nil {
				resp, _ := BuildHATEOAS(links, Status{500, err.Error()}, nil, nil)
				ctx.String(500, resp)
				return
			}

			status := Status{200, "provokes listed successfully."}
			resp, _ := BuildHATEOAS(links, status, data, nil)
			ctx.String(200, resp)
			return

		} else if correctness == "false" {

			data, err := db.QueryProvokes(false)
			if err != nil {
				resp, _ := BuildHATEOAS(links, Status{500, err.Error()}, nil, nil)
				ctx.String(500, resp)
				return
			}

			status := Status{200, "provokes listed successfully."}
			resp, _ := BuildHATEOAS(links, status, data, nil)
			ctx.String(200, resp)
			return
		}

		status := Status{400, "correctness should be true or false."}
		resp, _ := BuildHATEOAS(links, status, nil, nil)
		ctx.String(400, resp)
		return
	}

	data, err := db.ListProvokes()
	if err != nil {
		resp, _ := BuildHATEOAS(links, Status{500, err.Error()}, nil, nil)
		ctx.String(500, resp)
		return
	}

	status := Status{200, "provokes listed successfully."}
	resp, _ := BuildHATEOAS(links, status, data, nil)
	ctx.String(200, resp)
}

// PostProvokesHandler handlers POST requests on route /provokes.
func PostProvokesHandler(ctx *gin.Context) {

	links := map[string]LinkDetail{
		"self": LinkDetail{"/v1/provokes"},
	}

	var provoke db.Provoke
	err := ctx.BindJSON(&provoke)
	if err != nil {
		resp, _ := BuildHATEOAS(links, Status{400, err.Error()}, nil, nil)
		ctx.String(400, resp)
		return
	}

	err = db.CreateProvoke(provoke.Correct, provoke.Message)
	if err != nil {
		resp, _ := BuildHATEOAS(links, Status{400, err.Error()}, nil, nil)
		ctx.String(400, resp)
		return
	}

	status := Status{201, "provoke created successfully."}
	resp, _ := BuildHATEOAS(links, status, provoke, nil)
	ctx.String(201, resp)
}
