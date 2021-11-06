package hash

import (
	"fmt"
	"github.com/Vallghall/CRYMEIN/internal/alphabet"
	"github.com/Vallghall/CRYMEIN/internal/hash"
	"github.com/Vallghall/CRYMEIN/internal/rsa"
)

func Hash() {
	fmt.Println("Введите текст для хеширования:")
	var txt string
	fmt.Scan(&txt)
	primes := rsa.NewPrimesFromInput()
	hashed := hash.Transform(alphabet.ToRussianIndexes(txt), primes)
	fmt.Printf("Хеш-образ: %v", hashed)
}
