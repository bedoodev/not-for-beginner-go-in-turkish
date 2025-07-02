package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(strings.NewReader("Hello, bufio package!\nHow are you doing?"))

	// Reading byte slice
	data := make([]byte, 20) // create 20 bytes
	n, err := reader.Read(data)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	// Read 20 bytes: Hello, bufio package%
	// Read 22 bytes: Hello, bufio package!
	// Read 16 bytes: Hello, bufio pac%
	fmt.Printf("Read %d bytes: %s\n", n, data[:n])

	line, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading string:", err)
		return
	}
	fmt.Println("Read string:", line) // Read string: !

	writer := bufio.NewWriter(os.Stdout)

	// Writing byte slice
	data2 := []byte("Hello, bufio package!\n")
	nn2, err := writer.Write(data2)

	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", nn2) // Wrote 22 bytes

	data3 := []byte("How are you doing?\n")
	nn3, err := writer.Write(data3)

	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", nn3) // Wrote 18 bytes

	// 	Flush the buffer to ensure all data is written to os.Stdout
	//	Hello, bufio package!
	//	How are you doing?
	if err := writer.Flush(); err != nil {
		fmt.Println("Flushing error:", err)
		return
	}

	// Writing string
	str := "This is a string.\n"
	intN, err := writer.WriteString(str)

	if err != nil {
		fmt.Println("String writer error", err)
		return
	}

	fmt.Printf("Wrote %d bytes\n", intN)

	if err := writer.Flush(); err != nil { // This is a string.
		fmt.Println("Flushing error:", err)
		return
	}

	// ********* ********* ***********

	FlushMail()
	// Output
	/*
		******************************************************************************
		Hello,

			My name is Bedirhan, I am a software engineer who interested mostly backend and data things.

		Kind Regards,
		Bedirhan Sahin
	*/

}
