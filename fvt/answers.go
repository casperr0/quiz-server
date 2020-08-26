package fvt

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ccns/quiz-server/config"
)

// VerifyGetAnswers verify the GET method of the route /answers.
func VerifyGetAnswers() {

	fmt.Print(fmt.Sprintf(config.Config.FVT.Topic, "VerifyGetAnswers"))

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "All Answers"))
	url := "http://0.0.0.0:8080/v1/answers"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(fmt.Sprintf("$ GET %s\n", url))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(bodyBytes))

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "Query Answers by Player"))
	url = fmt.Sprintf("http://0.0.0.0:8080/v1/answers?player=%s", testPlayerName)
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(fmt.Sprintf("$ GET %s\n", url))
	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(bodyBytes))

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "Query Answers by Quiz"))
	url = fmt.Sprintf("http://0.0.0.0:8080/v1/answers?quiz=%d", testQuizNumber)
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(fmt.Sprintf("$ GET %s\n", url))
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

// VerifyPostAnswers verify the POST method of the route /answers.
func VerifyPostAnswers() {

	fmt.Print(fmt.Sprintf(config.Config.FVT.Topic, "VerifyPostAnswers"))
	url := "http://0.0.0.0:8080/v1/answers"
	payload := `
	{
		"player_name":"%s",
		"quiz_number":%d,
		"correct":true
	}
	`
	jsonStr := []byte(fmt.Sprintf(payload, testPlayerName, testQuizNumber))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "New Answer"))
	fmt.Print(fmt.Sprintf("$ POST %s\n", url))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(bodyBytes))

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "Duplicate Answer"))
	fmt.Print(fmt.Sprintf("$ POST %s\n", url))
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
