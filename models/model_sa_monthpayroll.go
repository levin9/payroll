package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

//Monthpayroll
type Monthpayroll struct {
	Id                   string     `gorm:"column:ID;primary_key" form:"id" json:"id" comment:"ID" sql:"varchar(50),PRI"`
	Tenantid             string     `gorm:"column:TenantId" form:"tenant_id" json:"tenant_id" comment:"租户编号" sql:"varchar(50)"`
	Monthid              string     `gorm:"column:MonthId" form:"month_id" json:"month_id" comment:"月份编号" sql:"varchar(50)"`
	Personid             string     `gorm:"column:PersonId" form:"person_id" json:"person_id" comment:"雇员编号" sql:"varchar(50)"`
	Realname             string     `gorm:"column:RealName" form:"real_name" json:"real_name" comment:"姓名" sql:"varchar(50)"`
	Deptname             string     `gorm:"column:DeptName" form:"dept_name" json:"dept_name" comment:"部门" sql:"varchar(50)"`
	Joblevel             string     `gorm:"column:JobLevel" form:"job_level" json:"job_level" comment:"职级" sql:"varchar(50)"`
	Jobname              string     `gorm:"column:JobName" form:"job_name" json:"job_name" comment:"职务" sql:"varchar(50)"`
	ChngRemark           string     `gorm:"column:ChngRemark" form:"chng_remark" json:"chng_remark" comment:"移动记录" sql:"varchar(50)"`
	Monthlypayamount     float32    `gorm:"column:MonthlyPayAmount" form:"monthly_pay_amount" json:"monthly_pay_amount" comment:"月薪" sql:"decimal(10,0)"`
	Workday              float32    `gorm:"column:WorkDay" form:"work_day" json:"work_day" comment:"工作天数" sql:"decimal(10,4)"`
	Planworkday          float32    `gorm:"column:PlanWorkDay" form:"plan_work_day" json:"plan_work_day" comment:"应上班天数" sql:"decimal(10,4)"`
	Actualworkday        float32    `gorm:"column:ActualWorkDay" form:"actual_work_day" json:"actual_work_day" comment:"实际出勤天数" sql:"decimal(10,4)"`
	Basicsallaryamount   float32    `gorm:"column:BasicSallaryAmount" form:"basic_sallary_amount" json:"basic_sallary_amount" comment:"基本工资" sql:"decimal(10,4)"`
	Kpiamout             float32    `gorm:"column:KPIAmout" form:"kpiamout" json:"kpiamout" comment:"绩效工资" sql:"decimal(10,4)"`
	Trafficfee           float32    `gorm:"column:TrafficFee" form:"traffic_fee" json:"traffic_fee" comment:"交通补贴" sql:"decimal(10,4)"`
	Foodfee              float32    `gorm:"column:FoodFee" form:"food_fee" json:"food_fee" comment:"餐费补贴" sql:"decimal(10,4)"`
	Attfee               float32    `gorm:"column:AttFee" form:"att_fee" json:"att_fee" comment:"假期扣款" sql:"decimal(10,4)"`
	Attawardfee          float32    `gorm:"column:AttAwardFee" form:"att_award_fee" json:"att_award_fee" comment:"考勤奖励" sql:"decimal(10,4)"`
	Chidaocount          float32    `gorm:"column:ChiDaoCount" form:"chi_dao_count" json:"chi_dao_count" comment:"迟到次数" sql:"decimal(10,4)"`
	Chidaofee            float32    `gorm:"column:ChiDaoFee" form:"chi_dao_fee" json:"chi_dao_fee" comment:"迟到扣款" sql:"decimal(10,4)"`
	ZaoTuicount          float32    `gorm:"column:ZaoTuicount" form:"zaotui_count" json:"zaotui_count" comment:"早退次数" sql:"decimal(10,4)"`
	ZaoTuifee            float32    `gorm:"column:ZaoTuifee" form:"zaotui_fee" json:"zaotui_fee" comment:"早退扣款" sql:"decimal(10,4)"`
	Sickfee              float32    `gorm:"column:SickFee" form:"sick_fee" json:"sick_fee" comment:"病假扣款" sql:"decimal(10,4)"`
	Sicknum              float32    `gorm:"column:SickNum" form:"sick_num" json:"sick_num" comment:"病假时间" sql:"decimal(10,4)"`
	Shijiafee            float32    `gorm:"column:ShiJiaFee" form:"shi_jia_fee" json:"shi_jia_fee" comment:"事假扣款" sql:"decimal(10,4)"`
	Shijianum            float32    `gorm:"column:ShiJiaNum" form:"shi_jia_num" json:"shi_jia_num" comment:"事假时间" sql:"decimal(10,4)"`
	Annualleavenum       float32    `gorm:"column:AnnualLeaveNum" form:"annual_leave_num" json:"annual_leave_num" comment:"年假时间" sql:"decimal(10,4)"`
	Annualleavefee       float32    `gorm:"column:AnnualLeaveFee" form:"annual_leave_fee" json:"annual_leave_fee" comment:"年假扣费" sql:"decimal(10,4)"`
	Professionafee       float32    `gorm:"column:ProfessionaFee" form:"professiona_fee" json:"professiona_fee" comment:"职称补贴" sql:"decimal(10,4)"`
	Joballowance         float32    `gorm:"column:JobAllowance" form:"job_allowance" json:"job_allowance" comment:"岗位津贴" sql:"decimal(10,4)"`
	Assignfee            float32    `gorm:"column:AssignFee" form:"assign_fee" json:"assign_fee" comment:"外派项目补贴" sql:"decimal(10,4)"`
	Mobilefee            float32    `gorm:"column:MobileFee" form:"mobile_fee" json:"mobile_fee" comment:"通讯补贴" sql:"decimal(10,4)"`
	Shebaofee            float32    `gorm:"column:SheBaoFee" form:"she_bao_fee" json:"she_bao_fee" comment:"社保费" sql:"decimal(10,4)"`
	Housefee             float32    `gorm:"column:HouseFee" form:"house_fee" json:"house_fee" comment:"公积金费用" sql:"decimal(10,4)"`
	Specjoballowance     float32    `gorm:"column:SpecJobAllowance" form:"spec_job_allowance" json:"spec_job_allowance" comment:"特殊岗位津贴" sql:"decimal(10,4)"`
	Sangjiafee           float32    `gorm:"column:SangJiaFee" form:"sang_jia_fee" json:"sang_jia_fee" comment:"丧假" sql:"decimal(10,4)"`
	Hunjiafee            float32    `gorm:"column:HunJiaFee" form:"hun_jia_fee" json:"hun_jia_fee" comment:"婚假" sql:"decimal(10,4)"`
	Burvjiafee           float32    `gorm:"column:BuRvJiaFee" form:"bu_rv_jia_fee" json:"bu_rv_jia_fee" comment:"哺乳假" sql:"decimal(10,4)"`
	Chanjiafee           float32    `gorm:"column:ChanJiaFee" form:"chan_jia_fee" json:"chan_jia_fee" comment:"产假" sql:"decimal(10,4)"`
	Peichanjiafee        float32    `gorm:"column:PeiChanJiaFee" form:"pei_chan_jia_fee" json:"pei_chan_jia_fee" comment:"陪产假" sql:"decimal(10,4)"`
	Kongbaifee           float32    `gorm:"column:KongBaiFee" form:"kong_bai_fee" json:"kong_bai_fee" comment:"考勤空白扣费" sql:"decimal(10,4)"`
	Queqinfee            float32    `gorm:"column:QueQinFee" form:"que_qin_fee" json:"que_qin_fee" comment:"缺勤" sql:"decimal(10,4)"`
	Tiaoxiufee           float32    `gorm:"column:TiaoXiuFee" form:"tiao_xiu_fee" json:"tiao_xiu_fee" comment:"调休" sql:"decimal(10,4)"`
	Daixinjiafee         float32    `gorm:"column:DaiXinJiaFee" form:"dai_xin_jia_fee" json:"dai_xin_jia_fee" comment:"带薪假" sql:"decimal(10,4)"`
	Personaltax          float32    `gorm:"column:PersonalTax" form:"personal_tax" json:"personal_tax" comment:"个人所得税" sql:"decimal(10,4)"`
	Wagespayableamount   float32    `gorm:"column:WagesPayableAmount" form:"wages_payable_amount" json:"wages_payable_amount" comment:"应发工资" sql:"decimal(10,4)"`
	Actualpayableamount  float32    `gorm:"column:ActualPayableAmount" form:"actual_payable_amount" json:"actual_payable_amount" comment:"实发工资" sql:"decimal(10,4)"`
	Adjustfee            float32    `gorm:"column:AdjustFee" form:"adjust_fee" json:"adjust_fee" comment:"调整项" sql:"decimal(10,4)"`
	Freetaxamount        float32    `gorm:"column:FreeTaxAmount" form:"free_tax_amount" json:"free_tax_amount" comment:"每月免税金额" sql:"decimal(10,4)"`
	Incometaxamount      float32    `gorm:"column:IncomeTaxAmount" form:"income_tax_amount" json:"income_tax_amount" comment:"代扣个税" sql:"decimal(10,4)"`
	Standardshebaoamount float32    `gorm:"column:StandardSheBaoAmount" form:"standard_she_bao_amount" json:"standard_she_bao_amount" comment:"社保基数" sql:"decimal(10,4)"`
	Standardhouseamount  float32    `gorm:"column:StandardHouseAmount" form:"standard_house_amount" json:"standard_house_amount" comment:"公积金基数" sql:"decimal(10,4)"`
	Custsubfee2          float32    `gorm:"column:CustSubFee2" form:"cust_sub_fee2" json:"cust_sub_fee2" comment:"自定义减项1" sql:"decimal(10,4)"`
	Custsubfee1          float32    `gorm:"column:CustSubFee1" form:"cust_sub_fee1" json:"cust_sub_fee1" comment:"自定义减项2" sql:"decimal(10,4)"`
	Custaddfee2          float32    `gorm:"column:CustAddFee2" form:"cust_add_fee2" json:"cust_add_fee2" comment:"自定义加项1" sql:"decimal(10,4)"`
	Custaddfee1          float32    `gorm:"column:CustAddFee1" form:"cust_add_fee1" json:"cust_add_fee1" comment:"自定义加项2" sql:"decimal(10,4)"`
	Extinfo              string     `gorm:"column:ExtInfo" form:"ext_info" json:"ext_info" comment:"扩展字段" sql:"varchar(1000)"`
	FEnabledmark         int        `gorm:"column:F_EnabledMark" form:"f_enabled_mark" json:"f_enabled_mark" comment:"有效标志0否1是" sql:"int(11)"`
	FCreateuserid        string     `gorm:"column:F_CreateUserId" form:"f_create_user_id" json:"f_create_user_id" comment:"创建人ID" sql:"varchar(50)"`
	FModifyusername      string     `gorm:"column:F_ModifyUserName" form:"f_modify_user_name" json:"f_modify_user_name" comment:"编辑人" sql:"varchar(50)"`
	FModifyuserid        string     `gorm:"column:F_ModifyUserId" form:"f_modify_user_id" json:"f_modify_user_id" comment:"编辑人ID" sql:"varchar(50)"`
	FDeletemark          string     `gorm:"column:F_DeleteMark" form:"f_delete_mark" json:"f_delete_mark" comment:"删除标记0否1是" sql:"varchar(200)"`
	FSortcode            int        `gorm:"column:F_SortCode" form:"f_sort_code" json:"f_sort_code" comment:"排序码" sql:"int(11)"`
	FModifydate          *time.Time `gorm:"column:F_ModifyDate" form:"f_modify_date" json:"f_modify_date,omitempty" comment:"编辑日期" sql:"datetime"`
	FCreateusername      string     `gorm:"column:F_CreateUserName" form:"f_create_user_name" json:"f_create_user_name" comment:"创建人" sql:"varchar(50)"`
	FCreatedate          *time.Time `gorm:"column:F_CreateDate" form:"f_create_date" json:"f_create_date,omitempty" comment:"创建时间" sql:"datetime"`
	FDescription         string     `gorm:"column:F_Description" form:"f_description" json:"f_description" comment:"备注" sql:"varchar(200)"`
}

//TableName
func (m *Monthpayroll) TableName() string {
	return "sa_monthpayroll"
}

//One,获取一条记录
func (m *Monthpayroll) One() (one *Monthpayroll, err error) {
	one = &Monthpayroll{}
	err = crudOne(m, one)
	return
}

//All，根据条件
func (m *Monthpayroll) All(q *PaginationQuery) (list *[]Monthpayroll, total uint, err error) {
	list = &[]Monthpayroll{}
	total, err = crudAll(m, q, list)
	return
}

func (m *Monthpayroll) FindAll(q *PaginationQuery) (list *[]Monthpayroll, total uint, err error) {
	list = &[]Monthpayroll{}
	total, err = crudGetAll(m, q, list)
	return
}

//Update，更新
func (m *Monthpayroll) Update() (err error) {
	where := Monthpayroll{Id: m.Id}
	//m.Id = 0

	return crudUpdate(m, where)
}

//Create，新增
func (m *Monthpayroll) Create() (err error) {
	//m.Id = 0

	return mysqlDB.Create(m).Error
}

//Delete，删除
func (m *Monthpayroll) Delete() (err error) {
	//if m.Id == 0 {
	if m == nil {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
