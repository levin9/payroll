package blls

import (
	"fmt"
	"payroll/models"
	"strconv"
	"strings"
)

type OperatorPayrollHandler struct {
}

//CalculateMonthPayroll
func (p *OperatorPayrollHandler) CalculateMonthPayroll(para CalcParameter) string {
	enginee := ExpressParseEnginee{}
	payrollService := models.TenantFormulaService{}
	rulllist, err := payrollService.GetTenantMap(para.TenantId)
	if err != nil {
		return "falure"
	}
	//加载配置参数
	provider := models.VariableProvider{}
	standVarList := provider.Load(para.TenantId)
	for k, s := range standVarList {
		fmt.Println("映射关系", k, "=>", s)
	}
	userDataList := payrollService.GetAllPayrollList(para.TenantId, para.MonthId)
	for _, userData := range *userDataList {
		enginee.VarList = p.changeToMap(userData, standVarList)
		dao := new(models.UpdatePayrollDao)
		dao.AddId(userData.ID)

		var formulaResult float64
		var formulaText string
		for _, r := range *rulllist {
			if r.IsHandCalc == 1 {
				continue
			}
			if r.TenantFormula == "" || r.FieldName == "" {
				continue
			}
			formulaResult, formulaText, err = enginee.Evaluate(r.TenantFormula)
			if err != nil {
				//break;
				fmt.Println("错误：", r.TenantPayrollName, "=>", r.TenantFormula, "=>", formulaText)
			}
			//更新变量里的值
			enginee.VarList[standVarList[r.FieldName]] = strconv.FormatFloat(float64(formulaResult), 'f', -3, 32)

			dao.AddField(r.FieldName, strconv.FormatFloat(float64(formulaResult), 'f', -3, 32), r.TenantPayrollName, formulaText)
			if r.PayrollId == "" {
				//计算个税
				dao.AddField(`IncomeTaxAmount`, "1234.56", "个人所得税", "")
			}
		}

		//计算应发工资

		//计算实发工资
		dao.SaveChange()
		break
	}
	return "success"
}

