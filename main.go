package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Greetings and Salutations")
	decodeBase64("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9")
	decodeBase64("eyJVc2VySUQiOjEsIkVtYWlsIjoiY29uc3RlbGxhdGlvbkBnbWFpbC5mciIsImV4cCI6MTY1MjE5NDA0OX0")
	split := strings.Split("aaa.bonjour.zzz", ".")

	fmt.Println(split)
	fmt.Println(split[1])
}

// Base 64 decoding test
func decodeBase64(s string) {
	fmt.Println("\n::: String to decode:", s)
	// tk := &models.Token{}
	
	// tokenString := r.Header["Bearer"][0]
	// claims := jwt.MapClaims{}
	// fmt.Println(":::: Decode ::::", jwt.ParseWithClaims(r.Header["Bearer"][0], tk, func(token *jwt.Token)))
	
	// db64, err := base64.URLEncoding.DecodeString(s)
	db64, err := base64.RawStdEncoding.DecodeString(s)
	fmt.Println(":: Decode:", string(db64))

	if err != nil {
		fmt.Println(":: Decode Error:", err)
	}
	// fmt.Println(":::: Decode ::::", jwt.ParseWithClaims(r.Header["Bearer"][0], claims, func(t *jwt.Token) ))
}
