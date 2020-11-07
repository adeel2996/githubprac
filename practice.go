package main

import (
	"fmt"
)

func calc(x, y int) (out1, out2 int) {
	out1 = x + y
	out2 = x - y
	return
}

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done

	//this is just to show that this is in feature branch
	fmt.Println("in feature branch")
}
