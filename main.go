package main

import (
	"fmt"
	"time"
)

func task(control chan bool) {

}

func main() {
	/* limit, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("Input must be an int")
	}*/

	timeout := 1 * time.Second
	ansChannel := make(chan int, 1)

	go func() {
		sum := 0
		timeoutChannel := time.After(timeout)
		for {
			select {
			case <-timeoutChannel:
				ansChannel <- sum
				// return
				fmt.Printf("I counted to: %v\n", sum)
				// break
			default:
				sum++
				//	fmt.Printf("Incremented to: %v\n", sum)
			}
		}
		// fmt.Println("Iteration!")
		fmt.Println("I ran")
	}()

	fmt.Printf("Answer: %v\n", <-ansChannel)
}
