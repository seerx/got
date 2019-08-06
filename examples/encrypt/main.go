package main

import (
	"fmt"

	"github.com/seerx/got"
)

func main() {
	nonce := "37b8e8a308c3540409245f6d"
	key := "A99900Key-32Charddters12345jdiks"
	plainText := "172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88"
	cipherText := got.AESEncrypt(plainText, key, nonce)
	newPlain := got.AESDecrypt(cipherText, key, nonce)

	fmt.Println("plain:", plainText)
	fmt.Println("cipher:", cipherText)
	fmt.Println("plain:", newPlain)
}
