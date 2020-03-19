package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

//Attcategorytenant
type Attcategorytenant struct {
	dal                 BaseDal
	Holidaycategoryid   string     `gorm:"column:HolidayCategoryId;primary_key" form:"holiday_category_id" json:"holiday_category_id" comment:"Id" sql:"varchar(50),PRI"`
	Tenantid            string     `gorm:"column:TenantId" form:"tenant_id" json:"tenant_id" comment:"租户编号" sql:"varchar(50)"`
	Holidaycode         string     `gorm:"column:HolidayCode" form:"holiday_code" json:"holiday_code" comment:"编号" sql:"varchar(20)"`
	Tenantholidayname   string     `gorm:"column:TenantHolidayName" form:"tenant_holiday_name" json:"tenant_holiday_name" comment:"名称" sql:"varchar(50)"`
	Tenantformula       string     `gorm:"column:TenantFormula" form:"tenant_formula" json:"tenant_formula" comment:"公式" sql:"varchar(200)"`
	Tenantformularemark string     `gorm:"column:TenantFormulaRemark" form:"tenant_formula_remark" json:"tenant_formula_remark" comment:"公式备注" sql:"varchar(200)"`
	Isonslip            int        `gorm:"column:IsOnSlip" form:"is_on_slip" json:"is_on_slip" comment:"是否显示在工资单上" sql:"int(11)"`
	Isused              int        `gorm:"column:IsUsed" form:"is_used" json:"is_used" comment:"是否启用" sql:"int(11)"`
	Sqltenantformula    string     `gorm:"column:SQLTenantFormula" form:"sqltenant_formula" json:"sqltenant_formula" comment:"SQL公式" sql:"varchar(200)"`
	FEnabledmark        int        `gorm:"column:F_EnabledMark" form:"f_enabled_mark" json:"f_enabled_mark" comment:"有效标志0否1是" sql:"int(11)"`
	FCreateuserid       string     `gorm:"column:F_CreateUserId" form:"f_create_user_id" json:"f_create_user_id" comment:"创建人ID" sql:"varchar(50)"`
	FModifyusername     string     `gorm:"column:F_ModifyUserName" form:"f_modify_user_name" json:"f_modify_user_name" comment:"编辑人" sql:"varchar(50)"`
	FModifyuserid       string     `gorm:"column:F_ModifyUserId" form:"f_modify_user_id" json:"f_modify_user_id" comment:"编辑人ID" sql:"varchar(50)"`
	FDeletemark         int        `gorm:"column:F_DeleteMark" form:"f_delete_mark" json:"f_delete_mark" comment:"删除标记0否1是" sql:"int(11)"`
	FSortcode           int        `gorm:"column:F_SortCode" form:"f_sort_code" json:"f_sort_code" comment:"排序码" sql:"int(11)"`
	FModifydate         *time.Time `gorm:"column:F_ModifyDate" form:"f_modify_date" json:"f_modify_date,omitempty" comment:"编辑日期" sql:"datetime"`
	FCreateusername     string     `gorm:"column:F_CreateUserName" form:"f_create_user_name" json:"f_create_user_name" comment:"创建人" sql:"varchar(50)"`
	FCreatedate         *time.Time `gorm:"column:F_CreateDate" form:"f_create_date" json:"f_create_date,omitempty" comment:"创建时间" sql:"datetime"`
	FDescription        string     `gorm:"column:F_Description" form:"f_description" json:"f_description" comment:"备注" sql:"varchar(200)"`
}

//TableName
func (m *Attcategorytenant) TableName() string {
	return "sa_attcategorytenant"
}

//One,获取一条记录
func (m *Attcategorytenant) One() (one *Attcategorytenant, err error) {
	one = &Attcategorytenant{}
	err = crudOne(m, one)
	return
}

//All，根据条件
func (m *Attcategorytenant) All(q *PaginationQuery) (list *[]Attcategorytenant, total uint, err error) {
	list = &[]Attcategorytenant{}
	total, err = crudAll(m, q, list)
	return
}

//Update，更新
func (m *Attcategorytenant) Update() (err error) {
	where := Attcategorytenant{Holidaycategoryid: m.Holidaycategoryid}
	//m.Id = 0

	return crudUpdate(m, where)
}

//Create，新增
func (m *Attcategorytenant) Create() (err error) {
	//m.Id = 0

	return MysqlDB.Create(m).Error
}

//Delete，删除
func (m *Attcategorytenant) Delete() (err error) {
	//if m.Id == 0 {
	if m == nil {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
