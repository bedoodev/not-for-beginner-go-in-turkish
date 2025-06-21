> **Not** ▪ `Print*` ailesi doğrudan **STDOUT**’a yazar.\
> **Not** ▪ `F*` ailesi belirtilen `io.Writer`’a yazar.\
> **Not** ▪ `S*` ailesi ekran yerine **string** döndürür.\
> **Not** ▪ `Append*` (Go 1.19+) fonksiyonları var olan `[]byte` dilimine ekleme yapar; bellek tahsis etmeden çalışır.

---

## 1 | Yazdırma Fonksiyonları (`Print`)

| Fonksiyon | Görev              | Kısa Örnek              | Çıktı  |
| --------- | ------------------ | ----------------------- | ------ |
| `Print`   | Boşluksuz yazdırır | `fmt.Print("Go")`       | `Go`   |
| `Println` | Boşluk + `\n`      | `fmt.Println("Go")`     | `Go\n` |
| `Printf`  | Formatlı yazdırır  | `fmt.Printf("%02d", 7)` | `07`   |

---

## 2 | Writer’a Yazdırma (`F*`)

| Fonksiyon  | Görev         | Örnek (→ `os.Stdout`)         | Çıktı  |
| ---------- | ------------- | ----------------------------- | ------ |
| `Fprint`   | Boşluksuz     | `fmt.Fprint(w, "Hi")`         | `Hi`   |
| `Fprintln` | Boşluk + `\n` | `fmt.Fprintln(w, "Hi")`       | `Hi\n` |
| `Fprintf`  | Formatlı      | `fmt.Fprintf(w, "0x%X", 255)` | `0xFF` |

---

## 3 | String Döndüren (`S*`)

| Fonksiyon  | Görev         | Örnek                       | Dönen Değer |
| ---------- | ------------- | --------------------------- | ----------- |
| `Sprint`   | Birleştirir   | `fmt.Sprint("A",1)`         | `"A1"`      |
| `Sprintln` | Boşluk + `\n` | `fmt.Sprintln("A",1)`       | `"A 1\n"`   |
| `Sprintf`  | Formatlı      | `fmt.Sprintf("%.1f", 3.14)` | `"3.1"`     |

---

## 4 | `[]byte`’e Ekleyen (`Append*`)

| Fonksiyon  | Görev          | Örnek                        | Sonuç (`[]byte`) |
| ---------- | -------------- | ---------------------------- | ---------------- |
| `Append`   | Boşluksuz ekle | `fmt.Append(nil, 1)`         | `[]byte("1")`    |
| `Appendln` | Boşluk + `\n`  | `fmt.Appendln(nil, 1)`       | `[]byte("1\n")`  |
| `Appendf`  | Formatlı       | `fmt.Appendf(nil, "%02d",5)` | `[]byte("05")`   |

---

## 5 | Hata Üretme

| Fonksiyon | Görev                  | Örnek                      | Çıktı              |
| --------- | ---------------------- | -------------------------- | ------------------ |
| `Errorf`  | Formatlı `error` döner | `fmt.Errorf("kod %d",404)` | `error("kod 404")` |

---

## 6 | Girdi Okuma (`Scan` Ailesi)

| Fonksiyon | Kaynak      | Formatlı mı? | Tipik Kullanım             | Sonuç                   |
| --------- | ----------- | ------------ | -------------------------- | ----------------------- |
| `Scan`    | STDIN       | ❌           | `fmt.Scan(&a,&b)`          | Girilen değerler atanır |
| `Scanln`  | STDIN       | ❌           | `fmt.Scanln(&s)`           | Satır sonuna kadar okur |
| `Scanf`   | STDIN       | ✅           | `fmt.Scanf("%d-%d",&m,&d)` | Format eşleşirse atanır |
| `Fscan*`  | `io.Reader` | Çeşitli      | `fmt.Fscan(r,&x)`          | Okunan değer atanır     |
| `Sscan*`  | `string`    | Çeşitli      | `fmt.Sscan("3 4",&x,&y)`   | `x=3,y=4`               |

---

## 7 | Sık Kullanılan Format Verb’leri

```
%v  varsayılan gösterim      %+v alan adlarıyla        %#v Go sözdizimi
%T  tür bilgisi              %d  ondalık               %b  ikili
%x  hex (küçük)              %X  hex (büyük)           %f  ondalık float
%s  byte slice → string      %q  çift tırnaklı         %p  pointer adresi
```

---

## 🚀 Derleme / Çalıştırma

```bash
# standart çalışma
$ go run main.go
```

> **İpucu:** Kod bloklarında `\n` gerçek satır sonunu temsil eder; terminalde yeni satır olarak görünür.
