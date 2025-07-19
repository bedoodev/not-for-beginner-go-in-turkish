package main

import (
	"fmt"
	"time"
)

func main() {

	var err error
	fmt.Println("Before sayHello function.")
	go sayHello()
	fmt.Println("After sayHello function.")

	go func() {
		err = doWork()

	}()
	// err = go doWork() // This is not accepted

	go printNumbers()
	go printLetters()

	time.Sleep(2 * time.Second)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Work Completed Successfully.")
	}

	/*
		Output:
			Before sayHello function.
			After sayHello function.
			Letter: a 2025-07-19 15:27:28.549057 +0300 +03 m=+0.000225293
			Number: 1 2025-07-19 15:27:28.549055 +0300 +03 m=+0.000223835
			Letter: b 2025-07-19 15:27:28.750391 +0300 +03 m=+0.201570668
			Number: 2 2025-07-19 15:27:28.750365 +0300 +03 m=+0.201543960
			Letter: c 2025-07-19 15:27:28.951547 +0300 +03 m=+0.402737251
			Number: 3 2025-07-19 15:27:28.951522 +0300 +03 m=+0.402712543
			Number: 4 2025-07-19 15:27:29.152684 +0300 +03 m=+0.603885085
			Letter: d 2025-07-19 15:27:29.152709 +0300 +03 m=+0.603910251
			Letter: e 2025-07-19 15:27:29.353836 +0300 +03 m=+0.805047585
			Number: 5 2025-07-19 15:27:29.353866 +0300 +03 m=+0.805077835
			Hello from goroutines.
			Error: This is a dummy error from doWork function.
	*/
}

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from goroutines.")
}

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i, time.Now())
		time.Sleep(200 * time.Millisecond)

	}
}

func printLetters() {
	for _, letter := range "abcde" {
		fmt.Println("Letter:", string(letter), time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error {
	// Simulate a work
	time.Sleep(1 * time.Second)
	return fmt.Errorf("This is a dummy error from doWork function.")
}
