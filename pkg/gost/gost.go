package gost

import (
	"fmt"
	"github.com/Vallghall/CRYMEIN/internal/alphabet"
	"github.com/Vallghall/CRYMEIN/internal/gost"
	"os"
)

func GOST() {
	key, _ := os.ReadFile("plaintext.txt")
	fmt.Printf("Последовательность символов ключа\n%s\n", key)
	fmt.Print("Введите текс: ")
	var plainText string
	fmt.Scanln(&plainText)

	rusText := alphabet.ToRussianBytes(plainText)
	rusKey := alphabet.ToRussianBytes(string(key))
	encrypted := gost.Encrypt(rusText, rusKey)
	fmt.Printf("Итого получаем:\n%08b\n", encrypted)
}
