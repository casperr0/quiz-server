package main

import (
	"fmt"

	_ "github.com/ccns/quiz-server/config"
	"github.com/ccns/quiz-server/db"
)

func main() {

	db.CreateOfficer("rain", "0114")
	db.CreateOfficer("rain2", "0114")
	db.CreateOfficer("rain3", "0114")
	fmt.Println(db.ListOfficers())
}
