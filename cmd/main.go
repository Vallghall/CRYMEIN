package main

import (
	"github.com/Vallghall/CRYMEIN/pkg/des"
	"github.com/Vallghall/CRYMEIN/pkg/gost"
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
	default:
		des.DES()
	}
}
