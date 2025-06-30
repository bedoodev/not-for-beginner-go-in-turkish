# Go `math/rand` Nedir

Rastgele sayılar oyunlar, test verisi veya simülasyonlar için vazgeçilmezdir. Go’nun standart kütüphanesindeki `math/rand` paketi, **kripto‑güvenli olmayan** (ama hızlı) sözde rastgele sayı üretimi yapar.

---

## 1. Temel Rastgele Tamsayı

```go
n := rand.Intn(10) // 0‑9 dâhil (üst sınır hariç)
fmt.Println("0‑9 arası:", n)

// [80, 120] arası 5 sayı
for i := 0; i < 5; i++ {
    fmt.Printf("80‑120 arası: %d\n", rand.Intn(41)+80)
}
```

> **İpucu:** Üst sınır her zaman **hariç** tutulur: `rand.Intn(k)` → `[0,k)`.

---

## 2. Rastgele Float ve Bool

```go
fmt.Printf("0 ile 1 arasında: %.4f\n", rand.Float64())

if rand.Intn(2) == 0 {
    fmt.Println("Yazı")
} else {
    fmt.Println("Tura")
}
```

---

## 3. Seeding

Aynı tohum (seed) ile her çalıştırmada aynı diziyi elde edersiniz. Testler için idealdir.

```go
src := rand.NewSource(99)
r := rand.New(src)
fmt.Println("Tekrarlanabilir:", r.Intn(100)) // her zaman 53
```

Gerçek uygulamalarda farklı sonuçlar isteyenler için:

```go
rand.Seed(time.Now().UnixNano()) // genellikle main() içinde bir kere çağrılır
```

---

## 4. Karıştırma (Shuffle) & Rastgele Seçim

```go
fruits := []string{"elma", "armut", "muz", "kiraz"}
rand.Shuffle(len(fruits), func(i, j int) { fruits[i], fruits[j] = fruits[j], fruits[i] })
fmt.Println("Karıştırılmış:", fruits)

pick := fruits[rand.Intn(len(fruits))]
fmt.Println("Seçilen meyve:", pick)
```

---

## 5. Zar Oyunu (Mini Örnek)

```go
dice1, dice2 := rand.Intn(6)+1, rand.Intn(6)+1
fmt.Printf("Zarlar: %d + %d = %d\n", dice1, dice2, dice1+dice2)
```

> **Not:** Gerçek projelerde kullanıcı etkileşimi (`fmt.Scan`) ile basit menüler oluşturabilirsiniz.

---

## 6. Pseudo vs Gerçek Rastgele

- `math/rand` **hızlı** ama kripto‑güvenli değildir.
- Şifreleme, token, anahtar üretimi gibi alanlarda `crypto/rand` kullanın:

```go
n, _ := rand.Int(rand.Reader, big.NewInt(100)) // 0‑99 arası kripto‑güvenli sayı
```

---

## 7. Özet

- `rand.Intn`, `Float32/64` → Temel sayı üretimi
- `rand.Seed` veya `rand.NewSource` → Tekrarlanabilirlik
- `Shuffle`, `Perm` → Koleksiyon karıştırma
- **Kriptografi** için `crypto/rand`

Bu rehber, Go ile rastgele sayı üretimine hızlı bir başlangıç sunar. Başka örneklere ihtiyaç duyarsanız lütfen belirtin!
