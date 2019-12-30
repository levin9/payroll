package models

import (
	"errors"
	"time"
	
)

var _ = time.Thursday
//GMenuAction
type GMenuAction struct {
	
	Id   uint     `gorm:"column:id;primary_key" form:"id" json:"id" comment:"" sql:"int(10) unsigned,PRI"`
	CreatedAt   *time.Time     `gorm:"column:created_at" form:"created_at" json:"created_at,omitempty" comment:"" sql:"timestamp"`
	UpdatedAt   *time.Time     `gorm:"column:updated_at" form:"updated_at" json:"updated_at,omitempty" comment:"" sql:"timestamp"`
	DeletedAt   *time.Time     `gorm:"column:deleted_at" form:"deleted_at" json:"deleted_at,omitempty" comment:"" sql:"timestamp,MUL"`
	MenuId   string     `gorm:"column:menu_id" form:"menu_id" json:"menu_id" comment:"" sql:"varchar(36),MUL"`
	Code   string     `gorm:"column:code" form:"code" json:"code" comment:"" sql:"varchar(50),MUL"`
	Name   string     `gorm:"column:name" form:"name" json:"name" comment:"" sql:"varchar(50)"`
}
//TableName
func (m *GMenuAction) TableName() string {
	return "g_menu_action"
}
//One,获取一条记录
func (m *GMenuAction) One() (one *GMenuAction, err error) {
	one = &GMenuAction{}
	err = crudOne(m, one)
	return
}
//All，根据条件
func (m *GMenuAction) All(q *PaginationQuery) (list *[]GMenuAction, total uint, err error) {
	list = &[]GMenuAction{}
	total, err = crudAll(m, q, list)
	return
}
//Update，更新
func (m *GMenuAction) Update() (err error) {
	where := GMenuAction{  Id: m.Id}
	//m.Id = 0
	
	return crudUpdate(m, where)
}
//Create，新增
func (m *GMenuAction) Create() (err error) {
	//m.Id = 0
    
	return mysqlDB.Create(m).Error
}
//Delete，删除
func (m *GMenuAction) Delete() (err error) {
	//if m.Id == 0 {
	if m==nil {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
