## 1 | Generics Nedir?

Go’da **generics**, fonksiyon veya tip tanımına **tip parametresi** ekleyerek, farklı tiplerle tekrar eden kodu tek noktada tanımlamamıza olanak verir.

```go
// T, fonksiyon çağrısında belirlenen herhangi bir tip olabilir
func Identity[T any](v T) T {
    return v
}
```

- `T` → tip parametresi
- `any` → sınırsız (boş interface gibi) temel kısıtlama

---

## 2 | Tip Parametreli Fonksiyon

```go
func Sum[T int | int32 | int64](a, b T) T { // basit union constraint
    return a + b
}

func main() {
    fmt.Println(Sum[int](3, 4))   // 7
    fmt.Println(Sum(5, 6))        // Tip çıkarımı – 11
}
```

> **Not:** Basit örneklerde tip çıkarımı (type inference) sayesinde `Sum(5,6)` çağrısı yeterlidir.

---

## 3 | Generics ile Struct Tanımı

```go
type Box[T any] struct {
    Value T
}

func (b Box[T]) Get() T {
    return b.Value
}

func main() {
    intBox := Box[int]{Value: 42}
    fmt.Println(intBox.Get()) // 42

    strBox := Box[string]{Value: "Go"}
    fmt.Println(strBox.Get()) // Go
}
```

---

## 4 | Basit Constraint Kullanımı

En yaygın **yerleşik kısıt**: `comparable` → `==`, `!=` operatörlerini destekleyen tipler.

```go
func IndexOf[T comparable](slice []T, target T) int {
    for i, v := range slice {
        if v == target { return i }
    }
    return -1
}
```

---

## 5 | Instantiation (Nesne Oluşturma)

```go
n := Identity[int](10) // açık tip belirtimi
s := Identity("Go")   // çıkarımlı tip
_ = n; _ = s
```

---

## 🧠 Özet

| Başlık     | Ana Fikir                    |
| ---------- | ---------------------------- |
| Tip Param. | `func Foo[T any](v T)`       |
| any        | Kısıtsız tip                 |
| comparable | `==` / `!=` destekler        |
| Union      | \`T int int64 int32 \`       |
| Struct     | `type Box[T] struct { ... }` |
