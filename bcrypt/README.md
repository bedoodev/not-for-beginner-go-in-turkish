# Go ile SHA-256/SHA-512 Hashleme
## 1. Hash Fonksiyonları Nedir?

Hash fonksiyonları, girdi olarak aldıkları bir veriden sabit uzunlukta, genellikle rastgele dağılan ve tek yönlü bir hash üretirler.\
SHA-256 ve SHA-512 en çok tercih edilen modern hash fonksiyonlarındandır.

**Temel llikler:**

- Aynı girdi her zaman aynı hash'i üretir.
- Farklı girdiler neredeyse her zaman farklı hash üretir.
- Hash'ten orijinal veriye ulaşmak mümkün değildir (tek yönlü).
- Küçük bir değişiklik, tamamen farklı bir hash sonucu doğurur.

### **Örnek: SHA-256 ve SHA-512 Kullanımı**

```go
import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main() {
	password := "testPass=1"
	hash256 := sha256.Sum256([]byte(password))
	hash512 := sha512.Sum512([]byte(password))

	fmt.Printf("SHA-256: %x\n", hash256)
	fmt.Printf("SHA-512: %x\n", hash512)
}
```

Çıktı olarak, her iki algoritma da parolanın kendine özgü bir hashini üretir.

---

## 2. Parola Hashleme Neden Gerekli?

Parolaların hash'lenmeden saklanması, ciddi güvenlik açıklarına yol açar.\
Bir veri tabanındaki parolalar çalınırsa, hash'lenmiş ve salt'lanmış parolalar saldırgan için neredeyse işe yaramazdır.

**Hash'li Parola Saklama Avantajları:**

- Parolanın aslı sistemde tutulmaz, sadece hash'i tutulur.
- Aynı parolayı kullanan farklı kullanıcıların hash'leri birbirinden farklı olur (salt kullanımı sayesinde).

---

## 3. Salt Nedir ve Neden Kullanılır?

**Salt**, hash fonksiyonuna eklenen rastgele bir veridir.\
Amaç, aynı parolayı kullanan iki kişinin hash'inin de farklı olmasını sağlamaktır.\
Salt sayesinde "Rainbow Table" gibi önceden hazırlanmış saldırı teknikleri etkisiz hale gelir.

### **Go'da Salt Üretimi:**

```go
import (
	"crypto/rand"
	"io"
)

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16) // 16 baytlık rastgele bir salt
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}
```

---

## 4. Parola Hashleme ve Doğrulama Akışı

Aşağıda örnek Go kodu ile bir parola nasıl hash'lenir ve doğrulama sırasında nasıl kontrol edilir gösterilmiştir.

### **Adımlar:**

1. Kullanıcı kayıt olurken: Parola + salt alınır → Hash'lenir → Salt ve hash veritabanında saklanır.
2. Kullanıcı giriş yaparken: Girilen parola, kayıtlı salt ile yeniden hash'lenir ve veritabanındaki hash ile karşılaştırılır.

### **Go ile Parola Hashleme Fonksiyonu:**

```go
import (
	"crypto/sha256"
	"encoding/base64"
)

func hashPassword(p string, salt []byte) string {
	saltedPassword := append(salt, []byte(p)...)
	hashedPassword := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hashedPassword[:])
}
```

---

## 5. Uygulama: Parola Kayıt ve Doğrulama Senaryosu

Aşağıdaki örnekte, kullanıcıya ait bir parola hash'lenip salt ile saklanmakta ve ardından yanlış parola ile giriş denenmektedir.

```go
package main

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {

	var password string = "testPass=1"

	hash256 := sha256.Sum256([]byte(password))
	hash512 := sha512.Sum512([]byte(password))

	fmt.Println("Password:", password)
	fmt.Printf("Hashed Password 256 as String: %x\n", hash256)
	fmt.Printf("Hashed Password 512 as String: %x\n", hash512)

	fmt.Print("\n***************** Login App ******************\n\n")

	salt, err := generateSalt()
	if err != nil {
		fmt.Println("Error generating salt:", err)
		return
	}
	p := "test2000!"
	signUpHash := hashPassword(p, salt)
	saltStr := base64.StdEncoding.EncodeToString(salt)
	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Unable to decode salt:", err)
		return
	}
	loginHash := hashPassword("passwordForFailScenerio", decodedSalt)

	fmt.Println("Salt String:", saltStr)
	fmt.Println("Sign Up Hashed Password:", signUpHash)
	fmt.Println("Decoded Salt", decodedSalt)
	fmt.Println("Login Hashed Password:", loginHash)

	if signUpHash == loginHash {
		fmt.Println("Password is correct. You are logged in.")
	} else {
		fmt.Println("Login Failed. Plase check user credentials.")
	}
}

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func hashPassword(p string, salt []byte) string {
	saltedPassword := append(salt, []byte(p)...)
	hashedPassword := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hashedPassword[:])
}
```

---

## 6. Açıklamalar ve Güvenlik İpuçları

- **Hash algoritması olarak SHA-256 veya daha üstü kullanılması önerilir.**
- **Salt**, her kullanıcı ve her parola için **benzersiz** ve **rastgele** üretilmelidir.
- Parola hash'leri, doğrudan karşılaştırılır, asla çözülmeye çalışılmaz.
- Daha ileri güvenlik için bcrypt, scrypt veya Argon2 gibi yavaş hash algoritmaları da kullanılabilir.

---

## 7. Kısaca

- Parolalar asla düz metin olarak saklanmaz.
- Salt ile birlikte hash'lenip, hash ve salt birlikte saklanır.
- Kullanıcı girişinde girilen parola ve saklanan salt ile tekrar hash hesaplanır ve doğrulanır.

---

**Daha ileri bilgi veya kod örneği için: **[**Go resmi dökümantasyonuna göz atın**](https://pkg.go.dev/crypto)
