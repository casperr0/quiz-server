package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	provokes []Provoke = []Provoke{
		Provoke{
			ID:      1,
			Correct: true,
			Message: "correct provoke 1",
		},
		Provoke{
			ID:      2,
			Correct: true,
			Message: "correct provoke 2",
		},
		Provoke{
			ID:      3,
			Correct: true,
			Message: "correct provoke 3",
		},
		Provoke{
			ID:      4,
			Correct: false,
			Message: "incorrect provoke 1",
		},
		Provoke{
			ID:      5,
			Correct: false,
			Message: "incorrect provoke 2",
		},
		Provoke{
			ID:      6,
			Correct: false,
			Message: "incorrect provoke 3",
		},
	}
)

func TestCreateProvoke(t *testing.T) {

	for _, p := range provokes {
		err := CreateProvoke(p.Correct, p.Message)
		assert.Nil(t, err)
	}
}

func TestListProvokes(t *testing.T) {

	result, err := QueryProvokes(true)
	assert.Nil(t, err)
	assert.Equal(t, provokes[:3], result)

	result, err = QueryProvokes(false)
	assert.Nil(t, err)
	assert.Equal(t, provokes[3:], result)
}
