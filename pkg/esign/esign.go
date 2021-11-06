package esign

import (
	"fmt"
	"github.com/Vallghall/CRYMEIN/internal/esign"
)

func ESign() {
	fmt.Println("Введите текст для шифрования:")
	var txt string
	fmt.Scan(&txt)
	if esign.Sign(txt) {
		fmt.Println("СОВПАДАЕТ")
	} else {
		fmt.Println("НЕ СОВПАДАЕТ")
	}
}
