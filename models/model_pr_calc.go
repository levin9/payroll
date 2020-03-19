package models

import (
	"encoding/json"
	"fmt"
)

type UpdatePayrollDao struct {
	Id       string
	Data     map[string]string
	Tips     map[string]string
	CreateBy string
}

func (o *UpdatePayrollDao) AddId(id string) {
	o.Id = id
	o.CreateBy = "fff"
}

func (o *UpdatePayrollDao) AddField(fieldName, fieldValue, fieldName_CN, formulaText string) {
	if o.Data == nil {
		o.Data = make(map[string]string)
		o.Tips = make(map[string]string)
	}
	o.Data[fieldName] = fieldValue
	o.Tips[fieldName_CN] = formulaText
}

func (o *UpdatePayrollDao) SaveChange() {

	sql := "update sa_monthpayroll set "
	for k, v := range o.Data {
		sql = sql + k + "=" + v + ","
	}
	formulaText, _ := json.Marshal(o.Tips)
	sql = sql + " ExtInfo='" + string(formulaText) + "'"
	sql = sql + " ,F_ModifyDate=now() ,F_ModifyUserName='" + o.CreateBy + "' "
	sql = sql + " where id='" + o.Id + "'"
	fmt.Println(sql)
	MysqlDB.Exec(sql)
}