func (p *OperatorPayrollHandler) getVarValue(name string, input models.CalcVarItem) string {
	switch name {
	case "MonthlyPayAmount":
		return strconv.FormatFloat(float64(input.MonthlyPayAmount), 'f', -3, 32)
	case "WorkDay":
		return strconv.FormatFloat(float64(input.WorkDay), 'f', -3, 32)
	case "PlanWorkDay":
		return strconv.FormatFloat(float64(input.PlanWorkDay), 'f', -3, 32)
	case "ActualWorkDay":
		return strconv.FormatFloat(float64(input.ActualWorkDay), 'f', -3, 32)
	case "KpiAmout":
		return strconv.FormatFloat(float64(input.KpiAmout), 'f', -3, 32)
	case "TrafficFee":
		return strconv.FormatFloat(float64(input.TrafficFee), 'f', -3, 32)
	case "FoodFee":
		return strconv.FormatFloat(float64(input.FoodFee), 'f', -3, 32)
	case "AttFee":
		return strconv.FormatFloat(float64(input.FoodFee), 'f', -3, 32)
	case "AttAwardFee":
		return strconv.FormatFloat(float64(input.AttAwardFee), 'f', -3, 32)
	case "ChiDaoFee":
		return strconv.FormatFloat(float64(input.ChiDaoFee), 'f', -3, 32)
	case "ChiDaoCount":
		return strconv.FormatFloat(float64(input.ChiDaoCount), 'f', -3, 32)
	case "ZaoTuiCount":
		return strconv.FormatFloat(float64(input.ZaoTuiCount), 'f', -3, 32)
	case "ZaoTuiFee":
		return strconv.FormatFloat(float64(input.ZaoTuiFee), 'f', -3, 32)
	case "SickFee":
		return strconv.FormatFloat(float64(input.FoodFee), 'f', -3, 32)
	case "SickNum":
		return strconv.FormatFloat(float64(input.FoodFee), 'f', -3, 32)
	case "ShiJiaFee":
		return strconv.FormatFloat(float64(input.FoodFee), 'f', -3, 32)
	case "ShiJiaNum":
		return strconv.FormatFloat(float64(input.FoodFee), 'f', -3, 32)
	case "AnnualLeaveFee":
		return strconv.FormatFloat(float64(input.AnnualLeaveFee), 'f', -3, 32)
	case "ProfessionaFee":
		return strconv.FormatFloat(float64(input.ProfessionaFee), 'f', -3, 32)
	case "JobAllowance":
		return strconv.FormatFloat(float64(input.JobAllowance), 'f', -3, 32)
	case "AssignFee":
		return strconv.FormatFloat(float64(input.AssignFee), 'f', -3, 32)
	case "MobileFee":
		return strconv.FormatFloat(float64(input.AssignFee), 'f', -3, 32)
	case "SheBaoFee":
		return strconv.FormatFloat(float64(input.SheBaoFee), 'f', -3, 32)
	case "HouseFee":
		return strconv.FormatFloat(float64(input.HouseFee), 'f', -3, 32)
	case "SpecJobAllowance":
		return strconv.FormatFloat(float64(input.SpecJobAllowance), 'f', -3, 32)
	case "SangJiaFee":
		return strconv.FormatFloat(float64(input.SangJiaFee), 'f', -3, 32)
	case "HunJIaFee":
		return strconv.FormatFloat(float64(input.HunJIaFee), 'f', -3, 32)
	case "BuRvJiaFee":
		return strconv.FormatFloat(float64(input.BuRvJiaFee), 'f', -3, 32)
	case "ChanJiaFee":
		return strconv.FormatFloat(float64(input.ChanJiaFee), 'f', -3, 32)
	case "PeiChanJiaFee":
		return strconv.FormatFloat(float64(input.PeiChanJiaFee), 'f', -3, 32)
	case "KongBaiFee":
		return strconv.FormatFloat(float64(input.KongBaiFee), 'f', -3, 32)
	case "QueQinFee":
		return strconv.FormatFloat(float64(input.QueQinFee), 'f', -3, 32)
	case "TiaoXiuFee":
		return strconv.FormatFloat(float64(input.TiaoXiuFee), 'f', -3, 32)
	case "DaiXinJiaFee":
		return strconv.FormatFloat(float64(input.DaiXinJiaFee), 'f', -3, 32)
	case "WagesPayableAmount":
		return strconv.FormatFloat(float64(input.WagesPayableAmount), 'f', -3, 32)
	case "ActualPayableAmount":
		return strconv.FormatFloat(float64(input.ActualPayableAmount), 'f', -3, 32)
	case "AdjustFee":
		return strconv.FormatFloat(float64(input.AdjustFee), 'f', -3, 32)
	case "FreeTaxAmount":
		return strconv.FormatFloat(float64(input.FreeTaxAmount), 'f', -3, 32)
	case "IncomeTaxAmount":
		return strconv.FormatFloat(float64(input.IncomeTaxAmount), 'f', -3, 32)
	case "CustAddFee1":
		return strconv.FormatFloat(float64(input.CustAddFee1), 'f', -3, 32)
	case "CustAddFee2":
		return strconv.FormatFloat(float64(input.CustAddFee2), 'f', -3, 32)
	case "CustSubFee1":
		return strconv.FormatFloat(float64(input.CustSubFee1), 'f', -3, 32)
	case "CustSubFee2":
		return strconv.FormatFloat(float64(input.CustSubFee2), 'f', -3, 32)
	case "BasicSallaryAmount":
		return strconv.FormatFloat(float64(input.BasicSallaryAmount), 'f', -3, 32)
	case "BasicSallaryAmount01":
		return strconv.FormatFloat(float64(input.BasicSallaryAmount01), 'f', -3, 32)
	case "KpiAmout01":
		return strconv.FormatFloat(float64(input.KpiAmout01), 'f', -3, 32)
	case "TrafficFee01":
		return strconv.FormatFloat(float64(input.TrafficFee01), 'f', -3, 32)
	case "FoodFee01":
		return strconv.FormatFloat(float64(input.FoodFee01), 'f', -3, 32)
	case "ProfessionaFee01":
		return strconv.FormatFloat(float64(input.ProfessionaFee01), 'f', -3, 32)
	case "JobAllowance01":
		return strconv.FormatFloat(float64(input.JobAllowance01), 'f', -3, 32)
	case "AssignFee01":
		return strconv.FormatFloat(float64(input.AssignFee01), 'f', -3, 32)
	case "MobileFee01":
		return strconv.FormatFloat(float64(input.MobileFee01), 'f', -3, 32)
	case "SheBaoFee01":
		return strconv.FormatFloat(float64(input.SheBaoFee01), 'f', -3, 32)
	case "HouseFee01":
		return strconv.FormatFloat(float64(input.HouseFee01), 'f', -3, 32)
	case "AttAwardFee01":
		return strconv.FormatFloat(float64(input.AttAwardFee01), 'f', -3, 32)
	case "SpecJobAllowance01":
		return strconv.FormatFloat(float64(input.SpecJobAllowance01), 'f', -3, 32)
	case "CustAddFee101":
		return strconv.FormatFloat(float64(input.CustAddFee101), 'f', -3, 32)
	case "CustAddFee201":
		return strconv.FormatFloat(float64(input.CustAddFee201), 'f', -3, 32)
	case "CustSubFee101":
		return strconv.FormatFloat(float64(input.CustSubFee101), 'f', -3, 32)
	case "CustSubFee201":
		return strconv.FormatFloat(float64(input.CustSubFee201), 'f', -3, 32)
	case "AnnualLeaveNum":
		return strconv.FormatFloat(float64(input.AnnualLeaveNum), 'f', -3, 32)
		// case "WorkdayQty":
		// 	return strconv.FormatFloat(float64(input.WorkdayQty), 'f', -3, 32)
	}

	return "-99999"
}
func (p *OperatorPayrollHandler) changeToMap(input models.CalcVarItem, standVarList map[string]string) map[string]string {
	var result = make(map[string]string)
	filednamestring := `MonthlyPayAmount,WorkDay,PlanWorkDay,ActualWorkDay,BasicSallaryAmount,KpiAmout,TrafficFee,FoodFee,AttFee,AttAwardFee,ChiDaoFee,ChiDaoCount,ZaotuiCount,ZaotuiFee,SickFee,SickNum,ShiJiaFee,ShiJiaNum,AnnualLeaveNum,AnnualLeaveFee,ProfessionaFee,JobAllowance,AssignFee,MobileFee,SheBaoFee,HouseFee,SpecJobAllowance,SangJiaFee,HunJIaFee,BuRvJiaFee,ChanJiaFee,PeiChanJiaFee,KongBaiFee,QueQinFee,TiaoXiuFee,DaiXinJiaFee,WagesPayableAmount,ActualPayableAmount,AdjustFee,FreeTaxAmount,IncomeTaxAmount,CustAddFee1,CustAddFee2,CustSubFee1,CustSubFee2,BasicSallaryAmount01,KpiAmout01,TrafficFee01,FoodFee01,ProfessionaFee01,JobAllowance01,AssignFee01,MobileFee01,SheBaoFee01,HouseFee01,AttAwardFee01,SpecJobAllowance01,CustAddFee101,CustAddFee201,CustSubFee101,CustSubFee201,WorkdayQty,`
	fields := strings.Split(filednamestring, ",")
	for _, field := range fields {
		tempVar := p.getVarValue(field, input)
		if strings.HasPrefix(tempVar, "-") {
			tempVar = "(" + tempVar + ")"
		}
		result[standVarList[field]] = tempVar
		fmt.Println("添加的变量", field, "=>", standVarList[field], "=>", tempVar)
	}
	return result
}

