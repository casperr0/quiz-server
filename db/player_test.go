package db

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	players []Player = []Player{
		Player{
			ID:   1,
			Name: "rain",
		},
		Player{
			ID:   2,
			Name: "!)*HU#)(!*UQoudqo",
		},
		Player{
			ID:   3,
			Name: "",
		},
		Player{
			ID:   4,
			Name: "佔榜魔法師",
		},
	}
)

// Put the TestCreateQuiz here because test cases for player answer needs initial quizes.
func TestCreateQuiz(t *testing.T) {

	for _, q := range quizzes {
		err := CreateQuiz(q)
		assert.Nil(t, err)
		result, err := GetQuiz(q.Number)
		assert.Nil(t, err)
		assert.Equal(t, *result, q)
	}

	err := CreateQuiz(quizzes[0])
	tpl := "quiz number %d already exists"
	assert.Equal(t, err, fmt.Errorf(tpl, quizzes[0].Number))
}

func TestCreatePlayer(t *testing.T) {

	for _, p := range players {
		err := CreatePlayer(p.Name)
		assert.Nil(t, err)
		result, err := GetPlayer(p.Name)
		assert.Nil(t, err)
		assert.Equal(t, *result, p)
	}
}

func TestGetPlayer(t *testing.T) {

	for _, p := range players {
		result, err := GetPlayer(p.Name)
		assert.Nil(t, err)
		assert.Equal(t, *result, p)
	}
}

func TestRegisterAnswer(t *testing.T) {

	for iq, q := range quizzes {
		for ip, p := range players {
			err := RegisterAnswer(p.Name, q.ID, ip > iq)
			assert.Nil(t, err)
		}
	}
}

func TestListPlayersOrderByScore(t *testing.T) {

	result, err := ListPlayersOrderByScore()
	assert.Nil(t, err)
	for i, r := range result {
		if i > 0 {
			assert.GreaterOrEqual(t, result[i-1].Score, r.Score)
		}
	}
}

func TestQueryPlayerAnswersByQuiz(t *testing.T) {

	for iq, q := range quizzes {
		result, err := QueryPlayerAnswersByQuiz(q.ID)
		assert.Nil(t, err)
		for ir, r := range result {
			assert.Equal(t, r.Correct, ir > iq)
		}
	}
}

func TestQueryQuizAnswersByPlayer(t *testing.T) {

	for ip, p := range players {
		result, err := QueryQuizAnswersByPlayer(p.Name)
		assert.Nil(t, err)
		for ir, r := range result {
			assert.Equal(t, r.Correct, ir < ip)
		}
	}
}
