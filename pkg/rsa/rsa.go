package rsa

import (
	"fmt"
	"github.com/Vallghall/CRYMEIN/internal/alphabet"
	"github.com/Vallghall/CRYMEIN/internal/rsa"
)

func RSA() {
	fmt.Println("Введите текст для шифрования:")
	var txt string
	fmt.Scan(&txt)

	primes := rsa.NewPrimesFromInput()
	fmt.Printf("P = %v; \tQ = %v;\nn = %v; \t\uF06A = %v\n",
		primes.P(),
		primes.Q(),
		primes.N(),
		primes.Phi())
	kp := primes.GenerateKeyPair()
	fmt.Printf("Закрытый ключ: %v\n", kp.D())
	fmt.Printf("Открытый ключ: %v\n", kp.E())

	encrypted := kp.Encrypt(alphabet.ToRussianIndexes(txt))
	decrypted := kp.Decrypt(encrypted)
	fmt.Printf("Зашифрованные символы: %v\n", encrypted)
	fmt.Printf("Расшифрованные символы: %v\n", decrypted)
	fmt.Printf("Расшифрованная строка: %v\n", string(alphabet.ToRussianRunes(decrypted)))
}
