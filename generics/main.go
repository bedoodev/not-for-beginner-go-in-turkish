package main

import "fmt"

// func swap[T any](a, b T) (T, T) {
// 	return b, a
// }

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]

	return element, true
}

func (s Stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}

func (s Stack[T]) printAll() {
	if len(s.elements) == 0 {
		fmt.Println("The Stack is Empty.")
	} else {
		fmt.Println("Stack Elements:", s.elements)
	}
}

func main() {
	// x1, y1 := 1, 2
	// fmt.Println(swap(x1, y1)) // 2, 1

	// x2, y2 := "Ahmet", "Mehmet"
	// fmt.Println(swap(x2, y2)) // "Mehmet", "Ahmet"

	intStack := Stack[int]{}

	fmt.Println("Is stack empty:", intStack.isEmpty()) // Is stack empty: true

	intStack.push(1)
	intStack.push(2)
	intStack.push(3)

	intStack.printAll() // Stack Elements: [1 2 3]

	fmt.Println(intStack.pop()) // 3 true
	intStack.printAll()         // Stack Elements: [1 2]

	fmt.Println(intStack.pop())                        // 2 true
	fmt.Println("Is stack empty:", intStack.isEmpty()) // Is stack empty: false
	fmt.Println(intStack.pop())                        // 1 true
	fmt.Println(intStack.pop())                        // 0 false
	intStack.printAll()                                // The Stack is Empty.
	fmt.Println("Is stack empty:", intStack.isEmpty()) // Is stack empty: true

	fmt.Print("\n***********************************\n\n")

	stringStack := Stack[string]{}

	fmt.Println("Is stack empty:", intStack.isEmpty()) // Is stack empty: true

	stringStack.push("Ahmet")
	stringStack.push("Mehmet")
	stringStack.push("Veli")

	stringStack.printAll() // Stack Elements: [Ahmet Mehmet Veli]

	fmt.Println(stringStack.pop()) // Veli true
	stringStack.printAll()         // Stack Elements: [Ahmet Mehmet]

	fmt.Println(stringStack.pop())                        // Mehmet true
	fmt.Println("Is stack empty:", stringStack.isEmpty()) // Is stack empty: false
	fmt.Println(stringStack.pop())                        // Ahmet true
	fmt.Println(stringStack.pop())                        //  false
	stringStack.printAll()                                // The Stack is Empty.
	fmt.Println("Is stack empty:", stringStack.isEmpty()) // Is stack empty: true

}
