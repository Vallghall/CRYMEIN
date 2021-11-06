package esign

import (
	"fmt"
	"github.com/Vallghall/CRYMEIN/internal/alphabet"
	"github.com/Vallghall/CRYMEIN/internal/hash"
	"github.com/Vallghall/CRYMEIN/internal/rsa"
)

func Sign(txt string) bool {
	primes := rsa.NewPrimesFromInput()
	docHash := hash.Transform(alphabet.ToRussianIndexes(txt), primes)
	fmt.Printf("Первичный хэш:\n%v\n", docHash)
	kp := primes.GenerateKeyPair()
	signed := kp.Encrypt([]int64{docHash})
	fmt.Printf("Подпись:\n%v\n", signed[0])
	unsigned := kp.Decrypt(signed)[0]
	fmt.Printf("Вторичный хэш:\n%v\n", unsigned)
	return docHash == unsigned
}
