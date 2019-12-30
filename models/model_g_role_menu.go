package models

import (
	"errors"
	"time"
	
)

var _ = time.Thursday
//GRoleMenu
type GRoleMenu struct {
	
	Id   uint     `gorm:"column:id;primary_key" form:"id" json:"id" comment:"" sql:"int(10) unsigned,PRI"`
	CreatedAt   *time.Time     `gorm:"column:created_at" form:"created_at" json:"created_at,omitempty" comment:"" sql:"timestamp"`
	UpdatedAt   *time.Time     `gorm:"column:updated_at" form:"updated_at" json:"updated_at,omitempty" comment:"" sql:"timestamp"`
	DeletedAt   *time.Time     `gorm:"column:deleted_at" form:"deleted_at" json:"deleted_at,omitempty" comment:"" sql:"timestamp,MUL"`
	RoleId   string     `gorm:"column:role_id" form:"role_id" json:"role_id" comment:"" sql:"varchar(36),MUL"`
	MenuId   string     `gorm:"column:menu_id" form:"menu_id" json:"menu_id" comment:"" sql:"varchar(36),MUL"`
	Action   string     `gorm:"column:action" form:"action" json:"action" comment:"" sql:"varchar(2048)"`
	Resource   string     `gorm:"column:resource" form:"resource" json:"resource" comment:"" sql:"varchar(2048)"`
}
//TableName
func (m *GRoleMenu) TableName() string {
	return "g_role_menu"
}
//One,获取一条记录
func (m *GRoleMenu) One() (one *GRoleMenu, err error) {
	one = &GRoleMenu{}
	err = crudOne(m, one)
	return
}
//All，根据条件
func (m *GRoleMenu) All(q *PaginationQuery) (list *[]GRoleMenu, total uint, err error) {
	list = &[]GRoleMenu{}
	total, err = crudAll(m, q, list)
	return
}
//Update，更新
func (m *GRoleMenu) Update() (err error) {
	where := GRoleMenu{  Id: m.Id}
	//m.Id = 0
	
	return crudUpdate(m, where)
}
//Create，新增
func (m *GRoleMenu) Create() (err error) {
	//m.Id = 0
    
	return mysqlDB.Create(m).Error
}
//Delete，删除
func (m *GRoleMenu) Delete() (err error) {
	//if m.Id == 0 {
	if m==nil {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
