package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/ccns/quiz-server/db"
)

// QuizLoad describe the struct of external quizzes file
type QuizLoad struct {
	Number      int      `json:"number"`
	Description string   `json:"description"`
	Score       int      `json:"score"`
	Options     []string `json:"options"`
	Answer      string   `json:"answer"`
	Tags        []string `json:"tags"`
}

var quizzes []QuizLoad

func init() {

	readQuizzes()
	loadQuizzes()
}

func readQuizzes() {

	jsonFile, err := os.Open("quizzes.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &quizzes)
}

func loadQuizzes() {

	for _, q := range quizzes {
		quiz := db.Quiz{
			Number:      q.Number,
			Description: q.Description,
			Score:       q.Score,
			OptionA:     q.Options[0],
			OptionB:     q.Options[1],
			OptionC:     q.Options[2],
			OptionD:     q.Options[3],
			Answer:      q.Answer,
		}
		err := db.CreateQuiz(quiz)
		if err != nil {
			log.Fatal(err)
		}

		for _, t := range q.Tags {
			err := db.RegisterTag(q.Number, t)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
