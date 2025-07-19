## Goroutine nedir?

- Go’da concurrency’nin temeli goroutine’dir.
- Her goroutine, Go runtime tarafından yönetilen “hafif” bir iş parçacığıdır. Başlangıçta 2 KB stack ile başlar, gerektiğinde büyür.
- Milyonlarca goroutine başlatabilirsin çünkü işletim sistemi thread’i kadar ağır değiller.
- Goroutine, bir programın çalışma zamanında paralel olarak çalışan bir işlemdir.
  Goroutine ile çağrılan bir fonksiyon, ana program akışından bağımsız olarak arka planda çalışmaya başlar ve işini bitirdiğinde kendi kendine sona erer. Program akışını durdurmaz veya engellemez.

- Goroutine’ler bittiğinde Go’nun garbage collector’ı otomatik olarak kaynakları temizler.
- Concurrency sırasında memory management de otomatik yürütülür.

- Goroutine, `M:N Scheduling Model` kullanır. `M:N` scheduling model M tane user-level thread’in (örneğin goroutine) N tane OS-level thread (işletim sistemi thread’i) üzerinde çalıştırılmasıdır. Yani M tane görev, N tane işletim sistemi thread’i üzerinde dinamik olarak dağıtılır ve çalıştırılır.
- Go runtime, hangi goroutine’in hangi thread’de çalışacağını, hangi thread’in hangi CPU çekirdeğinde olacağını en verimli şekilde ayarlar.

### Avantajları

1. **Yüksek verim:** Çok az thread ile çok fazla iş yapılır.
2. **Verimli kaynak kullanımı:** OS thread’leri ağırdır, ama goroutine çok hafiftir.
3. **Otomatik denge:** Go runtime en uygun thread ve çekirdek kullanımını sağlar.
4. **Kolay geliştirme:** Programcı olarak thread yönetimini düşünmezsin, her şey arka planda yönetilir.

### Sık yapılan hatalar.

#### 1. Avoiding Goroutine Leaks (Goroutine Sızıntılarından Kaçınmak)

Bir goroutine, içinde sonsuza kadar beklerse (örneğin bir döngüde hiç çıkmazsa veya bir koşulu hiç sağlamazsa) RAM’de kalır, program şişer.

```go
func neverEnds() {
    for {
        // Sonsuz döngü, bu goroutine asla bitmez.
        time.Sleep(time.Hour)
    }
}

func main() {
    go neverEnds()
    time.Sleep(2 * time.Second) // Ana program kısa bekleyip biter
}

// Doğru Kullanım:

func endsAfterTenSeconds() {
    start := time.Now()
    for {
        if time.Since(start) > 10*time.Second {
            return // Goroutine düzgün şekilde biter
        }
        time.Sleep(time.Second)
    }
}
```

#### 2. Limiting Goroutine Creation (Goroutine Oluşumunu Sınırlamak)

Bir döngü ile kontrolsüz şekilde binlerce/milyonlarca goroutine başlatırsan sistem kilitlenebilir.

```go
func main() {
    for i := 0; i < 1000000; i++ {
        go func() {
            time.Sleep(time.Second)
        }()
    }
    time.Sleep(2 * time.Second)
}

// Doğru Kullanım:

func main() {
    for i := 0; i < 10; i++ { // Maksimum 10 goroutine başlatıyoruz
        go func(i int) {
            time.Sleep(time.Second)
            fmt.Println("Goroutine:", i)
        }(i)
    }
    time.Sleep(2 * time.Second)

}
```

#### 3. Synchronization (Eşzamanlama)

Birden çok goroutine aynı anda bir veriye erişirse veri bozulabilir (race condition).

```go
var count int

func increase() {
    for i := 0; i < 1000; i++ {
        count++
    }
}

func main() {
    go increase()
    go increase()
    time.Sleep(1 * time.Second)
    fmt.Println(count)
}

// Beklenen sonuç: 2000
// Gerçek sonuç: Her çalıştırmada farklı olabilir! (Çünkü iki goroutine aynı anda count değişkenini değiştiriyor, arada bazı değerler “kayboluyor.”)
```
