package main

import (
	"fmt"
	"sync"
)

func main() {
	n := 10
	evenChannel, oddChannel := make(chan int), make(chan int)

	var wg sync.WaitGroup

	// Goroutine to send values to channels
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= n; i++ {
			if i%2 == 0 {
				evenChannel <- i
			} else {
				oddChannel <- i
			}
		}
		close(evenChannel)
		close(oddChannel)
	}()

	// Goroutine to receive values from channels and print them
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case i, ok := <-evenChannel:
				if ok {
					fmt.Println("i", i)
				} else {
					evenChannel = nil
				}
			case j, ok := <-oddChannel:
				if ok {
					fmt.Println("j", j)
				} else {
					oddChannel = nil
				}
			}

			// Exit the loop when both channels are nil
			if evenChannel == nil && oddChannel == nil {
				break
			}
		}
	}()

	wg.Wait()
}
