package rsa

import (
	"fmt"
	"github.com/Vallghall/CRYMEIN/internal/rsa"
	"log"
	"math/big"
)

func RSA() {
	p, err := getPrimeFromInput("P")
	if err != nil {
		log.Fatalln(err)
	}
	q, err := getPrimeFromInput("Q")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("P = %v; \tQ = %v;\n", p, q)

	rsa.Encrypt([]rune{}, p.Int64(), q.Int64())

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
