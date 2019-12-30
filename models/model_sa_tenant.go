package models

import (
	"errors"
	"time"
	
)

var _ = time.Thursday
//Tenant
type Tenant struct {
	
	Tenantid   string     `gorm:"column:TenantID;primary_key" form:"tenant_id" json:"tenant_id" comment:"租户编号" sql:"varchar(50),PRI"`
	Fullname   string     `gorm:"column:FullName" form:"full_name" json:"full_name" comment:"全称" sql:"varchar(200)"`
	Simplename   string     `gorm:"column:SimpleName" form:"simple_name" json:"simple_name" comment:"简称" sql:"varchar(10)"`
	Companyid   string     `gorm:"column:CompanyID" form:"company_id" json:"company_id" comment:"公司系统编号" sql:"varchar(50)"`
	Contactname   string     `gorm:"column:ContactName" form:"contact_name" json:"contact_name" comment:"联系人" sql:"varchar(50)"`
	Contacttitle   string     `gorm:"column:ContactTitle" form:"contact_title" json:"contact_title" comment:"称谓" sql:"varchar(50)"`
	Contactemail   string     `gorm:"column:ContactEmail" form:"contact_email" json:"contact_email" comment:"联系人邮件" sql:"varchar(200)"`
	Contactphone   string     `gorm:"column:ContactPhone" form:"contact_phone" json:"contact_phone" comment:"联系电话" sql:"varchar(20)"`
	Contactmobile   string     `gorm:"column:ContactMobile" form:"contact_mobile" json:"contact_mobile" comment:"联系人手机" sql:"varchar(20)"`
	Payrollday   int     `gorm:"column:PayrollDay" form:"payroll_day" json:"payroll_day" comment:"薪资周期" sql:"int(10)"`
	Hrsource   string     `gorm:"column:HRSource" form:"hrsource" json:"hrsource" comment:"薪资周期" sql:"varchar(10)"`
	Attsource   string     `gorm:"column:AttSource" form:"att_source" json:"att_source" comment:"人事来源" sql:"varchar(10)"`
	Buyshebao   string     `gorm:"column:BuySheBao" form:"buy_she_bao" json:"buy_she_bao" comment:"社保设置" sql:"varchar(10)"`
	Dingid   string     `gorm:"column:DingId" form:"ding_id" json:"ding_id" comment:"钉钉ID" sql:"varchar(50)"`
	Corpid   string     `gorm:"column:Corpid" form:"corpid" json:"corpid" comment:"Corpid" sql:"varchar(50)"`
	Corpsecret   string     `gorm:"column:Corpsecret" form:"corpsecret" json:"corpsecret" comment:"Corpsecret" sql:"varchar(200)"`
	State   int     `gorm:"column:State" form:"state" json:"state" comment:"租用状态" sql:"int(11)"`
	Shebaorate   float32     `gorm:"column:SheBaoRate" form:"she_bao_rate" json:"she_bao_rate" comment:"社保费率" sql:"decimal(10,2)"`
	Houserate   float32     `gorm:"column:HouseRate" form:"house_rate" json:"house_rate" comment:"公积金费率" sql:"decimal(10,2)"`
	Tenantendtime   *time.Time     `gorm:"column:TenantEndTime" form:"tenant_end_time" json:"tenant_end_time,omitempty" comment:"结束时间" sql:"datetime"`
	Tenantstarttime   *time.Time     `gorm:"column:TenantStartTime" form:"tenant_start_time" json:"tenant_start_time,omitempty" comment:"开始时间" sql:"datetime"`
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
func (m *Tenant) TableName() string {
	return "sa_tenant"
}
//One,获取一条记录
func (m *Tenant) One() (one *Tenant, err error) {
	one = &Tenant{}
	err = crudOne(m, one)
	return
}
//All，根据条件
func (m *Tenant) All(q *PaginationQuery) (list *[]Tenant, total uint, err error) {
	list = &[]Tenant{}
	total, err = crudAll(m, q, list)
	return
}
//Update，更新
func (m *Tenant) Update() (err error) {
	where := Tenant{  Tenantid: m.Tenantid}
	//m.Id = 0
	
	return crudUpdate(m, where)
}
//Create，新增
func (m *Tenant) Create() (err error) {
	//m.Id = 0
    
	return mysqlDB.Create(m).Error
}
//Delete，删除
func (m *Tenant) Delete() (err error) {
	//if m.Id == 0 {
	if m==nil {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
