package alphabet

import (
	"bytes"
	"strconv"
	"strings"
)

func ToRussianBytes(original string) []byte {
	return rusToRSCII(strings.TrimSpace(strings.TrimRight(original, "\n")))
}

func rusToRSCII(word string) []byte {
	bb := bytes.Buffer{}
	for _, sym := range word {
		if sym != 32 {
			bb.WriteByte(byte(sym%1040 + 192))
		} else {
			bb.WriteByte(byte(32))
		}
	}
	return bb.Bytes()
}

func rSCIIToRus(word string) string {
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
