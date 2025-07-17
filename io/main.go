package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func readFromReader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		log.Fatalln("Error reading from reader.", err)
	}
	fmt.Println(string(buf[:n]))
}

func writeToWriter(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error writing to writer.", err)
	}
}

func closeResource(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatalln("Error closing resource.", err)
	}
}

func bufferExample() {
	var buf bytes.Buffer // stack
	buf.WriteString("Hello buffer!")
	fmt.Println(buf.String())

}

func multiReaderExample() {
	r1 := strings.NewReader("Hello")
	r2 := strings.NewReader(" World!")
	mr := io.MultiReader(r1, r2)
	buf := new(bytes.Buffer) // heap
	_, err := buf.ReadFrom(mr)
	if err != nil {
		log.Fatalln("Error reading from multi reader.", err)
	}

	fmt.Println(buf.String())

}

func pipeExample() {
	reader, writer := io.Pipe()
	go func() {
		writer.Write([]byte("Hello Pipe!"))
		writer.Close()
	}()
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	fmt.Println(buf.String())
}

func writeToFile(filepath, data string) {
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Error opening file.", err)
	}
	defer closeResource(file)

	_, err = file.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error writing to file.", err)
	}
	// Type(value)
	// writer := io.Writer(file)
	// _, err = writer.Write([]byte(data))
	// if err != nil {
	// 	log.Fatalln("Error writing to file.", err)
	// }
}

func main() {
	fmt.Println("******* Read from Reader *******")
	readFromReader(strings.NewReader("Hello Reader!"))

	fmt.Println("******* Write to Writer *******")
	var writer bytes.Buffer
	writeToWriter(&writer, "Hello Writer!")
	fmt.Println(writer.String())

	fmt.Println("******* Buffer Example *******")
	bufferExample()

	fmt.Println("******* Multi Reader Example *******")
	multiReaderExample()

	fmt.Println("******* Pipe Example *******")
	pipeExample()

	filepath := "io/io.txt"
	writeToFile(filepath, "Hello File!\n")

	myResource := MyResource{name: "testResource"}
	closeResource(myResource)
}

type MyResource struct {
	name string
}

func (m MyResource) Close() error {
	fmt.Printf("Closing resource %s\n", m.name)
	return nil
}
