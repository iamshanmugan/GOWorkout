package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {

	for j := range jobs {
		fmt.Println("worker", id, "startedjob", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numofjobs = 5
	jobs := make(chan int, numofjobs)
	results := make(chan int, numofjobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	for j := 1; j <= numofjobs; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= numofjobs; a++ {
		<-results
	}
}
