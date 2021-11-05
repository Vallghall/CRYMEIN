package rsa

import "fmt"

func RSACipher(txt string) ([]int64, []int64) {
	primes := NewPrimesFromInput()
	fmt.Printf("P = %v; \tQ = %v;\nn = %v; \t\uF06A = %v\n",
		primes.P(),
		primes.Q(),
		primes.N(),
		primes.Phi())
	kp := primes.GenerateKeyPair()
	fmt.Printf("Закрытый ключ: %v\n", kp.PrivateKey.d)
	fmt.Printf("Открытый ключ: %v\n", kp.PublicKey.e)

	encrypted := Encrypt(txt, kp.PublicKey)
	decrypted := Decrypt(encrypted, kp.PrivateKey)
	return encrypted, decrypted
}

func Encrypt(txt string, pk *PublicKey) []int64 {
	return pk.Encrypt(txt)
}

func Decrypt(enc []int64, pk *PrivateKey) []int64 {
	return pk.Decrypt(enc)
}
