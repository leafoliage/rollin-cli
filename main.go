package main

import (
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/leafoliage/rollin-cli/models"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for _, request := range os.Args[1:] {
		diceStrs := parseRequest(request)
		result := models.Result{}
		for _, diceStr := range diceStrs {
			if !isValidDiceStr(diceStr) {
				continue
			}
			dice, err := parseDiceStr(diceStr)
			if err != nil {
				continue
			}
			newRes := dice.Roll()
			result.Append(&newRes)
		}
		result.Show()
	}
}

func parseRequest(reqStr string) (diceStrs []string) {
	diceStrs = strings.Split(reqStr, "+")
	return
}

func parseDiceStr(diceStr string) (dice models.Dice, err error) {
	parsed := strings.Split(diceStr, "d")
	dice.Amount, err = strconv.Atoi(parsed[0])
	if err != nil {
		return
	}
	dice.Side, err = strconv.Atoi(parsed[1])
	return
}

func isValidDiceStr(diceStr string) bool {
	matched, err := regexp.Match("\\d+d\\d+", []byte(diceStr))
	return err == nil && matched
}
