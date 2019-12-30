package models

import (
	"errors"
	"time"
	
)

var _ = time.Thursday
//Bankaccount
type Bankaccount struct {
	
	Id   string     `gorm:"column:Id;primary_key" form:"id" json:"id" comment:"编号" sql:"varchar(50),PRI"`
	Personid   string     `gorm:"column:PersonId" form:"person_id" json:"person_id" comment:"员工编号" sql:"varchar(50)"`
	Itemindex   int     `gorm:"column:ItemIndex" form:"item_index" json:"item_index" comment:"序号" sql:"int(11)"`
	Bankname   string     `gorm:"column:BankName" form:"bank_name" json:"bank_name" comment:"开户行" sql:"varchar(50)"`
	Bankbranch   string     `gorm:"column:BankBranch" form:"bank_branch" json:"bank_branch" comment:"开户分行" sql:"varchar(50)"`
	Accountno   string     `gorm:"column:AccountNo" form:"account_no" json:"account_no" comment:"账号" sql:"varchar(50)"`
	Accountname   string     `gorm:"column:AccountName" form:"account_name" json:"account_name" comment:"账号姓名" sql:"varchar(50)"`
	FEnabledmark   int     `gorm:"column:F_EnabledMark" form:"f_enabled_mark" json:"f_enabled_mark" comment:"有效标志0否1是" sql:"int(11)"`
	FCreateuserid   string     `gorm:"column:F_CreateUserId" form:"f_create_user_id" json:"f_create_user_id" comment:"创建人ID" sql:"varchar(50)"`
	FModifyusername   string     `gorm:"column:F_ModifyUserName" form:"f_modify_user_name" json:"f_modify_user_name" comment:"编辑人" sql:"varchar(50)"`
	FModifyuserid   string     `gorm:"column:F_ModifyUserId" form:"f_modify_user_id" json:"f_modify_user_id" comment:"编辑人ID" sql:"varchar(50)"`
	FDeletemark   int     `gorm:"column:F_DeleteMark" form:"f_delete_mark" json:"f_delete_mark" comment:"删除标记0否1是" sql:"int(11)"`
	FSortcode   int     `gorm:"column:F_SortCode" form:"f_sort_code" json:"f_sort_code" comment:"排序码" sql:"int(11)"`
	FModifydate   *time.Time     `gorm:"column:F_ModifyDate" form:"f_modify_date" json:"f_modify_date,omitempty" comment:"编辑日期" sql:"datetime"`
	FCreateusername   string     `gorm:"column:F_CreateUserName" form:"f_create_user_name" json:"f_create_user_name" comment:"创建人" sql:"varchar(50)"`
	FCreatedate   *time.Time     `gorm:"column:F_CreateDate" form:"f_create_date" json:"f_create_date,omitempty" comment:"创建时间" sql:"datetime"`
	FDescription   string     `gorm:"column:F_Description" form:"f_description" json:"f_description" comment:"备注" sql:"varchar(100)"`
}
//TableName
func (m *Bankaccount) TableName() string {
	return "sa_bankaccount"
}
//One,获取一条记录
func (m *Bankaccount) One() (one *Bankaccount, err error) {
	one = &Bankaccount{}
	err = crudOne(m, one)
	return
}
//All，根据条件
func (m *Bankaccount) All(q *PaginationQuery) (list *[]Bankaccount, total uint, err error) {
	list = &[]Bankaccount{}
	total, err = crudAll(m, q, list)
	return
}
//Update，更新
func (m *Bankaccount) Update() (err error) {
	where := Bankaccount{  Id: m.Id}
	//m.Id = 0
	
	return crudUpdate(m, where)
}
//Create，新增
func (m *Bankaccount) Create() (err error) {
	//m.Id = 0
    
	return mysqlDB.Create(m).Error
}
//Delete，删除
func (m *Bankaccount) Delete() (err error) {
	//if m.Id == 0 {
	if m==nil {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
