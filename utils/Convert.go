package utils

import (
	"fmt"
	"strconv"

	uuid "github.com/satori/go.uuid"
)

func Print() {
	fmt.Println("hello,utils")
}

//float32 转 String工具类，保留3位小数
func Float32ToString(input_num float32) string {
	// to convert a float number to a string
	return strconv.FormatFloat(float64(input_num), 'f', 6, 64)
}

//float64 转 String工具类，保留3位小数
func Float64ToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

//字符转32位float，为空，为0
func ToFloat32(str string) float32 {
	if str == "" {
		return 0
	}
	result, err := strconv.ParseFloat(str, 32)
	if err != nil {
		fmt.Println("TODO")
	}
	return float32(result)
}

//字符转64位float，为空，为0
func ToFloat64(str string) float64 {
	if str == "" {
		return 0
	}
	result, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("TODO")
	}
	return result
}

//NewID
func NewID() string {
	id := uuid.NewV4()
	return id.String()
}

// func ToDate(str string, format string) float64 {
// 	if str == "" {
// 		return 0
// 	}
// 	result, err := strconv.ParseFloat(str, 64)
// 	if err != nil {
// 		fmt.Println("TODO")
// 	}
// 	return result
// }

// func ToDateWithFormat(str string, format string) float64 {
// 	if str == "" {
// 		return 0
// 	}
// 	result, err := strconv.ParseFloat(str, 64)
// 	if err != nil {
// 		fmt.Println("TODO")
// 	}
// 	return result
// }
