# Go'da Environment Variables

Go programlarında **environment variables** (çevre değişkenleri), uygulamaların dış ortamdan bilgi almasını, konfigürasyonların koddan bağımsız tutulmasını ve güvenliğin artırılmasını sağlar.

---

## Environment Variables Neden Kullanılır?

- **Taşınabilirlik:** Farklı ortamda (development, staging, production) uygulama ayarlarını kodu değiştirmeden yönetebilirsin.
- **Güvenlik:** API anahtarları, parola gibi gizli bilgileri kodda yazmak yerine environment variable ile saklayabilirsin.
- **Kolay Yönetim:** Deployment ve konfigürasyon işlemlerini basitleştirir.
- **Docker/Kubernetes:** Modern konteyner teknolojileri environment variable kullanımını teşvik eder.

Özetle, environment variable'lar bir uygulamanın çalıştığı ortam hakkında bilgi almasını sağlar ve uygulamanın dışarıdan kontrol edilmesini mümkün kılar.

---

## Environment Variables ile Çalışmak (Go'da)

Aşağıda, Go'da environment variable okuma, ayarlama, silme ve tüm environment variable'ları listeleme işlemleri örneklenmiştir. Değişken, fonksiyon ve teknik terimler İngilizce olarak bırakılmıştır.

```go
package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    user := os.Getenv("USER")
    home := os.Getenv("HOME")
    fmt.Println("User:", user) // User: bedirhansahin
    fmt.Println("Home:", home) // Home: /Users/bedirhansahin

    err := os.Setenv("FRUIT", "Apple")
    if err != nil {
        fmt.Println("Error setting environment variable:", err)
        return
    }

    fmt.Println("Name of Fruit:", os.Getenv("FRUIT")) // Name of Fruit: Apple

    // Tüm environment variable isimlerini listelemek:
    for _, e := range os.Environ() {
        kvpair := strings.SplitN(e, "=", 2)
        fmt.Println(kvpair[0])
    }

    err = os.Unsetenv("FRUIT")
    if err != nil {
        fmt.Println("Error unsetting environment variable:", err)
        return
    }

    fmt.Println("Name of Fruit:", os.Getenv("FRUIT")) // Name of Fruit:
}
```

### Kullanım ve Açıklamalar:

- `os.Getenv(key)`: Bir environment variable'ın değerini alır.
- `os.Setenv(key, value)`: Environment variable ayarlar.
- `os.Unsetenv(key)`: Environment variable siler.
- `os.Environ()`: Tüm environment variable'ları döner.

---
