package handler

import (
	"fmt"
	"math/rand"

	"github.com/gin-gonic/gin"

	"github.com/ccns/quiz-server/db"
)

// Profile describe the schema of playser profile.
type Profile struct {
	ID       int    `json:"_id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Platform string `json:"platform"`
	Rank     int    `json:"rank"`
	Score    int    `json:"score"`
	Last     int    `json:"last"`
}

// GetPlayersHandler handles get requests on route /players.
func GetPlayersHandler(ctx *gin.Context) {

	links := map[string]LinkDetail{
		"self": LinkDetail{"/v1/players"},
	}

	data, err := db.ListPlayers()
	if err != nil {
		status := Status{500, err.Error()}
		resp, _ := BuildHATEOAS(links, status, nil, nil)
		ctx.String(500, resp)
		return
	}

	status := Status{200, "players listed successfully."}
	resp, _ := BuildHATEOAS(nil, status, data, nil)
	ctx.String(200, resp)
}

// PostPlayersHandler handles post requests on route /players.
func PostPlayersHandler(ctx *gin.Context) {

	links := map[string]LinkDetail{
		"self": LinkDetail{"/v1/players"},
	}

	var player db.Player
	err := ctx.BindJSON(&player)
	if err != nil {
		status := Status{400, err.Error()}
		resp, _ := BuildHATEOAS(links, status, nil, nil)
		ctx.String(400, resp)
		return
	}
	links["player"] = LinkDetail{fmt.Sprintf("/v1/players/%s", player.Name)}

	err = db.CreatePlayer(player.Name, player.Nickname, player.Platform)
	if err != nil {
		if err.Error() == fmt.Sprintf("player %s already existed", player.Name) {
			resp, _ := BuildHATEOAS(links, Status{409, err.Error()}, player, nil)
			ctx.String(409, resp)
			return
		}
		resp, _ := BuildHATEOAS(links, Status{500, err.Error()}, player, nil)
		ctx.String(500, resp)
		return
	}
	links["feed"] = LinkDetail{fmt.Sprintf("/v1/players/%s/feed", player.Name)}

	status := Status{201, fmt.Sprintf("player %s created successfully.", player.Name)}
	resp, _ := BuildHATEOAS(links, status, player, nil)
	ctx.String(201, resp)
}

// GetPlayerHandler handles get requests on route /players/:player_name.
func GetPlayerHandler(ctx *gin.Context) {

	links := map[string]LinkDetail{
		"list": LinkDetail{"/v1/players"},
	}

	playerName := ctx.Param("player_name")
	if playerName == "" {
		resp, _ := BuildHATEOAS(links, Status{400, "player_name not specified."}, nil, nil)
		ctx.String(400, resp)
		return
	}

	playerData, err := db.GetPlayer(playerName)
	if err != nil {
		resp, _ := BuildHATEOAS(links, Status{400, err.Error()}, nil, nil)
		ctx.String(400, resp)
		return
	}
	data := Profile{
		ID:       playerData.ID,
		Name:     playerData.Name,
		Nickname: playerData.Nickname,
		Platform: playerData.Platform,
	}

	playerList, err := db.ListPlayersRank()
	data.Rank = len(playerList) + 1
	for i, p := range playerList {
		if p.Name == playerData.Name {
			data.Rank = i + 1
			data.Score = p.Score
		}
	}

	playerAnswer, _ := db.QueryQuizAnswersByPlayer(playerData.Name)
	quizList, _ := db.ListQuizzes()
	data.Last = len(quizList) - len(playerAnswer)
	links["self"] = LinkDetail{fmt.Sprintf("/v1/players/%s", playerName)}
	links["answers"] = LinkDetail{fmt.Sprintf("/v1/answers?player=%s", playerName)}

	resp, _ := BuildHATEOAS(links, Status{200, "player accessed successfully."}, data, nil)
	ctx.String(200, resp)
}

// DeletePlayerHandler handles delete requests on route /players/:player_name.
func DeletePlayerHandler(ctx *gin.Context) {

	links := map[string]LinkDetail{
		"list": LinkDetail{"/players"},
	}

	playerName := ctx.Param("player_name")
	if playerName == "" {
		resp, _ := BuildHATEOAS(links, Status{400, "player_name not specified."}, nil, nil)
		ctx.String(400, resp)
		return
	}

	data, err := db.GetPlayer(playerName)
	if err != nil {
		resp, _ := BuildHATEOAS(links, Status{400, err.Error()}, nil, nil)
		ctx.String(400, resp)
		return
	}
	db.DeletePlayer(playerName)
	resp, _ := BuildHATEOAS(links, Status{200, "player deleted successfully."}, data, nil)
	ctx.String(200, resp)
}

// GetPlayerFeedHandler handles GET requests on route /players/:player_name/feed.
func GetPlayerFeedHandler(ctx *gin.Context) {

	links := map[string]LinkDetail{
		"list": LinkDetail{"/v1/players"},
	}

	playerName := ctx.Param("player_name")
	if playerName == "" {
		resp, _ := BuildHATEOAS(links, Status{400, "player_name not specified."}, nil, nil)
		ctx.String(400, resp)
		return
	}
	links["player"] = LinkDetail{fmt.Sprintf("/v1/players/%s", playerName)}
	links["rand"] = LinkDetail{fmt.Sprintf("/v1/players/%s/rand", playerName)}
	links["self"] = LinkDetail{fmt.Sprintf("/v1/players/%s/feed", playerName)}

	numbers, err := db.FeedQuizzes(playerName)
	if err != nil {
		if err.Error() == "no quiz left" {
			resp, _ := BuildHATEOAS(links, Status{200, err.Error()}, nil, nil)
			ctx.String(200, resp)
			return
		}
		resp, _ := BuildHATEOAS(links, Status{500, err.Error()}, nil, nil)
		ctx.String(500, resp)
		return
	}
	randomQuizNumber := numbers[rand.Intn(len(numbers))].ID

	data, err := db.GetQuiz(randomQuizNumber)
	if err != nil {
		resp, _ := BuildHATEOAS(nil, Status{400, err.Error()}, nil, nil)
		ctx.String(400, resp)
		return
	}
	resp, _ := BuildHATEOAS(links, Status{200, "player fed successfully."}, data, nil)
	ctx.String(200, resp)
}

// GetPlayerRandHandler handles GET requests on route /players/:player_name/rand.
func GetPlayerRandHandler(ctx *gin.Context) {

	links := map[string]LinkDetail{
		"list": LinkDetail{"/v1/quizzes"},
	}
	playerName := ctx.Param("player_name")
	if playerName == "" {
		resp, _ := BuildHATEOAS(links, Status{400, "player_name not specified."}, nil, nil)
		ctx.String(400, resp)
		return
	}
	links["player"] = LinkDetail{fmt.Sprintf("/v1/players/%s", playerName)}
	links["feed"] = LinkDetail{fmt.Sprintf("/v1/players/%s/feed", playerName)}
	links["self"] = LinkDetail{fmt.Sprintf("/v1/players/%s/rand", playerName)}

	quizzes, err := db.ListQuizzes()
	if err != nil {
		resp, _ := BuildHATEOAS(links, Status{500, err.Error()}, nil, nil)
		ctx.String(500, resp)
		return
	}
	randomQuizNumber := quizzes[rand.Intn(len(quizzes))].ID

	data, err := db.GetQuiz(randomQuizNumber)
	if err != nil {
		resp, _ := BuildHATEOAS(nil, Status{400, err.Error()}, nil, nil)
		ctx.String(400, resp)
		return
	}
	resp, _ := BuildHATEOAS(links, Status{200, "quiz accessed successfully."}, data, nil)
	ctx.String(200, resp)
}
