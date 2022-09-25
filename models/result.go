package models

import "fmt"

type Result struct {
	total int
	array []int
}

func (res *Result) Show() {
	fmt.Println(res.total, res.array)
}
