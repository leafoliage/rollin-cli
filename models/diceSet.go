package models

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type DiceSet struct {
	Set   []Dice
	Score int
}

var errInvalidDiceFormat = errors.New("invalid dice format")
var errInvalidDice = errors.New("invalid dice")

func (diceSet *DiceSet) Import(request string) {

	side, amount, err := validateRequest(request)
	if err != nil {
		fmt.Println(errInvalidDiceFormat)
		return
	}

	for i := 0; i < amount; i++ {
		diceSet.Set = append(diceSet.Set, Dice{Side: side})
	}

}

func (diceSet *DiceSet) Roll() {

	diceSet.Score = 0

	for i := range diceSet.Set {
		dice := &diceSet.Set[i]
		dice.Roll()
		diceSet.Score += dice.Score
	}

}

func (diceSet *DiceSet) Print(showScore bool) {

	if showScore {
		fmt.Printf("%d ", diceSet.Score)
	} else {
		fmt.Print("? ")
	}

	fmt.Printf("[")
	for _, dice := range diceSet.Set {
		fmt.Printf("%d ", dice.Score)
	}
	fmt.Printf("\b]\n")

}

func (diceSet *DiceSet) ScoreAnimation() {

	time.Sleep(time.Millisecond * 100)
	saveCursorPos()

	for i := 0; i < 25; i++ {

		rollBackCursorPos()
		diceSet.Roll()
		diceSet.Print(false)
		time.Sleep(time.Millisecond * time.Duration(50+i*10))

	}

	rollBackCursorPos()
	diceSet.Print(true)
	time.Sleep(time.Millisecond * 100)
}

func (diceSet *DiceSet) Empty() bool {
	return len(diceSet.Set) == 0
}

func validateRequest(request string) (side int, amount int, err error) {

	matched, err := regexp.Match("\\d+d\\d+", []byte(request))
	if err != nil || !matched {
		return 0, 0, errInvalidDiceFormat
	}

	parsed := strings.Split(request, "d")

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

func saveCursorPos() {
	fmt.Print("\033[s")
}

func rollBackCursorPos() {
	fmt.Print("\033[u\033[K")
}
