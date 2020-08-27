package fvt

import "net/http"

const (
	testProvokeMessage string = "test provoke message"
	testTagName        string = "Engineering"
	testPlayerName     string = "testplayer"
	testQuizNumber     int    = 999
	testTag            string = "TestTag"
)

var (
	client *http.Client = &http.Client{}
)
