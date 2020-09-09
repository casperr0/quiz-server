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

	fmt.Printf(config.Config.FVT.Topic, "VerifyGetAnswers")

	fmt.Printf(config.Config.FVT.Section, "All Answers")
	url := "http://0.0.0.0:8080/v1/answers"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(config.Config.FVT.Detail, "method and url")
	fmt.Printf("```\n$ GET %s\n```\n", url)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(config.Config.FVT.Detail, "example response")
	fmt.Printf("```\n%s\n```\n", string(bodyBytes))

	fmt.Printf(config.Config.FVT.Section, "Query Answers by Player")
	url = fmt.Sprintf("http://0.0.0.0:8080/v1/answers?player=%s", testPlayerName)
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(config.Config.FVT.Detail, "method and url")
	fmt.Printf("```\n$ GET %s\n```\n", url)
	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(config.Config.FVT.Detail, "example response")
	fmt.Printf("```\n%s\n```\n", string(bodyBytes))

	fmt.Printf(config.Config.FVT.Section, "Query Answers by Quiz")
	url = fmt.Sprintf("http://0.0.0.0:8080/v1/answers?quiz=%d", testQuizNumber)
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(config.Config.FVT.Detail, "method and url")
	fmt.Printf("```\n$ GET %s\n```\n", url)
	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(config.Config.FVT.Detail, "example response")
	fmt.Printf("```\n%s\n```\n", string(bodyBytes))
}

// VerifyPostAnswers verify the POST method of the route /answers.
func VerifyPostAnswers() {

	fmt.Printf(config.Config.FVT.Topic, "VerifyPostAnswers")
	url := "http://0.0.0.0:8080/v1/answers"
	payload := `{
	"player_name":"%s",
	"quiz_number":%d,
	"correct":true
}`
	jsonStr := fmt.Sprintf(payload, testPlayerName, testQuizNumber)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(config.Config.FVT.Section, "New Answer")
	fmt.Printf(config.Config.FVT.Detail, "method and url")
	fmt.Printf("```\n$ POST %s\n```\n", url)
	fmt.Printf(config.Config.FVT.Detail, "example payload")
	fmt.Printf("```\n%s\n```\n", jsonStr)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(config.Config.FVT.Detail, "example response")
	fmt.Printf("```\n%s\n```\n", string(bodyBytes))

	fmt.Printf(config.Config.FVT.Section, "Duplicate Answer")
	fmt.Printf(config.Config.FVT.Detail, "method and url")
	fmt.Printf("```\n$ POST %s\n```\n", url)
	fmt.Printf(config.Config.FVT.Detail, "example payload")
	fmt.Printf("```\n%s\n```\n", jsonStr)
	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(config.Config.FVT.Detail, "example response")
	fmt.Printf("```\n%s\n```\n", string(bodyBytes))
}
