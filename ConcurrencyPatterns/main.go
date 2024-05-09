package main

import "fmt"

func main() {
	wait := make(chan bool)
	count := 0

	go func() {
		count++
		wait <- true
	}()

	<-wait

	fmt.Println("count", count)
}
