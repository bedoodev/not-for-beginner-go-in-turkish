## 1 | `error` Tipi Nedir?

`error` yerleşik (built‑in) bir **interface**’tir:

```go
type error interface {
    Error() string
}
```

Bir fonksiyon hata döndürmek istediğinde imza genellikle şu şekildedir:

```go
func Foo() (T, error) { /* ... */ }
```

---

## 2 | Hata Oluşturma

| Yöntem       | Kod Örneği                                 | Açıklama                      |
| ------------ | ------------------------------------------ | ----------------------------- |
| `errors.New` | `err := errors.New("bozuk dosya")`         | Sabit (sentinel) hata metni   |
| `fmt.Errorf` | `err := fmt.Errorf("%s bulunamadı", file)` | Dinamik metin, format desteği |

```go
import (
    "errors"
    "fmt"
)
```

---

## 3 | Temel Hata Kontrolü

```go
f, err := os.Open("data.csv")
if err != nil {
    // Hata yolunu ele al
    log.Fatal(err)
}
// Devam...
```

> **Kural:** Go’da istisna yoktur; her hatayı _hemen kontrol et_ (`if err != nil`).

---

## 4 | Sentinel (Sabit) Hatalar

Belirli durumları tanımlamak için **paket‑dışı erişilebilir** değişkenler tanımlanabilir:

```go
var ErrNotFound = errors.New("kayıt bulunamadı")

func Find(id int) (Item, error) {
    if _, ok := db[id]; !ok {
        return Item{}, ErrNotFound
    }
    return db[id], nil
}
```

Kullanıcı kodu:

```go
item, err := service.Find(3)
if errors.Is(err, service.ErrNotFound) {
    // 404 dön
}
```

---

## 5 | Hata Sarmalama (Wrapping)

Ayrıntıyı kaybetmeden bağlam eklemek için `%w` kullanılır:

```go
if err := json.Unmarshal(raw, &cfg); err != nil {
    return fmt.Errorf("json çözümleme hatası: %w", err)
}
```

Kontrol:

```go
if errors.Is(err, io.EOF) { /* ... */ }
```

---

## 6 | Özel Hata Tipi

```go
type ParseErr struct {
    Line int
    Msg  string
}
func (e ParseErr) Error() string {
    return fmt.Sprintf("satır %d: %s", e.Line, e.Msg)
}
```

Kullanım:

```go
return ParseErr{Line: 42, Msg: "beklenmeyen karakter"}
```

---

## 🧠 Özet

| Konu         | Ana Fikir                                         |
| ------------ | ------------------------------------------------- |
| `error`      | Tek methodlı interface → `Error() string`         |
| Oluşturma    | `errors.New`, `fmt.Errorf`                        |
| Kontrol      | `if err != nil { ... }`                           |
| Sentinel     | Paket‑geniş                                       |
| Sarmalama    | `%w` + `errors.Is/Unwrap`                         |
| Custom Error | `type MyErr struct{}` + `Error()` implementasyonu |
