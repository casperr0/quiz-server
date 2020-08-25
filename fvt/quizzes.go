package fvt

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ccns/quiz-server/config"
)

// VerifyGetQuizzes verify the function of player endpoints.
func VerifyGetQuizzes() {

	resp, err := http.Get("http://0.0.0.0:8080/v1/quizzes")
	if err != nil {
		log.Fatal(err)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(fmt.Sprintf(config.Config.FVT.Topic, "VerifyGetQuizzes"))
	fmt.Print(string(bodyBytes))
}
