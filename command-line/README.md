# Go'da Command Line Arguments ve Subcommand Kullanımı

Go programlarında komut satırı argümanları (command line arguments) ve flag yönetimi, hem temel CLI araçları geliştirmek hem de profesyonel uygulamalar yazmak için çok önemlidir. Bu dökümanda, iki ana başlıkta Go'nun standart kütüphaneleriyle command line arguments nasıl işlenir, flag nasıl kullanılır ve subcommand yapısı nasıl oluşturulur; bunları gerçekçi ve açıklamalı örneklerle öğreneceksin.

---

## 1. Temel: os.Args ile Command Line Arguments Okuma

Go'da `os.Args` dizisi, program çalıştırılırken verilen tüm argümanlara erişmemizi sağlar.  
`os.Args[0]` her zaman programın kendisini (yoluyla birlikte) gösterir.  
Kalanlar ise sırasıyla verilen argümanlardır.

### Örnek:

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    for i, arg := range os.Args {
        fmt.Printf("Argument-%d: %s\n", i, arg)
    }
}
```

#### Komut:

```bash
go run main.go Hello Golang World
```

#### Çıktı:

```
Argument-0: /tmp/go-build…/exe/main
Argument-1: Hello
Argument-2: Golang
Argument-3: World
```

> **Not:** Bu yöntemle, tüm argümanlara **doğrudan** erişebilirsin. Ancak flag gibi anahtar-değer çiftlerini elle ayrıştırman gerekir.

---

## 2. flag Paketi ile Argument Yönetimi

Go'nun `flag` paketi, command line üzerinden parametre almak ve bunlara default (varsayılan) değerler vermek için kullanılır.

### Temel Kullanım:

```go
package main

import (
    "flag"
    "fmt"
)

func main() {
    var city string
    var age int
    var active bool

    flag.StringVar(&city, "city", "Istanbul", "Your city")
    flag.IntVar(&age, "age", 30, "Your age")
    flag.BoolVar(&active, "active", false, "Are you active?")

    flag.Parse()

    fmt.Printf("City: %s, Age: %d, Active: %t\n", city, age, active)
}
```

#### Komut:

```bash
go run main.go -city Ankara -age 28 -active=true
```

#### Çıktı:

```
City: Ankara, Age: 28, Active: true
```

---

## 3. Subcommand (Alt Komut) Yapısı

CLI araçlarında, `git commit`, `docker run` gibi farklı **subcommand**'lar kullanılır.  
Go'da `flag.NewFlagSet` ile bu yapıyı kolayca kurabilirsin.

### Örnek Kod:

```go
package main

import (
    "flag"
    "fmt"
    "os"
)

func main() {
    // Subcommand tanımları
    greetCmd := flag.NewFlagSet("greet", flag.ExitOnError)
    greetName := greetCmd.String("name", "Visitor", "Name to greet")

    sumCmd := flag.NewFlagSet("sum", flag.ExitOnError)
    num1 := sumCmd.Int("num1", 0, "First number")
    num2 := sumCmd.Int("num2", 0, "Second number")

    if len(os.Args) < 2 {
        fmt.Println("Bir subcommand girilmeli: greet veya sum")
        os.Exit(1)
    }

    switch os.Args[1] {
    case "greet":
        greetCmd.Parse(os.Args[2:])
        fmt.Printf("Hello, %s!\n", *greetName)
    case "sum":
        sumCmd.Parse(os.Args[2:])
        fmt.Printf("Sum: %d\n", *num1 + *num2)
    default:
        fmt.Println("Bilinmeyen subcommand:", os.Args[1])
        os.Exit(1)
    }
}
```

#### Kullanım:

1. **greet Subcommand:**

   ```bash
   go run main.go greet -name "Bedirhan"
   ```

   Çıktı:  
   `Hello, Bedirhan!`

2. **sum Subcommand:**
   ```bash
   go run main.go sum -num1 12 -num2 7
   ```
   Çıktı:  
   `Sum: 19`

---

## 4. Kapanış ve Pratik Bilgiler

- `os.Args` ile argümanlar üzerinde tam kontrol sağlarsın, ama flag kullanımı için elle parsing yapmak gerekir.
- `flag` paketiyle, anahtar-değer çiftlerini kolayca yönetebilirsin.
- `flag.NewFlagSet` ile çok sayıda komutlu (subcommand’lı) profesyonel CLI uygulamaları geliştirebilirsin.
- Flag'lerde boşluklu argümanlar kullanırken **tırnak** kullanmak şarttır:  
  `-name "Bedirhan Sahin"`

---
