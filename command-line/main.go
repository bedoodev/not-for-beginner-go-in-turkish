package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// **************************** Args and Flags ****************************

	// fmt.Println("Command:", os.Args[0])

	// for i, arg := range os.Args {
	// 	fmt.Printf("Argument-%d: %s\n", i, arg)
	// }
	// run the following command:
	// > go run command-line/main.go hello world from go
	// output should be:
	// > Argument-0: /var/folders/x2/hfz9539d7xb0fgz9564prby40000gn/T/go-build3499650752/b001/exe/main
	// > Argument-1: hello
	// > Argument-2: world
	// > Argument-3: from
	// > Argument-4: go

	// var name string
	// var age int
	// var male bool

	// flag.StringVar(&name, "name", "Somebody", "Name of the user")
	// flag.IntVar(&age, "age", 20, "Age of the user")
	// flag.BoolVar(&male, "male", true, "Gender of the user")

	// // run the following command:
	// // > go run command-line/main.go -name "Bedirhan Şahin" -age 26 -male true
	// flag.Parse()
	// fmt.Printf("Name: %s, Age: %d, Male: %t\n", name, age, male) // Name: Bedirhan Şahin, Age: 26, Male: true

	// **************************** Sub Commands ****************************

	subCommand1 := flag.NewFlagSet("firstSub", flag.ExitOnError)
	subCommand2 := flag.NewFlagSet("secondSub", flag.ExitOnError)

	firstFlag := subCommand1.Bool("processing", false, "Command Processing Status")
	secondFlag := subCommand1.Int("bytes", 1024, "Byte length of result")

	flagsc2 := subCommand2.String("language", "Go", "Enter your language")

	if len(os.Args) < 2 {
		fmt.Println("This program requires additional arguments")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "firstSub":
		subCommand1.Parse(os.Args[2:])
		fmt.Println("subCommand1:")
		fmt.Println("Processing:", *firstFlag)
		fmt.Println("Bytes:", *secondFlag)

	case "secondSub":
		subCommand2.Parse(os.Args[2:])
		fmt.Println("subCommand2:")
		fmt.Println("Language:", *flagsc2)

	default:
		fmt.Println("no subcommand entered!")
		os.Exit(1)
	}

	// > go run command-line/main.go firstSub -processing=true -bytes=2048
	// > subCommand1:
	// > Processing: true
	// > Bytes: 2048

	// > go run command-line/main.go secondSub -language=Python
	// > subCommand2:
	// > Language: Python
}
