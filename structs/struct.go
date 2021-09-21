package main

import "fmt"

type bd struct {
	date  int
	month string
	year  int
}

type person struct {
	name     string
	dialogue string
	birthday bd
}

func (p person) show() {
	fmt.Printf("%+v", p)
}

func (p *person) chng(ns string) {
	(*p).birthday.month = "June"
}

func main() {
	niharu := person{
		name:     "Niharuuu",
		dialogue: "adithyaa",
		birthday: bd{
			date:  19,
			month: "September",
			year:  2000,
		},
	}
	niharu.show()
	niharu.chng("december")
	niharu.show()
}
