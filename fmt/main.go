package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Print("Print")             // Print
	fmt.Println(" Println")        //  Print \n hali
	fmt.Printf("Printf: %d\n", 42) // Printf: 42

	//────────────────────────────── 2) Fprint ailesi (io.Writer)
	w := os.Stdout
	fmt.Fprint(w, " Fprint")               //  Fprint
	fmt.Fprintln(w, " Fprintln")           //  Fprintln\n
	fmt.Fprintf(w, "Fprintf: 0x%X\n", 255) // Fprintf: 0xFF

	//────────────────────────────── 3) Sprint ailesi (string döndürür)
	s1 := fmt.Sprint("Sprint", 1)           // "Sprint1"
	s2 := fmt.Sprintln("Sprintln", 2)       // "Sprintln 2\n"
	s3 := fmt.Sprintf("Sprintf: %.1f", 3.2) // "Sprintf: 3.2"

	fmt.Println(s1)
	fmt.Print(s2)
	fmt.Println(s3)

	//────────────────────────────── 4) Append ailesi (Go 1.19+)
	var buf []byte
	buf = fmt.Append(buf, "A")         // "A"
	buf = fmt.Appendln(buf, "B")       // "A B\n"
	buf = fmt.Appendf(buf, " C=%d", 3) // "A B\n C=3"
	fmt.Println(string(buf))

	//────────────────────────────── 5) Errorf
	innerErr := fmt.Errorf("inner error")
	err := fmt.Errorf("Errorf demo: %w", innerErr)
	fmt.Println(err) // Errorf demo: inner error

	//────────────────────────────── 6) Okuma (Scan ailesi)

	// Fscan – strings.Reader üzerinden
	var a, b int
	input := strings.NewReader("7 8\n")
	fmt.Fscan(input, &a, &b)
	fmt.Printf("Fscan read: %d %d\n", a, b) // Fscan read: 7 8

	// Fscanln – satır sonuna kadar
	var w1, w2 string
	input2 := strings.NewReader("foo bar\n")
	fmt.Fscanln(input2, &w1, &w2)
	fmt.Printf("Fscanln read: %s %s\n", w1, w2) // Fscanln read: foo bar

	// Fscanf – formatlı okuma
	var num1, num2 int
	input3 := strings.NewReader("10/11")
	fmt.Fscanf(input3, "%d/%d", &num1, &num2)
	fmt.Printf("Fscanf read: %d %d\n", num1, num2) // Fscanf read: 10 11

	// Sscan – string kaynağından boşluk ayraçlı
	var sA, sB int
	fmt.Sscan("3 4", &sA, &sB)
	fmt.Printf("Sscan read: %d %d\n", sA, sB) // Sscan read: 3 4

	// Sscanln – string kaynağından satır sonuna kadar
	var txt string
	fmt.Sscanln("golang", &txt)
	fmt.Println("Sscanln read:", txt) // Sscanln read: golang

	// Sscanf – formatlı string okuma
	var day, month int
	fmt.Sscanf("01-02", "%02d-%02d", &day, &month)
	fmt.Printf("Sscanf read: %d %d\n", day, month) // Sscanf read: 1 2

	// String formatting
	fmt.Printf("%5d\n", 55)     //    55 -> min 5 char
	fmt.Printf("%5d\n", 559898) // 559898

	fmt.Printf("|%10s|\n", "Hello") // |     Hello|
	fmt.Printf("|%-10s|\n", "Hello") // |Hello     |

}
