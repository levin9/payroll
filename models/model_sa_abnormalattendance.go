package models

import (
	"errors"
	"time"
	
)

var _ = time.Thursday
//Abnormalattendance
type Abnormalattendance struct {
	
	Id   string     `gorm:"column:Id;primary_key" form:"id" json:"id" comment:"ID" sql:"varchar(50),PRI"`
	Tenantid   string     `gorm:"column:TenantId" form:"tenant_id" json:"tenant_id" comment:"租户编号" sql:"varchar(50)"`
	Personid   string     `gorm:"column:PersonId" form:"person_id" json:"person_id" comment:"雇员编号" sql:"varchar(50)"`
	Atttype   string     `gorm:"column:AttType" form:"att_type" json:"att_type" comment:"考勤类别" sql:"varchar(50)"`
	Attname   string     `gorm:"column:AttName" form:"att_name" json:"att_name" comment:"类别名称" sql:"varchar(50)"`
	Starttime   *time.Time     `gorm:"column:StartTime" form:"start_time" json:"start_time,omitempty" comment:"开始时间" sql:"datetime"`
	Endtime   *time.Time     `gorm:"column:EndTime" form:"end_time" json:"end_time,omitempty" comment:"结束时间" sql:"datetime"`
	Amount   float32     `gorm:"column:Amount" form:"amount" json:"amount" comment:"数值" sql:"decimal(10,0)"`
	Bizno   string     `gorm:"column:BizNo" form:"biz_no" json:"biz_no" comment:"单号" sql:"varchar(50)"`
	Fee   float32     `gorm:"column:fee" form:"fee" json:"fee" comment:"考勤扣费" sql:"decimal(10,0)"`
	Rulename   string     `gorm:"column:RuleName" form:"rule_name" json:"rule_name" comment:"扣费规则" sql:"varchar(50)"`
	Formulatext   string     `gorm:"column:FormulaText" form:"formula_text" json:"formula_text" comment:"扣费公式" sql:"varchar(50)"`
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
func (m *Abnormalattendance) TableName() string {
	return "sa_abnormalattendance"
}
//One,获取一条记录
func (m *Abnormalattendance) One() (one *Abnormalattendance, err error) {
	one = &Abnormalattendance{}
	err = crudOne(m, one)
	return
}
//All，根据条件
func (m *Abnormalattendance) All(q *PaginationQuery) (list *[]Abnormalattendance, total uint, err error) {
	list = &[]Abnormalattendance{}
	total, err = crudAll(m, q, list)
	return
}
//Update，更新
func (m *Abnormalattendance) Update() (err error) {
	where := Abnormalattendance{  Id: m.Id}
	//m.Id = 0
	
	return crudUpdate(m, where)
}
//Create，新增
func (m *Abnormalattendance) Create() (err error) {
	//m.Id = 0
    
	return mysqlDB.Create(m).Error
}
//Delete，删除
func (m *Abnormalattendance) Delete() (err error) {
	//if m.Id == 0 {
	if m==nil {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
