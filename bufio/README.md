# Go'da `bufio` nedir?

`bufio`, Go'nun standart kütüphanelerinden biridir ve **buffered I/O** (arabellekli girdi‑çıktı) işlemlerini kolaylaştırır.  
Doğrudan akımdan (stream) okuma‑yazma yapmak yerine, veriyi geçici bir *buffer* içinde tutarak **daha az sistem çağrısı (syscall)** yapar; bu da performansı artırır ve kodu sadeleştirir.


* **Byte slice okuma**
* **Satır okuma (`ReadString`)**
* **Byte slice yazma**
* **String yazma ve `Flush`**


---

## 1. Byte Slice Okuma

Aşağıdaki örnek 12 byte okur ve ekrana yazdırır:

```go
package main

import (
    "bufio"
    "fmt"
    "strings"
)

func main() {
    kaynak := "Merhaba, dünya!\nBufio ile örnek."
    reader := bufio.NewReader(strings.NewReader(kaynak))

    veri := make([]byte, 12)        // 12 byte'lık dilim oluştur
    n, err := reader.Read(veri)     // ilk 12 byte'ı oku
    if err != nil {
        fmt.Println("Okuma hatası:", err)
        return
    }

    fmt.Printf("%d byte okundu: %s\n", n, veri[:n])
}
```

Çıktı:

```
12 byte okundu: Merhaba, dün
```

---

## 2. Satır Okuma (`ReadString`)

```go
line, err := reader.ReadString('\n')
```

`ReadString` belirtilen ayırıcıya (*delimiter*) kadar okur ve satırı döndürür.

```go
// Devam eden main() içinde:
satir, err := reader.ReadString('\n')
if err != nil {
    fmt.Println("Satır okuma hatası:", err)
    return
}
fmt.Print("Satır:", satir) // Satır: Bufio ile örnek.
```

---

## 3. Byte Slice Yazma

```go
package main

import (
    "bufio"
    "bytes"
    "fmt"
)

func main() {
    var hedef bytes.Buffer          // Bellek içi hedef
    writer := bufio.NewWriter(&hedef)

    data := []byte("Selam bufio paketi!\n")
    n, err := writer.Write(data)    // Byte dilimini yaz
    if err != nil {
        fmt.Println("Yazma hatası:", err)
        return
    }
    fmt.Printf("%d byte yazıldı\n", n)

    writer.Flush()                  // Buffer'ı boşalt
    fmt.Print(hedef.String())       // Selam bufio paketi!
}
```

---

## 4. String Yazma ve `Flush`

`WriteString` doğrudan string alır; `Flush` buffer'daki veriyi gerçek çıktıya gönderir:

```go
writer.WriteString("Bugün nasılsın?\n")
if err := writer.Flush(); err != nil {
    fmt.Println("Flush hatası:", err)
}
```
---
## 5. Mail Örneği
`flushing.go` dosyasında bulunan `FlushAllContent` fonksiyonunu çağırarak 
`go run .` komutu ile çalıştırıp çıktıyı inceleyebilirsiniz.

---

## Kapanış 🚀

* `Reader` ve `Writer` tipi ile küçük parça okuma‑yazma işlemlerini **buffer** sayesinde toplu hâlde yapabilirsiniz.  
* `Scanner` kullanarak satır satır veya kelime kelime okuma da mümkündür (detaylara girmedik).  
* `Flush` çağrısını **unutmayın**; aksi hâlde yazdığınız veriler buffer'da kalabilir.

Artık basit metinlerle pratik yapabilir; her adımı terminalde test ederek `bufio`’nun sağladığı kolaylığı görebilirsiniz!

