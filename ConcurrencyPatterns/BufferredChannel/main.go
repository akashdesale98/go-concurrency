package main

import (
	"fmt"
)

func producer(channel chan int) {
	defer close(channel) // Close the channel when the producer is done
	for i := 0; i < 10; i++ {
		fmt.Println("Producing", i)
		channel <- i
	}
}

func consumer(channel chan int) {
	for v := range channel {
		fmt.Println("Consuming", v*v)
	}
}

func main() {
	channel := make(chan int, 10)

	go producer(channel)

	consumer(channel)

}
