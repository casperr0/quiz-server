package fvt

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ccns/quiz-server/config"
)

// VerifyGetRank verify the function of rank endpoints.
func VerifyGetRank() {

	fmt.Printf(config.Config.FVT.Topic, "VerifyGetRank")

	fmt.Print(fmt.Sprintf(config.Config.FVT.Section, "Rank of active player"))
	url := "http://0.0.0.0:8080/v1/rank"
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
