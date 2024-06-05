package main

import (
	"fmt"
	"time"
)

type Semaphore struct {
	permits chan struct{}
}

func NewSemaPhore(n int) *Semaphore {
	return &Semaphore{permits: make(chan struct{}, n)}
}

func (s *Semaphore) Acquire() {
	s.permits <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.permits
}

func main() {
	sem := NewSemaPhore(3)

	for i := 0; i < 10; i++ {
		go func(id int) {
			sem.Acquire()
			fmt.Printf("Goroutine %d acquired the semaphore\n", id)
			time.Sleep(2 * time.Second)
			fmt.Printf("Goroutine %d releasing the semaphore\n", id)
			sem.Release()
		}(i)
	}

	time.Sleep(10 * time.Second)
}
