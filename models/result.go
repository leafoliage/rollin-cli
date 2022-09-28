package models

import (
	"fmt"
	"math/rand"
)

type Result struct {
	Total int
	Array []int
}

func Roll(dice *Dice) *Result {
	res := &Result{}
	for i := 0; i < dice.Amount; i++ {
		score := rand.Int()%(dice.Side-1) + 1
		res.Total += score
		res.Array = append(res.Array, score)
	}
	return res
}

func (res *Result) Show() {
	if !res.Empty() {
		fmt.Println(res.Total, res.Array)
	}
}

func (res *Result) Append(newRes *Result) {
	res.Total += newRes.Total
	res.Array = append(res.Array, newRes.Array...)
}

func (res *Result) Empty() bool {
	return len(res.Array) == 0 || res.Total == 0
}
