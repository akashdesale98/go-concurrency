package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Fan-In, Fan-Out
// Fan-In refers to a technique in which you join data from multiple inputs into a single entity
// Fan-Out means to divide the data from a single source into multiple smaller chunks.

func updatePosition(name string) <-chan string {
	posChannel := make(chan string)

	go func() {
		for i := 0; ; i++ {
			posChannel <- fmt.Sprintf("%s : %d", name, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return posChannel
}

func fanIn(channel1, channel2 <-chan string) <-chan string {
	fanInChannel := make(chan string)

	go func() {
		for {
			fanInChannel <- <-channel1
		}
	}()

	go func() {
		for {
			fanInChannel <- <-channel2
		}
	}()

	return fanInChannel
}

func main() {

	// positionChannel := fanIn(updatePosition("John"), updatePosition("Dev"))

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(<-positionChannel)
	// }
	// fmt.Println("Done with getting updates on positions.")

	var myNumbers [10]int

	for i := 0; i < 10; i++ {
		myNumbers[i] = i
	}

	outChannel := channelGenerator(myNumbers)

	// fan-out
	channel1 := double(outChannel)
	channel2 := double(outChannel)

	inChannel := fanInV2(channel1, channel2)

	for i := 0; i < len(myNumbers); i++ {
		fmt.Println(<-inChannel)
	}
}

func channelGenerator(nums [10]int) <-chan string {
	channel := make(chan string)

	go func() {
		for _, v := range nums {
			channel <- strconv.Itoa(v)
		}
		close(channel)
	}()

	return channel
}

func double(outChannel <-chan string) <-chan string {
	fanOutchannel := make(chan string)
	go func() {
		for v := range outChannel {
			num, _ := strconv.Atoi(v)
			fanOutchannel <- fmt.Sprintf("%d * 2 = %d", num, num*num)
		}
		close(fanOutchannel)
	}()

	return fanOutchannel
}

func fanInV2(chan1, chan2 <-chan string) <-chan string {
	inChannel := make(chan string)

	go func() {
		for {
			select {
			case msg1 := <-chan1:
				inChannel <- msg1
			case msg2 := <-chan2:
				inChannel <- msg2
			}
		}
	}()

	return inChannel
}
