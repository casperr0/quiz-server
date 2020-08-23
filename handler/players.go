package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/ccns/quiz-server/db"
)

// GetPlayersHandler handles get requests on route /players.
func GetPlayersHandler(ctx *gin.Context) {

	data, err := db.ListPlayersOrderByScore()
	if err != nil {
		status := Status{500, err.Error()}
		resp, _ := BuildHATEOAS(nil, status, nil, nil)
		ctx.String(500, resp)
		return
	}

	status := Status{200, "players listed successfully."}
	resp, _ := BuildHATEOAS(nil, status, data, nil)
	ctx.String(200, resp)
}

// PostPlayersHandler handles post requests on route /players.
func PostPlayersHandler(ctx *gin.Context) {

	var player db.Player
	err := ctx.BindJSON(&player)
	if err != nil {
		status := Status{400, err.Error()}
		resp, _ := BuildHATEOAS(nil, status, nil, nil)
		ctx.String(400, resp)
		return
	}

	err = db.CreatePlayer(player.Name)
	if err != nil {
		status := Status{500, err.Error()}
		if err == fmt.Errorf("player %s already existed", player.Name) {
			status = Status{409, err.Error()}
		}
		resp, _ := BuildHATEOAS(nil, status, player, nil)
		ctx.String(500, resp)
		return
	}

	status := Status{201, fmt.Sprintf("player %s created successfully.", player.Name)}
	resp, _ := BuildHATEOAS(nil, status, player, nil)
	ctx.String(201, resp)
}

// DeletePlayerHandler handles delete requests on route /players/<player_name>.
func DeletePlayerHandler(ctx *gin.Context) {

	playerName := ctx.Param("player_name")
	if playerName == "" {
		status := Status{400, "player_name not specified."}
		resp, _ := BuildHATEOAS(nil, status, nil, nil)
		ctx.String(400, resp)
		return
	}

	data, err := db.GetPlayer(playerName)
	if err != nil {
		status := Status{400, err.Error()}
		resp, _ := BuildHATEOAS(nil, status, nil, nil)
		ctx.String(400, resp)
		return
	}
	db.DeletePlayer(playerName)
	status := Status{200, fmt.Sprintf("player %s deleted successfully.", playerName)}
	resp, _ := BuildHATEOAS(nil, status, data, nil)
	ctx.String(200, resp)
}
