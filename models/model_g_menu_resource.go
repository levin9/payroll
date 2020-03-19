package models

import (
	"errors"
	"time"
	
)

var _ = time.Thursday
//GMenuResource
type GMenuResource struct {
	
	Id   uint     `gorm:"column:id;primary_key" form:"id" json:"id" comment:"" sql:"int(10) unsigned,PRI"`
	CreatedAt   *time.Time     `gorm:"column:created_at" form:"created_at" json:"created_at,omitempty" comment:"" sql:"timestamp"`
	UpdatedAt   *time.Time     `gorm:"column:updated_at" form:"updated_at" json:"updated_at,omitempty" comment:"" sql:"timestamp"`
	DeletedAt   *time.Time     `gorm:"column:deleted_at" form:"deleted_at" json:"deleted_at,omitempty" comment:"" sql:"timestamp,MUL"`
	MenuId   string     `gorm:"column:menu_id" form:"menu_id" json:"menu_id" comment:"" sql:"varchar(36),MUL"`
	Code   string     `gorm:"column:code" form:"code" json:"code" comment:"" sql:"varchar(50),MUL"`
	Name   string     `gorm:"column:name" form:"name" json:"name" comment:"" sql:"varchar(50)"`
	Method   string     `gorm:"column:method" form:"method" json:"method" comment:"" sql:"varchar(50)"`
	Path   string     `gorm:"column:path" form:"path" json:"path" comment:"" sql:"varchar(255)"`
}
//TableName
func (m *GMenuResource) TableName() string {
	return "g_menu_resource"
}
//One,获取一条记录
func (m *GMenuResource) One() (one *GMenuResource, err error) {
	one = &GMenuResource{}
	err = crudOne(m, one)
	return
}
//All，根据条件
func (m *GMenuResource) All(q *PaginationQuery) (list *[]GMenuResource, total uint, err error) {
	list = &[]GMenuResource{}
	total, err = crudAll(m, q, list)
	return
}
//Update，更新
func (m *GMenuResource) Update() (err error) {
	where := GMenuResource{  Id: m.Id}
	//m.Id = 0
	
	return crudUpdate(m, where)
}
//Create，新增
func (m *GMenuResource) Create() (err error) {
	//m.Id = 0
    
	return MysqlDB.Create(m).Error
}
//Delete，删除
func (m *GMenuResource) Delete() (err error) {
	//if m.Id == 0 {
	if m==nil {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
