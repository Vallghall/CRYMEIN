package stable

import (
	"fmt"
	"strconv"
)

type sTable [4][16]byte

func (s *sTable) value(row, column uint8) string {
	return fmt.Sprintf("%04b", s[row][column])
}

type sBox struct {
	s1 *sTable
	s2 *sTable
	s3 *sTable
	s4 *sTable
	s5 *sTable
	s6 *sTable
	s7 *sTable
	s8 *sTable
}

func NewSBox() *sBox {
	return &sBox{
		s1: &s1,
		s2: &s2,
		s3: &s3,
		s4: &s4,
		s5: &s5,
		s6: &s6,
		s7: &s7,
		s8: &s8,
	}
}

func (sb *sBox) GetS1Value(str string) string {
	row, _ := strconv.ParseInt(str[:1]+str[5:6], 2, 8)
	col, _ := strconv.ParseInt(str[1:5], 2, 8)
	return sb.s1.value(uint8(row), uint8(col))
}

func (sb *sBox) GetS2Value(str string) string {
	row, _ := strconv.ParseInt(str[:1]+str[5:6], 2, 8)
	col, _ := strconv.ParseInt(str[1:5], 2, 8)
	return sb.s2.value(uint8(row), uint8(col))
}

func (sb *sBox) GetS3Value(str string) string {
	row, _ := strconv.ParseInt(str[:1]+str[5:6], 2, 8)
	col, _ := strconv.ParseInt(str[1:5], 2, 8)
	return sb.s3.value(uint8(row), uint8(col))
}

func (sb *sBox) GetS4Value(str string) string {
	row, _ := strconv.ParseInt(str[:1]+str[5:6], 2, 8)
	col, _ := strconv.ParseInt(str[1:5], 2, 8)
	return sb.s4.value(uint8(row), uint8(col))
}

func (sb *sBox) GetS5Value(str string) string {
	row, _ := strconv.ParseInt(str[:1]+str[5:6], 2, 8)
	col, _ := strconv.ParseInt(str[1:5], 2, 8)
	return sb.s5.value(uint8(row), uint8(col))
}

func (sb *sBox) GetS6Value(str string) string {
	row, _ := strconv.ParseInt(str[:1]+str[5:6], 2, 8)
	col, _ := strconv.ParseInt(str[1:5], 2, 8)
	return sb.s6.value(uint8(row), uint8(col))
}

func (sb *sBox) GetS7Value(str string) string {
	row, _ := strconv.ParseInt(str[:1]+str[5:6], 2, 8)
	col, _ := strconv.ParseInt(str[1:5], 2, 8)
	return sb.s7.value(uint8(row), uint8(col))
}

func (sb *sBox) GetS8Value(str string) string {
	row, _ := strconv.ParseInt(str[:1]+str[5:6], 2, 8)
	col, _ := strconv.ParseInt(str[1:5], 2, 8)
	return sb.s8.value(uint8(row), uint8(col))
}

var s1 = sTable{
	{14, 4, 13, 1, 2, 15, 11, 8, 3, 10, 6, 12, 5, 9, 0, 7},
	{0, 15, 7, 4, 14, 2, 13, 1, 10, 6, 12, 11, 9, 5, 3, 8},
	{4, 1, 14, 8, 13, 6, 2, 11, 15, 12, 9, 7, 3, 10, 5, 0},
	{15, 12, 8, 2, 4, 9, 1, 7, 5, 11, 3, 14, 10, 0, 6, 13},
}

var s2 = sTable{
	{15, 1, 8, 14, 6, 11, 3, 4, 9, 7, 2, 13, 12, 0, 5, 10},
	{3, 13, 4, 7, 15, 2, 8, 14, 12, 0, 1, 10, 6, 9, 11, 5},
	{0, 14, 7, 11, 10, 4, 13, 1, 5, 8, 12, 6, 9, 3, 2, 15},
	{13, 8, 10, 1, 3, 15, 4, 2, 11, 6, 7, 12, 0, 5, 14, 9},
}

var s3 = sTable{
	{10, 0, 9, 14, 6, 3, 15, 5, 1, 13, 12, 7, 11, 4, 2, 8},
	{13, 7, 0, 9, 3, 4, 6, 10, 2, 8, 5, 14, 12, 11, 15, 1},
	{13, 6, 4, 9, 8, 15, 3, 0, 11, 1, 2, 12, 5, 10, 14, 7},
	{1, 10, 13, 0, 6, 9, 8, 7, 4, 15, 14, 3, 11, 5, 2, 12},
}

var s4 = sTable{
	{7, 13, 14, 3, 0, 6, 9, 10, 1, 2, 8, 5, 11, 12, 4, 15},
	{13, 8, 11, 5, 6, 15, 0, 3, 4, 7, 2, 12, 1, 10, 14, 9},
	{10, 6, 9, 0, 12, 11, 7, 13, 15, 1, 3, 14, 5, 2, 8, 4},
	{3, 15, 0, 6, 10, 1, 13, 8, 9, 4, 5, 11, 12, 7, 2, 14},
}

var s5 = sTable{
	{2, 12, 4, 1, 7, 10, 11, 6, 8, 5, 3, 15, 13, 0, 14, 9},
	{14, 11, 2, 12, 4, 7, 13, 1, 5, 0, 15, 10, 3, 9, 8, 6},
	{4, 2, 1, 11, 10, 13, 7, 8, 15, 9, 12, 5, 6, 3, 0, 14},
	{11, 8, 12, 7, 1, 14, 2, 13, 6, 15, 0, 9, 10, 4, 5, 3},
}

var s6 = sTable{
	{12, 1, 10, 15, 9, 2, 6, 8, 0, 13, 3, 4, 14, 7, 5, 11},
	{10, 15, 4, 2, 7, 12, 9, 5, 6, 1, 13, 14, 0, 11, 3, 8},
	{9, 14, 15, 5, 2, 8, 12, 3, 7, 0, 4, 10, 1, 13, 11, 6},
	{4, 3, 2, 12, 9, 5, 15, 10, 11, 14, 1, 7, 6, 0, 8, 13},
}

var s7 = sTable{
	{4, 11, 2, 14, 15, 0, 8, 13, 3, 12, 9, 7, 5, 10, 6, 1},
	{13, 0, 11, 7, 4, 9, 1, 10, 14, 3, 5, 12, 2, 15, 8, 6},
	{1, 4, 11, 13, 12, 3, 7, 14, 10, 15, 6, 8, 0, 5, 9, 2},
	{6, 11, 13, 8, 1, 4, 10, 7, 9, 5, 0, 15, 14, 2, 3, 12},
}

var s8 = sTable{
	{13, 2, 8, 4, 6, 15, 11, 1, 10, 9, 3, 14, 5, 0, 12, 7},
	{1, 15, 13, 8, 10, 3, 7, 4, 12, 5, 6, 11, 0, 14, 9, 2},
	{7, 11, 4, 1, 9, 12, 14, 2, 0, 6, 10, 13, 15, 3, 5, 8},
	{2, 1, 14, 7, 4, 10, 8, 13, 15, 12, 9, 0, 3, 5, 6, 11},
}
