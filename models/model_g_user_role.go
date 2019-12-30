package models

import (
	"errors"
	"time"
	
)

var _ = time.Thursday
//GUserRole
type GUserRole struct {
	
	Id   uint     `gorm:"column:id;primary_key" form:"id" json:"id" comment:"" sql:"int(10) unsigned,PRI"`
	CreatedAt   *time.Time     `gorm:"column:created_at" form:"created_at" json:"created_at,omitempty" comment:"" sql:"timestamp"`
	UpdatedAt   *time.Time     `gorm:"column:updated_at" form:"updated_at" json:"updated_at,omitempty" comment:"" sql:"timestamp"`
	DeletedAt   *time.Time     `gorm:"column:deleted_at" form:"deleted_at" json:"deleted_at,omitempty" comment:"" sql:"timestamp,MUL"`
	UserId   string     `gorm:"column:user_id" form:"user_id" json:"user_id" comment:"" sql:"varchar(36),MUL"`
	RoleId   string     `gorm:"column:role_id" form:"role_id" json:"role_id" comment:"" sql:"varchar(36),MUL"`
}
//TableName
func (m *GUserRole) TableName() string {
	return "g_user_role"
}
//One,获取一条记录
func (m *GUserRole) One() (one *GUserRole, err error) {
	one = &GUserRole{}
	err = crudOne(m, one)
	return
}
//All，根据条件
func (m *GUserRole) All(q *PaginationQuery) (list *[]GUserRole, total uint, err error) {
	list = &[]GUserRole{}
	total, err = crudAll(m, q, list)
	return
}
//Update，更新
func (m *GUserRole) Update() (err error) {
	where := GUserRole{  Id: m.Id}
	//m.Id = 0
	
	return crudUpdate(m, where)
}
//Create，新增
func (m *GUserRole) Create() (err error) {
	//m.Id = 0
    
	return mysqlDB.Create(m).Error
}
//Delete，删除
func (m *GUserRole) Delete() (err error) {
	//if m.Id == 0 {
	if m==nil {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
