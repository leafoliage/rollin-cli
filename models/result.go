package models

import "fmt"

type Result struct {
	Total int
	Array []int
}

func (res *Result) Show() {
	fmt.Println(res.Total, res.Array)
}

func (res *Result) Append(newRes *Result) {
	res.Total += newRes.Total
	res.Array = append(res.Array, newRes.Array...)
}
