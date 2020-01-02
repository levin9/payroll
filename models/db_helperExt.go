package models

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

func DoQuery(sqlInfo string, args ...interface{}) ([]map[string]interface{}, error) {
	conn := fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=True&loc=Local", viper.GetString("mysql.user"),
		viper.GetString("mysql.password"), viper.GetString("mysql.addr"), viper.GetString("mysql.database"),
		viper.GetString("mysql.charset"))

	db, err := sql.Open("mysql", conn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query(sqlInfo, args...)

	if err != nil {
		return nil, err
	}
	columns, _ := rows.Columns()
	columnLength := len(columns)
	cache := make([]interface{}, columnLength) //临时存储每行数据
	for index, _ := range cache {              //为每一列初始化一个指针
		var a interface{}
		cache[index] = &a
	}
	var list []map[string]interface{} //返回的切片
	for rows.Next() {
		_ = rows.Scan(cache...)

		item := make(map[string]interface{})
		for i, data := range cache {
			item[columns[i]] = *data.(*interface{}) //取实际类型
		}
		list = append(list, item)
	}
	_ = rows.Close()
	return list, nil
}

// func ExecuteMap(sqlInfo string, args ...interface{}) ([]map[string]interface{}, error) {
// 	rows := mysqlDB.Raw(sqlInfo, args...)
// 	columns := getFieldNames(sqlInfo)

// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	//columns, _ := rows.Columns()
// 	columnLength := len(columns)
// 	cache := make([]interface{}, columnLength) //临时存储每行数据
// 	for index, _ := range cache {              //为每一列初始化一个指针
// 		var a interface{}
// 		cache[index] = &a
// 	}
// 	var list []map[string]interface{} //返回的切片
// 	for rows.Next() {
// 		_ = rows.Scan(cache...)

// 		item := make(map[string]interface{})
// 		for i, data := range cache {
// 			item[columns[i]] = *data.(*interface{}) //取实际类型
// 		}
// 		list = append(list, item)
// 	}
// 	_ = rows.Close()
// 	return list, nil
// }

func getFieldNames(sql string) []string {

	sql = strings.ToLower(sql)
	//fmt.Println(strings.Index(sql, "from"))
	sql = substring(sql, 7, strings.Index(sql, "from"))

	arr := strings.Split(sql, ",")
	result := make([]string, len(arr))

	for key, a := range arr {
		field := strings.Split(a, "as")
		//fmt.Println(strings.Trim(field[len(field)-1], " "))
		result[key] = strings.Trim(field[len(field)-1], " ")
	}

	return result
}

//获取source的子串,如果start小于0或者end大于source长度则返回""
//start:开始index，从0开始，包括0
//end:结束index，以end结束，但不包括end
func substring(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)

	if start < 0 || end > length || start > end {
		return ""
	}

	if start == 0 && end == length {
		return source
	}

	return string(r[start:end])
}
