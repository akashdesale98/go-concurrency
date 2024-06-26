package main

import "testing"

func BenchmarkProcessData(b *testing.B) {
	numbers := []int{
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

	for i := 0; i < b.N; i++ {
		// ConcurrencyWithoutWorkerPool(numbers)
		// Running tool: C:\Program Files\Go\bin\go.exe test -benchmem -run=^$ -bench ^BenchmarkProcessData$ WorkerPool

		// goos: windows
		// goarch: amd64
		// pkg: WorkerPool
		// cpu: AMD Ryzen 7 6800HS with Radeon Graphics
		// BenchmarkProcessData-16    	   26392	     45388 ns/op	   15597 B/op	     312 allocs/op
		// PASS
		// ok  	WorkerPool	1.687s

		ConcurrencyWithWorkerPool(numbers)
		// Running tool: C:\Program Files\Go\bin\go.exe test -benchmem -run=^$ -bench ^BenchmarkProcessData$ WorkerPool

		// goos: windows
		// goarch: amd64
		// pkg: WorkerPool
		// cpu: AMD Ryzen 7 6800HS with Radeon Graphics
		// BenchmarkProcessData-16    	   20168	     58304 ns/op	    8624 B/op	      19 allocs/op
		// PASS
		// ok  	WorkerPool	1.805s
	}
}
