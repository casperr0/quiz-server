package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/ccns/quiz-server/config"
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

var (
	client *http.Client = &http.Client{}
)

// LoadAll will load all external staticdata
func LoadAll(prod bool) {

	dir := config.Config.Load.DevDir
	if prod {
		dir = config.Config.Load.ProdDir
	}
	loadProvokes(dir)
	loadTags(dir)
	loadQuizzes(dir)
	loadPlayers(dir)
}

func fireRecord(url, jsonStr string) {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(bodyBytes) + "\n")
}

func loadProvokes(dataDir string) {

	jsonFile, err := os.Open(dataDir + "provokes.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	provokes := ProvokeLoad{}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &provokes)

	url := "http://0.0.0.0:8080/v1/provokes"
	payload := `{"correct":%t,"message":"%s"}`
	for _, p := range provokes.Correct {
		jsonStr := fmt.Sprintf(payload, true, p)
		fireRecord(url, jsonStr)
	}
	for _, p := range provokes.Incorrect {
		jsonStr := fmt.Sprintf(payload, false, p)
		fireRecord(url, jsonStr)
	}
}

func loadTags(dataDir string) {

	jsonFile, err := os.Open(dataDir + "tags.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	tags := TagLoad{}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &tags)

	url := "http://0.0.0.0:8080/v1/tags"
	payload := `{"name":"%s"}`
	for _, t := range tags.Names {
		jsonStr := fmt.Sprintf(payload, t)
		fireRecord(url, jsonStr)
	}
}

func loadQuizzes(dataDir string) {

	jsonFile, err := os.Open(dataDir + "quizzes.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	quizzes := []QuizLoad{}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &quizzes)

	url := "http://0.0.0.0:8080/v1/quizzes"
	payload := `{
	"number":%d,
	"description":"%s",
	"score":%d,
	"optionA":"%s",
	"optionB":"%s",
	"optionC":"%s",
	"optionD":"%s",
	"answer":"%s"
}`
	for _, q := range quizzes {
		jsonStr := fmt.Sprintf(
			payload,
			q.Number,
			q.Description,
			q.Score,
			q.OptionA,
			q.OptionB,
			q.OptionC,
			q.OptionD,
			q.Answer,
		)
		fireRecord(url, jsonStr)

		tagURL := fmt.Sprintf("http://0.0.0.0:8080/v1/quizzes/%d/tags", q.Number)
		tagPayload := `{"name":"%s"}`
		for _, t := range q.Tags {
			jsonStr := fmt.Sprintf(tagPayload, t)
			fireRecord(tagURL, jsonStr)
		}
	}
}

func loadPlayers(dataDir string) {

	jsonFile, err := os.Open(dataDir + "players.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	players := []PlayerLoad{}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &players)

	url := "http://0.0.0.0:8080/v1/players"
	payload := `{"name":"%s"}`
	for _, p := range players {

		jsonStr := fmt.Sprintf(payload, p.Name)
		fireRecord(url, jsonStr)

		answerURL := "http://0.0.0.0:8080/v1/answers"
		answerPayload := `{"player_name":"%s","quiz_number":%d,"correct":%t}`
		for _, c := range p.CorrectQuizNumbers {
			jsonStr := fmt.Sprintf(answerPayload, p.Name, c, true)
			fireRecord(answerURL, jsonStr)
		}
		for _, c := range p.IncorrectQuizNumbers {
			jsonStr := fmt.Sprintf(answerPayload, p.Name, c, false)
			fireRecord(answerURL, jsonStr)
		}
	}
}
