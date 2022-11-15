package models

import (
	"fmt"
)

type Result struct {
	Total int
	Array []int
}

func (res *Result) Show() {
	if !res.Empty() {
		fmt.Println(res.Total, res.Array)
	}
}

func (res *Result) Merge(newRes *Result) {
	res.Total += newRes.Total
	res.Array = append(res.Array, newRes.Array...)
}

func (res *Result) Empty() bool {
	return len(res.Array) == 0 || res.Total == 0
}
