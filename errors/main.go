package main

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("math: square root of negative number.")
	}
	// Compute the square root
	return math.Pow(num, 0.5), nil
}

func process[T any](data []T) error {
	if len(data) == 0 {
		return errors.New("data cannot be empty to process")
	}
	return nil
}

type myError struct {
	message string
}

func (m *myError) Error() string {
	return fmt.Sprintf("Error: %s", m.message)
}

func errorProcess() error {
	return &myError{message: "Custom error message"}
}

func readData() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("readData: %w", err)
	}
	return nil
}

func readConfig() error {
	return errors.New("Error Config")
}

func main() {
	var num float64 = 25.0
	var numNegative float64 = -5.0

	sqrtNum, err := sqrt(num)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("The square root of %v => %v\n", num, sqrtNum) // The square root of 25 => 5

	// **********************************************************
	sqrtNum, err = sqrt(numNegative)
	if err != nil {
		fmt.Println(err) // math: square root of negative number.
	}
	fmt.Printf("The square root of %v => %v\n", numNegative, sqrtNum)

	// **********************************************************
	data := []int{}
	data2 := []string{"coin A", "coin B", "coin C"}
	if err := process(data); err != nil {
		fmt.Println(err) // data cannot be empty to process
	}
	if err := process(data2); err != nil {
		fmt.Println(err) // nil
	}

	// **********************************************************

	err1 := errorProcess()
	if err1 != nil {
		fmt.Println(err1) // Error: Custom error message
	}

	// **********************************************************

	if err := readData(); err != nil {
		fmt.Println(err) // readData: Error Config
	}
}
