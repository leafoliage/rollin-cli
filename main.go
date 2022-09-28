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

	for _, request := range os.Args[1:] {
		result := models.Result{}
		diceStrs := parseRequest(request)

		for _, diceStr := range diceStrs {
			dice, err := models.NewDice(diceStr)
			if err != nil {
				fmt.Println(err)
				continue
			}

			newRes := models.Roll(dice)
			result.Append(newRes)
		}

		result.Show()
	}
}

func parseRequest(reqStr string) (diceStrs []string) {
	diceStrs = strings.Split(reqStr, "+")
	return
}
