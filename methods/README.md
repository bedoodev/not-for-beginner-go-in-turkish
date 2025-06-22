## 1 | Method Nedir?

Go’da method, bir **receiver** ile ilişkilendirilmiş fonksiyondur. Receiver çoğunlukla bir struct örneğidir ve nokta (`.`) sözdizimiyle çağrılır.

```go
type Circle struct {
    R float64
}

// Alan hesaplayan method (value receiver)
func (c Circle) Area() float64 {
    return 3.14 * c.R * c.R
}

func main() {
    c := Circle{R: 2.0}
    fmt.Println(c.Area()) // 12.56
}
```

---

## 2 | Value Receiver ⇄ Pointer Receiver

| Receiver | Çalışma Biçimi                 | Ne Zaman Kullanılır?                |
| -------- | ------------------------------ | ----------------------------------- |
| `(v T)`  | Struct’un **kopyası** üzerinde | Küçük yapılar, değişiklik gerekmez  |
| `(v *T)` | **Orijinal** veri üzerinde     | Büyük yapılar, mutasyon gerekiyorsa |

```go
func (c Circle) ScaleCopy(f float64) Circle { // value ⇒ kopya
    c.R *= f
    return c            // çağıranın verisi değişmez
}

func (c *Circle) ScaleInPlace(f float64) {    // pointer ⇒ yerinde
    c.R *= f            // doğrudan orijinali değiştirir
}
```

---

## 3 | Method Zincirleme (Chaining) Örneği

Pointer receiver kullanan method’lar kendilerini döndürerek ardışık çağrılabilir.

```go
package main

import "fmt"

type Counter struct{ n int }

func (c *Counter) Add(k int) *Counter {
    c.n += k
    return c
}

func (c Counter) Value() int { return c.n }

func main() {
    v := new(Counter).Add(2).Add(3).Value()
    fmt.Println(v) // 5
}
```

---
