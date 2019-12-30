package blls

import (
	"fmt"
	"payroll/models"
)

type OperatorPayrollHandler struct {
}

//CalculateMonthPayroll
func (p *OperatorPayrollHandler) CalculateMonthPayroll(para CalcParameter) string {
	//load personal
	//tenantBll := models.TenantFormula{}
	personal := models.Monthpayroll{}
	personalData, _, _ := personal.All(nil)

	payrollService := models.TenantFormulaService{}
	rulllist, err := payrollService.GetTenantMap(para.TenantId)
	if err != nil {
		return "falure"
	}
	//change user variable
	for _, person := range *personalData {
		fmt.Println(person)
	}
	//calc payroll
	for _, person := range *personalData {
		for _, r := range *rulllist {
			fmt.Println(r.TenantFormula)
		}
		fmt.Println(person)
	}
	//update database

	fmt.Println(rulllist)

	//personalData := payrollService.

	//tenantBll.getUserData()
	//payrollHandler

	//load var
	//fmt.Println(tenantBll)

	//load formula
	fmt.Println(personalData)
	fmt.Println(para.MonthId, para.TenantId)
	return "success"
}

//得到用户的考勤，薪资数据
func (p *OperatorPayrollHandler) getUserData(rullList []models.PayrollMapItem, payrollData models.Monthpayroll, attData models.Attendance) map[string]float64 {
	empData := make(map[string]float64)
	for r := range rullList {
		fmt.Println(r)
	}
	return empData
}
