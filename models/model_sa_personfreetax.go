package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

//Personfreetax
type Personfreetax struct {
	Id              string     `gorm:"column:Id;primary_key" form:"id" json:"id" comment:"编号" sql:"varchar(50),PRI"`
	Personid        string     `gorm:"column:PersonId" form:"person_id" json:"person_id" comment:"员工编号" sql:"varchar(50)"`
	Tenantid        string     `gorm:"column:TenantId" form:"tenant_id" json:"tenant_id" comment:"租户编号" sql:"varchar(50)"`
	Znjy            float32    `gorm:"column:ZNJY" form:"znjy" json:"znjy" comment:"子女教育" sql:"decimal(10,0)"`
	Jxjy            float32    `gorm:"column:JXJY" form:"jxjy" json:"jxjy" comment:"继续教育" sql:"decimal(10,0)"`
	Zfdk            float32    `gorm:"column:ZFDK" form:"zfdk" json:"zfdk" comment:"住房贷款" sql:"decimal(10,0)"`
	Zfzj            float32    `gorm:"column:ZFZJ" form:"zfzj" json:"zfzj" comment:"住房租金" sql:"decimal(10,0)"`
	Sylr            float32    `gorm:"column:SYLR" form:"sylr" json:"sylr" comment:"赡养老人" sql:"decimal(10,0)"`
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
func (m *Personfreetax) TableName() string {
	return "sa_personfreetax"
}

//One,获取一条记录
func (m *Personfreetax) One() (one *Personfreetax, err error) {
	one = &Personfreetax{}
	err = crudOne(m, one)
	return
}

//All，根据条件
func (m *Personfreetax) All(q *PaginationQuery) (list *[]Personfreetax, total uint, err error) {
	list = &[]Personfreetax{}
	total, err = crudAll(m, q, list)
	return
}

//Update，更新
func (m *Personfreetax) Update() (err error) {
	where := Personfreetax{Id: m.Id}
	//m.Id = 0

	return crudUpdate(m, where)
}

//Create，新增
func (m *Personfreetax) Create() (err error) {
	//m.Id = 0
	//m.FCreatedate = time.Now().Location()

	return MysqlDB.Create(m).Error
}

//Delete，删除
func (m *Personfreetax) Delete() (err error) {
	//if m.Id == 0 {
	if m == nil {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