//得到用户的考勤，薪资数据
// func (p *OperatorPayrollHandler) userDataToMap(empSrcData map[string]interface{}, standVarList map[string]string) map[string]string {
// 	empData := make(map[string]string)

// 	for k := range standVarList {
// 		if empSrcData[k] != 0 {
// 			empData[standVarList[k]] = toString(empSrcData[k]) //Tofloat64(empSrcData[k])
// 		}
// 	}

// 	return empData
// }

func toString(obj interface{}) string {
	switch obj.(type) {
	case string:
		return obj.(string)
	case float32:
	case float64:
		return strconv.FormatFloat(obj.(float64), 'f', -3, 64)

	}
	return "0.01"
}

//Tofloat64
func Tofloat64(inter interface{}) float64 {
	if inter == nil {
		return 0.01
	}
	//fmt.Println("type: ", reflect.TypeOf(inter))
	//fmt.Println("value: ", reflect.ValueOf(inter))

	switch inter.(type) {
	case string:
		break
	case int8:
	case int:
	case uint8:
		return float64(inter.(int))
	case uint:
	case uint32:
	case uint64:
		return float64(inter.(int))
	case int32:
	case int64:
		return float64(inter.(int))
	case float32:
	case float64:
		return float64(inter.(float64))
	case nil:
		return 0
	default:
		return -99999
	}

	return -9999
}
