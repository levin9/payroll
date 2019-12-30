package models

import (
	"errors"
	"time"
	
)

var _ = time.Thursday
//Attcategory
type Attcategory struct {
	
	Id   string     `gorm:"column:Id;primary_key" form:"id" json:"id" comment:"Id" sql:"varchar(50),PRI"`
	Holidaycode   string     `gorm:"column:HolidayCode" form:"holiday_code" json:"holiday_code" comment:"编号" sql:"varchar(20)"`
	Holidayname   string     `gorm:"column:HolidayName" form:"holiday_name" json:"holiday_name" comment:"名称" sql:"varchar(50)"`
	Calctype   int     `gorm:"column:CalcType" form:"calc_type" json:"calc_type" comment:"计算类别" sql:"int(255)"`
	Isspeccalc   int     `gorm:"column:IsSpecCalc" form:"is_spec_calc" json:"is_spec_calc" comment:"特殊计算规则" sql:"int(1)"`
	Isonslip   int     `gorm:"column:IsOnSlip" form:"is_on_slip" json:"is_on_slip" comment:"是否显示在工资单上" sql:"int(1)"`
	Attfieldname   string     `gorm:"column:AttFieldName" form:"att_field_name" json:"att_field_name" comment:"考勤存储的字段" sql:"varchar(50)"`
	Fieldname   string     `gorm:"column:FieldName" form:"field_name" json:"field_name" comment:"存储字段" sql:"varchar(50)"`
	Formula   string     `gorm:"column:Formula" form:"formula" json:"formula" comment:"公式" sql:"varchar(200)"`
	Formularemark   string     `gorm:"column:FormulaRemark" form:"formula_remark" json:"formula_remark" comment:"公式备注" sql:"varchar(200)"`
	Sqlformula   string     `gorm:"column:SqlFormula" form:"sql_formula" json:"sql_formula" comment:"SQL公式，没有用的字段" sql:"varchar(255)"`
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
func (m *Attcategory) TableName() string {
	return "sa_attcategory"
}
//One,获取一条记录
func (m *Attcategory) One() (one *Attcategory, err error) {
	one = &Attcategory{}
	err = crudOne(m, one)
	return
}
//All，根据条件
func (m *Attcategory) All(q *PaginationQuery) (list *[]Attcategory, total uint, err error) {
	list = &[]Attcategory{}
	total, err = crudAll(m, q, list)
	return
}
//Update，更新
func (m *Attcategory) Update() (err error) {
	where := Attcategory{  Id: m.Id}
	//m.Id = 0
	
	return crudUpdate(m, where)
}
//Create，新增
func (m *Attcategory) Create() (err error) {
	//m.Id = 0
    
	return mysqlDB.Create(m).Error
}
//Delete，删除
func (m *Attcategory) Delete() (err error) {
	//if m.Id == 0 {
	if m==nil {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
