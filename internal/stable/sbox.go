package stable

import (
	"strconv"
)

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
