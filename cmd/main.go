package main

import (
	"github.com/Vallghall/CRYMEIN/pkg/des"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		os.Stderr.Write([]byte("No Args provided"))
		os.Exit(2)
	}
	switch os.Args[1] {
	case "DES":
		des.DES()
	default:
		des.DES()
	}
}
