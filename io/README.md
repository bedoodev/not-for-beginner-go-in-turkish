# IO Nedir?

- io Go’nun standart kütüphanesindeki bir pakettir.
- Dosya okuma/yazma, ekrana yazma, ağ üzerinden veri alışverişi, vb. işlemleri kapsar.
- Farklı kaynaklardan (dosya, ağ, bellek) veri almak (input) ve vermek (output) için ortak arabirimler (interface’ler) sunar.

### En önemli 2 interface:

- io.Reader: Bir kaynaktan veri okuyabilen her şey.
- io.Writer: Bir kaynağa veri yazabilen her şey.

```go
package main

import (
    "fmt"
    "io"
    "strings"
)

func main() {
    reader := strings.NewReader("Merhaba dünya!")
    buf := make([]byte, 8)
    n, _ := reader.Read(buf)
    fmt.Println(string(buf[:n])) // "Merhaba "
}
```

# Buffer Nedir?

- Buffer veriyi geçici olarak depolayan bir ara katmandır.
- `Okuma` ve `Yazma` işlemlerinde veriyi biriktirip topluca işlemek için kullanılır.
- Örneğin: Bir dosyadan veya ağdan bir veriyi parça parça almak yavaş olacağından
  dolayı, buffer ile verileri biriktirip daha az sayıda büyük işlem yaparak performansı
  arttırabiliriz.

## Örnek

```go
package main

import (
    "bytes"
    "fmt"
)

func main() {
    var buf bytes.Buffer

    buf.WriteString("Merhaba, ")
    buf.WriteString("Dünya!") // Buffer'a yazıyoruz.

    fmt.Println(buf.String()) // Buffer'daki tüm veriyi yazdırır: "Merhaba, Dünya!"
}
```

```go
package main

import (
    "bufio"
    "os"
)

func main() {
    f, _ := os.Create("deneme.txt")
    defer f.Close()

    writer := bufio.NewWriter(f)
    writer.WriteString("Buffered yazma örneği!\n")
    writer.Flush() // Buffer'ı diske (dosyaya) boşalt!
}

```

# Pipe Nedir?

- Bir pipe bir işlemin çıktısını başka bir işleme doğrudan bağlamak için kullanılır.

```go
package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    reader, writer := io.Pipe()

    // Yazma işlemi (go routine ile)
    go func() {
        writer.Write([]byte("Hello from pipe!\n"))
        writer.Close()
    }()

    // Okuma işlemi
    io.Copy(os.Stdout, reader)
}
```

> Ne yaptık?

- io.Pipe() ile bir reader ve writer oluşturduk.
- writer.Write(...) ile boruya veri yazdık.
- reader ucundan bu veriyi okuduk ve ekrana bastık.
- İki taraf da eş zamanlı çalışıyor.

## Buffer ve Pipe Kullanım Yerleri

### Buffer

- Veriyi toplu işlemek (ör. büyük dosya okuma/yazma)
- String veya byte’ları biriktirmek

### Pipe

- Bir işlemin çıktısını diğerine aktarmak
- Ör: Bir process bir şey üretir, diğeri anında okur
- Aynı program içinde paralel (go routine) veri akışı
