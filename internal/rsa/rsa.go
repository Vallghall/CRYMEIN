package rsa

import (
	"fmt"
)

func Encrypt(txt string, primes *Primes) ([]int64, []int64) {
	kp := primes.GenerateKeyPair()
	fmt.Printf("P = %v; \tQ = %v;\t\uF06A = %v;\n", primes.P(), primes.Q(), primes.Phi())
	fmt.Printf("Закрытый ключ: %v\n", kp.PrivateKey.d)
	fmt.Printf("Открытый ключ: %v\n", kp.PublicKey.e)

	encrypted := kp.Encrypt(txt)
	decrypted := kp.Decrypt(encrypted)

	return encrypted, decrypted
}
