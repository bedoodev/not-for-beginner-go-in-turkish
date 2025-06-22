## 1 | Struct Nedir?

`struct`, ilişkili alanları (field) tek bir tip altında gruplamak için kullanılır. Sınıf benzeri ancak gömülü davranışsız bir veri yapısıdır.

```go
// Kişi adında basit bir struct tanımı
type Person struct {
    Name string
    Age  int
}
```

---

## 2 | Değişken Oluşturma

```go
var p1 Person                 // zero‑value alanlar: "", 0
p2 := Person{"Ada", 30}      // sıralı literalle
p3 := Person{Name: "Ali"}    // adlandırılmış literal (Age=0)
```

---

## 3 | Alanlara Erişim & Güncelleme

```go
fmt.Println(p2.Name)  // Ada
p2.Age++              // 31
```

---

## 4 | Pointer ile Çalışmak

Struct büyükse pointer vererek kopya maliyetinden kaçınırız.

```go
func birthday(pp *Person) {
    pp.Age++
}

birthday(&p3) // Age artırıldı
```

---

## 5 | Method Tanımlamak

Method’lar, struct’a bağlı fonksiyonlardır.

```go
func (p Person) Greet() string {
    return "Merhaba " + p.Name
}

fmt.Println(p2.Greet()) // Merhaba Ada
```

> **Değer vs Pointer Receiver**
>
> - `(p Person)` → kopya üzerinde çalışır.
> - `(p *Person)` → orijinal veriyi değiştirir.

---

## 6 | Embedding

Bir struct, başka bir struct’ı embed edebilir.; bu, "basit kalıtım" sunar.

```go
type Address struct {
    City string
    Zip  int
}

type Employee struct {
    Person   // anonim embedded
    Position string
    Address  // birden çok da gömülebilir
}

e := Employee{
    Person: Person{"Eda", 28},
    Position: "Developer",
    Address: Address{"İstanbul", 34000},
}

fmt.Println(e.Name, e.City) // Eda İstanbul
```

---

## 7 | Anonymous Structs

```go
user := struct {
    username string
    email    string
}{username: "bedoodev", email: "bedoodev@gmail.com"}
fmt.Println(user) // {bedoodev bedoodev@gmail.com}
```

---

## 8 | Struct Tag’leri (örn. JSON)

```go
type Product struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
}

b, _ := json.Marshal(Product{1, "Book"})
fmt.Println(string(b)) // {"id":1,"title":"Book"}
```
