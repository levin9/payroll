package models

import (
	"errors"
	"time"
	
)

var _ = time.Thursday
//Personpayroll
type Personpayroll struct {
	
	Id   string     `gorm:"column:Id;primary_key" form:"id" json:"id" comment:"ID" sql:"varchar(50),PRI"`
	Tenantid   string     `gorm:"column:TenantId" form:"tenant_id" json:"tenant_id" comment:"租户编号" sql:"varchar(50)"`
	Personid   string     `gorm:"column:PersonId" form:"person_id" json:"person_id" comment:"雇员编号" sql:"varchar(50)"`
	Monthlypayamount   float32     `gorm:"column:MonthlyPayAmount" form:"monthly_pay_amount" json:"monthly_pay_amount" comment:"月薪" sql:"decimal(10,0)"`
	Validdate   *time.Time     `gorm:"column:ValidDate" form:"valid_date" json:"valid_date,omitempty" comment:"生效日期" sql:"datetime(6)"`
	Invaliddate   *time.Time     `gorm:"column:InvalidDate" form:"invalid_date" json:"invalid_date,omitempty" comment:"失效日期" sql:"datetime(6)"`
	Basicsallaryamount   float32     `gorm:"column:BasicSallaryAmount" form:"basic_sallary_amount" json:"basic_sallary_amount" comment:"基本工资" sql:"decimal(10,4)"`
	Kpiamout   float32     `gorm:"column:KPIAmout" form:"kpiamout" json:"kpiamout" comment:"绩效工资" sql:"decimal(10,4)"`
	Trafficfee   float32     `gorm:"column:TrafficFee" form:"traffic_fee" json:"traffic_fee" comment:"交通补贴" sql:"decimal(10,4)"`
	Foodfee   float32     `gorm:"column:FoodFee" form:"food_fee" json:"food_fee" comment:"餐费补贴" sql:"decimal(10,4)"`
	Attawardfee   float32     `gorm:"column:AttAwardFee" form:"att_award_fee" json:"att_award_fee" comment:"全勤奖" sql:"decimal(10,4)"`
	Attfee   float32     `gorm:"column:AttFee" form:"att_fee" json:"att_fee" comment:"考勤奖励" sql:"decimal(10,4)"`
	Professionafee   float32     `gorm:"column:ProfessionaFee" form:"professiona_fee" json:"professiona_fee" comment:"职称补贴" sql:"decimal(10,4)"`
	Joballowance   float32     `gorm:"column:JobAllowance" form:"job_allowance" json:"job_allowance" comment:"岗位津贴" sql:"decimal(10,4)"`
	Assignfee   float32     `gorm:"column:AssignFee" form:"assign_fee" json:"assign_fee" comment:"外派项目补贴" sql:"decimal(10,4)"`
	Mobilefee   float32     `gorm:"column:MobileFee" form:"mobile_fee" json:"mobile_fee" comment:"通讯补贴" sql:"decimal(10,4)"`
	Shebaofee   float32     `gorm:"column:SheBaoFee" form:"she_bao_fee" json:"she_bao_fee" comment:"社保标准" sql:"decimal(10,4)"`
	Housefee   float32     `gorm:"column:HouseFee" form:"house_fee" json:"house_fee" comment:"住房公积金" sql:"decimal(10,4)"`
	Custaddfee2   float32     `gorm:"column:CustAddFee2" form:"cust_add_fee2" json:"cust_add_fee2" comment:"自定义加项2" sql:"decimal(10,4)"`
	Custaddfee1   float32     `gorm:"column:CustAddFee1" form:"cust_add_fee1" json:"cust_add_fee1" comment:"自定义加项1" sql:"decimal(10,4)"`
	Custsubfee2   float32     `gorm:"column:CustSubFee2" form:"cust_sub_fee2" json:"cust_sub_fee2" comment:"自定义减项1" sql:"decimal(10,4)"`
	Custsubfee1   float32     `gorm:"column:CustSubFee1" form:"cust_sub_fee1" json:"cust_sub_fee1" comment:"自定义减项2" sql:"decimal(10,4)"`
	Islast   int     `gorm:"column:IsLast" form:"is_last" json:"is_last" comment:"是否最新" sql:"int(11)"`
	Specjoballowance   float32     `gorm:"column:SpecJobAllowance" form:"spec_job_allowance" json:"spec_job_allowance" comment:"特殊岗位津贴" sql:"decimal(10,4)"`
	Shebaoaccountno   string     `gorm:"column:SheBaoAccountNo" form:"she_bao_account_no" json:"she_bao_account_no" comment:"社保账号" sql:"varchar(50)"`
	Houseaccountno   string     `gorm:"column:HouseAccountNo" form:"house_account_no" json:"house_account_no" comment:"住房公积金账号" sql:"varchar(50)"`
	FEnabledmark   int     `gorm:"column:F_EnabledMark" form:"f_enabled_mark" json:"f_enabled_mark" comment:"有效标志0否1是" sql:"int(11)"`
	FCreateuserid   string     `gorm:"column:F_CreateUserId" form:"f_create_user_id" json:"f_create_user_id" comment:"创建人ID" sql:"varchar(50)"`
	FModifyusername   string     `gorm:"column:F_ModifyUserName" form:"f_modify_user_name" json:"f_modify_user_name" comment:"编辑人" sql:"varchar(50)"`
	FModifyuserid   string     `gorm:"column:F_ModifyUserId" form:"f_modify_user_id" json:"f_modify_user_id" comment:"编辑人ID" sql:"varchar(50)"`
	FDeletemark   int     `gorm:"column:F_DeleteMark" form:"f_delete_mark" json:"f_delete_mark" comment:"删除标记0否1是" sql:"int(11)"`
	FSortcode   int     `gorm:"column:F_SortCode" form:"f_sort_code" json:"f_sort_code" comment:"排序码" sql:"int(11)"`
	FModifydate   *time.Time     `gorm:"column:F_ModifyDate" form:"f_modify_date" json:"f_modify_date,omitempty" comment:"编辑日期" sql:"datetime"`
	FCreateusername   string     `gorm:"column:F_CreateUserName" form:"f_create_user_name" json:"f_create_user_name" comment:"创建人" sql:"varchar(50)"`
	FCreatedate   *time.Time     `gorm:"column:F_CreateDate" form:"f_create_date" json:"f_create_date,omitempty" comment:"创建时间" sql:"datetime"`
	FDescription   string     `gorm:"column:F_Description" form:"f_description" json:"f_description" comment:"备注" sql:"varchar(200)"`
}
//TableName
func (m *Personpayroll) TableName() string {
	return "sa_personpayroll"
}
//One,获取一条记录
func (m *Personpayroll) One() (one *Personpayroll, err error) {
	one = &Personpayroll{}
	err = crudOne(m, one)
	return
}
//All，根据条件
func (m *Personpayroll) All(q *PaginationQuery) (list *[]Personpayroll, total uint, err error) {
	list = &[]Personpayroll{}
	total, err = crudAll(m, q, list)
	return
}
//Update，更新
func (m *Personpayroll) Update() (err error) {
	where := Personpayroll{  Id: m.Id}
	//m.Id = 0
	
	return crudUpdate(m, where)
}
//Create，新增
func (m *Personpayroll) Create() (err error) {
	//m.Id = 0
    
	return mysqlDB.Create(m).Error
}
//Delete，删除
func (m *Personpayroll) Delete() (err error) {
	//if m.Id == 0 {
	if m==nil {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
