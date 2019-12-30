package models

import (
	"errors"
	"time"
	
)

var _ = time.Thursday
//GMenu
type GMenu struct {
	
	Id   uint     `gorm:"column:id;primary_key" form:"id" json:"id" comment:"" sql:"int(10) unsigned,PRI"`
	CreatedAt   *time.Time     `gorm:"column:created_at" form:"created_at" json:"created_at,omitempty" comment:"" sql:"timestamp"`
	UpdatedAt   *time.Time     `gorm:"column:updated_at" form:"updated_at" json:"updated_at,omitempty" comment:"" sql:"timestamp"`
	DeletedAt   *time.Time     `gorm:"column:deleted_at" form:"deleted_at" json:"deleted_at,omitempty" comment:"" sql:"timestamp,MUL"`
	RecordId   string     `gorm:"column:record_id" form:"record_id" json:"record_id" comment:"" sql:"varchar(36),MUL"`
	Name   string     `gorm:"column:name" form:"name" json:"name" comment:"" sql:"varchar(50),MUL"`
	Sequence   int     `gorm:"column:sequence" form:"sequence" json:"sequence" comment:"" sql:"int(11),MUL"`
	Icon   string     `gorm:"column:icon" form:"icon" json:"icon" comment:"" sql:"varchar(255)"`
	Router   string     `gorm:"column:router" form:"router" json:"router" comment:"" sql:"varchar(255)"`
	Hidden   int     `gorm:"column:hidden" form:"hidden" json:"hidden" comment:"" sql:"int(11),MUL"`
	ParentId   string     `gorm:"column:parent_id" form:"parent_id" json:"parent_id" comment:"" sql:"varchar(36),MUL"`
	ParentPath   string     `gorm:"column:parent_path" form:"parent_path" json:"parent_path" comment:"" sql:"varchar(518),MUL"`
	Creator   string     `gorm:"column:creator" form:"creator" json:"creator" comment:"" sql:"varchar(36)"`
}
//TableName
func (m *GMenu) TableName() string {
	return "g_menu"
}
//One,获取一条记录
func (m *GMenu) One() (one *GMenu, err error) {
	one = &GMenu{}
	err = crudOne(m, one)
	return
}
//All，根据条件
func (m *GMenu) All(q *PaginationQuery) (list *[]GMenu, total uint, err error) {
	list = &[]GMenu{}
	total, err = crudAll(m, q, list)
	return
}
//Update，更新
func (m *GMenu) Update() (err error) {
	where := GMenu{  Id: m.Id}
	//m.Id = 0
	
	return crudUpdate(m, where)
}
//Create，新增
func (m *GMenu) Create() (err error) {
	//m.Id = 0
    
	return mysqlDB.Create(m).Error
}
//Delete，删除
func (m *GMenu) Delete() (err error) {
	//if m.Id == 0 {
	if m==nil {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
