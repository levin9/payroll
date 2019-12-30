package models

import (
	"errors"
	"time"
	
)

var _ = time.Thursday
//Operatelog
type Operatelog struct {
	
	Id   string     `gorm:"column:Id;primary_key" form:"id" json:"id" comment:"编号" sql:"varchar(50),PRI"`
	Tenantid   string     `gorm:"column:TenantId" form:"tenant_id" json:"tenant_id" comment:"租户编号" sql:"varchar(50)"`
	Monthid   string     `gorm:"column:MonthId" form:"month_id" json:"month_id" comment:"月份" sql:"varchar(50)"`
	Operatename   string     `gorm:"column:OperateName" form:"operate_name" json:"operate_name" comment:"操作名称" sql:"varchar(50)"`
	Operateindex   int     `gorm:"column:OperateIndex" form:"operate_index" json:"operate_index" comment:"操作序号" sql:"int(3)"`
	Statename   string     `gorm:"column:StateName" form:"state_name" json:"state_name" comment:"状态名称" sql:"varchar(50)"`
	State   int     `gorm:"column:State" form:"state" json:"state" comment:"状态" sql:"int(11)"`
	Num   float32     `gorm:"column:Num" form:"num" json:"num" comment:"耗时" sql:"decimal(10,0)"`
	Timestampnum   float32     `gorm:"column:TimeStampNum" form:"time_stamp_num" json:"time_stamp_num" comment:"时间戳" sql:"decimal(18,0)"`
	Chnghistory   string     `gorm:"column:ChngHistory" form:"chng_history" json:"chng_history" comment:"改变值" sql:"varchar(2000)"`
	FEnabledmark   int     `gorm:"column:F_EnabledMark" form:"f_enabled_mark" json:"f_enabled_mark" comment:"有效标志0否1是" sql:"int(11)"`
	FCreateuserid   string     `gorm:"column:F_CreateUserId" form:"f_create_user_id" json:"f_create_user_id" comment:"创建人ID" sql:"varchar(50)"`
	FModifyusername   string     `gorm:"column:F_ModifyUserName" form:"f_modify_user_name" json:"f_modify_user_name" comment:"编辑人" sql:"varchar(50)"`
	FModifyuserid   string     `gorm:"column:F_ModifyUserId" form:"f_modify_user_id" json:"f_modify_user_id" comment:"编辑人ID" sql:"varchar(50)"`
	FDeletemark   int     `gorm:"column:F_DeleteMark" form:"f_delete_mark" json:"f_delete_mark" comment:"删除标记0否1是" sql:"int(11)"`
	FSortcode   int     `gorm:"column:F_SortCode" form:"f_sort_code" json:"f_sort_code" comment:"排序码" sql:"int(11)"`
	FModifydate   *time.Time     `gorm:"column:F_ModifyDate" form:"f_modify_date" json:"f_modify_date,omitempty" comment:"编辑日期" sql:"datetime"`
	FCreateusername   string     `gorm:"column:F_CreateUserName" form:"f_create_user_name" json:"f_create_user_name" comment:"创建人" sql:"varchar(50)"`
	FCreatedate   *time.Time     `gorm:"column:F_CreateDate" form:"f_create_date" json:"f_create_date,omitempty" comment:"创建时间" sql:"datetime"`
	FDescription   string     `gorm:"column:F_Description" form:"f_description" json:"f_description" comment:"备注" sql:"varchar(2000)"`
}
//TableName
func (m *Operatelog) TableName() string {
	return "sa_operatelog"
}
//One,获取一条记录
func (m *Operatelog) One() (one *Operatelog, err error) {
	one = &Operatelog{}
	err = crudOne(m, one)
	return
}
//All，根据条件
func (m *Operatelog) All(q *PaginationQuery) (list *[]Operatelog, total uint, err error) {
	list = &[]Operatelog{}
	total, err = crudAll(m, q, list)
	return
}
//Update，更新
func (m *Operatelog) Update() (err error) {
	where := Operatelog{  Id: m.Id}
	//m.Id = 0
	
	return crudUpdate(m, where)
}
//Create，新增
func (m *Operatelog) Create() (err error) {
	//m.Id = 0
    
	return mysqlDB.Create(m).Error
}
//Delete，删除
func (m *Operatelog) Delete() (err error) {
	//if m.Id == 0 {
	if m==nil {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
