package des

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/Vallghall/CRYMEIN/internal/stable"
	"strconv"
	"strings"
)

var replacementPositions = [64]byte{
	58, 50, 42, 34, 26, 18, 10, 2,
	60, 52, 44, 36, 28, 20, 12, 4,
	62, 54, 46, 38, 30, 22, 14, 6,
	64, 56, 48, 40, 32, 24, 16, 8,
	57, 49, 41, 33, 25, 17, 9, 1,
	59, 51, 43, 35, 27, 19, 11, 3,
	61, 53, 45, 37, 29, 21, 13, 5,
	63, 55, 47, 39, 31, 23, 15, 7,
}

var extensionBlock = [48]byte{
	32, 1, 2, 3, 4, 5,
	4, 5, 6, 7, 8, 9,
	8, 9, 10, 11, 12, 13,
	12, 13, 14, 15, 16, 17,
	16, 17, 18, 19, 20, 21,
	20, 21, 22, 23, 24, 25,
	24, 25, 26, 27, 28, 29,
	28, 29, 30, 31, 32, 1,
}

var replacement32Positions = [32]byte{
	16, 7, 20, 21, 29, 12, 28, 17,
	1, 15, 23, 26, 5, 18, 31, 10,
	2, 8, 24, 14, 32, 27, 3, 9,
	19, 13, 30, 6, 22, 11, 4, 25,
}

var finalReplacementPositions = [64]byte{
	40, 8, 48, 16, 56, 24, 64, 32,
	39, 7, 47, 15, 55, 23, 63, 31,
	38, 6, 46, 14, 54, 22, 62, 30,
	37, 5, 45, 13, 53, 21, 61, 29,
	36, 4, 44, 12, 52, 20, 60, 28,
	35, 3, 43, 11, 51, 19, 59, 27,
	34, 2, 42, 10, 50, 18, 58, 26,
	33, 1, 41, 9, 49, 17, 57, 25,
}

func Encrypt(textBytes, keyBytes []byte) []byte {
	buf := new(bytes.Buffer)

	text := binary.BigEndian.Uint64(textBytes)
	fmt.Printf("Биты исходного текста    : %b\n", text)

	key := binary.BigEndian.Uint64(keyBytes)
	fmt.Printf("Биты заданного ключа     : %b\n", key)

	factKey := shortenKey(keyBytes, 7)
	fmt.Printf("Биты фактич. ключа шифра : %b\n", factKey)

	binary.Write(buf, binary.BigEndian, factKey)
	factKeyBytes := buf.Bytes()
	factKeyBytes = append(factKeyBytes, factKeyBytes[1:]...)
	roundKey := shortenKey(keyBytes, 6) >> 1
	fmt.Printf("Биты ключа раунда        : %b\n", roundKey)

	pk := permuteBlock(text, replacementPositions[:], 64)
	fmt.Printf("Биты перемешанного текста: %b\n", pk)

	L0, R0 := uint32(pk>>32), uint32(pk)
	fmt.Printf("Биты левой части : %b\n", L0)
	fmt.Printf("Биты правой части: %b\n", R0)

	expR := permuteBlock(uint64(R0), extensionBlock[:], 32)
	fmt.Printf("Расширенный R: %b\n", expR)

	whitener := expR ^ roundKey
	fmt.Printf("Результат XOR: %b\n", whitener)

	s := fmt.Sprintf("%b", whitener)
	for len(s) != 48 {
		s = "0" + s
	}
	s = sBoxing(s)
	fmt.Printf("Результат S-боксинга: %s\n", s)

	sInt64, _ := strconv.ParseInt(s, 2, 64)
	newR0 := permuteBlock(uint64(sInt64), replacement32Positions[:], 32)
	fmt.Printf("Перестановка результата S-боксинга: %b\n", newR0)

	R1 := uint32(newR0) ^ L0
	preFinal := fmt.Sprintf("%b", R0) + fmt.Sprintf("%b", R1)
	fmt.Printf("Результат конкатенации L1 и R1: %s\n", preFinal)

	finalSrc, _ := strconv.ParseInt(preFinal, 2, 64)
	final := permuteBlock(uint64(finalSrc), finalReplacementPositions[:], 64)
	buf.Reset()
	binary.Write(buf, binary.BigEndian, final)

	return buf.Bytes()
}

func permuteBlock(src uint64, permutation []uint8, size uint8) (block uint64) {
	for position, n := range permutation {
		bit := (src >> (size - n)) & 1
		block |= bit << uint((len(permutation)-1)-position)
	}
	return
}

func shortenKey(key []byte, size int) (shortKey uint64) {
	for position, n := range key {
		shortBlock := uint64(n) >> 1
		shortKey |= shortBlock << uint(size*(7-position))
	}
	return
}

func sBoxing(xored string) string {
	builder := strings.Builder{}
	sb := stable.NewSBox()

	builder.WriteString(sb.GetS1Value(xored[:6]))
	builder.WriteString(sb.GetS2Value(xored[6:12]))

	builder.WriteString(sb.GetS3Value(xored[12:18]))
	builder.WriteString(sb.GetS4Value(xored[18:24]))

	builder.WriteString(sb.GetS5Value(xored[24:30]))
	builder.WriteString(sb.GetS6Value(xored[30:36]))

	builder.WriteString(sb.GetS7Value(xored[36:42]))
	builder.WriteString(sb.GetS8Value(xored[42:48]))

	return builder.String()
}
