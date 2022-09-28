package models

import "math/rand"

type Dice struct {
	Side   int
	Amount int
}

func (dice *Dice) Roll() (res Result) {
	for i := 0; i < dice.Amount; i++ {
		score := rand.Int() % dice.Side
		res.Total += score
		res.Array = append(res.Array, score)
	}
	return res
}
