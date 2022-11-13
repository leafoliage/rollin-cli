package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/leafoliage/rollin-cli/models"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	requests := os.Args[1:]

	for _, request := range requests {
		var totalResult models.Result
		diceStrs := parseRequest(request)

		for _, diceStr := range diceStrs {
			dice, err := models.NewDice(diceStr)
			if err != nil {
				fmt.Println(err)
				break
			}

			newResult := dice.Roll()
			totalResult.Append(newResult)
		}

		totalResult.Show()
	}
}

func parseRequest(reqStr string) (diceStrs []string) {
	diceStrs = strings.Split(reqStr, "+")
	return
}
