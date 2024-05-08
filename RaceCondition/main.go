package main

import (
	"fmt"
	"sync"
)

func deposit1(balance *int, wg *sync.WaitGroup, amount int) {
	*balance += amount
	wg.Done()
}

func withdraw1(balance *int, amount int) {
	*balance -= amount
}

func example1() {
	balance := 100

	var wg sync.WaitGroup

	wg.Add(1)
	go deposit1(&balance, &wg, 10)
	wg.Wait()

	withdraw1(&balance, 50)

	fmt.Println("balance1", balance)
}

func deposit2(balance int, send chan int, amount int) {
	balance += amount
	send <- balance
}

func withdraw2(balance *int, amount int) {
	*balance -= amount
}

func example2() {
	balance := 100

	send := make(chan int)
	go deposit2(balance, send, 10)
	balance = <-send

	withdraw2(&balance, 50)

	fmt.Println("balance2", balance)
}

func main() {
	example1()
	example2()
}
