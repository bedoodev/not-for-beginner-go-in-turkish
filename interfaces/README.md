## 1 | Interface Nedir?

Bir **interface**, bir veya daha fazla method imzası içeren, **davranış sözleşmesi** tanımlayan soyut bir tiptir. Bir tip, gereken methodları **otomatik** (explicit anahtar kelime olmadan) sağlayarak bu interface’i uygular.

```go
// Söylenebilir (Speaker) davranışı tanımı
type Speaker interface {
    Speak() string
}
```

---

## 2 | Interface’i Uygulayan Struct

```go
type Person struct{ Name string }
func (p Person) Speak() string { return "Merhaba, ben " + p.Name }

func greet(s Speaker) {
    fmt.Println(s.Speak())
}

func main() {
    p := Person{"Ada"}
    greet(p)  // Person, Speaker interface’ini karşılar
}
```

> **Not:** `Person` tipine herhangi bir `implements` bildirimi eklemedik; method imzası uyuştuğu için derleyici otomatik olarak eşleşmeyi yapar.

---

## 3 | Değer (Value) ve Pointer Receiver’lar

Bir struct’ın **method set’i**, value ya da pointer receiver kullanımına göre değişir.

- `(v T)` → value + pointer değişkenleri sağlar.
- `(v *T)` → **yalnızca pointer** değişkenini sağlar.

```go
type counter int
func (c *counter) Inc() { *c++ }

type Incer interface { Inc() }

var v counter = 0
// greetInc(v)  // derlenmez, çünkü value `counter` Inc() methoduna sahip değil.
ptr := &v
ptr.Inc()      // OK – pointer, method set’i sağlar
```

---

## 4 | Boş Interface – `interface{}`

`interface{}` hiçbir method içermez; dolayısıyla **her tip** onu uygular. Genellikle jenerik olmayan koleksiyon ya da JSON gibi dinamik veri taşıyıcılarında kullanılır.

```go
var any interface{}
any = 42
any = "metin"
```

> **Uyarı:** Sık kullanmak, tip güvenliğini zayıflatır.

---

## 5 | Tip Ataması (Type Assertion)

```go
var x interface{} = "Go"
str, ok := x.(string) // ok==true → str=="Go"
```

- `ok` bool değeri, atamanın başarılı olup olmadığını gösterir. Hata fırlatmadan kontrol sağlar.

---

## 6 | Tip Switch

```go
switch v := any.(type) {
case int:
    fmt.Println("int", v)
case string:
    fmt.Println("string", v)
default:
    fmt.Println("bilinmeyen tip")
}
```

---

## 7 | Gerçek Dünya Örneği – `io.Writer`

`fmt.Fprintf` gibi fonksiyonlar, **interface tabanlı** soyutlama kullanır.

```go
func main() {
    // os.Stdout da io.Writer’ı uygular
    fmt.Fprintln(os.Stdout, "Yazdırıldı!")

    // Buffer da aynı interface’i uygular
    var buf bytes.Buffer
    fmt.Fprintln(&buf, "Buffered")
    fmt.Println(buf.String())
}
```

---

## 8 | Nil Interface Tuzağı

Boş method setli interface değişkeni `nil` değeri içerse bile **tip bilgisi yoksa** `nil` değildir.

```go
var w io.Writer      // hem tip hem değer nil → gerçek nil
var buf *bytes.Buffer = nil
var w2 io.Writer = buf // tip: *bytes.Buffer, değer: nil  ➔ w2 != nil
```

Her ikisi nil gibi görünse de **yalnızca ilk değişken** gerçek nil’dir.

---

## 🧠 Özet

| Konu          | Ana fikir                                       |
| ------------- | ----------------------------------------------- |
| Tanım         | Interface method imzaları koleksiyonudur        |
| Uygulama      | Tipler, gereken methodları tanımlayarak uygular |
| Boş interface | Tüm tipleri kabul eder, tip güvencesi yok       |
| Assertion     | `v, ok := x.(T)` tip kontrolü                   |
| Type switch   | Çeşitli tiplere göre dallanma                   |
