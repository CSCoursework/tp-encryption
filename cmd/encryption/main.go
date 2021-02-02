package main

import (
	"fmt"

	"github.com/CSCoursework/tp-encryption/pkg/caesar"
)

func main() {
	p := "Hello world"
	k := 5
	c := caesar.Encrypt(p, k)
	d := caesar.Decrypt(c, k)

	fmt.Println("plaintext", p)
	fmt.Println("key", k)
	fmt.Println("ciphertext", c)
	fmt.Println("decryption", d)
}
