package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

//Tenantadjustitem
type Tenantadjustitem struct {
	Id              string     `gorm:"column:Id;primary_key" form:"id" json:"id" comment:"编号" sql:"varchar(50),PRI"`
	Tenantid        string     `gorm:"column:TenantId" form:"tenant_id" json:"tenant_id" comment:"商户编号" sql:"varchar(50)"`
	Personid        string     `gorm:"column:PersonId" form:"person_id" json:"person_id" comment:"员工Id" sql:"varchar(50)"`
	Personname      string     `gorm:"column:PersonName" form:"person_name" json:"person_name" comment:"员工姓名" sql:"varchar(50)"`
	Deptname        string     `gorm:"column:DeptName" form:"dept_name" json:"dept_name" comment:"部门名称" sql:"varchar(50)"`
	Monthid         string     `gorm:"column:MonthId" form:"month_id" json:"month_id" comment:"月份编号" sql:"varchar(50)"`
	Amount          float32    `gorm:"column:Amount" form:"amount" json:"amount" comment:"调整金额" sql:"decimal(50,0)"`
	Adjestremark    string     `gorm:"column:AdjestRemark" form:"adjest_remark" json:"adjest_remark" comment:"原因" sql:"varchar(500)"`
	FEnabledmark    int        `gorm:"column:F_EnabledMark" form:"f_enabled_mark" json:"f_enabled_mark" comment:"有效标志0否1是" sql:"int(11)"`
	FCreateuserid   string     `gorm:"column:F_CreateUserId" form:"f_create_user_id" json:"f_create_user_id" comment:"创建人ID" sql:"varchar(50)"`
	FModifyusername string     `gorm:"column:F_ModifyUserName" form:"f_modify_user_name" json:"f_modify_user_name" comment:"编辑人" sql:"varchar(50)"`
	FModifyuserid   string     `gorm:"column:F_ModifyUserId" form:"f_modify_user_id" json:"f_modify_user_id" comment:"编辑人ID" sql:"varchar(50)"`
	FDeletemark     int        `gorm:"column:F_DeleteMark" form:"f_delete_mark" json:"f_delete_mark" comment:"删除标记0否1是" sql:"int(11)"`
	FSortcode       int        `gorm:"column:F_SortCode" form:"f_sort_code" json:"f_sort_code" comment:"排序码" sql:"int(11)"`
	FModifydate     *time.Time `gorm:"column:F_ModifyDate" form:"f_modify_date" json:"f_modify_date,omitempty" comment:"编辑日期" sql:"datetime"`
	FCreateusername string     `gorm:"column:F_CreateUserName" form:"f_create_user_name" json:"f_create_user_name" comment:"创建人" sql:"varchar(50)"`
	FCreatedate     *time.Time `gorm:"column:F_CreateDate" form:"f_create_date" json:"f_create_date,omitempty" comment:"创建时间" sql:"datetime"`
	FDescription    string     `gorm:"column:F_Description" form:"f_description" json:"f_description" comment:"备注" sql:"varchar(200)"`
}

//TableName
func (m *Tenantadjustitem) TableName() string {
	return "sa_tenantadjustitem"
}

//One,获取一条记录
func (m *Tenantadjustitem) One() (one *Tenantadjustitem, err error) {
	one = &Tenantadjustitem{}
	err = crudOne(m, one)
	return
}

//All，根据条件
func (m *Tenantadjustitem) All(q *PaginationQuery) (list *[]Tenantadjustitem, total uint, err error) {
	list = &[]Tenantadjustitem{}
	total, err = crudAll(m, q, list)
	return
}
func (m *Tenantadjustitem) FindAll(q *PaginationQuery) (list *[]Tenantadjustitem, total uint, err error) {
	list = &[]Tenantadjustitem{}
	total, err = crudGetAll(m, q, list)
	return
}
func (m *Tenantadjustitem) FindByMap(map[string]interface{}) (list *[]Tenantadjustitem, err error) {
	list = &[]Tenantadjustitem{}

	return
}

//Update，更新
func (m *Tenantadjustitem) Update() (err error) {
	where := Tenantadjustitem{Id: m.Id}
	//m.Id = 0

	return crudUpdate(m, where)
}

//Create，新增
func (m *Tenantadjustitem) Create() (err error) {
	//m.Id = 0

	return MysqlDB.Create(m).Error
}

//Delete，删除
func (m *Tenantadjustitem) Delete() (err error) {
	//if m.Id == 0 {
	if m == nil {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
