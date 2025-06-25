package main

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"strings"
)

func main() {
	// tmp1, err := template.New("welcome").Parse("Welcome {{.name}}! How are you doing?\n")
	// if err != nil {
	// 	panic(err)
	// }

	// data := map[string]interface{}{
	// 	"name": "Joe",
	// }

	// if err := tmp1.Execute(os.Stdout, data); err != nil {
	// 	panic(err)
	// } // Welcome Joe! How are you doing?

	// ******************************************************

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Println(name)

	templates := map[string]string{
		"welcome":      "Welcome {{.name}}! How are you doing?",
		"notification": "{{.name}}, you have a new notification: {{.notification}}",
		"error":        "Oops! An error occured: {{.errorMessage}}",
	}

	parsedTemplates := make(map[string]*template.Template)
	for name, tmpl := range templates {
		parsedTemplates[name] = template.Must(template.New(name).Parse(tmpl))
	}

	for {
		fmt.Println("\nMenu: ")
		fmt.Println("1. Join")
		fmt.Println("2. Get Notification")
		fmt.Println("3. Get Error")
		fmt.Println("4. Exit")
		fmt.Println("Choose an option: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string]interface{}
		var tmpl *template.Template

		switch choice {
		case "1":
			tmpl = parsedTemplates["welcome"]
			data = map[string]interface{}{"name": name}

		case "2":
			tmpl = parsedTemplates["notification"]
			fmt.Println("Enter a notification: ")
			notification, _ := reader.ReadString('\n')
			notification = strings.TrimSpace(notification)
			data = map[string]interface{}{"name": name, "notification": notification}

		case "3":
			tmpl = parsedTemplates["error"]
			fmt.Println("Enter an error: ")
			errorMessage, _ := reader.ReadString('\n')
			errorMessage = strings.TrimSpace(errorMessage)
			data = map[string]interface{}{"errorMessage": errorMessage}

		case "4":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Please select a valid option.")
			continue
		}

		if err := tmpl.Execute(os.Stdout, data); err != nil {
			fmt.Println("Error executing template:", err)
		}
	}

	// ******************************************************

}
