package models

import (
	"errors"
	"time"
	
)

var _ = time.Thursday
//Basemonth
type Basemonth struct {
	
	Id   string     `gorm:"column:ID;primary_key" form:"id" json:"id" comment:"编号" sql:"varchar(50),PRI"`
	Monthid   string     `gorm:"column:MonthId" form:"month_id" json:"month_id" comment:"月份编号" sql:"varchar(10)"`
	Firstcode   string     `gorm:"column:FirstCode" form:"first_code" json:"first_code" comment:"年份" sql:"varchar(5)"`
	Secondcode   string     `gorm:"column:SecondCode" form:"second_code" json:"second_code" comment:"月份" sql:"varchar(5)"`
	Monthname   string     `gorm:"column:MonthName" form:"month_name" json:"month_name" comment:"名称" sql:"varchar(50)"`
	Tenantid   string     `gorm:"column:TenantID" form:"tenant_id" json:"tenant_id" comment:"租户编号" sql:"varchar(50)"`
	Startdate   *time.Time     `gorm:"column:StartDate" form:"start_date" json:"start_date,omitempty" comment:"开始时间" sql:"datetime"`
	Enddate   *time.Time     `gorm:"column:EndDate" form:"end_date" json:"end_date,omitempty" comment:"结束时间" sql:"datetime"`
	Shebaoduedate   *time.Time     `gorm:"column:SheBaoDueDate" form:"she_bao_due_date" json:"she_bao_due_date,omitempty" comment:"当月购买社保的日期" sql:"datetime"`
	Workdayqty   int     `gorm:"column:WorkdayQty" form:"workday_qty" json:"workday_qty" comment:"工作天数" sql:"int(11)"`
	State   int     `gorm:"column:State" form:"state" json:"state" comment:"状态" sql:"int(11)"`
	Statename   string     `gorm:"column:StateName" form:"state_name" json:"state_name" comment:"状态Name" sql:"varchar(50)"`
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
func (m *Basemonth) TableName() string {
	return "sa_basemonth"
}
//One,获取一条记录
func (m *Basemonth) One() (one *Basemonth, err error) {
	one = &Basemonth{}
	err = crudOne(m, one)
	return
}
//All，根据条件
func (m *Basemonth) All(q *PaginationQuery) (list *[]Basemonth, total uint, err error) {
	list = &[]Basemonth{}
	total, err = crudAll(m, q, list)
	return
}
//Update，更新
func (m *Basemonth) Update() (err error) {
	where := Basemonth{  Id: m.Id}
	//m.Id = 0
	
	return crudUpdate(m, where)
}
//Create，新增
func (m *Basemonth) Create() (err error) {
	//m.Id = 0
    
	return mysqlDB.Create(m).Error
}
//Delete，删除
func (m *Basemonth) Delete() (err error) {
	//if m.Id == 0 {
	if m==nil {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
