package utils

import (
	"strings"
	"time"
)

//去掉空白字符
func Trim(str string) string {
	return strings.TrimLeft(strings.Trim(str, " "), " ")
}

//字符串联接
func Contact(str string, args ...string) string {
	// for _, arg := range args {
	//     result += arg
	// }
	result := str
	for _, arg := range args {
		result = result + arg
	}
	return result
}

func ParseExcelDate(t1 string) (time.Time, error) {
	timeTemplate3 := "2006-01-02"
	result, err := time.ParseInLocation(t1, timeTemplate3, time.Local)
	if err != nil {
		return time.Now(), err
	}
	return result, nil
}

func ParseTimestamp(t1 string) (int64, error) {
	t2, err := ParseExcelDate(t1)
	if err != nil {
		return 0, err
	}
	return t2.UnixNano(), nil
}
