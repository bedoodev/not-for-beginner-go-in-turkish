package main

import "fmt"

func main() {

	a := 10   // a adında bir tamsayı değişkeni tanımlanır ve değeri 10 atanır
	ptr := &a // ptr, a'nın bellek adresini tutan bir pointer olur

	fmt.Println(a) // a'nın mevcut değeri ekrana yazdırılır (10)

	modifyValue(ptr) // Bellek adresi doğrudan fonksiyona gönderilir; a'nın değeri fonksiyon içinde artar

	fmt.Println(a) // a'nın yeni değeri yazdırılır (11)

}

// modifyValue fonksiyonu, bir pointer alır ve işaret ettiği değerin değerini 1 artırır
func modifyValue(ptr *int) {
	*ptr++ // pointer üzerinden erişilen değer 1 artırılır
}
