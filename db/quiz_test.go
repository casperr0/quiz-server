package db

import (
	"testing"

	"github.com/ccns/quiz-server/config"
	"github.com/stretchr/testify/assert"
)

var (
	quizes []Quiz = []Quiz{
		Quiz{
			ID:          1,
			Description: "quiz 1 description",
			Score:       1,
			OptionA:     "A",
			OptionB:     "B",
			OptionC:     "C",
			OptionD:     "D",
			Answer:      "A",
		},
		Quiz{
			ID:          2,
			Description: "OI!HONLQ!)@()!*&",
			Score:       -1,
			OptionA:     ")!(H)",
			OptionB:     "(Q*WGO!HO@",
			OptionC:     ")A(H)IHQO",
			OptionD:     "(*QGWU!@J",
			Answer:      "B",
		},
		Quiz{
			ID:          3,
			Description: "",
			Score:       999,
			OptionA:     "",
			OptionB:     "",
			OptionC:     "",
			OptionD:     "",
			Answer:      "C",
		},
	}
)

func TestCreateQuiz(t *testing.T) {

	for _, q := range quizes {
		err := CreateQuiz(q)
		assert.Nil(t, err)
		result, err := GetQuiz(q.ID)
		assert.Nil(t, err)
		assert.Equal(t, *result, q)
	}
}

func TestGetQuiz(t *testing.T) {

	for _, q := range quizes {
		result, err := GetQuiz(q.ID)
		assert.Nil(t, err)
		assert.Equal(t, *result, q)
	}
}

func TestListQuizes(t *testing.T) {

	result, err := ListQuizes()
	assert.Nil(t, err)
	for i, r := range result {
		assert.Equal(t, quizes[i], r, "Quiz not match.")
	}
}

func TestGetTag(t *testing.T) {

	for _, tag := range config.Config.Quiz.DefaultTags {
		result, err := GetTag(tag)
		assert.Nil(t, err)
		assert.Equal(t, result.Name, tag)
	}
}

func TestListTags(t *testing.T) {

	result, err := ListTags()
	assert.Nil(t, err)
	tags := config.Config.Quiz.DefaultTags
	for i, r := range result {
		assert.Equal(t, tags[i], r.Name, "Tag not match.")
	}
}

func TestRegisterTag(t *testing.T) {

	tags := config.Config.Quiz.DefaultTags
	for _, q := range quizes {
		RegisterTag(q.ID, tags[0])
	}
	for _, tag := range tags {
		RegisterTag(quizes[0].ID, tag)
	}

	resultQuizes, err := QueryQuizes(tags[0])
	assert.Nil(t, err)
	for i, q := range resultQuizes {
		assert.Equal(t, quizes[i], q)
	}

	resultTags, err := QueryTags(quizes[0].ID)
	assert.Nil(t, err)
	for i, tag := range resultTags {
		assert.Equal(t, tags[i], tag.Name)
	}
}

func TestDeregisterTag(t *testing.T) {

	// deregister the tag with id 0 from all quizes
	tags := config.Config.Quiz.DefaultTags
	for _, q := range quizes {
		DeregisterTag(q.ID, tags[0])
	}

	// check the first quiz has all tags exclude the tag with id 0
	resultTags, err := QueryTags(quizes[0].ID)
	assert.Nil(t, err)
	for i, tag := range resultTags {
		assert.Equal(t, tags[i+1], tag.Name)
	}

	// deregister all tags with ID greater then 3 from the first quiz
	for _, tag := range tags[3:] {
		DeregisterTag(quizes[0].ID, tag)
	}

	// check the first quiz has only two tags with id 1 and 2
	resultTags, err = QueryTags(quizes[0].ID)
	assert.Nil(t, err)
	for i, tag := range tags[1:3] {
		assert.Equal(t, tag, resultTags[i].Name)
	}
}
