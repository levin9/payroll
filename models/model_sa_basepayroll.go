package models

import (
	"errors"
	"time"
	
)

var _ = time.Thursday
//Basepayroll
type Basepayroll struct {
	
	Pid   string     `gorm:"column:PID;primary_key" form:"pid" json:"pid" comment:"编号" sql:"varchar(50),PRI"`
	Standardname   string     `gorm:"column:StandardName" form:"standard_name" json:"standard_name" comment:"标准名称" sql:"varchar(50)"`
	Fieldname   string     `gorm:"column:FieldName" form:"field_name" json:"field_name" comment:"字段名称" sql:"varchar(50)"`
	Formula   string     `gorm:"column:Formula" form:"formula" json:"formula" comment:"日薪公式" sql:"varchar(50)"`
	Formulamem   string     `gorm:"column:FormulaMem" form:"formula_mem" json:"formula_mem" comment:"公式备注" sql:"varchar(50)"`
	Ismustused   int     `gorm:"column:IsMustUsed" form:"is_must_used" json:"is_must_used" comment:"必须启用" sql:"int(11)"`
	Isadd   int     `gorm:"column:IsADD" form:"is_add" json:"is_add" comment:"是否加减项" sql:"int(11)"`
	FEnabledmark   int     `gorm:"column:F_EnabledMark" form:"f_enabled_mark" json:"f_enabled_mark" comment:"有效标志0否1是" sql:"int(11)"`
	FCreateuserid   string     `gorm:"column:F_CreateUserId" form:"f_create_user_id" json:"f_create_user_id" comment:"创建人ID" sql:"varchar(50)"`
	FModifyusername   string     `gorm:"column:F_ModifyUserName" form:"f_modify_user_name" json:"f_modify_user_name" comment:"编辑人" sql:"varchar(50)"`
	FModifyuserid   string     `gorm:"column:F_ModifyUserId" form:"f_modify_user_id" json:"f_modify_user_id" comment:"编辑人ID" sql:"varchar(50)"`
	FDeletemark   int     `gorm:"column:F_DeleteMark" form:"f_delete_mark" json:"f_delete_mark" comment:"删除标记0否1是" sql:"int(11)"`
	FSortcode   int     `gorm:"column:F_SortCode" form:"f_sort_code" json:"f_sort_code" comment:"排序码" sql:"int(11)"`
	FModifydate   *time.Time     `gorm:"column:F_ModifyDate" form:"f_modify_date" json:"f_modify_date,omitempty" comment:"编辑日期" sql:"datetime"`
	FCreateusername   string     `gorm:"column:F_CreateUserName" form:"f_create_user_name" json:"f_create_user_name" comment:"创建人" sql:"varchar(50)"`
	FCreatedate   *time.Time     `gorm:"column:F_CreateDate" form:"f_create_date" json:"f_create_date,omitempty" comment:"创建时间" sql:"datetime"`
	FDescription   string     `gorm:"column:F_Description" form:"f_description" json:"f_description" comment:"备注" sql:"varchar(200)"`
}
//TableName
func (m *Basepayroll) TableName() string {
	return "sa_basepayroll"
}
//One,获取一条记录
func (m *Basepayroll) One() (one *Basepayroll, err error) {
	one = &Basepayroll{}
	err = crudOne(m, one)
	return
}
//All，根据条件
func (m *Basepayroll) All(q *PaginationQuery) (list *[]Basepayroll, total uint, err error) {
	list = &[]Basepayroll{}
	total, err = crudAll(m, q, list)
	return
}
//Update，更新
func (m *Basepayroll) Update() (err error) {
	where := Basepayroll{  Pid: m.Pid}
	//m.Id = 0
	
	return crudUpdate(m, where)
}
//Create，新增
func (m *Basepayroll) Create() (err error) {
	//m.Id = 0
    
	return MysqlDB.Create(m).Error
}
//Delete，删除
func (m *Basepayroll) Delete() (err error) {
	//if m.Id == 0 {
	if m==nil {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
