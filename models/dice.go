package models

import (
	"errors"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

type Dice struct {
	Side   int
	Amount int
}

var errInvalidDiceFormat = errors.New("invalid dice format")
var errInvalidDice = errors.New("invalid dice")

func NewDice(diceStr string) (*Dice, error) {
	side, amount, err := validateDiceStr(diceStr)
	if err != nil {
		return nil, err
	}

	dice := &Dice{Side: side, Amount: amount}
	return dice, nil
}

func (dice *Dice) Roll(res *Result) {
	for i := 0; i < dice.Amount; i++ {
		score := rand.Int()%(dice.Side-1) + 1
		res.Total += score
		res.Array = append(res.Array, score)
	}
}

func validateDiceStr(diceStr string) (side int, amount int, err error) {
	matched, err := regexp.Match("\\d+d\\d+", []byte(diceStr))
	if err != nil || !matched {
		return 0, 0, errInvalidDiceFormat
	}

	parsed := strings.Split(diceStr, "d")

	amount, err = strconv.Atoi(parsed[0])
	if err != nil {
		return 0, 0, errInvalidDiceFormat
	}

	side, err = strconv.Atoi(parsed[1])
	if err != nil {
		return 0, 0, errInvalidDiceFormat
	}

	if amount <= 0 || side <= 0 {
		return 0, 0, errInvalidDice
	}

	return side, amount, nil
}
