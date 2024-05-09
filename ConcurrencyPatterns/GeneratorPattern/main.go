package main

import "fmt"

// Generators return the next value in a sequence each time they are called.
// This means that each value is available as an output before the generator computes the next value.
// Hence, this pattern is used to introduce parallelism in our program.

func foo() <-chan string {
	channel := make(chan string)

	go func() {
		for i := 0; ; i++ {
			channel <- fmt.Sprintf("Counter at : %d", i)
		}
	}()

	return channel
}

func updatePosition(name string) <-chan string {
	positionChannel := make(chan string)

	go func() {
		for i := 0; ; i++ {
			positionChannel <- fmt.Sprintf("%s : %d", name, i)
		}
	}()

	return positionChannel
}

func fibonacci(n int) chan int {
	mychannel := make(chan int)
	go func() {
		k := 0
		for i, j := 0, 1; k < n; k++ {
			mychannel <- i
			i, j = i+j, i

		}
		close(mychannel)
	}()
	return mychannel
}

func main() {
	// channel := foo()

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(<-channel)
	// }

	// fmt.Println("Done with counter")

	// posChannel1, posChannel2 := updatePosition("John"), updatePosition("Jane")

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(<-posChannel1)
	// 	fmt.Println(<-posChannel2)
	// 	// select {
	// 	// case johnWins := <-posChannel1:
	// 	// 	fmt.Println("Winner is", johnWins)
	// 	// case janeWins := <-posChannel2:
	// 	// 	fmt.Println("Winner is", janeWins)
	// 	// default:
	// 	// 	fmt.Println("No one wins")
	// 	// }
	// }

	// fmt.Println("Done with getting updates on positions.")

	for i := range fibonacci(10) {
		//do anything with the nth term while the fibonacci()
		//is computing the next term
		fmt.Println(i)
	}
}
