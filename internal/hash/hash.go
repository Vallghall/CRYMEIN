package hash

import (
	"github.com/Vallghall/CRYMEIN/internal/alphabet"
	"github.com/Vallghall/CRYMEIN/internal/rsa"
	"math/big"
)

const H0 = 8

func Transform(txtStr string, primes *rsa.Primes) int64 {
	txt := alphabet.ToRussianIndexes(txtStr)
	n := primes.N()
	h := big.NewInt(H0)
	for _, m := range txt {
		h = hashFunc(h, big.NewInt(m), n)
	}
	return h.Int64()
}

func hashFunc(h, m, n *big.Int) *big.Int {
	res := new(big.Int)
	return res.Mod(res.Add(h, m).Mul(res, res), n)
}
