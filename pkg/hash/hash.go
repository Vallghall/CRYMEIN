package hash

import (
	"fmt"
	"github.com/Vallghall/CRYMEIN/internal/hash"
	"github.com/Vallghall/CRYMEIN/internal/rsa"
)

func Hash() {
	fmt.Println("Введите текст для шифрования:")
	var txt string
	fmt.Scan(&txt)
	primes := rsa.NewPrimesFromInput()
	hashed := hash.Transform(txt, primes)
	fmt.Printf("Хеш-образ: %v", hashed)
}
