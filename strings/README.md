## 1 | `strings` Paketi Nedir?

`strings` paketi, UTF‑8 string’ler üzerinde kolay arama, kesme, birleştirme ve dönüştürme işlemleri sağlayan saf‑Go yardımcı fonksiyonlar topluluğudur.

```go
import "strings"
```

---

## 2 | Arama & Karşılaştırma

| Fonksiyon                 | Ne Yapar?                         | Örnek Kod (çıktı yorumda)                      |
| ------------------------- | --------------------------------- | ---------------------------------------------- |
| `Contains`                | Alt dize var mı?                  | `strings.Contains("golang", "go") // true`     |
| `HasPrefix` / `HasSuffix` | Baş/son kontrolü                  | `strings.HasPrefix("http://", "http") // true` |
| `Index` / `LastIndex`     | İlk / son pozisyon                | `strings.Index("banana", "na") // 2`           |
| `Count`                   | Alt dize tekrar sayısı            | `strings.Count("cheese", "e") // 3`            |
| `Compare`                 | Sözlük sırası (-1/0/1)            | `strings.Compare("a", "b") // -1`              |
| `EqualFold`               | Harf büyüklüğüne duyarsız eşitlik | `strings.EqualFold("Go", "gO") // true`        |

---

## 3 | Dönüştürme İşlemleri

```go
s := " GoLang "
fmt.Println(strings.ToLower(s))      // " golang "
fmt.Println(strings.ToUpper(s))      // " GOLANG "
fmt.Println(strings.TrimSpace(s))    // "GoLang"
fmt.Println(strings.Replace(s, " ", "_", -1)) // "_GoLang_"
fmt.Println(strings.Repeat("*", 5)) // "*****"
```

> **Not:** `Replace`’deki son parametre _adet_ (`-1` → tüm eşleşmeleri değiştir).

---

## 4 | Kesme & Birleştirme

```go
csv := "a,b,c"
parts := strings.Split(csv, ",")   // ["a" "b" "c"]
joined := strings.Join(parts, ";") // "a;b;c"

line := "  foo bar   baz "
fields := strings.Fields(line)       // ["foo" "bar" "baz"]

trimmed := strings.Trim(line, " ")  // "foo bar   baz"
```

---

## 5 | `strings.Builder`

[`strings.Builder`](https://pkg.go.dev/strings#Builder), birçok parçadan **tek bir string** üretirken her adımda yeni bellek tahsisi (allocation) yapmaktan kaçınmak için tasarlanmış hafif bir yapıdır.  
`+` veya `fmt.Sprintf` yerine `Builder` kullanmak, özellikle döngü içinde büyüyen metinlerde **daha az kopya, daha az GC** anlamına gelir.

### Neden Kullanılır?

- **Performans:** Art arda birleştirmelerde slice-copy yerine tek tampon kullanır.
- **Kullanım Kolaylığı:** Sıfır değerle hazır; `WriteString`, `WriteRune`, `WriteByte`, `fmt.Fprint` vb. destekler.

### Temel Örnek

```go
var b strings.Builder          // sıfır değer yeterli
b.WriteString("Go")
b.WriteRune(' ')               // tek rune ekle
fmt.Fprint(&b, "rocks!")       // fmt ailesiyle uyumlu

result := b.String()           // "Go rocks!"
b.Reset()                      // sıfırla ve tekrar kullan
```

---

## 🧠 Özet

| Kategori       | Öne Çıkan Fonksiyonlar                             |
| -------------- | -------------------------------------------------- |
| Arama          | `Contains`, `Index`, `HasPrefix` / `HasSuffix`     |
| Karşılaştırma  | `Compare`, `EqualFold`, `Count`                    |
| Dönüştürme     | `ToUpper`, `ToLower`, `Title`, `Repeat`, `Replace` |
| Kesme/Birleşim | `Split`, `Join`, `Fields`, `Trim`, `TrimSpace`     |

---
