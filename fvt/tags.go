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

	fmt.Print(fmt.Sprintf(config.Config.FVT.Topic, "VerifyGetTags"))
	url := "http://0.0.0.0:8080/v1/tags"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "All Tags"))
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

// VerifyPostTags verify the POST method of the route /tags.
func VerifyPostTags() {

	fmt.Print(fmt.Sprintf(config.Config.FVT.Topic, "VerifyPostTags"))
	url := "http://0.0.0.0:8080/v1/tags"
	payload := `
	{
		"name":"%s"
	}
	`
	jsonStr := []byte(fmt.Sprintf(payload, testTag))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "New Tag"))
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

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "Duplicate Tag"))
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
