package main

import (
	"fmt"
	"sync"
)

type Data struct {
	number int
	sqaure float64
}

func main() {

	nums := []int{
		1, 2, 3, 4, 5,
		6, 7, 8, 9, 10,
		11, 12, 13, 14, 15,
		16, 17, 18, 19, 20,
		21, 22, 23, 24, 25,
		26, 27, 28, 29, 30,
		31, 32, 33, 34, 35,
		36, 37, 38, 39, 40,
		41, 42, 43, 44, 45,
		46, 47, 48, 49, 50,
		51, 52, 53, 54, 55,
		56, 57, 58, 59, 60,
		61, 62, 63, 64, 65,
		66, 67, 68, 69, 70,
		71, 72, 73, 74, 75,
		76, 77, 78, 79, 80,
		81, 82, 83, 84, 85,
		86, 87, 88, 89, 90,
		91, 92, 93, 94, 95,
		96, 97, 98, 99, 100,
		111, 102, 103, 104, 105,
		106, 107, 108, 109, 110,
		111, 112, 113, 114, 115,
		116, 117, 118, 119, 120,
		121, 122, 123, 124, 125,
		126, 127, 128, 129, 130,
		131, 132, 133, 134, 135,
		136, 137, 138, 139, 140,
		141, 142, 143, 144, 145,
		146, 147, 148, 149, 150,
	}

	fmt.Println("ConcurrencyWithoutWorkerPool ::", ConcurrencyWithWorkerPool(nums))
}

func ConcurrencyWithoutWorkerPool(nums []int) []Data {

	output := make([]Data, 0)
	outputCh := make(chan Data)

	var wg sync.WaitGroup

	for i := 0; i < len(nums); i++ {
		v := nums[i]
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			outputCh <- Data{
				number: v,
				sqaure: float64(v) * float64(v),
			}
		}(v)
	}

	go func() {
		wg.Wait()
		close(outputCh)
	}()

	for v := range outputCh {
		output = append(output, v)
	}

	return output
}

func WorkerPool(nums []int) <-chan int {
	workerChan := make(chan int)

	go func() {
		for _, v := range nums {
			workerChan <- v
		}
		close(workerChan)
	}()

	return workerChan
}

func ConcurrencyWithWorkerPool(nums []int) []Data {
	output := make([]Data, 0)
	outputChan := make(chan Data)

	workerChan := WorkerPool(nums)
	workerCount := 5
	var wg sync.WaitGroup

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range workerChan {
				outputChan <- Data{
					number: v,
					sqaure: float64(v) * float64(v),
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(outputChan)
	}()

	for v := range outputChan {
		output = append(output, v)
	}

	return output
}
