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
	testPlayerName string = "testplayer"
)

var (
	client *http.Client = &http.Client{}
)

// VerifyGetPlayers verify the function of player endpoints.
func VerifyGetPlayers() {

	fmt.Print(fmt.Sprintf(config.Config.FVT.Topic, "VerifyGetPlayers"))
	url := "http://0.0.0.0:8080/v1/players"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "All Players"))
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
}

// VerifyPostPlayers verify the function of player endpoints.
func VerifyPostPlayers() {

	fmt.Print(fmt.Sprintf(config.Config.FVT.Topic, "VerifyPostPlayers"))
	url := "http://0.0.0.0:8080/v1/players"
	jsonStr := []byte(fmt.Sprintf(`{"Name":"%s"}`, testPlayerName))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "New Player"))
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

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "Duplicate Player"))
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

// VerifyDeletePlayer verify the delete function of player endpoints.
func VerifyDeletePlayer() {

	fmt.Print(fmt.Sprintf(config.Config.FVT.Topic, "VerifyPostPlayers"))
	url := fmt.Sprintf("http://0.0.0.0:8080/v1/players/%s", testPlayerName)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "Existed Player"))
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

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "Non-existed Player"))
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
