package blls

import (
	"fmt"
	"payroll/models"
	"payroll/utils"
	"time"
)

func UpdatePersonal(input PersonalInput) {

	person := models.Personal{}
	tax := new(models.Personfreetax)
	bank := new(models.Bankaccount)
	sallary := new(models.Personpayroll)
	fmt.Println("开始查询用户信息")
	//if err := db.Where("id=?", 5).First(&stu).Error; err != nil {
	if err := models.MysqlDB.Where("Personid=?", input.Personid).First(&person).Error; err != nil {
		fmt.Println(person)
		fmt.Println("没有找到用户")
	}
	fmt.Println("查询用户信息结束。")
	fmt.Println(input.Personid)
	fmt.Println(person.Realname)
	fmt.Println(person.Deptname)

	//if flag := models.MysqlDB.Where("Personid=?", input.Personid).First(person).RecordNotFound(); flag {
	if person.Personid != "" {
		if err := models.MysqlDB.Where("Personid=?", input.Personid).FirstOrCreate(&tax); err != nil {

		}
		if err := models.MysqlDB.Where("Personid=?", input.Personid).FirstOrCreate(&bank); err != nil {

		}
		if err := models.MysqlDB.Where("Personid=?", input.Personid).FirstOrCreate(&sallary); err != nil {

		}
	}

	person.Personid = input.Personid
	person.Tenantid = input.Tenantid
	person.Userid = input.Userid
	person.Realname = input.Realname
	person.Mobile = input.Mobile
	person.Jobno = input.Jobno
	person.Deptname = input.Deptname
	person.Jobname = input.Jobname
	person.Joblevel = input.Joblevel
	person.Joindate = input.Joindate
	person.Lastdate = input.Lastdate
	person.Regulardate = input.Regulardate
	person.Sallaryamount = input.Sallaryamount
	person.Currentfreetax = input.Currentfreetax
	//person.FModifydate = time.Now()

	tax.Znjy = input.Znjy
	tax.Jxjy = input.Jxjy
	tax.Zfdk = input.Zfdk
	tax.Zfzj = input.Zfzj
	tax.Sylr = input.Sylr

	bank.Itemindex = input.Itemindex
	bank.Bankname = input.Bankname
	bank.Bankbranch = input.Bankbranch
	bank.Accountno = input.Accountno
	bank.Accountname = input.Accountname

	sallary.Monthlypayamount = input.Monthlypayamount
	sallary.Validdate = input.Validdate
	sallary.Invaliddate = input.Invaliddate
	sallary.Basicsallaryamount = input.Basicsallaryamount
	sallary.KpiAmout = input.KpiAmout
	sallary.Trafficfee = input.Trafficfee
	sallary.Foodfee = input.Foodfee
	sallary.Attawardfee = input.Attawardfee
	sallary.Attfee = input.Attfee
	sallary.Professionafee = input.Professionafee

	sallary.Joballowance = input.Joballowance
	sallary.Assignfee = input.Assignfee
	sallary.Mobilefee = input.Mobilefee
	sallary.Shebaofee = input.Shebaofee
	sallary.Housefee = input.Housefee
	sallary.Custaddfee2 = input.Custaddfee2
	sallary.Custaddfee1 = input.Custaddfee1
	sallary.Custaddfee1 = input.Custsubfee2
	sallary.Custaddfee1 = input.Custsubfee1
	sallary.Islast = input.Islast
	sallary.Specjoballowance = input.Specjoballowance
	sallary.Shebaoaccountno = input.Shebaoaccountno
	sallary.Houseaccountno = input.Houseaccountno

	if person.Personid == "" {
		person.Create()
	} else {
		person.Update()
	}

	//如果对象为空，则personid为空，否则酒存在
	if tax.Personid == "" {
		tax.Id = utils.NewID()
		tax.Personid = person.Personid
		tax.Tenantid = input.Tenantid
		tax.Create()
	} else {
		tax.Update()
	}

	if bank.Personid == "" {
		bank.Id = utils.NewID()
		bank.Personid = person.Personid
		//bank.Tenantid = input.Tenantid
		bank.Create()
	} else {
		bank.Update()
	}

	if sallary.Personid == "" {
		sallary.Id = utils.NewID()
		sallary.Personid = person.Personid
		sallary.Tenantid = input.Tenantid
		sallary.Create()
	} else {
		sallary.Update()
	}

	fmt.Println("hello")
}

