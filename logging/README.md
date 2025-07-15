# Go'da Logging (Loglama) ve log Paketi Kullanımı

Go'da loglama (logging), uygulamanın çalışma sırasında olan biteni takip etmek, hata ve uyarıları kaydetmek, uygulama davranışını izlemek için temel bir pratiktir. Go'nun standart `log` paketi; mesajları, seviyeleri ve çıktıları özelleştirilebilir şekilde yönetmek için basit ama güçlü araçlar sunar.

---

## 1. Basit Loglama

En temel haliyle log mesajı göndermek için doğrudan `log.Println` fonksiyonu kullanılır:

```go
package main

import (
    "log"
)

func main() {
    log.Println("This is a log message")
    // Örnek çıktı: 2025/07/15 13:56:28 This is a log message
}
```

Standart olarak, tarih ve saat bilgisi ile mesajınız stdout'a (terminal) yazılır.

---

## 2. Prefix ve Flag ile Log Formatını Özelleştirme

`log.SetPrefix()` fonksiyonu ile log mesajlarının başına özel bir etiket ekleyebilirsiniz.  
`log.SetFlags()` ile tarih, saat, mikro-saniye, dosya ismi gibi ekstra bilgiler kontrol edilebilir.

```go
package main

import (
    "log"
)

func main() {
    log.SetPrefix("INFO: ")
    log.Println("This is first info message.")
    log.Println("This is second info message.")
    // Çıktı: INFO: 2025/07/15 13:56:47 This is first info message.

    log.SetFlags(log.Ldate)
    log.Println("This is a log message with only date.")
    // Çıktı: INFO: 2025/07/15 This is a log message with only date.

    log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
    log.Println("This is a custom configured log message.")
    // Çıktı: INFO: 2025/07/15 14:02:49.671418 main.go:17: This is a custom configured log message.
}
```

### Sık Kullanılan Log Flag'leri

- `log.Ldate` : Sadece tarih
- `log.Ltime` : Sadece saat
- `log.Lmicroseconds`: Mikro-saniye hassasiyeti
- `log.Lshortfile`: Logu tetikleyen kodun dosya adı ve satırı
- `log.Llongfile`: Tam dosya yolu ve satırı
- Flag'ler `|` ile birlikte kullanılabilir.

---

## 3. Farklı Log Seviyeleri İçin Logger Nesneleri

Go'nun log paketi ile birden fazla farklı logger oluşturup, her biri için farklı ön ekler, flag'ler ve çıktılar tanımlayabilirsin. Böylece uygulama seviyesinde info, warning, error gibi ayrımlar yapılabilir.

```go
package main

import (
    "log"
    "os"
)

var (
    infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    warnLogger  = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
    errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)

func main() {
    infoLogger.Println("This is an info logger.")   // INFO: 2025/07/15 14:05:46 main.go:22: This is an info logger.
    warnLogger.Println("This is a warning logger.") // WARNING: 2025/07/15 14:05:46 main.go:23: This is a warning logger.
    errorLogger.Println("This is an error logger.") // ERROR: 2025/07/15 14:05:46 main.go:24: This is an error logger.
}
```

---

## 4. Dosyaya Log Yazmak

Log mesajlarını dosyaya yazmak için `os.OpenFile` ile bir dosya açıp, logger'ın çıktısını (`SetOutput`) bu dosyaya yönlendirebilirsin.

```go
package main

import (
    "log"
    "os"
)

func main() {
    file, err := os.OpenFile("logging/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("Failed to open log file: %v", err)
    }
    defer file.Close()

    debugLogger := log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
    debugLogger.Println("This is a debug message.")
}
```

Aynı şekilde, diğer logger'ların da çıktısını dosyaya yönlendirebilirsin:

```go
infoLogger.SetOutput(file)
warnLogger.SetOutput(file)
errorLogger.SetOutput(file)
```

---

## 5. Özet ve İpuçları

- Log mesajları uygulamanın çalışmasını izlemek ve hata ayıklamak için kritik öneme sahiptir.
- `log` paketi basit ve hızlı loglama için yeterlidir. Büyük ve karmaşık projelerde ise "zap", "logrus", "zerolog" gibi üçüncü parti paketler tercih edilebilir.
- Prefix ve flag'lerle log formatı kolayca özelleştirilebilir.
- Farklı seviyelerde logger nesneleri ile temiz ve okunabilir loglar oluşturabilirsin.
- Dosyaya log yazmak, özellikle prod ortamında kayıt tutmak için şarttır.

---