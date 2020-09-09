package fvt

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ccns/quiz-server/config"
)

// VerifyGetPlayers verify the function of player endpoints.
func VerifyGetPlayers() {

	fmt.Printf(config.Config.FVT.Topic, "VerifyGetPlayers")
	url := "http://0.0.0.0:8080/v1/players"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(config.Config.FVT.Section, "All Players")
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

// VerifyPostPlayers verify the function of player endpoints.
func VerifyPostPlayers() {

	fmt.Printf(config.Config.FVT.Topic, "VerifyPostPlayers")
	url := "http://0.0.0.0:8080/v1/players"
	paylaod := `{
	"name":"%s"
}`
	jsonStr := fmt.Sprintf(paylaod, testPlayerName)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(config.Config.FVT.Section, "New Player")
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

	fmt.Printf(config.Config.FVT.Section, "Duplicate Player")
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

// VerifyDeletePlayer verify the DELETE method on the route /player/:player_name.
func VerifyDeletePlayer() {

	fmt.Printf(config.Config.FVT.Topic, "VerifyDeletePlayers")
	url := fmt.Sprintf("http://0.0.0.0:8080/v1/players/%s", testPlayerName)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(config.Config.FVT.Section, "Existed Player")
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

	fmt.Printf(config.Config.FVT.Section, "Non-existed Player")
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

// VerifyGetPlayerFeed verify the GET method on the route /player/:player_name/feed.
func VerifyGetPlayerFeed() {

	fmt.Printf(config.Config.FVT.Topic, "VerifyGetPlayerFeed")

	fmt.Printf(config.Config.FVT.Section, "Finished Player")
	url := fmt.Sprintf("http://0.0.0.0:8080/v1/players/%s/feed", "RainrainWu")
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

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "Ongoing Player"))
	url = fmt.Sprintf("http://0.0.0.0:8080/v1/players/%s/feed", testPlayerName)
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

// VerifyGetPlayerRand verify the GET method on the route /player/:player_name/rand.
func VerifyGetPlayerRand() {

	fmt.Printf(config.Config.FVT.Topic, "VerifyGetPlayerRand")

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "Any Player"))
	url := fmt.Sprintf("http://0.0.0.0:8080/v1/players/%s/rand", "RainrainWu")
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
}
