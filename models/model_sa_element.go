package models

import (
	"errors"
	"time"
	
)

var _ = time.Thursday
//Element
type Element struct {
	
	Id   int     `gorm:"column:Id;primary_key" form:"id" json:"id" comment:"" sql:"int(11),PRI"`
	Elemcode   string     `gorm:"column:ElemCode" form:"elem_code" json:"elem_code" comment:"" sql:"varchar(50)"`
	Category   string     `gorm:"column:Category" form:"category" json:"category" comment:"" sql:"varchar(50)"`
	Elemname   string     `gorm:"column:ElemName" form:"elem_name" json:"elem_name" comment:"" sql:"varchar(50)"`
	Canrename   int     `gorm:"column:CanRename" form:"can_rename" json:"can_rename" comment:"" sql:"tinyint(255)"`
	Reftable   string     `gorm:"column:RefTable" form:"ref_table" json:"ref_table" comment:"" sql:"varchar(50)"`
	Reffield   string     `gorm:"column:RefField" form:"ref_field" json:"ref_field" comment:"" sql:"varchar(50)"`
	FSortcode   int     `gorm:"column:F_SortCode" form:"f_sort_code" json:"f_sort_code" comment:"" sql:"int(50)"`
}
//TableName
func (m *Element) TableName() string {
	return "sa_element"
}
//One,获取一条记录
func (m *Element) One() (one *Element, err error) {
	one = &Element{}
	err = crudOne(m, one)
	return
}
//All，根据条件
func (m *Element) All(q *PaginationQuery) (list *[]Element, total uint, err error) {
	list = &[]Element{}
	total, err = crudAll(m, q, list)
	return
}
//Update，更新
func (m *Element) Update() (err error) {
	where := Element{  Id: m.Id}
	//m.Id = 0
	
	return crudUpdate(m, where)
}
//Create，新增
func (m *Element) Create() (err error) {
	//m.Id = 0
    
	return mysqlDB.Create(m).Error
}
//Delete，删除
func (m *Element) Delete() (err error) {
	//if m.Id == 0 {
	if m==nil {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
