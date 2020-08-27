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
	testTag string = "TestTag"
)

// VerifyGetTags verify the GET method of the route /tags.
func VerifyGetTags() {

	fmt.Printf(config.Config.FVT.Topic, "VerifyGetTags")
	url := "http://0.0.0.0:8080/v1/tags"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(config.Config.FVT.Section, "All Tags")
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

// VerifyPostTags verify the POST method of the route /tags.
func VerifyPostTags() {

	fmt.Printf(config.Config.FVT.Topic, "VerifyPostTags")
	url := "http://0.0.0.0:8080/v1/tags"
	payload := `{
	"name":"%s"
}`
	jsonStr := fmt.Sprintf(payload, testTag)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(config.Config.FVT.Section, "New Tag")
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

	fmt.Printf(config.Config.FVT.Section, "Duplicate Tag")
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
