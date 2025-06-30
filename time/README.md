# Go `time` Nedir

Go’nun `time` paketi tarih‑saat işlemlerinin çoğunu kolayca yapmanızı sağlar.

## 1. Anlık Zaman (Şu An)

`time.Now()` sistem saatini döndürür.

```go
now := time.Now()
fmt.Printf("Şu an (yerel): %s\n", now.Format("02-01-2006 15:04:05"))
```

---

## 2. Belirli Zaman Oluşturma

`time.Date` ile istenen tarih‑saat değeri üretilebilir.

```go
istanbul, _ := time.LoadLocation("Europe/Istanbul")
d := time.Date(2024, time.December, 31, 23, 59, 30, 0, istanbul)
fmt.Println("Yılbaşı:", d)
```

---

## 3. Zaman Ayrıştırma (Parse)

Bir metni zamana dönüştürmek için **şablon (layout)** ve veri verilir.

```go
layout := "02/01/2006 15:04"
parsed, err := time.Parse(layout, "18/03/2023 08:45")
if err != nil { log.Fatal(err) }
fmt.Println("Parse sonucu:", parsed)
```

> **Not:** Şablon, 2 Ocak 2006 saat 15:04:05 MST örneğine göre yazılır.

---

## 4. Zaman Biçimlendirme (Format)

Bir `time.Time` değerini okunabilir metne çevirmek için `Format` kullanılır.

```go
fmt.Println(parsed.Format("Monday, 02 Jan 2006 15:04"))
```

---

## 5. Zaman Hesaplamaları

Süre (duration) ekleme‑çıkarma işlemleri:

```go
twoHoursLater := parsed.Add(2 * time.Hour)
fmt.Println("+2 saat:", twoHoursLater)

fark := twoHoursLater.Sub(parsed)
fmt.Println("Süre farkı:", fark) // 2h0m0s
```

---

## 6. Yuvarlama (Round)

Dakika veya saat gibi en yakın birime yuvarlamak:

```go
rounded := now.Round(time.Minute)
fmt.Println("Dakikaya yuvarla:", rounded)
```

---

## 7. Saat Dilimine Dönüştürme

```go
newYork, _ := time.LoadLocation("America/New_York")
fmt.Println("New York:", d.In(newYork))
```

---

## 8. Kesme (Truncate)

Alt birimlerini sıfırlayarak kırpma:

```go
dayStart := now.Truncate(24 * time.Hour)
fmt.Println("Günün başlangıcı:", dayStart)
```

---

## 9. Zaman Karşılaştırma

```go
gelecek := now.Add(48 * time.Hour)
fmt.Println("Şu an gelecekten önce mi?", now.Before(gelecek)) // true
```

---

## 10. Unix Zamanı (Epoch)

```go
epoch := now.Unix()
fmt.Println("Epoch (saniye):", epoch)
fmt.Println("Tekrar zaman:", time.Unix(epoch, 0))
```

---

## Özet

- `time.Date` → Elle tarih oluşturma
- `time.Parse` / `Format` → Metin ↔︎ zaman
- `Add`, `Sub` → Hesaplama
- `Round`, `Truncate` → Yuvarlama / Kesme
- `LoadLocation` + `In` → Saat dilimi
