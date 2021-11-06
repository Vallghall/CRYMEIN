package des

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/Vallghall/CRYMEIN/internal/stable"
	"strconv"
	"strings"
)

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
	roundKey := shortenKey(factKeyBytes, 7) >> 1
	fmt.Printf("Биты ключа раунда        : %b\n", roundKey)

	pk := permuteBlock(text, initialPermutationPositions[:], 64)
	fmt.Printf("Биты перемешанного текста: %b\n", pk)

	L0, R0 := uint32(pk>>32), uint32(pk)
	fmt.Printf("Биты левой части : %b\n", L0)
	fmt.Printf("Биты правой части: %b\n", R0)

	expR := permuteBlock(uint64(R0), rightHalfExtensionPositions[:], 32)
	fmt.Printf("Расширенный R: %b\n", expR)

	whitener := expR ^ roundKey
	fmt.Printf("Результат XOR: %048b\n", whitener)

	s := fmt.Sprintf("%048b", whitener)
	s = sBoxing(s)
	fmt.Printf("Результат S-боксинга: %s\n", s)

	sInt64, _ := strconv.ParseInt(s, 2, 64)
	newR0 := permuteBlock(uint64(sInt64), replacement32Positions[:], 32)
	fmt.Printf("Перестановка результата S-боксинга: %b\n", newR0)

	R1 := uint32(newR0) ^ L0
	preFinal := (uint64(R0) << 32) | uint64(R1)
	fmt.Printf("Результат конкатенации L1 и R1: %b\n", preFinal)

	final := permuteBlock(preFinal, finalReplacementPositions[:], 64)
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
