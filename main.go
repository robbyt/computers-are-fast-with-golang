package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"strconv"
	"time"
)

func task(control chan bool) {

}

func main() {
	// disable gc
	debug.SetGCPercent(-1)

	// arg parser
	if len(os.Args) != 2 {
		panic("Must provide a timeout value in seconds")
	}

	timeoutInput, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		panic("Input must be an int or float")
	}

	// after timeout, this will report the answer
	ansChannel := make(chan int, 1)

	i := 0
	timeout := time.Duration(timeoutInput) * time.Second
	timeoutChannel := time.After(timeout)
	// start backround calculation
	go func() {
		for {
			select {
			case <-timeoutChannel:
				ansChannel <- i
				break
			default:
				i++
				// more stuff could be done here...
			}
		}
	}()

	// block, waiting for timeout and answer
	output := <-ansChannel
	fmt.Printf("Total: %v\n", output)

	perSec := float64(output) / timeoutInput
	fmt.Printf("Average per-second: %f\n", perSec)
}
