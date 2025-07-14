package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Writing Files
	file, err := os.Create("readAndWrite/files/README.md")
	if err != nil {
		fmt.Println("Error Creating the file:", err)
		return
	}

	defer file.Close()

	data := []byte("## Hello\n")
	if _, err := file.Write(data); err != nil {
		fmt.Println("Error Writing to the file:", err)
		return
	}

	body := []byte(`
My name is Bedirhan.
I am a software engineer at adin.ai in Turkey. 
I am professional Python developer 
and professional go enthusiast.😁`)
	file.Write(body)

	strBody := "\n\nI don't look for any new job."
	_, err = file.WriteString(strBody)
	if err != nil {
		fmt.Println("Error Writing to the file:", err)
		return
	}

	// Reading Files
	readedFile, err := os.ReadFile("readAndWrite/files/README.md")
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}
	fmt.Printf("%s", readedFile)
	/*
		## Hello

		My name is Bedirhan
		I am a software engineer at adin.ai in Turkey.
		I am professional Python developer
		and professional go enthusiast.😁

		I don't look for any new job.%
	*/

	// Open and Read File line by line
	// pfd         poll.FD
	// name        string
	// dirinfo     atomic.Pointer[dirInfo] // nil unless directory being read
	// nonblock    bool                    // whether we set nonblocking mode
	// stdoutOrErr bool                    // whether this is stdout or stderr
	// appendMode  bool
	openedFile, err := os.Open("readAndWrite/files/README.md")
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}
	defer openedFile.Close()
	fmt.Println("File Name:", openedFile.Name()) // File Name: os/files/README.md

	scanner := bufio.NewScanner(openedFile)

	var i int32 = 0
	for scanner.Scan() {
		i++
		line := scanner.Text()
		fmt.Printf("Line-%d: %s\n", i, line)

	}
}
