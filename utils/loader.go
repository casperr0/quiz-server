package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/ccns/quiz-server/db"
)

// ProvokeLoad describe the struct of external provoke file
type ProvokeLoad struct {
	Correct   []string `json:"correct"`
	Incorrect []string `json:"incorrect"`
}

// TagLoad describe the struct of external tags file
type TagLoad struct {
	Names []string `json:"names"`
}

// QuizLoad describe the struct of external quizzes file
type QuizLoad struct {
	Number      int      `json:"number"`
	Description string   `json:"description"`
	Score       int      `json:"score"`
	OptionA     string   `json:"option_a"`
	OptionB     string   `json:"option_b"`
	OptionC     string   `json:"option_c"`
	OptionD     string   `json:"option_d"`
	Answer      string   `json:"answer"`
	Tags        []string `json:"tags"`
}

// PlayerLoad describe the struct of external players file
type PlayerLoad struct {
	Name                 string `json:"name"`
	CorrectQuizNumbers   []int  `json:"correct_quiz_numbers"`
	IncorrectQuizNumbers []int  `json:"incorrect_quiz_numbers"`
}

// LoadAll will load all external staticdata
func LoadAll() {

	loadProvokes()
	loadTags()
	loadQuizzes()
	loadPlayers()
}

func loadProvokes() {

	jsonFile, err := os.Open("data/provokes.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	provokes := ProvokeLoad{}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &provokes)

	for _, p := range provokes.Correct {
		err := db.CreateProvoke(true, p)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, p := range provokes.Incorrect {
		err := db.CreateProvoke(false, p)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func loadTags() {

	jsonFile, err := os.Open("data/tags.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	tags := TagLoad{}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &tags)

	for _, t := range tags.Names {
		err := db.CreateTag(t)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func loadQuizzes() {

	jsonFile, err := os.Open("data/quizzes.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	quizzes := []QuizLoad{}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &quizzes)

	for _, q := range quizzes {

		quiz := db.Quiz{
			Number:      q.Number,
			Description: q.Description,
			Score:       q.Score,
			OptionA:     q.OptionA,
			OptionB:     q.OptionB,
			OptionC:     q.OptionC,
			OptionD:     q.OptionD,
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

func loadPlayers() {

	jsonFile, err := os.Open("data/players.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	players := []PlayerLoad{}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &players)

	for _, p := range players {

		err := db.CreatePlayer(p.Name)
		if err != nil {
			log.Fatal(err)
		}

		for _, c := range p.CorrectQuizNumbers {
			err := db.RegisterAnswer(p.Name, c, true)
			if err != nil {
				log.Fatal(err)
			}
		}

		for _, c := range p.IncorrectQuizNumbers {
			err := db.RegisterAnswer(p.Name, c, false)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
