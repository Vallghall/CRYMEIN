package rsa

import (
	"fmt"
	"github.com/Vallghall/CRYMEIN/internal/alphabet"
)

func Encrypt(txt string, primes *primes) {
	kp := primes.GenerateKeyPair()
	fmt.Printf("P = %v; \tQ = %v;\t\uF06A = %v;\n", primes.P(), primes.Q(), primes.Phi())
	fmt.Printf("Закрытый ключ: %v\n", kp.PrivateKey.d)
	fmt.Printf("Открытый ключ: %v\n", kp.PublicKey.e)

	encrypted := kp.PublicKey.Encrypt(txt)
	fmt.Printf("Зашифрованные символы: %v\n", encrypted)

	decrypted := kp.Decrypt(encrypted)
	fmt.Printf("Расшифрованные символы: %v\n", decrypted)
	fmt.Printf("Расшифрованная строка: %v\n", string(alphabet.ToRussianRunes(decrypted)))
}
