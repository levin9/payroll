package utils

import "fmt"

type ConvertUtil struct {
}

func (c *ConvertUtil) Print() {
	fmt.Println("hello world")
}

// //Interface2String
// func Interface2String(inter interface{}) string {

// 	switch inter.(type) {

// 	case string:
// 		fmt.Println("string", inter.(string))
// 		break
// 	case int:
// 		fmt.Println("int", inter.(int))
// 		break
// 	case float64:
// 		fmt.Println("float64", inter.(float64))
// 		break
// 	}
// 	return ""
// }
