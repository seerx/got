package main

import (
	"fmt"

	"github.com/seerx/got"
	got2 "github.com/seerx/got/pkg/got"
)

func main() {
	nonce := got2.UUID()[:24] //  "37b8e8a308c3540409245f6ddjdfhkhakhfasjhdfahdkjsahhak"
	key := "A99900Key-32Charddters12345jdiks"

	fmt.Println("key:", len(key), "n:", len(nonce))

	plainText := "172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88172.10.99.88"
	cipherText, _ := got.AESEncrypt(plainText, key, nonce)
	newPlain, _ := got.AESDecrypt(cipherText, key, nonce)

	fmt.Println("plain:", plainText)
	fmt.Println("cipher:", cipherText)
	fmt.Println("plain:", newPlain)
}
