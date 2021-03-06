package main

import (
	"github.com/Vallghall/CRYMEIN/pkg/des"
	"github.com/Vallghall/CRYMEIN/pkg/esign"
	"github.com/Vallghall/CRYMEIN/pkg/gost"
	"github.com/Vallghall/CRYMEIN/pkg/hash"
	"github.com/Vallghall/CRYMEIN/pkg/rsa"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		os.Stderr.Write([]byte("No Args provided"))
		os.Exit(2)
	}
	switch args[1] {
	case "DES":
		des.DES()
	case "GOST":
		gost.GOST()
	case "RSA":
		rsa.RSA()
	case "HASH":
		hash.Hash()
	case "ESIGN":
		esign.ESign()
	default:
		des.DES()
	}
}
