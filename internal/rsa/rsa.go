package rsa

import (
	"fmt"
)

func Encrypt(txt []rune, p, q int64) {
	phi := (p - 1) * (q - 1)
	fmt.Printf("P = %v; \tQ = %v; \t\uF06A = %v;\n", p, q, phi)

	d := findPrivateKey(phi)
	fmt.Printf("Закрытый ключ: %v\n", d)

	e := findPublicKey(d, phi)
	fmt.Printf("Открытый ключ: %v\n", e)

}

func findPublicKey(d, phi int64) int64 {
	var e int64 = 2
	for {
		if (d*e-1)%phi == 0 {
			return e
		}
		e++
	}
}

func findPrivateKey(phi int64) int64 {
	var d int64 = 2
	for {
		if isCoPrime(d, phi) {
			return d
		}
		d++
	}
}

func isCoPrime(a, b int64) bool {
	if a == b {
		return a == 1
	}

	if a > b {
		return isCoPrime(a-b, b)
	} else {
		return isCoPrime(a, b-a)
	}
}
