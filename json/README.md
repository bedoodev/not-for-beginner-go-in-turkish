# Go ile JSON İşlemleri

Bu proje, Go dilinde JSON veri formatı ile çalışma örnekleri sunar. Kodda struct'lardan JSON'a çevirme (`marshal`), JSON'dan struct'a çevirme (`unmarshal`), slice (dizi) ve map ile JSON işlemleri, ayrıca bilinmeyen JSON yapılarını dinamik olarak yönetme gibi pratikler örneklenmektedir.

---

## JSON Nedir? Ne İçin Kullanılır?

**JSON** (JavaScript Object Notation), insan tarafından kolayca okunabilen ve makineler tarafından kolayca oluşturulup ayrıştırılabilen bir veri değişim formatıdır. Platform ve dil bağımsızdır. Modern yazılım dünyasında;

- Web API'larda veri alışverişi,
- Yapılandırma dosyalarında (ör. `config.json`),
- Mikroservisler arasında iletişimde,
- Mobil ve web uygulamalarında sunucu/istemci veri transferinde,

gibi hemen her alanda kullanılır.

JSON'un avantajları:

- Okunabilir ve sade bir yapıya sahip olması
- Birçok dil tarafından (Go, Python, Java, JS, vb.) desteklenmesi
- Nested (iç içe) veri yapılarını kolayca ifade edebilmesi

---

## Özellikler

- Struct'tan JSON'a veri dönüştürme (`json.Marshal`)
- JSON'dan struct'a veri dönüştürme (`json.Unmarshal`)
- `omitempty` etiketi kullanımı
- İç içe struct (nested struct) örneği
- Slice (dizi) ile JSON işlemleri
- Map ile bilinmeyen JSON yapısı yönetimi

---

## Farklı Örneklerle JSON İşlemleri

### 1. Kullanıcı Bilgisini Struct'tan JSON'a Dönüştürme

```go
user := User{Username: "gopher42", Score: 1200}
data, _ := json.Marshal(user)
fmt.Println(string(data))
// Çıktı: {"username":"gopher42","score":1200}
```

---

### 2. İç içe Struct ile JSON

```go
profile := Profile{
    Name:    "Ada Lovelace",
    Contact: Contact{Email: "ada@example.com", Phone: "+90 555 111 22 33"},
}
jsonProfile, _ := json.Marshal(profile)
fmt.Println(string(jsonProfile))
// Çıktı: {"name":"Ada Lovelace","contact":{"email":"ada@example.com","phone":"+90 555 111 22 33"}}
```

---

### 3. JSON'dan Struct'a Dönüştürme

```go
data := `{"city": "Bursa", "population": 3500000}`
var city City
json.Unmarshal([]byte(data), &city)
fmt.Println(city)
// Çıktı: {Bursa 3500000}
```

---

### 4. Slice (Dizi) ile JSON

```go
colors := []string{"red", "green", "blue"}
jsonColors, _ := json.Marshal(colors)
fmt.Println(string(jsonColors))
// Çıktı: ["red","green","blue"]
```

---

### 5. Bilinmeyen JSON Yapısını Map ile Yönetmek

```go
jsonData := `{"brand": "Tesla", "year": 2022, "electric": true}`
var car map[string]interface{}
json.Unmarshal([]byte(jsonData), &car)
fmt.Println(car)
// Çıktı: map[brand:Tesla year:2022 electric:true]
```
