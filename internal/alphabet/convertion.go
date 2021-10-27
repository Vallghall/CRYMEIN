package alphabet

import (
	"fmt"
	"strconv"
	"strings"
)

func RusToRSCII(word string) string {
	sb := strings.Builder{}
	for _, sym := range word {
		if sym != 32 {
			sb.WriteString(fmt.Sprintf("%08b ", sym%1040+192))
		} else {
			sb.WriteString(fmt.Sprintf("%08b ", 32))
		}
	}
	return strings.TrimSpace(sb.String())
}

func RSCIIToRus(word string) string {
	sb := strings.Builder{}
	for _, sym := range strings.Split(word, " ") {
		letter, _ := strconv.ParseInt(sym, 2, 32)
		if letter == 32 {
			sb.WriteRune(rune(letter))
		} else {
			sb.WriteRune(rune(letter - 192 + 1040))
		}
	}
	return sb.String()
}
