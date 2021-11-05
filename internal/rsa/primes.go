package rsa

import (
	"fmt"
	"log"
	"math/big"
)

type primes struct {
	p   *big.Int
	q   *big.Int
	n   *big.Int
	phi *big.Int
}

func (p primes) Phi() *big.Int {
	return p.phi
}

func (p primes) P() *big.Int {
	return p.p
}

func (p primes) Q() *big.Int {
	return p.q
}

func (p primes) N() *big.Int {
	return p.n
}

func (p *primes) generatePhi() {
	a, b, one := new(big.Int), new(big.Int), big.NewInt(1)
	p.phi = new(big.Int).Mul(a.Sub(p.P(), one), b.Sub(p.Q(), one))
}

func NewPrimesFromInput() *primes {
	p, err := getPrimeFromInput("P")
	if err != nil {
		log.Fatalln(err)
	}
	q, err := getPrimeFromInput("Q")
	if err != nil {
		log.Fatalln(err)
	}
	n := new(big.Int).Mul(p, q)
	return &primes{
		p: p,
		q: q,
		n: n,
	}
}

func getPrimeFromInput(vrbl string) (*big.Int, error) {
	bi := new(big.Int)
	var temp string
	fmt.Printf("Введите число %s: ", vrbl)
	fmt.Scan(&temp)
	if _, err := fmt.Sscan(temp, bi); err != nil {
		return nil, err
	}

	if !bi.ProbablyPrime(0) {
		return nil, fmt.Errorf("%v is not a prime", bi)
	}
	return bi, nil
}
