# Go `net/url` Nedir?

Web uygulamaları geliştirirken URL çözümleme (parse), sorgu parametreleri okuma‑yazma ve güvenli kodlama/çözme işlemleri sıkça gerekir. Go’nun `net/url` paketi bu ihtiyaçları karşılar.

## 1. URL Ayrıştırma (Parse)

`url.Parse` fonksiyonu, ham (raw) URL’yi bileşenlerine ayırır.

```go
raw := "http://blog.gopher.dev:9090/articles/42?tag=go&lang=tr#comments"

u, err := url.Parse(raw)
if err != nil { log.Fatal(err) }

fmt.Println("Şema:", u.Scheme)  // http
fmt.Println("Host:", u.Host)    // blog.gopher.dev:9090
fmt.Println("Port:", u.Port())  // 9090
fmt.Println("Yol:", u.Path)     // /articles/42
fmt.Println("Sorgu:", u.RawQuery)      // tag=go&lang=tr
fmt.Println("Parçalar:", u.Fragment)    // comments
```

> **Not:** `url.Parse` hata döndürebilir; mutlaka kontrol edin.

---

## 2. Sorgu Parametreleri Okuma

```go
params := u.Query()
fmt.Println("Tüm parametreler:", params)        // map[lang:[tr] tag:[go]]
fmt.Println("Dil:", params.Get("lang"))         // tr
fmt.Println("Etiket(ler):", params["tag"])     // [go]
```

Listeden aynı anahtara birden fazla değer varsa, `params["key"]` slice olarak döner.

---

## 3. URL Oluşturma (Build) & Sorgu Ekleme

```go
base := url.URL{
    Scheme: "https",
    Host:   "api.myapp.io",
    Path:   "/v1/search",
}
q := base.Query()
q.Set("q", "gopher")
q.Set("page", "2")
q.Add("page", "3") // aynı anahtara ikinci değer eklemek için Add()

base.RawQuery = q.Encode()

fmt.Println("Tam URL:", base.String())
// https://api.myapp.io/v1/search?page=2&page=3&q=gopher
```

> **İpucu:** `Set` aynı anahtarı overwrite ederken, `Add` yeni değer ekler.

---

## 4. Güvenli Kodlama & Çözme (Escape / Unescape)

Bazen yol veya parametrelerde boşluk, Türkçe karakter gibi özel semboller kodlanmalıdır.

```go
safePath := url.PathEscape("özel klasör/ödev 1.pdf")
fmt.Println(safePath)   // %C3%B6zel%20klas%C3%B6r/%C3%B6dev%201.pdf

decoded, _ := url.PathUnescape(safePath)
fmt.Println(decoded)   // özel klasör/ödev 1.pdf
```

Benzer şekilde sorgu değeri için:

```go
v := url.QueryEscape("ada & deniz") // ada+%26+deniz
```

---

## 5. Göreli → Mutlak URL Birleştirme

```go
base, _ := url.Parse("https://docs.example.org/guide/")
rel, _ := url.Parse("intro.html#top")
full := base.ResolveReference(rel)
fmt.Println(full.String()) // https://docs.example.org/guide/intro.html#top
```

---

## 6. Kullanıcı Bilgisi (Userinfo)

```go
u := url.URL{
    Scheme: "ftp",
    Host:   "ftp.example.net",
    Path:   "/pub/data.zip",
    User:   url.UserPassword("alice", "s3cret"),
}
fmt.Println(u.String()) // ftp://alice:s3cret@ftp.example.net/pub/data.zip
```

> **Uyarı:** Parola içeren URL’leri log’lara yazmaktan kaçının.

---

## 7. Toplu Özet

| İşlem                    | Fonksiyon(lar)              |
| ------------------------ | --------------------------- |
| URL ayrıştırma           | `url.Parse`                 |
| Sorgu okuma‑yazma        | `Values` + `Query()`        |
| URL oluşturma            | `url.URL` yapısı            |
| Escape / Unescape        | `PathEscape`, `QueryEscape` |
| Göreli URL birleştirme   | `ResolveReference`          |
| Kullanıcı bilgisi ekleme | `url.User`, `UserPassword`  |
