# Go'da `encoding/base64` nedir

`encoding/base64`, Go'nun standart kütüphanelerinden biridir ve ikili (binary) veriyi **Base64** metin formatına çevirir.  
Bu sayede resim, PDF veya özel karakterli veriler **JSON**, **XML** ya da **HTTP** gövdesi gibi metin tabanlı ortamlarda güvenle taşınabilir.

---

## 1. Neden Base64?

- **Metin uyumluluğu:** İkili veriyi yalnızca `A–Z`, `a–z`, `0–9`, `+`, `/` ve `=` karakterleriyle temsil eder.
- **Deterministik çıktı:** Aynı girdi → aynı Base64 sonucu.
- **Kolay geri dönüş:** `Decode` fonksiyonları ile veriyi eski hâline getirmek basittir.

---

## 2. Temel Kavramlar

| Kavram           | Açıklama                                                         |
| ---------------- | ---------------------------------------------------------------- |
| `StdEncoding`    | Klasik Base64 alfabesini kullanır (`+`, `/` karakterleri dahil). |
| `URLEncoding`    | URL güvenli sürüm; `+` yerine `-`, `/` yerine `_` kullanır.      |
| `EncodeToString` | `[]byte` → Base64 metin.                                         |
| `DecodeString`   | Base64 metin → orijinal `[]byte`.                                |

---

## 3. Basit Örnekler

### 3.1. Veriyi Base64'e Kodlamak

```go
package main

import (
    "encoding/base64"
    "fmt"
)

func main() {
    veri := []byte("Merhaba, Base64!")

    kodlanmis := base64.StdEncoding.EncodeToString(veri)
    fmt.Println("Kodlanmış:", kodlanmis) // TWVyYWJhYSwgQmFzZTY0IQ==
}
```

### 3.2. Base64 Metnini Çözmek

```go
cozulen, err := base64.StdEncoding.DecodeString("TWVyYWJhYSwgQmFzZTY0IQ==")
if err != nil {
    panic(err)
}
fmt.Println("Çözülen:", string(cozulen)) // Merhaba, Base64!
```

### 3.3. URL Güvenli Kodlama

```go
data := []byte("Go/URL Örneği")
urlSafe := base64.URLEncoding.EncodeToString(data)
fmt.Println("URL Güvenli:", urlSafe) // R28_VVJMI8O8cm5la8Sx
```

Bu çıktı `+` veya `/` içermediğinden, **query parametreleri** veya **cookie** değerleri içinde sorunsuz kullanılabilir.

---

## 4. Padding (`=`) Hakkında

Base64, veriyi 24‑bit bloklar hâlinde işler.  
Girdi byte sayısı 3'ün katı değilse eksik kısım `=` karakteriyle doldurulur (padding). Örneğin:

- `"Go"` → `R28=`
- `"G"` → `Rw==`

> `URLEncoding` kullanırken padding karakterleri isteğe bağlıdır; ancak çoğu senaryoda bırakılması önerilir.

---

## 5. Özet 🚀

- Kodlamak için **`EncodeToString`**, çözmek için **`DecodeString`** yeterli.
- `StdEncoding` genel kullanım içindir; `URLEncoding` URL ve benzeri ortamlarda **daha güvenli** karakter seti sunar.
- Çıktı her zaman **tekrar çözülebilir** ve her koşulda aynı sonucu üretir.
