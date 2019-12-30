package models

import (
	"errors"
	"time"
	
)

var _ = time.Thursday
//GRole
type GRole struct {
	
	Id   uint     `gorm:"column:id;primary_key" form:"id" json:"id" comment:"" sql:"int(10) unsigned,PRI"`
	CreatedAt   *time.Time     `gorm:"column:created_at" form:"created_at" json:"created_at,omitempty" comment:"" sql:"timestamp"`
	UpdatedAt   *time.Time     `gorm:"column:updated_at" form:"updated_at" json:"updated_at,omitempty" comment:"" sql:"timestamp"`
	DeletedAt   *time.Time     `gorm:"column:deleted_at" form:"deleted_at" json:"deleted_at,omitempty" comment:"" sql:"timestamp,MUL"`
	RecordId   string     `gorm:"column:record_id" form:"record_id" json:"record_id" comment:"" sql:"varchar(36),MUL"`
	Name   string     `gorm:"column:name" form:"name" json:"name" comment:"" sql:"varchar(100),MUL"`
	Sequence   int     `gorm:"column:sequence" form:"sequence" json:"sequence" comment:"" sql:"int(11),MUL"`
	Memo   string     `gorm:"column:memo" form:"memo" json:"memo" comment:"" sql:"varchar(200)"`
	Creator   string     `gorm:"column:creator" form:"creator" json:"creator" comment:"" sql:"varchar(36)"`
}
//TableName
func (m *GRole) TableName() string {
	return "g_role"
}
//One,获取一条记录
func (m *GRole) One() (one *GRole, err error) {
	one = &GRole{}
	err = crudOne(m, one)
	return
}
//All，根据条件
func (m *GRole) All(q *PaginationQuery) (list *[]GRole, total uint, err error) {
	list = &[]GRole{}
	total, err = crudAll(m, q, list)
	return
}
//Update，更新
func (m *GRole) Update() (err error) {
	where := GRole{  Id: m.Id}
	//m.Id = 0
	
	return crudUpdate(m, where)
}
//Create，新增
func (m *GRole) Create() (err error) {
	//m.Id = 0
    
	return mysqlDB.Create(m).Error
}
//Delete，删除
func (m *GRole) Delete() (err error) {
	//if m.Id == 0 {
	if m==nil {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
