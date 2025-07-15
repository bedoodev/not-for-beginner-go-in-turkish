# Go'da Zap Logger ile Loglama ve Özelleştirme

Go ekosisteminde popüler, hızlı ve esnek bir logging paketi olan [Uber Zap](https://github.com/uber-go/zap), yüksek performanslı ve yapılandırılabilir loglama sağlar. Zap, JSON tabanlı structured logging destekler ve log formatı üzerinde detaylı kontrol sunar.

---

## 1. Temel Kullanım

```go
package main

import (
    "log"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

func main() {
    logger, err := zap.NewProduction()
    if err != nil {
        log.Println("Error in initializing Zap logger.")
    }
    defer logger.Sync()

    cfg := zap.NewProductionConfig()
    cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05") // Tarih ve saat formatı özelleştirildi
    cfg.DisableStacktrace = true // Stacktrace çıktısı tamamen kapatıldı
    logger, err = cfg.Build()
    if err != nil {
        log.Println("Error in building Zap logger.")
    }

    logger.Info("User logged in.", zap.String("username", "Bedirhan"), zap.String("method", "GET"))
    logger.Error("Error in processing request.", zap.String("error", "Database connection failed."))
    logger.Warn("Slow query detected.", zap.Duration("duration", 1_000_000_000))
}
```

---

## 2. Açıklamalar

- `zap.NewProductionConfig()`: Production ortamı için varsayılan ayarları içerir. İstersen `zap.NewDevelopmentConfig()` ile development ayarları da kullanabilirsin.
- `cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(...)`: Zaman damgasının formatı Go'nun `time.Format` stilinde özelleştirilebilir. Örnekte `YYYY-MM-DD HH:MM:SS` formatı kullanıldı.
- `cfg.DisableStacktrace = true`: Hata seviyesinden itibaren otomatik olarak eklenen stacktrace loglarını tamamen kapatır.
- `zap.String`, `zap.Duration` gibi fonksiyonlar ile log mesajına anahtar-değer (field) ekleyebilirsin. Böylece loglar yapısal ve makine tarafından kolayca okunabilir olur.
- Log seviyesi (info, error, warn, vb.) metodlarla belirlenir.

---

## 3. Neden Özelleştirme Yapılır?

- **Zaman formatı**: Logları başka sistemlerle entegre etmek veya insan tarafından okunmasını kolaylaştırmak için.
- **Stacktrace**: Sadece gerçekten ihtiyaç olduğunda eklenmesi istenebilir, fazla ayrıntıdan kaçınmak için kapatılabilir.
- **Structured logging**: Key-value formatı sayesinde loglar sistemler arası taşınabilir ve kolayca sorgulanabilir.

---

## 4. Farklı Özelleştirmeler

- Stacktrace anahtarının adını değiştirmek için:
  ```go
  cfg.EncoderConfig.StacktraceKey = "trace"
  ```
- JSON yerine konsola renkli (human-readable) log için:
  ```go
  cfg.Encoding = "console"
  ```
- Sadece belirli seviyede stacktrace eklemek için:
  ```go
  cfg.Level = zap.NewAtomicLevelAt(zapcore.WarnLevel) // Örnek
  cfg.DisableStacktrace = false
  cfg.EncoderConfig.StacktraceKey = "stack"
  cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
  ```
