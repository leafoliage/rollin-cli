package models

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Dice struct {
	Side   int
	Amount int
}

func NewDice(diceStr string) (*Dice, error) {
	if !isValidDiceStr(diceStr) {
		return nil, errors.New("invalid dice format")
	}

	parsed := strings.Split(diceStr, "d")

	amount, err := strconv.Atoi(parsed[0])
	if err != nil {
		return nil, err
	}

	side, err := strconv.Atoi(parsed[1])
	if err != nil {
		return nil, err
	}

	dice := &Dice{Side: side, Amount: amount}
	return dice, nil
}

func isValidDiceStr(diceStr string) bool {
	matched, err := regexp.Match("\\d+d\\d+", []byte(diceStr))
	return err == nil && matched
}
