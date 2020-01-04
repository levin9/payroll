package models

import (
	"strings"
)

type VariableProvider struct {
}

func (o *VariableProvider) GetAll(TenantId string) map[string]string {
	var list = &[]Element{}
	result := make(map[string]string)
	//list, total, err := mdl.All(query)
	//fmt.Println("eeeeeeeeeeeeeeeeeee")
	mysqlDB.Raw(" Select * from sa_element").Scan(&list)
	for _, v := range *list {
		result[v.Elemcode] = v.Elemname
		//fmt.Println(v.Elemname)
	}

	return result
}

func (o *VariableProvider) Load(TenantId string) map[string]string {
	result := make(map[string]string)
	result["WorkDay"] = "工作日总数"
	result["PlanWorkDay"] = "应出勤天数"
	result["ActualWorkDay"] = "当月出勤天数"
	result["ChiDaoCount"] = "迟到早退"
	result["0"] = "不扣费"
	result["ChiDaoNum"] = "迟到时长"
	result["MonthlyPayAmount"] = "月薪"
	result["FreeTaxAmount"] = "个税抵扣"
	//result["SheBaoFee"] = "社保扣费"
	payrollService := TenantFormulaService{}
	rulllist, err := payrollService.GetTenantMap(TenantId)
	if err != nil {
		return nil
	}
	for _, r := range *rulllist {
		//社保扣费、住房扣费
		if r.IsHandCalc != 0 && r.IsHandCalc == 1 {
			result[r.FieldName] = r.TenantPayrollName
			continue
		}
		//明显扣费，考勤扣款
		if r.FieldName == "AttFee" {
			result["AttFee"] = r.TenantPayrollName
			continue
		}

		if strings.ContainsAny(r.FieldName, "CustSub") {
			result[r.FieldName] = "当月" + r.TenantPayrollName
			result[r.StandardFieldName] = r.TenantPayrollName
			continue
		}

		result[r.FieldName] = "应发" + r.TenantPayrollName
		if r.StandardFieldName != "" {
			result[r.StandardFieldName] = r.TenantPayrollName
		}
	}

	attCfgs, err1 := payrollService.GetTenantAttCfg(TenantId)
	if err1 != nil {
		return nil
	}
	for _, cfg := range *attCfgs {
		result[cfg.FieldName] = cfg.TenantHolidayName + "扣费"
		result[cfg.AttFieldName] = cfg.TenantHolidayName + "时长"
	}
	return result
}
