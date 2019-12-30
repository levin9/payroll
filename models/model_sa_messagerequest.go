package models

import (
	"errors"
	"time"
	
)

var _ = time.Thursday
//Messagerequest
type Messagerequest struct {
	
	Id   string     `gorm:"column:Id;primary_key" form:"id" json:"id" comment:"编号" sql:"varchar(50),PRI"`
	Rquestid   string     `gorm:"column:RquestId" form:"rquest_id" json:"rquest_id" comment:"请求编号" sql:"varchar(50)"`
	Tenantid   string     `gorm:"column:TenantId" form:"tenant_id" json:"tenant_id" comment:"商户编号" sql:"varchar(50)"`
	Requserid   string     `gorm:"column:ReqUserId" form:"req_user_id" json:"req_user_id" comment:"请求用户ID" sql:"varchar(200)"`
	Requsername   string     `gorm:"column:ReqUserName" form:"req_user_name" json:"req_user_name" comment:"请求用户姓名" sql:"varchar(200)"`
	Recids   string     `gorm:"column:RecIds" form:"rec_ids" json:"rec_ids" comment:"接收Id" sql:"varchar(500)"`
	Recmobileids   string     `gorm:"column:RecMobileIds" form:"rec_mobile_ids" json:"rec_mobile_ids" comment:"接收者姓名" sql:"varchar(50)"`
	Textbody   string     `gorm:"column:TextBody" form:"text_body" json:"text_body" comment:"消息正文" sql:"varchar(200)"`
	Plantime   *time.Time     `gorm:"column:PlanTime" form:"plan_time" json:"plan_time,omitempty" comment:"计划时间" sql:"datetime(4)"`
	State   int     `gorm:"column:State" form:"state" json:"state" comment:"计划状态" sql:"int(11)"`
	Trynum   int     `gorm:"column:TryNum" form:"try_num" json:"try_num" comment:"尝试次数" sql:"int(4)"`
	FEnabledmark   int     `gorm:"column:F_EnabledMark" form:"f_enabled_mark" json:"f_enabled_mark" comment:"有效标志0否1是" sql:"int(11)"`
	FCreateuserid   string     `gorm:"column:F_CreateUserId" form:"f_create_user_id" json:"f_create_user_id" comment:"创建人ID" sql:"varchar(50)"`
	FModifyusername   string     `gorm:"column:F_ModifyUserName" form:"f_modify_user_name" json:"f_modify_user_name" comment:"编辑人" sql:"varchar(50)"`
	FModifyuserid   string     `gorm:"column:F_ModifyUserId" form:"f_modify_user_id" json:"f_modify_user_id" comment:"编辑人ID" sql:"varchar(50)"`
	FDeletemark   int     `gorm:"column:F_DeleteMark" form:"f_delete_mark" json:"f_delete_mark" comment:"删除标记0否1是" sql:"int(11)"`
	FSortcode   int     `gorm:"column:F_SortCode" form:"f_sort_code" json:"f_sort_code" comment:"排序码" sql:"int(11)"`
	FModifydate   *time.Time     `gorm:"column:F_ModifyDate" form:"f_modify_date" json:"f_modify_date,omitempty" comment:"编辑日期" sql:"datetime(4)"`
	FCreateusername   string     `gorm:"column:F_CreateUserName" form:"f_create_user_name" json:"f_create_user_name" comment:"创建人" sql:"varchar(50)"`
	FCreatedate   *time.Time     `gorm:"column:F_CreateDate" form:"f_create_date" json:"f_create_date,omitempty" comment:"创建时间" sql:"datetime(4)"`
	FDescription   string     `gorm:"column:F_Description" form:"f_description" json:"f_description" comment:"备注" sql:"varchar(50)"`
}
//TableName
func (m *Messagerequest) TableName() string {
	return "sa_messagerequest"
}
//One,获取一条记录
func (m *Messagerequest) One() (one *Messagerequest, err error) {
	one = &Messagerequest{}
	err = crudOne(m, one)
	return
}
//All，根据条件
func (m *Messagerequest) All(q *PaginationQuery) (list *[]Messagerequest, total uint, err error) {
	list = &[]Messagerequest{}
	total, err = crudAll(m, q, list)
	return
}
//Update，更新
func (m *Messagerequest) Update() (err error) {
	where := Messagerequest{  Id: m.Id}
	//m.Id = 0
	
	return crudUpdate(m, where)
}
//Create，新增
func (m *Messagerequest) Create() (err error) {
	//m.Id = 0
    
	return mysqlDB.Create(m).Error
}
//Delete，删除
func (m *Messagerequest) Delete() (err error) {
	//if m.Id == 0 {
	if m==nil {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
