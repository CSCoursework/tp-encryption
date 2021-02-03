package main

import (
	"fmt"

	"github.com/CSCoursework/tp-encryption/pkg/caesar"
	"github.com/CSCoursework/tp-encryption/pkg/keyedcaesar"
	"github.com/CSCoursework/tp-encryption/pkg/vigenere"
)

func main() {

	fmt.Println("CAESAR")

	p := "Hello world"
	s := 5
	c := caesar.Encrypt(p, s)
	d := caesar.Decrypt(c, s)

	fmt.Println("plaintext", p)
	fmt.Println("key", s)
	fmt.Println("ciphertext", c)
	fmt.Println("decryption", d)

	fmt.Println("\nKEYED CAESAR")

	p = "Hello world"
	k := "teg"
	s = 5
	c, err := keyedcaesar.Encrypt(p, k, s)
	if err != nil {
		panic(err)
	}
	d, err = keyedcaesar.Decrypt(c, k, s)
	if err != nil {
		panic(err)
	}

	fmt.Println("plaintext", p)
	fmt.Println("key", s, k)
	fmt.Println("ciphertext", c)
	fmt.Println("decryption", d)

	fmt.Println("\nVIGENÈRE")

	p = "Hello world"
	k = "ermahgawd"
	c = vigenere.Encrypt(p, k)
	d = vigenere.Decrypt(c, k)

	fmt.Println("plaintext", p)
	fmt.Println("key", k)
	fmt.Println("ciphertext", c)
	fmt.Println("decryption", d)
}
