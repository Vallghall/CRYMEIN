package esign

import (
	"fmt"
	"github.com/Vallghall/CRYMEIN/internal/hash"
	"github.com/Vallghall/CRYMEIN/internal/rsa"
)

func Sign(txt string) bool {
	primes := rsa.NewPrimesFromInput()
	docHash := hash.Transform(txt, primes)
	fmt.Println(docHash)
	newPrimes := rsa.NewPrimesFromInput()
	kp := newPrimes.GenerateKeyPair()
	signed := rsa.Encrypt(string(rune(docHash)), kp.PublicKey)
	unsigned := rsa.Decrypt(signed, kp.PrivateKey)
	urunes := int64ToRunes(unsigned)
	hashSum := hash.Transform(string(urunes), primes)
	fmt.Println(hashSum)
	return docHash == hashSum
}

func int64ToRunes(n []int64) []rune {
	res := make([]rune, len(n), len(n))
	for i, val := range n {
		res[i] = rune(val)
	}
	return res
}
