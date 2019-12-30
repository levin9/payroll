package models

import (
	"errors"
	"time"
	
)

var _ = time.Thursday
//Payroll
type Payroll struct {
	
	Payrollid   string     `gorm:"column:PayrollId;primary_key" form:"payroll_id" json:"payroll_id" comment:"薪资项" sql:"varchar(50),PRI"`
	Standardname   string     `gorm:"column:StandardName" form:"standard_name" json:"standard_name" comment:"标准名称" sql:"varchar(50)"`
	Fieldname   string     `gorm:"column:FieldName" form:"field_name" json:"field_name" comment:"字段名称" sql:"varchar(50)"`
	Standardfieldname   string     `gorm:"column:StandardFieldName" form:"standard_field_name" json:"standard_field_name" comment:"实际字段名称" sql:"varchar(50)"`
	Defaultformula   string     `gorm:"column:DefaultFormula" form:"default_formula" json:"default_formula" comment:"日薪公式" sql:"varchar(200)"`
	Formularemark   string     `gorm:"column:FormulaRemark" form:"formula_remark" json:"formula_remark" comment:"公式备注" sql:"varchar(500)"`
	Sqlformula   string     `gorm:"column:SqlFormula" form:"sql_formula" json:"sql_formula" comment:"SQL公式" sql:"varchar(200)"`
	Ismustused   int     `gorm:"column:IsMustUsed" form:"is_must_used" json:"is_must_used" comment:"必须启用" sql:"int(11)"`
	Isaddition   int     `gorm:"column:IsAddition" form:"is_addition" json:"is_addition" comment:"加减项" sql:"int(11)"`
	Ishandcalc   int     `gorm:"column:IsHandCalc" form:"is_hand_calc" json:"is_hand_calc" comment:"手工计算项" sql:"int(11)"`
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
func (m *Payroll) TableName() string {
	return "sa_payroll"
}
//One,获取一条记录
func (m *Payroll) One() (one *Payroll, err error) {
	one = &Payroll{}
	err = crudOne(m, one)
	return
}
//All，根据条件
func (m *Payroll) All(q *PaginationQuery) (list *[]Payroll, total uint, err error) {
	list = &[]Payroll{}
	total, err = crudAll(m, q, list)
	return
}
//Update，更新
func (m *Payroll) Update() (err error) {
	where := Payroll{  Payrollid: m.Payrollid}
	//m.Id = 0
	
	return crudUpdate(m, where)
}
//Create，新增
func (m *Payroll) Create() (err error) {
	//m.Id = 0
    
	return mysqlDB.Create(m).Error
}
//Delete，删除
func (m *Payroll) Delete() (err error) {
	//if m.Id == 0 {
	if m==nil {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
