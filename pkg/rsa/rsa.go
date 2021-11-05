package rsa

import (
	"fmt"
	"github.com/Vallghall/CRYMEIN/internal/rsa"
)

func RSA() {
	fmt.Println("Введите текст для шифрования:")
	var txt string
	fmt.Scan(&txt)
	primes := rsa.NewPrimesFromInput()
	rsa.Encrypt(txt, primes)
}
