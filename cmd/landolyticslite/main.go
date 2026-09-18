package main

import (
	"fmt"

	landolyticslite "github.com/landopritcho/landolytics-lite"
)

func main() {
	scoreboard, err := landolyticslite.GetNFLScoreboard("20260917")
	if err != nil {
		fmt.Print("bruh")
	}
	fmt.Println(scoreboard)
}
