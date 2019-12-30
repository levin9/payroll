package models

import (
	"errors"
	"time"
	
)

var _ = time.Thursday
//Attendance
type Attendance struct {
	
	Id   string     `gorm:"column:ID;primary_key" form:"id" json:"id" comment:"编号" sql:"varchar(50),PRI"`
	Tenantid   string     `gorm:"column:TenantId" form:"tenant_id" json:"tenant_id" comment:"租户编号" sql:"varchar(50),MUL"`
	Personid   string     `gorm:"column:PersonId" form:"person_id" json:"person_id" comment:"雇员编号" sql:"varchar(50)"`
	Monthid   string     `gorm:"column:MonthId" form:"month_id" json:"month_id" comment:"月份编号" sql:"varchar(50)"`
	Chidaocount   float32     `gorm:"column:ChiDaoCount" form:"chi_dao_count" json:"chi_dao_count" comment:"迟到次数" sql:"decimal(11,2)"`
	Chidaonum   float32     `gorm:"column:ChiDaoNum" form:"chi_dao_num" json:"chi_dao_num" comment:"迟到时间" sql:"decimal(11,2)"`
	Zaotuicount   float32     `gorm:"column:ZaotuiCount" form:"zaotui_count" json:"zaotui_count" comment:"早退次数" sql:"decimal(11,2)"`
	Zaotuinum   float32     `gorm:"column:ZaotuiNum" form:"zaotui_num" json:"zaotui_num" comment:"早退时间" sql:"decimal(11,2)"`
	Tiaoxiunum   float32     `gorm:"column:TiaoxiuNum" form:"tiaoxiu_num" json:"tiaoxiu_num" comment:"调休" sql:"decimal(11,2)"`
	Shijianum   float32     `gorm:"column:ShiJiaNum" form:"shi_jia_num" json:"shi_jia_num" comment:"事假" sql:"decimal(11,2)"`
	Sicknum   float32     `gorm:"column:SickNum" form:"sick_num" json:"sick_num" comment:"病假" sql:"decimal(11,2)"`
	Burvjianum   float32     `gorm:"column:BuRvJiaNum" form:"bu_rv_jia_num" json:"bu_rv_jia_num" comment:"哺乳假" sql:"decimal(11,2)"`
	Peichanjianum   float32     `gorm:"column:PeiChanJiaNum" form:"pei_chan_jia_num" json:"pei_chan_jia_num" comment:"陪产假" sql:"decimal(11,2)"`
	Chanjianum   float32     `gorm:"column:ChanJiaNum" form:"chan_jia_num" json:"chan_jia_num" comment:"产假" sql:"decimal(11,2)"`
	Hunjianum   float32     `gorm:"column:HunJIaNum" form:"hun_jia_num" json:"hun_jia_num" comment:"婚假" sql:"decimal(11,2)"`
	Sangjianum   float32     `gorm:"column:SangJiaNum" form:"sang_jia_num" json:"sang_jia_num" comment:"丧假" sql:"decimal(11,2)"`
	Kuanggongnum   float32     `gorm:"column:KuanggongNum" form:"kuanggong_num" json:"kuanggong_num" comment:"旷工" sql:"decimal(11,2)"`
	Kongbainum   float32     `gorm:"column:KongBaiNum" form:"kong_bai_num" json:"kong_bai_num" comment:"考勤空白" sql:"decimal(11,2)"`
	Daixinshijianum   float32     `gorm:"column:DaiXinShiJiaNum" form:"dai_xin_shi_jia_num" json:"dai_xin_shi_jia_num" comment:"带薪事假" sql:"decimal(11,2)"`
	Planworkday   float32     `gorm:"column:PlanWorkDay" form:"plan_work_day" json:"plan_work_day" comment:"应上班天数" sql:"decimal(11,2)"`
	Actualworkday   float32     `gorm:"column:ActualWorkDay" form:"actual_work_day" json:"actual_work_day" comment:"实际出勤天数" sql:"decimal(11,2)"`
	Annualleavenum   float32     `gorm:"column:AnnualLeaveNum" form:"annual_leave_num" json:"annual_leave_num" comment:"年假" sql:"decimal(11,2)"`
	FEnabledmark   float32     `gorm:"column:F_EnabledMark" form:"f_enabled_mark" json:"f_enabled_mark" comment:"有效标志0否1是" sql:"decimal(11,2)"`
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
func (m *Attendance) TableName() string {
	return "sa_attendance"
}
//One,获取一条记录
func (m *Attendance) One() (one *Attendance, err error) {
	one = &Attendance{}
	err = crudOne(m, one)
	return
}
//All，根据条件
func (m *Attendance) All(q *PaginationQuery) (list *[]Attendance, total uint, err error) {
	list = &[]Attendance{}
	total, err = crudAll(m, q, list)
	return
}
//Update，更新
func (m *Attendance) Update() (err error) {
	where := Attendance{  Id: m.Id}
	//m.Id = 0
	
	return crudUpdate(m, where)
}
//Create，新增
func (m *Attendance) Create() (err error) {
	//m.Id = 0
    
	return mysqlDB.Create(m).Error
}
//Delete，删除
func (m *Attendance) Delete() (err error) {
	//if m.Id == 0 {
	if m==nil {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
