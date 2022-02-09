package strconvx

import (
	"strconv"
)

func MustStrToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func IntToStr(i int) string {
	return strconv.Itoa(i)
}

func StrToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func StrToInt32(s string) (int32, error) {
	num, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(num), nil
}

func StrToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func MustStrToInt32(s string) int32 {
	num, _ := StrToInt32(s)
	return num
}

func MustStrToInt64(s string) int64 {
	num, _ := StrToInt64(s)
	return num
}

func StrToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func MustStrToFloat64(s string) float64 {
	num, _ := StrToFloat64(s)
	return num
}
