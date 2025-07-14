# Go'da os ile File nedir?

## 1. Temel Paketler

| Paket   | Amaç                                     | Önemli Fonksiyonlar                   |
| ------- | ---------------------------------------- | ------------------------------------- |
| `os`    | Dosya ve dizin işlemleri                 | `Create`, `Open`, `ReadFile`, `Close` |
| `bufio` | Girdi/çıktı için tampon (buffer) desteği | `NewScanner`, `Scanner.Text`          |

---

## 2. Dosya Oluşturma ve Yazma (`os` paketi)

1. **Dosya Oluşturma**

   ```go
   file, err := os.Create("README.md")
   ```

   _Dosya yoksa oluşturur; varsa içeriğini sıfırlar._

2. **Veri Yazma**

   - **Ham bayt dizisi:**
     ```go
     file.Write([]byte("Merhaba"))
     ```
   - **Metin:**
     ```go
     file.WriteString("Merhaba Dünya")
     ```

3. **Kaynak Serbest Bırakma**
   ```go
   defer file.Close()
   ```
   _`defer`, fonksiyon bittiğinde `Close`'u otomatik çağırır._

---

## 3. Dosya Okuma Seçenekleri

### 3.1 Tüm Dosyayı Tek Seferde Okuma

```go
data, _ := os.ReadFile("README.md")
```

- **Avantaj:** Kısa dosyalar için hızlı ve basit.
- **Dezavantaj:** Büyük dosyalarda bellek tüketimi artar.

### 3.2 Satır Satır Okuma (`bufio.Scanner`)

```go
f, _ := os.Open("README.md")
scanner := bufio.NewScanner(f)
for scanner.Scan() {
    line := scanner.Text()
}
```

- **Avantaj:** Bellek dostu; satır satır işleme imkânı verir.
- **Kullanım Alanı:** Log dosyaları, büyük metin dosyaları.

---

## 4. Hata Yönetimi

Tüm dosya işlemlerinde **hata kontrolü** hayati öneme sahiptir:

```go
if err != nil {
    // uygun mesaj veya geri dönüş
}
```

> Unutmayın: Dosyaya erişim yetkisi, yol hataları veya disk sorunları sıkça karşılaşılan hata kaynaklarıdır.

---

## 5. Neden `bufio`?

- **Tamponlama (Buffering):** Veriyi bellekte küçük bloklar hâlinde tutarak I/O çağrılarını azaltır.
- **Performans:** Özellikle mekanik disklerde ve ağ üzerinden okuma/yazmada hız kazandırır.

---

## 6. Özet

| İşlem             | Fonksiyon              | Ne Zaman Kullanılır?                              |
| ----------------- | ---------------------- | ------------------------------------------------- |
| Dosya oluşturma   | `os.Create`            | Yeni dosya başlatmak veya eski dosyayı sıfırlamak |
| Dosyaya yazma     | `Write`, `WriteString` | Metin veya bayt verisi eklemek                    |
| Tüm dosyayı okuma | `os.ReadFile`          | Küçük/orta boyutlu dosyalar                       |
| Satır satır okuma | `bufio.Scanner`        | Büyük dosyalar, satır analizi gereken durumlar    |
