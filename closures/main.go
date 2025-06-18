package main

import "fmt"

func main() {
	// sequence := counter()

	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())

	substracter := func() func(int) int {

		countdown := 99
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()

	fmt.Println(substracter(1)) // 98
	fmt.Println(substracter(2)) // 96
	fmt.Println(substracter(3)) // 93
	fmt.Println(substracter(4)) // 89
	fmt.Println(substracter(5)) // 84

}

// func counter() func() int {
// 	i := 0
// 	fmt.Println("previous value of i:", i)

// 	return func() int {
// 		i++
// 		fmt.Println("added 1 to i")

// 		return i
// 	}
// }
