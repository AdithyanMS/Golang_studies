package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.animeout.xyz/")
	if err != nil {
		fmt.Println("Some error")
		os.Exit(1)
	}
	fmt.Println(resp)
}
