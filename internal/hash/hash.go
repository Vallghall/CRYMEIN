package hash

import (
	"github.com/Vallghall/CRYMEIN/internal/rsa"
	"math/big"
)

const H0 = 8

func Transform(txt []int64, primes *rsa.Primes) int64 {
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
