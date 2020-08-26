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
	testProvokeMessage = "test provoke message"
)

// VerifyGetProvokes verify the GET method of the route /provokes.
func VerifyGetProvokes() {

	fmt.Print(fmt.Sprintf(config.Config.FVT.Topic, "VerifyGetProvokes"))

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "All Provokes"))
	url := "http://0.0.0.0:8080/v1/provokes"
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

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "Query Provokes by Correctness"))
	url = "http://0.0.0.0:8080/v1/provokes?correct=true"
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

// VerifyPostProvokes verify the POST method of the route /provokes.
func VerifyPostProvokes() {

	fmt.Print(fmt.Sprintf(config.Config.FVT.Topic, "VerifyPostProvokes"))
	url := "http://0.0.0.0:8080/v1/provokes"
	payload := `
	{
		"correct":true,
		"message":"%s"
	}
	`
	jsonStr := []byte(fmt.Sprintf(payload, testProvokeMessage))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "New Provoke"))
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

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "Duplicate Provoke"))
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
