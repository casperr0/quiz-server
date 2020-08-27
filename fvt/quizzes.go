package fvt

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ccns/quiz-server/config"
)

// VerifyGetQuizzes verify the function of player endpoints.
func VerifyGetQuizzes() {

	fmt.Printf(config.Config.FVT.Topic, "VerifyGetQuizzes")

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "All Quizzes"))
	url := "http://0.0.0.0:8080/v1/quizzes"
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

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "Query Quizzes by Tag"))
	url = "http://0.0.0.0:8080/v1/quizzes?tag=Security"
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

// VerifyPostQuizzes verify the function of quizzes endpoints.
func VerifyPostQuizzes() {

	fmt.Printf(config.Config.FVT.Topic, "VerifyPostQuizzes")
	url := "http://0.0.0.0:8080/v1/quizzes"
	payload := `{
	"number":%d,
	"Description":"test description.",
	"Score":3,
	"OptionA":"test A",
	"OptionB":"test B",
	"OptionC":"test C",
	"OptionD":"test D",
	"Answer":"A"
}`
	jsonStr := fmt.Sprintf(payload, testQuizNumber)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(config.Config.FVT.Section, "New Quiz")
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

	fmt.Printf(config.Config.FVT.Section, "Duplicate Quiz")
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

// VerifyGetQuiz verify the get method of the route quizzes/:quiz_number.
func VerifyGetQuiz() {

	fmt.Printf(config.Config.FVT.Topic, "VerifyGetQuiz")
	url := fmt.Sprintf("http://0.0.0.0:8080/v1/quizzes/%d", testQuizNumber)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(config.Config.FVT.Section, "Particular Quiz")
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
}

// VerifyDeleteQuiz verify the delete method of the route quizzes/:quiz_number.
func VerifyDeleteQuiz() {

	fmt.Printf(config.Config.FVT.Topic, "VerifyDeletePlayer")
	url := fmt.Sprintf("http://0.0.0.0:8080/v1/quizzes/%d", testQuizNumber)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(config.Config.FVT.Section, "Existed Quiz")
	fmt.Printf(config.Config.FVT.Detail, "method and url")
	fmt.Printf("```\n$ DELETE %s\n```\n", url)
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

	fmt.Printf(config.Config.FVT.Section, "Non-existed Quiz")
	fmt.Printf(config.Config.FVT.Detail, "method and url")
	fmt.Printf("```\n$ DELETE %s\n```\n", url)
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

// VerifyGetQuizTags verify the GET method of the route quizzes/:quiz_number/tags.
func VerifyGetQuizTags() {

	fmt.Printf(config.Config.FVT.Topic, "VerifyGetQuizTags")
	url := fmt.Sprintf("http://0.0.0.0:8080/v1/quizzes/%d/tags", testQuizNumber)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(config.Config.FVT.Section, "All Quiz Tags")
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
}

// VerifyPostQuizTags verify the POST method of the route quizzes/:quiz_number/tags.
func VerifyPostQuizTags() {

	fmt.Printf(config.Config.FVT.Topic, "VerifyPostQuizTags")
	url := fmt.Sprintf("http://0.0.0.0:8080/v1/quizzes/%d/tags", testQuizNumber)
	payload := `{
	"name":"%s"
}`
	jsonStr := fmt.Sprintf(payload, testTagName)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(config.Config.FVT.Section, "New Quiz Tag")
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

	fmt.Printf(config.Config.FVT.Section, "Duplicate Quiz Tag")
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
