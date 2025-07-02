package main

import (
	"fmt"
	"net/url"
)

func main() {
	// [scheme://][userinfo@][host][:port][/path][?query][#fragment]

	rawURL := "https://example.com:8080/path?query=param#fragment"

	parsedURL, err := url.Parse(rawURL)

	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	fmt.Println("Scheme:", parsedURL.Scheme)            // Scheme: https
	fmt.Println("Host:", parsedURL.Host)                // Host: example.com:8080
	fmt.Println("Port:", parsedURL.Port())              // Port: 8080
	fmt.Println("Path:", parsedURL.Path)                // Path: /path
	fmt.Println("Raw Query:", parsedURL.RawQuery)       // Raw Query: query=param
	fmt.Println("Raw Fragment:", parsedURL.RawFragment) // Raw Fragment:

	fmt.Print("\n***************************************\n\n")

	rawURL1 := "https://example.com/path?name=John&lastName=Doe&age=30"
	parsedURL1, err := url.Parse(rawURL1)

	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	queryParams := parsedURL1.Query()
	fmt.Println("Query Params of rawURL1:", queryParams)   // Query Params of rawURL1: map[age:[30] lastName:[Doe] name:[John]]
	fmt.Println("Name:", queryParams.Get("name"))          // Name: John
	fmt.Println("Last Name:", queryParams.Get("lastName")) // Last Name: Doe
	fmt.Println("Age:", queryParams.Get("age"))            // Age: 30

	fmt.Print("\n***************************************\n\n")

	// Building URL
	baseURL := url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "/path",
	}
	query := baseURL.Query()
	query.Set("name", "John")
	query.Set("lastName", "Doe")
	query.Set("age", "25")
	baseURL.RawQuery = query.Encode()

	fmt.Println("Built URL:", baseURL.String()) // Built URL: https://example.com/path?age=25&lastName=Doe&name=John
}
