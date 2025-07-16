package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName    string  `json:"first_name"`
	Age          int     `json:"age,omitempty"`
	EmailAddress string  `json:"email"`
	Address      Address `json:"address"`
}

type Address struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type Employee struct {
	FullName   string `json:"full_name"`
	Age        int    `json:"age"`
	EmployeeID string `json:"emp_id"`
}

func main() {
	person := Person{FirstName: "Bedirhan", EmailAddress: "bedoodev@gmail.com"}
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
	// if omitempty => {"first_name":"Bedirhan","email":"bedoodev@gmail.com"}
	// else => {"first_name":"Bedirhan","age":0,"email":"bedoodev@gmail.com"}

	person2 := Person{
		FirstName:    "Bedirhan",
		Age:          26,
		EmailAddress: "bedoodev@gmail.com",
		Address: Address{
			City:    "Istanbul",
			Country: "Turkey",
		},
	}

	jsonData2, err := json.Marshal(person2)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println(string(jsonData2))

	jsonData3 := `{"full_name": "Bedirhan Şahin", "age": 26, "emp_id": "0001"}`
	employeeFromJson := Employee{}
	err = json.Unmarshal([]byte(jsonData3), &employeeFromJson)
	if err != nil {
		fmt.Println("Error unmarshalling from JSON:", err)
		return
	}

	fmt.Println(employeeFromJson)

	listOfCityCountry := []Address{
		{City: "Istanbul", Country: "Turkey"},
		{City: "Ankara", Country: "Turkey"},
		{City: "Izmir", Country: "Turkey"},
		{City: "Antalya", Country: "Turkey"},
		{City: "Edirne", Country: "Turkey"},
	}

	fmt.Println(listOfCityCountry)

	jsonList, err := json.Marshal(listOfCityCountry)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println(string(jsonList))

	// Handling Unknown JSON Structures
	jsonData4 := `{"name": "John", "salary": "2500"}`

	var data map[string]interface{}
	err = json.Unmarshal([]byte(jsonData4), &data)
	if err != nil {
		fmt.Println("Error unmarshalling from JSON:", err)
		return
	}
	fmt.Println(data)
}
