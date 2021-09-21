package main

import "fmt"

func show(m map[string]string) {
	for key, value := range m {
		fmt.Println(key, value)
	}
	fmt.Println("*****************************************")

}

func main() {
	best := map[string]string{}
	best["Niharuu"] = "The beeeeeeeeeest"
	best["others"] = "best"

	hobbies := map[string]string{
		"adi":    "games",
		"niharu": "Civil Service",
	}

	show(best)
	show(hobbies)

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

}
