package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
)

//PaginationQuery gin handler query binding struct
type PaginationQuery struct {
	Where     string `form:"where"`
	Fields    string `form:"fields"`
	Order     string `form:"order"`
	Offset    uint   `form:"offset"`
	Limit     uint   `form:"limit"`
	PageSize  uint   `form:"pageSize"`
	PageIndex uint   `form:"pageIndex"`
	//Conditions map[string]interface{}
}

// func (pq *PaginationQuery) AppendSql(template string, nVal interface{}) {
// 	if pq.Conditions == nil {
// 		pq.Conditions = make(map[string]interface{})
// 	}
// 	pq.Conditions[template] = nVal
// }

// func (pq *PaginationQuery) WhereIf() string {
// 	return fmt.Sprintf("w=%v_f=%s_o=%s_of=%d_l=%d", pq.Where, pq.Fields, pq.Order, pq.Offset, pq.Limit)
// }

//String to string
func (pq *PaginationQuery) String() string {
	return fmt.Sprintf("w=%v_f=%s_o=%s_of=%d_l=%d", pq.Where, pq.Fields, pq.Order, pq.Offset, pq.Limit)
}

func crudAll(m interface{}, q *PaginationQuery, list interface{}) (total uint, err error) {
	var tx *gorm.DB

	// fmt.Println(q.Fields)
	// fmt.Println(q.Where)
	// fmt.Println(q.Order)
	// fmt.Println(q.Offset)
	// fmt.Println(q.Limit)

	total, tx = getResourceCount(m, q)
	if q.Fields != "" {
		columns := strings.Split(q.Fields, ",")
		if len(columns) > 0 {
			tx = tx.Select(q.Fields)
		}
	}
	if q.Offset > 0 {
		if q.PageSize > 0 {
			q.Offset = q.PageSize * (q.PageIndex - 1)
		} else {
			q.Offset = 0
		}
		tx = tx.Offset(q.Offset)
	}
	if q.Offset > 0 {
		tx = tx.Offset(q.Offset)
	}
	if q.Limit <= 0 {
		if q.PageSize > 0 {
			q.Limit = q.PageSize
		} else {
			q.Limit = 15
		}
	}
	err = tx.Limit(q.Limit).Find(list).Error
	return
}

func crudGetAll(m interface{}, q *PaginationQuery, list interface{}) (total uint, err error) {
	var tx *gorm.DB
	fmt.Println("crud where", q.Where)
	total, tx = getSqlCount(m, q)
	if q.Fields != "" {
		columns := strings.Split(q.Fields, ",")
		if len(columns) > 0 {
			tx = tx.Select(q.Fields)
		}
	}
	if q.Order != "" {
		tx = tx.Order(q.Order)
	}
	if q.Offset <= 0 {
		if q.PageSize > 0 {
			q.Offset = q.PageSize * (q.PageIndex - 1)
		} else {
			q.Offset = 0
		}
		//tx = tx.Offset(q.Offset)
	}

	if q.Limit <= 0 {
		if q.PageSize > 0 {
			q.Limit = q.PageSize
		} else {
			q.Limit = 15
		}
	}
	//tx.Where(q.Where)
	// for key, value := range q.Conditions {
	// 	tx.Where(value, key)
	// 	fmt.Println(key)
	// 	fmt.Println(value)
	// }
	//fmt.Println("tx.where",tx.Where())
	//tx.Where("TenantId=?", "abcd")
	err = tx.Where(q.Where).Offset(q.Offset).Limit(q.Limit).Find(list).Error
	return
}

func crudOne(m interface{}, one interface{}) (err error) {
	if mysqlDB.Where(m).First(one).RecordNotFound() {
		return errors.New("resource is not found")
	}
	return nil
}

func crudUpdate(m interface{}, where interface{}) (err error) {
	db := mysqlDB.Model(where).Updates(m)
	if err = db.Error; err != nil {
		return
	}
	//if db.RowsAffected != 1 {
	//	return errors.New("id is invalid and resource is not found")
	//}
	return nil
}

func crudDelete(m interface{}) (err error) {
	//WARNING When delete a record, you need to ensure it’s primary field has value, and GORM will use the primary key to delete the record, if primary field’s blank, GORM will delete all records for the model
	//primary key must be not zero value
	db := mysqlDB.Delete(m)
	if err = db.Error; err != nil {
		return
	}
	if db.RowsAffected != 1 {
		return errors.New("resource is not found to destroy")
	}
	return nil
}
func getResourceCount(m interface{}, q *PaginationQuery) (uint, *gorm.DB) {
	var tx = mysqlDB.Model(m)
	conditions := strings.Split(q.Where, ",")
	for _, val := range conditions {
		w := strings.SplitN(val, ":", 2)
		if len(w) == 2 {
			bindKey, bindValue := w[0], w[1]
			if intV, err := strconv.ParseInt(bindValue, 10, 64); err == nil {
				// bind value is int
				field := fmt.Sprintf("`%s` > ?", bindKey)
				tx = tx.Where(field, intV)
			} else if fV, err := strconv.ParseFloat(bindValue, 64); err == nil {
				// bind value is float
				field := fmt.Sprintf("`%s` > ?", bindKey)
				tx = tx.Where(field, fV)
			} else if bindValue != "" {
				// bind value is string
				field := fmt.Sprintf("`%s` LIKE ?", bindKey)
				sV := fmt.Sprintf("%%%s%%", bindValue)
				tx = tx.Where(field, sV)
			}
		}
	}
	modelName := getType(m)
	rKey := redisPrefix + modelName + q.String() + "_count"
	v, err := mem.GetUint(rKey)
	if err != nil {
		var count uint
		tx.Count(&count)
		mem.Set(rKey, count)
		return count, tx
	}
	return v, tx
}

func getSqlCount(m interface{}, q *PaginationQuery) (uint, *gorm.DB) {
	var tx = mysqlDB.Model(m)
	tx.Where(q.Where)
	modelName := getType(m)
	rKey := redisPrefix + modelName + q.String() + "_count"
	v, err := mem.GetUint(rKey)
	if err != nil {
		var count uint
		tx.Where(q.Where).Count(&count)
		mem.Set(rKey, count)
		return count, tx
	}
	return v, tx
}

func getType(v interface{}) string {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	}
	return t.Name()
}
