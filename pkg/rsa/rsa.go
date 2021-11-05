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
	encrypted, decrypted := rsa.Encrypt(txt, primes)
	fmt.Printf("Зашифрованные символы: %v\n", encrypted)
	fmt.Printf("Расшифрованные символы: %v\n", decrypted)
	fmt.Printf("Расшифрованная строка: %v\n", string(alphabet.ToRussianRunes(decrypted)))
}