type PersonalInput struct {
	Personid           string     `form:"person_id" json:"person_id" comment:"雇员编号" `
	Tenantid           string     `form:"tenant_id" json:"tenant_id" comment:"租户编号" `
	Userid             string     `form:"user_id" json:"user_id" comment:"用户编号"`
	Realname           string     `form:"real_name" json:"real_name" comment:"用户姓名" `
	Mobile             string     `form:"mobile" json:"mobile" comment:"手机号"`
	Jobno              string     `form:"job_no" json:"job_no" comment:"工号" `
	Deptname           string     `form:"dept_name" json:"dept_name" comment:"部门名称"`
	Jobname            string     `form:"job_name" json:"job_name" comment:"职务" `
	Joblevel           string     `form:"job_level" json:"job_level" comment:"职级"`
	Joindate           *time.Time `form:"join_date" json:"join_date,omitempty" comment:"入职日期"`
	Lastdate           *time.Time `form:"last_date" json:"last_date,omitempty" comment:"离职日期"`
	Regulardate        *time.Time `form:"regular_date" json:"regular_date,omitempty" comment:"转正日期" `
	Sallaryamount      float32    `form:"sallary_amount" json:"sallary_amount" comment:"月薪"`
	Currentfreetax     float32    `form:"current_free_tax" json:"current_free_tax" comment:"免税金额"`
	Znjy               float32    `form:"znjy" json:"znjy" comment:"子女教育" `
	Jxjy               float32    `form:"jxjy" json:"jxjy" comment:"继续教育" `
	Zfdk               float32    `form:"zfdk" json:"zfdk" comment:"住房贷款" `
	Zfzj               float32    `form:"zfzj" json:"zfzj" comment:"住房租金" `
	Sylr               float32    `form:"sylr" json:"sylr" comment:"赡养老人" `
	Itemindex          int        `form:"item_index" json:"item_index" comment:"序号" `
	Bankname           string     `form:"bank_name" json:"bank_name" comment:"开户行" `
	Bankbranch         string     `form:"bank_branch" json:"bank_branch" comment:"开户分行" `
	Accountno          string     `form:"account_no" json:"account_no" comment:"账号" `
	Accountname        string     `form:"account_name" json:"account_name" comment:"账号姓名" `
	Monthlypayamount   float32    `form:"monthly_pay_amount" json:"monthly_pay_amount" comment:"月薪" `
	Validdate          *time.Time `form:"valid_date" json:"valid_date,omitempty" comment:"生效日期" `
	Invaliddate        *time.Time `form:"invalid_date" json:"invalid_date,omitempty" comment:"失效日期" `
	Basicsallaryamount float32    `form:"basic_sallary_amount" json:"basic_sallary_amount" comment:"基本工资" `
	KpiAmout           float32    `form:"KpiAmout" json:"KpiAmout" comment:"绩效工资" `
	Trafficfee         float32    `form:"traffic_fee" json:"traffic_fee" comment:"交通补贴" `
	Foodfee            float32    `form:"food_fee" json:"food_fee" comment:"餐费补贴" `
	Attawardfee        float32    `form:"att_award_fee" json:"att_award_fee" comment:"全勤奖" `
	Attfee             float32    `form:"att_fee" json:"att_fee" comment:"考勤奖励" `
	Professionafee     float32    `form:"professiona_fee" json:"professiona_fee" comment:"职称补贴" `
	Joballowance       float32    `form:"job_allowance" json:"job_allowance" comment:"岗位津贴" `
	Assignfee          float32    `form:"assign_fee" json:"assign_fee" comment:"外派项目补贴" `
	Mobilefee          float32    `form:"mobile_fee" json:"mobile_fee" comment:"通讯补贴" `
	Shebaofee          float32    `form:"she_bao_fee" json:"she_bao_fee" comment:"社保标准" `
	Housefee           float32    `form:"house_fee" json:"house_fee" comment:"住房公积金" `
	Custaddfee2        float32    `form:"cust_add_fee2" json:"cust_add_fee2" comment:"自定义加项2" `
	Custaddfee1        float32    `form:"cust_add_fee1" json:"cust_add_fee1" comment:"自定义加项1" `
	Custsubfee2        float32    `form:"cust_sub_fee2" json:"cust_sub_fee2" comment:"自定义减项1" `
	Custsubfee1        float32    `form:"cust_sub_fee1" json:"cust_sub_fee1" comment:"自定义减项2" `
	Islast             int        `form:"is_last" json:"is_last" comment:"是否最新" `
	Specjoballowance   float32    `form:"spec_job_allowance" json:"spec_job_allowance" comment:"特殊岗位津贴" `
	Shebaoaccountno    string     `form:"she_bao_account_no" json:"she_bao_account_no" comment:"社保账号" `
	Houseaccountno     string     `form:"house_account_no" json:"house_account_no" comment:"住房公积金账号" `
}

type PersonalBaseInput struct {
	Personid string `form:"person_id" json:"person_id" comment:"雇员编号" `
	Tenantid string `form:"tenant_id" json:"tenant_id" comment:"租户编号" `
}
