package handler

import (
	"github.com/ccns/quiz-server/db"
	"github.com/gin-gonic/gin"
)

// GetRankHandler handles get requests on route /rank.
func GetRankHandler(ctx *gin.Context) {

	links := map[string]LinkDetail{
		"self": LinkDetail{"/v1/rank"},
	}

	data, err := db.ListPlayersRank()
	if err != nil {
		status := Status{500, err.Error()}
		resp, _ := BuildHATEOAS(links, status, nil, nil)
		ctx.String(500, resp)
		return
	}

	resp, _ := BuildHATEOAS(nil, Status{200, "rank accessed successfully."}, data, nil)
	ctx.String(200, resp)
}
