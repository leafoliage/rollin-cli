package models

import (
	"math/rand"
)

type Dice struct {
	Side  int
	Score int
}

func (dice *Dice) Roll() {
	dice.Score = rand.Int()%(dice.Side-1) + 1
}
