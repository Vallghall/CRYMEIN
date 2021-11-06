package alphabet

import (
	"bytes"
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
