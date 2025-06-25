// strings_demo.go
// En sık kullanılan strings fonksiyonlarını tek bir programda örnekler.

package main

import (
	"fmt"
	"strings"
)

func main() {
	base := "  GoLang, go far!  "

	//— Arama & Karşılaştırma —//
	fmt.Println(strings.Contains(base, "Go"))       // true
	fmt.Println(strings.HasPrefix(base, "  Go"))    // true
	fmt.Println(strings.HasSuffix(base, "!  "))     // true
	fmt.Println(strings.Index(base, "go"))          // 9  (ilk eşleşme)
	fmt.Println(strings.LastIndex(base, "a"))       // 10 (son eşleşme)
	fmt.Println(strings.Count(base, "a"))           // 1
	fmt.Println(strings.Compare("apple", "banana")) // -1
	fmt.Println(strings.EqualFold("Go", "gO"))      // true

	//— Dönüştürme —//
	fmt.Println(strings.ToLower(base))                  // "  golang, go far!  "
	fmt.Println(strings.ToUpper(base))                  // "  GOLANG, GO FAR!  "
	fmt.Println(strings.Title("go makes you fly"))      // "Go Makes You Fly"
	fmt.Println(strings.TrimSpace(base))                // "GoLang, go far!"
	fmt.Println(strings.Trim(base, " "))                // "GoLang, go far!"
	fmt.Println(strings.Repeat("=", 5))                 // "====="
	fmt.Println(strings.Replace(base, "Go", "Run", -1)) // "  RunLang, run far!  " (hepsi)

	//— Kesme & Birleştirme —//
	csv := "a,b,c"
	fmt.Println(strings.Split(csv, ","))                    // ["a" "b" "c"]
	fmt.Println(strings.Join([]string{"x", "y", "z"}, "-")) // "x-y-z"
	sentence := "foo bar   baz"
	fmt.Println(strings.Fields(sentence)) // ["foo" "bar" "baz"]

	//- Builder Strings -//
	var builder strings.Builder

	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("Go!")

	result := builder.String()
	fmt.Println(result) // Hello, Go!

}
