package fvt

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ccns/quiz-server/config"
)

const (
	testQuizNumber int = 999
)

// VerifyGetQuizzes verify the function of player endpoints.
func VerifyGetQuizzes() {

	fmt.Print(fmt.Sprintf(config.Config.FVT.Topic, "VerifyGetQuizzes"))

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "All Quizzes"))
	url := "http://0.0.0.0:8080/v1/quizzes"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(fmt.Sprintf("$ %s\n", url))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(bodyBytes))

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "Query Quizzes"))
	url = "http://0.0.0.0:8080/v1/quizzes?tag=Security"
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(fmt.Sprintf("$ %s\n", url))
	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(bodyBytes))
}

// VerifyPostQuizzes verify the function of quizzes endpoints.
func VerifyPostQuizzes() {

	fmt.Print(fmt.Sprintf(config.Config.FVT.Topic, "VerifyPostQuizzes"))
	url := "http://0.0.0.0:8080/v1/quizzes"
	payload := `
	{
		"number":%d,
		"Description":"test description.",
		"Score":3,
		"OptionA":"test A",
		"OptionB":"test B",
		"OptionC":"test C",
		"OptionD":"test D",
		"Answer":"A"
	}
	`
	jsonStr := []byte(fmt.Sprintf(payload, testQuizNumber))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "New Quiz"))
	fmt.Print(fmt.Sprintf("$ %s\n", url))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(bodyBytes))

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "Duplicate Quiz"))
	fmt.Print(fmt.Sprintf("$ %s\n", url))
	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(bodyBytes))
}
