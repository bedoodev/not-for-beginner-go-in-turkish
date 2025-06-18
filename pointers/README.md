## 🔹 Pointer Nedir?

**Pointer**, bir değişkenin bellekteki adresini tutan yapıdır.
Bir fonksiyon alacağı argümanın değerini değiştirecekse kullanılır.

- `&x` → `x` değişkeninin adresini alır (reference)
- `*p` → `p` bir pointer ise, işaret ettiği değeri verir (dereference)

----

## 📌 Basit Pointer Örneği

```go
func main() {
	var a int = 10
	var ptr *int = &a // reference

	fmt.Println(a)
	fmt.Println(ptr)
    // fmt.Println(*ptr) // dereference

    /* var ptr2 *int
	 fmt.Println(ptr2) */  // nil
}

```

```bash
>> go run main.go
10
0x14000112020

>> go run main.go
10
0x14000104020

>> go run main.go
10
0x1400000e090
```

> Her `go run` çalıştırıldığında pointer (adres) değeri farklı olabilir.  
> Bu, Go runtime bellek yönetimine bağlı olarak değişir.  
> Adresler rastgele değildir ama her çalıştırmada farklı görünmeleri doğaldır.

## 📌 Pointer Kullanım Örneği

```go
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
```

```bash
>> go run main.go
10
11
```
