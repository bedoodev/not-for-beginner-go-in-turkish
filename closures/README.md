## 🔹 Scope Nedir?

Scope, bir değişkenin nerelerde geçerli ve erişilebilir olduğunu belirler.

Go dilinde üç temel scope tipi vardır:

1. **Function Scope:** Bir fonksiyon içinde tanımlanan değişkenler sadece o fonksiyon içinde geçerlidir.
2. **Block Scope:** `if`, `for`, vb. bloklar içinde tanımlanan değişkenler sadece o blok içinde erişilebilir.
3. **Package Scope:** Fonksiyon dışında tanımlanan değişkenler aynı paket içinde her yerden erişilebilir.

### Örnek:

```go
func main() {
    x := 10 // sadece main fonksiyonu içinde geçerli

    if x > 5 {
        y := 20 // sadece bu if bloğu içinde geçerli
        fmt.Println(x, y)
    }

    // fmt.Println(y) // HATA: y burada görünmez
}
```

## 🔹 Closure Nedir?

Closure, bir fonksiyonun kendi dışındaki değişkenlere erişmeye devam edebilmesi anlamına gelir. Yani bir closure, dış scope'taki değişkenleri “hatırlar”.

```go
func main() {
	// counter fonksiyonunu çağırarak bir closure elde ediyoruz.
	// Bu closure, i değişkenini hatırlayan bir fonksiyondur.
	sequence := counter()

	// closure'ı 4 kez çağırıyoruz. Her seferinde i bir artırılır.
	fmt.Println(sequence()) // 1
	fmt.Println(sequence()) // 2
	fmt.Println(sequence()) // 3
	fmt.Println(sequence()) // 4
}

// counter fonksiyonu bir closure döndürür.
// Bu closure, i değişkenini kendi içinde saklar ve her çağrıldığında i'yi 1 artırır.
func counter() func() int {
	i := 0 // i değişkeni, closure tarafından hatırlanacak
	fmt.Println("i'nin önceki değeri:", i)

	// Aşağıdaki fonksiyon, i'yi 1 artırır ve sonucu döner.
	return func() int {
		i++
		fmt.Println("i değişkenine 1 eklendi")

		return i
	}
}
```

```bash
>> go run main.go

previous value of i: 0
added 1 to i
1
added 1 to i
2
added 1 to i
3
added 1 to i
4
```
