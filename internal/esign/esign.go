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
	fmt.Println(docHash)
	newPrimes := rsa.NewPrimesFromInput()
	kp := newPrimes.GenerateKeyPair()
	signed := kp.Encrypt([]int64{docHash})
	fmt.Println(signed)
	unsigned := kp.Decrypt(signed)[0]
	fmt.Println(unsigned)
	return docHash == unsigned
}

func int64ToRunes(n []int64) []rune {
	res := make([]rune, len(n), len(n))
	for i, val := range n {
		res[i] = rune(val)
	}
	return res
}
