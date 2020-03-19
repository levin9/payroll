package models

import (
	"errors"
	"time"
	
)

var _ = time.Thursday
//GDemo
type GDemo struct {
	
	Id   uint     `gorm:"column:id;primary_key" form:"id" json:"id" comment:"" sql:"int(10) unsigned,PRI"`
	CreatedAt   *time.Time     `gorm:"column:created_at" form:"created_at" json:"created_at,omitempty" comment:"" sql:"timestamp"`
	UpdatedAt   *time.Time     `gorm:"column:updated_at" form:"updated_at" json:"updated_at,omitempty" comment:"" sql:"timestamp"`
	DeletedAt   *time.Time     `gorm:"column:deleted_at" form:"deleted_at" json:"deleted_at,omitempty" comment:"" sql:"timestamp,MUL"`
	RecordId   string     `gorm:"column:record_id" form:"record_id" json:"record_id" comment:"" sql:"varchar(36),MUL"`
	Code   string     `gorm:"column:code" form:"code" json:"code" comment:"" sql:"varchar(50),MUL"`
	Name   string     `gorm:"column:name" form:"name" json:"name" comment:"" sql:"varchar(100),MUL"`
	Memo   string     `gorm:"column:memo" form:"memo" json:"memo" comment:"" sql:"varchar(200)"`
	Status   int     `gorm:"column:status" form:"status" json:"status" comment:"" sql:"int(11),MUL"`
	Creator   string     `gorm:"column:creator" form:"creator" json:"creator" comment:"" sql:"varchar(36)"`
}
//TableName
func (m *GDemo) TableName() string {
	return "g_demo"
}
//One,获取一条记录
func (m *GDemo) One() (one *GDemo, err error) {
	one = &GDemo{}
	err = crudOne(m, one)
	return
}
//All，根据条件
func (m *GDemo) All(q *PaginationQuery) (list *[]GDemo, total uint, err error) {
	list = &[]GDemo{}
	total, err = crudAll(m, q, list)
	return
}
//Update，更新
func (m *GDemo) Update() (err error) {
	where := GDemo{  Id: m.Id}
	//m.Id = 0
	
	return crudUpdate(m, where)
}
//Create，新增
func (m *GDemo) Create() (err error) {
	//m.Id = 0
    
	return MysqlDB.Create(m).Error
}
//Delete，删除
func (m *GDemo) Delete() (err error) {
	//if m.Id == 0 {
	if m==nil {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
