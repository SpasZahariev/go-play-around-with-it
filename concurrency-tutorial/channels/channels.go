package main

import (
	"fmt"
)

func main() {

	var jobs chan int = make(chan int, 100)
	var results chan int = make(chan int, 100)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	for j := range results {
		fmt.Println(j)
	}
	close(results)

}

//will only ever receive stuff on this channel
// will only ever send results out of this channel
func worker(jobs <-chan int, results chan<- int) {

	for number := range jobs {

		results <- fib(number)
	}

}

func fib(number int) int {
	if number <= 1   {
		return number
	}
	return fib(number -1 ) + fib(number - 2)
}