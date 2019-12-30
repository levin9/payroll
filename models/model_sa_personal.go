package models

import (
	"errors"
	"time"
	
)

var _ = time.Thursday
//Personal
type Personal struct {
	
	Personid   string     `gorm:"column:PersonId;primary_key" form:"person_id" json:"person_id" comment:"雇员编号" sql:"varchar(50),PRI"`
	Tenantid   string     `gorm:"column:TenantId" form:"tenant_id" json:"tenant_id" comment:"租户编号" sql:"varchar(50)"`
	Userid   string     `gorm:"column:UserId" form:"user_id" json:"user_id" comment:"用户编号" sql:"varchar(50)"`
	Realname   string     `gorm:"column:RealName" form:"real_name" json:"real_name" comment:"用户姓名" sql:"varchar(50)"`
	Mobile   string     `gorm:"column:Mobile" form:"mobile" json:"mobile" comment:"手机号" sql:"varchar(59)"`
	Jobno   string     `gorm:"column:JobNo" form:"job_no" json:"job_no" comment:"工号" sql:"varchar(10)"`
	Deptname   string     `gorm:"column:DeptName" form:"dept_name" json:"dept_name" comment:"部门名称" sql:"varchar(200)"`
	Jobname   string     `gorm:"column:JobName" form:"job_name" json:"job_name" comment:"职务" sql:"varchar(50)"`
	Joblevel   string     `gorm:"column:JobLevel" form:"job_level" json:"job_level" comment:"职级" sql:"varchar(50)"`
	Joindate   *time.Time     `gorm:"column:JoinDate" form:"join_date" json:"join_date,omitempty" comment:"入职日期" sql:"datetime"`
	Lastdate   *time.Time     `gorm:"column:LastDate" form:"last_date" json:"last_date,omitempty" comment:"离职日期" sql:"datetime"`
	Regulardate   *time.Time     `gorm:"column:RegularDate" form:"regular_date" json:"regular_date,omitempty" comment:"转正日期" sql:"datetime"`
	Sallaryamount   float32     `gorm:"column:SallaryAmount" form:"sallary_amount" json:"sallary_amount" comment:"月薪" sql:"decimal(10,0)"`
	Currentfreetax   float32     `gorm:"column:CurrentFreeTax" form:"current_free_tax" json:"current_free_tax" comment:"免税金额" sql:"decimal(10,4)"`
	Dingid   string     `gorm:"column:DingId" form:"ding_id" json:"ding_id" comment:"钉钉编号" sql:"varchar(50)"`
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
func (m *Personal) TableName() string {
	return "sa_personal"
}
//One,获取一条记录
func (m *Personal) One() (one *Personal, err error) {
	one = &Personal{}
	err = crudOne(m, one)
	return
}
//All，根据条件
func (m *Personal) All(q *PaginationQuery) (list *[]Personal, total uint, err error) {
	list = &[]Personal{}
	total, err = crudAll(m, q, list)
	return
}
//Update，更新
func (m *Personal) Update() (err error) {
	where := Personal{  Personid: m.Personid}
	//m.Id = 0
	
	return crudUpdate(m, where)
}
//Create，新增
func (m *Personal) Create() (err error) {
	//m.Id = 0
    
	return mysqlDB.Create(m).Error
}
//Delete，删除
func (m *Personal) Delete() (err error) {
	//if m.Id == 0 {
	if m==nil {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
