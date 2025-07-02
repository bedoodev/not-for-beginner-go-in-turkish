package main

import (
	"bufio"
	"fmt"
	"os"
)

func FlushMail() {
	writer := bufio.NewWriter(os.Stdout)

	greeting := []byte("\tHello,\n\n")
	introducing := []byte("\t  My name is Bedirhan, I am a software engineer who interested mostly backend and data things.\n\n")
	bye := []byte("\tKind Regards,\n\tBedirhan Sahin\n")

	_, _ = writer.Write([]byte("\n\t******************************************************************************\n"))
	_, _ = writer.Write(greeting)
	_, _ = writer.Write(introducing)
	_, _ = writer.Write(bye)

	if err := writer.Flush(); err != nil {
		fmt.Println("Error flushing:", err)
		return
	}

}
