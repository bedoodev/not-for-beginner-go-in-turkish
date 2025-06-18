package main

import "fmt"

func main() {
	// fmt.Println(factorial(5))
	// fmt.Println(factorial(10))

	fmt.Println(sumOfDigits(123))
	fmt.Println(sumOfDigits(50123))

	fmt.Println(123 % 10)

}

func sumOfDigits(n int) int {

	// Base case: eğer n tek haneli bir sayı ise cevap n olur
	if n < 10 {
		return n
	}
	// Recursive case: son basamağı al ve kalan sayı için tekrar çağır
	return n%10 + sumOfDigits(n/10)
}

// func factorial(n int) int {
// 	// Base case: 0 değerinin faktöriyeli 1'dir
// 	if n == 0 {
// 		return 1
// 	}

// 	// Recursive case: n değerinin faktöriyeli n * factorial(n - 1)'dir
// 	return n * factorial(n-1)
// 	// n * (n - 1) * (n - 2) * ...... * 1
// }
