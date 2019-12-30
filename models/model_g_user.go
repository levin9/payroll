package models

import (
	"errors"
	"time"
	
)

var _ = time.Thursday
//GUser
type GUser struct {
	
	Id   uint     `gorm:"column:id;primary_key" form:"id" json:"id" comment:"" sql:"int(10) unsigned,PRI"`
	CreatedAt   *time.Time     `gorm:"column:created_at" form:"created_at" json:"created_at,omitempty" comment:"" sql:"timestamp"`
	UpdatedAt   *time.Time     `gorm:"column:updated_at" form:"updated_at" json:"updated_at,omitempty" comment:"" sql:"timestamp"`
	DeletedAt   *time.Time     `gorm:"column:deleted_at" form:"deleted_at" json:"deleted_at,omitempty" comment:"" sql:"timestamp,MUL"`
	RecordId   string     `gorm:"column:record_id" form:"record_id" json:"record_id" comment:"" sql:"varchar(36),MUL"`
	UserName   string     `gorm:"column:user_name" form:"user_name" json:"user_name" comment:"" sql:"varchar(64),MUL"`
	RealName   string     `gorm:"column:real_name" form:"real_name" json:"real_name" comment:"" sql:"varchar(64),MUL"`
	Password   string     `gorm:"column:password" form:"password" json:"password" comment:"" sql:"varchar(40)"`
	Email   string     `gorm:"column:email" form:"email" json:"email" comment:"" sql:"varchar(255),MUL"`
	Phone   string     `gorm:"column:phone" form:"phone" json:"phone" comment:"" sql:"varchar(20),MUL"`
	Status   int     `gorm:"column:status" form:"status" json:"status" comment:"" sql:"int(11),MUL"`
	Creator   string     `gorm:"column:creator" form:"creator" json:"creator" comment:"" sql:"varchar(36)"`
}
//TableName
func (m *GUser) TableName() string {
	return "g_user"
}
//One,获取一条记录
func (m *GUser) One() (one *GUser, err error) {
	one = &GUser{}
	err = crudOne(m, one)
	return
}
//All，根据条件
func (m *GUser) All(q *PaginationQuery) (list *[]GUser, total uint, err error) {
	list = &[]GUser{}
	total, err = crudAll(m, q, list)
	return
}
//Update，更新
func (m *GUser) Update() (err error) {
	where := GUser{  Id: m.Id}
	//m.Id = 0
	
	return crudUpdate(m, where)
}
//Create，新增
func (m *GUser) Create() (err error) {
	//m.Id = 0
    
	return mysqlDB.Create(m).Error
}
//Delete，删除
func (m *GUser) Delete() (err error) {
	//if m.Id == 0 {
	if m==nil {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
