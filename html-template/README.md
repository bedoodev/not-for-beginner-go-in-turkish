# Go `html/template`

Go’nun standart kütüphanesindeki **`html/template`** paketi, _HTML üreten_ web-uygulamalarında **XSS’e dayanıklı** bir şablon motoru sunar.  
`text/template` ile aynı sözdizimini kullanır; farkı, kontekst-duyarlı kaçış (escaping) mekanizmasıdır.

## 💡 Neden Öğrenilmeli?

| Ne zaman?                         | Neden kritik?                                                         |
| --------------------------------- | --------------------------------------------------------------------- |
| `net/http` + HTML sayfaları       | Framework’ler (Gin, Echo…) doğrudan `html/template` bekler.           |
| E-posta şablonu veya admin paneli | Otomatik XSS koruması, string birleştirmeye göre çok daha güvenlidir. |
| Statik site / rapor üretimi       | Harici template motoru eklemeden saf-Go çözüm sağlar.                 |

> Sadece JSON API/CLI geliştiriyorsanız önceliğiniz düşük olabilir, fakat hata sayfaları veya e-posta çıktısı için temel bilgiyi öğrenmek uzun vadede vakit kazandırır.

## 🚀 Hızlı Çalıştırma

```bash
go run main.go
```

Örnek oturum ⤵︎

```text
Enter your name: Ada

Menu:
1. Join
2. Get Notification
3. Get Error
4. Exit
Choose an option: 1
Welcome Ada! How are you doing?
```

---

## 📌 Temel Adımlar

| Adım                 | Kod Satırı                                     | Açıklama                          |
| -------------------- | ---------------------------------------------- | --------------------------------- |
| **İsim al**          | `reader.ReadString('\n')`                      | Kullanıcının adını alır.          |
| **Şablonları derle** | `template.Must(template.New(name).Parse(src))` | Üç metni bellekte derler.         |
| **Menü göster**      | `fmt.Println` + sonsuz döngü                   | Klasik CLI menüsü.                |
| **Veri haritası**    | `data := map[string]interface{}{...}`          | Şablona geçilecek alanlar.        |
| **Şablon yürüt**     | `tmpl.Execute(os.Stdout, data)`                | Çıktıyı doğrudan terminale yazar. |

---

## ↔️ Menü Seçenekleri

| Seçim                  | Çıktı Örneği                                                 |
| ---------------------- | ------------------------------------------------------------ |
| `1` – **Join**         | `Welcome Ada! How are you doing?`                            |
| `2` – **Notification** | `Ada, you have a new notification: Server restart scheduled` |
| `3` – **Error**        | `Oops! An error occured: Disk quota exceeded`                |
| `4` – **Exit**         | Programdan çıkar.                                            |
