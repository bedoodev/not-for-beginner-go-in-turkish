package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	user := os.Getenv("USER")
	home := os.Getenv("HOME")
	fmt.Println("User:", user) // User: bedirhansahin
	fmt.Println("Home:", home) // Home: /Users/bedirhansahin

	err := os.Setenv("FRUIT", "Apple")
	if err != nil {
		fmt.Println("Error setting environment variable:", err)
		return
	}

	fmt.Println("Name of Fruit:", os.Getenv("FRUIT")) // Name of Fruit: Apple

	// for _, e := range os.Environ() {
	// 	kvpair := strings.SplitN(e, "=", 2)
	// 	fmt.Println(kvpair[0])
	// } // To see all environment Variable

	err = os.Unsetenv("FRUIT")
	if err != nil {
		fmt.Println("Error unsetting environment variable:", err)
		return
	}

	fmt.Println("Name of Fruit:", os.Getenv("FRUIT")) // Name of Fruit:

	fmt.Println("-------------------------------------------------")

	str := "a=b=c=d=e"
	fmt.Println(strings.SplitN(str, "=", -1)) // [a b c d e]
	fmt.Println(strings.SplitN(str, "=", 0))  // []
	fmt.Println(strings.SplitN(str, "=", 1))  // [a=b=c=d=e]
	fmt.Println(strings.SplitN(str, "=", 2))  // [a b=c=d=e]
	fmt.Println(strings.SplitN(str, "=", 3))  // [a b c=d=e]
	fmt.Println(strings.SplitN(str, "=", 4))  // [a b c d=e]
	fmt.Println(strings.SplitN(str, "=", 5))  // [a b c d e]

}
