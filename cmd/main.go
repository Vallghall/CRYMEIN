package main

import (
	"bufio"
	"fmt"
	"github.com/Vallghall/CRYMEIN/internal/alphabet"
	_ "github.com/Vallghall/CRYMEIN/internal/des"
	"os"
	"strings"
)

func main() {
	fmt.Print("Введите фразу: ")
	word, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	cip := alphabet.RusToRSCII(strings.TrimSpace(strings.TrimRight(word, "\n")))
	fmt.Println(cip)
	fmt.Println(alphabet.RSCIIToRus(cip))
}
