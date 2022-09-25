package models

import "math/rand"

type Dice struct {
	Side   int
	Amount int
}

func (dice *Dice) Roll() (res Result) {
	for i := 0; i < dice.Amount; i++ {
		score := rand.Int() % dice.Side
		res.total += score
		res.array = append(res.array, score)
	}
	return res
}
