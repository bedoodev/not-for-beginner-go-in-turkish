## 🔹 Recursion Nedir?

Recursion, yani özyineleme, bir fonksiyonun **kendi kendini çağırması**dır. Genellikle problemleri daha küçük alt parçalara bölerek çözmek için kullanılır.

Bir özyinelemeli fonksiyonun **iki temel bileşeni** vardır:

1. **Base case:** Fonksiyonun duracağı ve tekrar çağırmayı bırakacağı koşul.
2. **Recursive case:** Fonksiyonun kendi kendini çağırdığı durum.

### 📌 Örnek: Faktöriyel Hesaplama

```go
func main() {
	fmt.Println(factorial(5))
	fmt.Println(factorial(10))

}

func factorial(n int) int {
	// Base case: 0 değerinin faktöriyeli 1'dir
	if n == 0 {
		return 1
	}

	// Recursive case: n değerinin faktöriyeli n * factorial(n - 1)'dir
	return n * factorial(n-1)
	// n * (n - 1) * (n - 2) * ...... * 1
}

```

```bash
>> go run main.go
120
3628800
```

### 📌 Örnek: Verilen Sayının Rakamlarının Toplamını Hesaplama

```go
func sumOfDigits(n int) int {

	// Base case: eğer n tek haneli bir sayı ise cevap n olur
	if n < 10 {
		return n
	}
	// Recursive case: son basamağı al ve kalan sayı için tekrar çağır
	return n%10 + sumOfDigits(n/10)
}

```

```bash
>> go run main.go
6
11
```
