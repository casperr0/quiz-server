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
	testQuizNumber int    = 999
	testTagName    string = "Engineering"
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

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "Query Quizzes by Tag"))
	url = "http://0.0.0.0:8080/v1/quizzes?tag=Security"
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

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "Duplicate Quiz"))
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

// VerifyGetQuiz verify the get method of the route quizzes/:quiz_number.
func VerifyGetQuiz() {

	fmt.Print(fmt.Sprintf(config.Config.FVT.Topic, "VerifyGetQuiz"))
	url := fmt.Sprintf("http://0.0.0.0:8080/v1/quizzes/%d", testQuizNumber)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "Particular Quiz"))
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
}

// VerifyDeleteQuiz verify the delete method of the route quizzes/:quiz_number.
func VerifyDeleteQuiz() {

	fmt.Print(fmt.Sprintf(config.Config.FVT.Topic, "VerifyDeletePlayer"))
	url := fmt.Sprintf("http://0.0.0.0:8080/v1/quizzes/%d", testQuizNumber)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "Existed Quiz"))
	fmt.Print(fmt.Sprintf("$ DELETE %s\n", url))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(bodyBytes))

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "Non-existed Quiz"))
	fmt.Print(fmt.Sprintf("$ DELETE %s\n", url))
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

// VerifyGetQuizTags verify the GET method of the route quizzes/:quiz_number/tags.
func VerifyGetQuizTags() {

	fmt.Print(fmt.Sprintf(config.Config.FVT.Topic, "VerifyGetQuizTags"))
	url := fmt.Sprintf("http://0.0.0.0:8080/v1/quizzes/%d/tags", testQuizNumber)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "All Quiz Tags"))
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
}

// VerifyPostQuizTags verify the POST method of the route quizzes/:quiz_number/tags.
func VerifyPostQuizTags() {

	fmt.Print(fmt.Sprintf(config.Config.FVT.Topic, "VerifyPostQuizTags"))
	url := fmt.Sprintf("http://0.0.0.0:8080/v1/quizzes/%d/tags", testQuizNumber)
	payload := `{"name":"%s"}`
	jsonStr := []byte(fmt.Sprintf(payload, testTagName))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "New Quiz Tag"))
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

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "Deuplicate Quiz Tag"))
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
