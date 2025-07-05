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

	fmt.Println("Password:", password) // testPass=1

	fmt.Println("Hashed Password 256:", hash256)               // Hashed Password 256: [168 52 150 115 127 175 124 190 0 237 101 180 28 100 129 160 21 34 29 2 50 179 53 128 175 173 106 105 119 123 133 249]
	fmt.Printf("Hashed Password 256 as String: %x\n", hash256) // Hashed Password 256 as String: a83496737faf7cbe00ed65b41c6481a015221d0232b33580afad6a69777b85f9

	fmt.Println("Hashed Password 512:", hash512)               // Hashed Password 512: [16 73 159 63 1 224 128 245 245 171 181 14 31 59 181 87 149 133 135 114 34 4 157 40 80 147 13 5 39 73 123 91 221 252 243 3 155 65 200 250 25 121 63 176 178 32 26 18 15 26 195 230 230 110 103 145 194 0 86 66 73 111 64 135]
	fmt.Printf("Hashed Password 512 as String: %x\n", hash512) // Hashed Password 512 as String: 10499f3f01e080f5f5abb50e1f3bb5579585877222049d2850930d0527497b5bddfcf3039b41c8fa19793fb0b2201a120f1ac3e6e66e6791c2005642496f4087

	// ***************** Login App *****************
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

// Function to hash password
func hashPassword(p string, salt []byte) string {
	saltedPassword := append(salt, []byte(p)...)
	hashedPassword := sha256.Sum256(saltedPassword)

	return base64.StdEncoding.EncodeToString(hashedPassword[:])
}
