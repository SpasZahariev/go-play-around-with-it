package main

import (
	"fmt"
	"time"
)

func main() {

	go repeat("hello")
	go repeat("spas")
	go repeat("how")
	repeat("u doing")

}

func repeat(word string) {
	for {
		fmt.Println(word)
		time.Sleep(1 * time.Second)
	}
}