package gost

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/bits"
)

var substitutionTable = [8][16]byte{
	{4, 10, 9, 2, 13, 8, 0, 14, 6, 11, 1, 12, 7, 15, 5, 3},
	{14, 11, 4, 12, 6, 13, 15, 10, 2, 3, 8, 1, 0, 7, 5, 9},
	{5, 8, 1, 13, 10, 3, 4, 2, 14, 15, 12, 7, 6, 0, 9, 11},
	{7, 13, 10, 1, 0, 8, 9, 15, 14, 4, 6, 12, 11, 2, 5, 3},
	{6, 12, 7, 1, 5, 15, 13, 8, 4, 10, 9, 14, 0, 3, 11, 2},
	{4, 11, 10, 0, 7, 2, 1, 13, 3, 6, 8, 5, 9, 12, 15, 14},
	{13, 11, 4, 1, 3, 15, 5, 9, 0, 10, 14, 7, 6, 8, 2, 12},
	{1, 15, 13, 0, 5, 7, 10, 4, 9, 2, 3, 14, 6, 11, 8, 12},
}

func Encrypt(textBytes, keyBytes []byte) []byte {
	//БоберВасяНатачиваетЗолотойОрешек
	bb := new(bytes.Buffer)
	firstSubKey := make([]byte, 4, 4)
	copy(firstSubKey, keyBytes)
	X0 := binary.BigEndian.Uint32(firstSubKey)
	fmt.Printf("Сравнение битов ключа и подключа:\n%08b\n%08b\n", keyBytes, X0)
	L0, R0 := binary.BigEndian.Uint32(textBytes[:4]), binary.BigEndian.Uint32(textBytes[4:])
	fmt.Printf("Сравнение битов текста и половин текста:\n%08b\n%08b%08b\n", binary.BigEndian.Uint64(textBytes), L0, R0)

	transformationFuncResult := f(X0, R0)
	fmt.Printf("Циклически сдвинутый результат S-таблицы:\n%08b\n", transformationFuncResult)

	R1 := transformationFuncResult ^ L0
	fmt.Printf("R1:\n%08b\n", R1)

	binary.Write(bb, binary.BigEndian, (uint64(R0)<<32)|uint64(R1))

	return bb.Bytes()
}

func f(X0, R0 uint32) uint32 {
	sum, _ := bits.Add32(R0, X0, 0)
	fmt.Printf("Мод суммы R0 и X0 по модулю 32:\n%08b\n", sum)

	adrs := makeSubstitutionTableAddresses(sum)
	fmt.Printf("Адресса для таблицы подстановок:\n%08b\n", adrs)

	sValues := valuesFromSubstitutionTable(adrs)
	fmt.Printf("Значения S-таблицы:\n%08b\n", sValues)

	return bits.RotateLeft32(sValues, 11)
}

func makeSubstitutionTableAddresses(n uint32) []uint32 {
	res := make([]uint32, 8, 8)
	for i := 0; i < 8; i++ {
		res[i] = (n >> (i * 4)) & 15
	}
	return res
}

func valuesFromSubstitutionTable(adrs []uint32) uint32 {
	res := make([]byte, 8, 8)
	for col, row := range adrs {
		res[7-col] = substitutionTable[col][row]
	}
	return byteSliceToUint32(res)
}

func byteSliceToUint32(src []byte) (dst uint32) {
	for i, b := range src {
		dst |= uint32(b) << ((7 - i) * 4)
	}
	return
}
