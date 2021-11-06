package des

import (
	"bufio"
	"fmt"
	"github.com/Vallghall/CRYMEIN/internal/alphabet"
	"github.com/Vallghall/CRYMEIN/internal/des"
	"os"
)

func DES() {
	fmt.Print("Введите исходный текст: ")
	reader := bufio.NewReader(os.Stdin)
	plainText, _ := reader.ReadString('\n')
	fmt.Print("Введите ключ: ")
	key, _ := reader.ReadString('\n')
	rusText := alphabet.ToRussianBytes(plainText)
	rusKey := alphabet.ToRussianBytes(key)
	encrypted := des.Encrypt(rusText, rusKey)
	fmt.Printf("Финальный результат: %08b\n", encrypted)
}
