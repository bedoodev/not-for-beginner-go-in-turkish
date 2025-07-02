package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("He~lo, Base64 Encoding")

	// Encode to Base64
	//! Every time we run the program, output is going to be same
	encodedData := base64.StdEncoding.EncodeToString(data)
	fmt.Println("Encoded Data:", encodedData) // Encoded Data: SGV+bG8sIEJhc2U2NCBFbmNvZGluZw==

	// Decode from base64
	decodedData, _ := base64.StdEncoding.DecodeString(encodedData)
	fmt.Println("Decoded Data:", decodedData)                // Decoded Data: [72 101 126 108 111 44 32 66 97 115 101 54 52 32 69 110 99 111 100 105 110 103]
	fmt.Println("Decoded Data String:", string(decodedData)) // Decoded Data String: He~lo, Base64 Encoding

	/* ***************************************** */
	// URL Safe, avoid `/` and `+`
	urlSafeEncodedData := base64.URLEncoding.EncodeToString(data)
	fmt.Println("URL Safe Encoded Data:", urlSafeEncodedData) // URL Safe Encoded Data: SGV-bG8sIEJhc2U2NCBFbmNvZGluZw==


}
