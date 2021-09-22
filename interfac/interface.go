package main

import "fmt"

type game interface {
	getReview() string
}
type riot struct{}
type mihoyo struct{}

func main() {
	valorant := riot{}
	genshin := mihoyo{}
	printReview(valorant)
	printReview(genshin)

}

func (r riot) getReview() string {
	return "Nice aan"
}

func (g mihoyo) getReview() string {
	return "single playerinu kollam"
}

func printReview(g game) {
	fmt.Println(g.getReview())
}
