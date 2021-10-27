package des

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

var ReplacementPositions = [64]byte{
	58, 50, 42, 34, 26, 18, 10, 2,
	60, 52, 44, 36, 28, 20, 12, 4,
	62, 54, 46, 38, 30, 22, 14, 6,
	64, 56, 48, 40, 32, 24, 16, 8,
	57, 49, 41, 33, 25, 17, 9, 1,
	59, 51, 43, 35, 27, 19, 11, 3,
	61, 53, 45, 37, 29, 21, 13, 5,
	63, 55, 47, 39, 31, 23, 15, 7,
}

var ExtensionBlock = [48]byte{
	32, 1, 2, 3, 4, 5,
	4, 5, 6, 7, 8, 9,
	8, 9, 10, 11, 12, 13,
	12, 13, 14, 15, 16, 17,
	16, 17, 18, 19, 20, 21,
	20, 21, 22, 23, 24, 25,
	24, 25, 26, 27, 28, 29,
	28, 29, 30, 31, 32, 1,
}

func Encrypt(textBytes, keyBytes []byte) []byte {
	buf := new(bytes.Buffer)

	text := binary.BigEndian.Uint64(textBytes)
	fmt.Printf("Биты исходного текста    : %08b\n", text)

	key := binary.BigEndian.Uint64(keyBytes)
	fmt.Printf("Биты заданного ключа     : %08b\n", key)

	factKey := shortenKey(keyBytes, 7)
	fmt.Printf("Биты фактич. ключа шифра : %08b\n", factKey)

	binary.Write(buf, binary.BigEndian, factKey)
	factKeyBytes := buf.Bytes()
	factKeyBytes = append(factKeyBytes, factKeyBytes[1:]...)
	roundKey := shortenKey(keyBytes, 6) >> 1
	fmt.Printf("Биты ключа раунда        : %08b\n", roundKey)

	pk := permuteBlock(text, ReplacementPositions[:], 64)
	fmt.Printf("Биты перемешанного текста: %08b\n", pk)

	L0, R0 := uint32(pk>>32), uint32(pk)
	fmt.Printf("Биты левой части : %08b\n", L0)
	fmt.Printf("Биты правой части: %08b\n", R0)

	expR := permuteBlock(uint64(R0), ExtensionBlock[:], 32)
	fmt.Printf("Расширенный R: %08b\n", expR)

	return nil
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

/*
func toByteSlice56(n uint64) []byte {
	res := make([]byte,7,7)
	for i := 0; i < 7; i++ {
		temp := binary.
	}
}
*/
