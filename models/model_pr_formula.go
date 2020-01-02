package models

type TenantPayrollCfg struct {
	PayrollId         string `gorm:"column:PayrollId"`
	TenantPayrollName string `gorm:"column:TenantPayrollName"`
	FieldName         string `gorm:"column:FieldName"`
	StandardFieldName string `gorm:"column:StandardFieldName"`
	TenantFormula     string `gorm:"column:TenantFormula"`
	IsHandCalc        int    `gorm:"column:IsHandCalc"`
}

type TenantAttCfg struct {
	HolidayCode       string `gorm:"column:HolidayCode"`
	TenantHolidayName string `gorm:"column:TenantHolidayName"`
	AttFieldName      string `gorm:"column:AttFieldName"`
	FieldName         string `gorm:"column:FieldName"`
	TenantFormula     string `gorm:"column:TenantFormula"`
}

type TenantFormulaService struct {
}

func (m *TenantFormulaService) GetTenantMap(TenantId string) (list *[]TenantPayrollCfg, err error) {
	//var tx *gorm.DB
	list = &[]TenantPayrollCfg{}
	sql := "select  a.PayrollId,b.TenantPayrollName,a.FieldName,a.StandardFieldName,b.TenantFormula,a.IsHandCalc  from sa_payroll a inner join sa_tenantpayroll b on a.PayrollId=b.PayrollId where TenantId is not NULL and b.IsUsed>0 ORDER BY b.F_SortCode asc "
	mysqlDB.Raw(sql).Scan(&list)

	return list, nil
}

func (m *TenantFormulaService) GetTenantAttCfg(TenantId string) (list *[]TenantAttCfg, err error) {
	//var tx *gorm.DB
	list = &[]TenantAttCfg{}
	sql := "select a.HolidayCode,b.TenantHolidayName,AttFieldName,a.FieldName,b.TenantFormula from sa_attcategory a inner join sa_attcategorytenant b on a.HolidayCode=b.HolidayCode where b.IsUsed=1 order by b.F_SortCode asc  "
	mysqlDB.Raw(sql).Scan(&list)

	return list, nil
}

func (m *TenantFormulaService) GetAllPayrollList(TenantId, MonthId string) (list *[]CalcVarItem) {
	//
	sql := ` select  a.ID,
	a.TenantId,
	a.MonthId,
	a.PersonId,
	a.RealName,
	a.DeptName,
	a.JobLevel,
	a.JobName,
	c.JoinDate,
	a.ExtInfo,
	if(a.MonthlyPayAmount is null,0,a.MonthlyPayAmount) as MonthlyPayAmount,
	if(a.PlanWorkDay  is null,0,a.PlanWorkDay ) as PlanWorkDay ,
	if(a.ActualWorkDay is null,0,a.ActualWorkDay ) as ActualWorkDay  ,
	if(a.BasicSallaryAmount is null,0,a.BasicSallaryAmount) as BasicSallaryAmount,
	if(a.KpiAmout is null,0,a.KpiAmout) as KpiAmout,
	if(a.TrafficFee is null,0,a.TrafficFee) as TrafficFee ,
	if(a.FoodFee is null,0,a.FoodFee) as FoodFee ,
	if(a.AttFee is null,0,a.AttFee) as AttFee ,
	if(a.AttAwardFee is null,0,a.AttAwardFee) as AttAwardFee ,
	if(a.ChiDaoFee is null,0,a.ChiDaoFee) as ChiDaoFee ,
	if(a.ChiDaoCount is null,0,a.ChiDaoCount) as ChiDaoCount ,
	if(a.ZaoTuiCount is null,0,a.ZaoTuiCount) as ZaoTuiCount ,
	if(a.ZaoTuiFee is null,0,a.ZaoTuiFee) as ZaoTuiFee ,
	if(a.SickFee is null,0,a.SickFee) as SickFee ,
	if(a.SickNum is null,0,a.SickNum) as SickNum ,
	if(a.ShiJiaFee is null,0,a.ShiJiaFee) as ShiJiaFee ,
	if(a.ShiJiaNum is null,0,a.ShiJiaNum) as ShiJiaNum ,
	if(a.AnnualLeaveNum is null,0,a.AnnualLeaveNum) as AnnualLeaveNum ,
	if(a.AnnualLeaveFee is null,0,a.AnnualLeaveFee ) as AnnualLeaveFee ,
	if(a.ProfessionaFee is null,0,a.ProfessionaFee) as ProfessionaFee ,
	if(a.JobAllowance is null,0,a.JobAllowance) as JobAllowance ,
	if(a.AssignFee is null,0,a.AssignFee) as AssignFee ,
	if(a.MobileFee is null,0,a.MobileFee) as MobileFee ,
	if(a.SheBaoFee is null,0,a.SheBaoFee) as SheBaoFee ,
	if(a.HouseFee is null,0,a.HouseFee) as HouseFee ,
	if(a.SpecJobAllowance is null,0,a.SpecJobAllowance) as SpecJobAllowance ,
	if(a.SangJiaFee is null,0,a.SangJiaFee) as SangJiaFee ,
	if(a.HunJIaFee is null,0,a.HunJIaFee) as HunJIaFee ,
	if(a.BuRvJiaFee is null,0,a.BuRvJiaFee) as BuRvJiaFee ,
	if(a.ChanJiaFee is null,0,a.ChanJiaFee) as ChanJiaFee ,
	if(a.PeiChanJiaFee is null,0,a.PeiChanJiaFee) as PeiChanJiaFee ,
	if(a.KongBaiFee is null,0,a.KongBaiFee) as KongBaiFee ,
	if(a.QueQinFee is null,0,a.QueQinFee) as QueQinFee ,
	if(a.TiaoXiuFee is null,0,a.TiaoXiuFee) as TiaoXiuFee ,
	if(a.DaiXinJiaFee is null,0,a.DaiXinJiaFee)as DaiXinJiaFee ,
	if(a.WagesPayableAmount is null,0,a.WagesPayableAmount) as WagesPayableAmount ,
	if(a.ActualPayableAmount is null,0,a.ActualPayableAmount) as ActualPayableAmount ,
	if(a.AdjustFee is null,0,a.AdjustFee) as AdjustFee ,
	if(a.FreeTaxAmount is null,0,a.FreeTaxAmount) as FreeTaxAmount ,
	if(a.IncomeTaxAmount is null,0,a.IncomeTaxAmount) as IncomeTaxAmount ,
	if(a.CustAddFee1 is null,0,a.CustAddFee1) as CustAddFee1 ,
	if(a.CustAddFee2 is null,0,a.CustAddFee2) as CustAddFee2 ,
	if(a.CustSubFee1 is null,0,a.CustSubFee1) as CustSubFee1 ,
	if(a.CustSubFee2 is null,0,a.CustSubFee2) as CustSubFee2 ,	
	if(b.BasicSallaryAmount is null,0,b.BasicSallaryAmount) as BasicSallaryAmount01,
	if(b.KpiAmout is null,0,b.KpiAmout) as KpiAmout01,
	if(b.TrafficFee is null,0,b.TrafficFee)  as TrafficFee01 ,
	if(b.FoodFee  is null,0,b.FoodFee) as FoodFee01,	                                   
	if(b.ProfessionaFee  is null,0,b.ProfessionaFee) as ProfessionaFee01,
	if(b.JobAllowance  is null,0,b.JobAllowance) as JobAllowance01,
	if(b.AssignFee  is null,0,b.AssignFee) as AssignFee01,
	if(b.MobileFee  is null,0,b.MobileFee) as MobileFee01,
	if(b.SheBaoFee  is null,0,b.SheBaoFee) as SheBaoFee01,
	if(b.HouseFee   is null,0,b.HouseFee) as HouseFee01,
	if(b.AttAwardFee  is null,0,b.AttAwardFee) as AttAwardFee01,
	if(b.SpecJobAllowance  is null,0,b.SpecJobAllowance) as SpecJobAllowance01,	
	if(b.CustAddFee1 is null,0,b.CustAddFee1) as CustAddFee101 ,
	if(b.CustAddFee2 is null,0,b.CustAddFee2) as CustAddFee201 ,
	if(b.CustSubFee1 is null,0,b.CustSubFee1) as CustSubFee101 ,
	if(b.CustSubFee2 is null,0,b.CustSubFee2) as CustSubFee201 ,
	if(d.WorkdayQty  is null,0,d.WorkdayQty) as WorkDay
from sa_monthpayroll a
inner join sa_personpayroll b on a.PersonId=b.PersonId and b.ValidDate<now() and b.InvalidDate>now()
inner join sa_personal c on c.PersonId=a.PersonId
inner join sa_basemonth d on d.MonthId=a.MonthId and d.TenantID=a.TenantId
where b.InvalidDate>d.StartDate   `
	list = &[]CalcVarItem{}
	mysqlDB.Raw(sql).Scan(&list)

	return list
}

func (m *TenantFormulaService) GetAllPayrollData(TenantId, MonthId string) []map[string]interface{} {
	//
	sql := ` select  a.ID,
	a.TenantId,
	a.MonthId,
	a.PersonId,
	a.RealName,
	a.DeptName,
	a.JobLevel,
	a.JobName,
	if(a.MonthlyPayAmount is null,0,a.MonthlyPayAmount) as MonthlyPayAmount,
	if(a.PlanWorkDay  is null,0,a.PlanWorkDay ) as PlanWorkDay ,
	if(a.ActualWorkDay is null,0,a.ActualWorkDay ) as ActualWorkDay  ,
	if(a.BasicSallaryAmount is null,0,a.BasicSallaryAmount) as BasicSallaryAmount,
	if(a.KpiAmout is null,0,a.KpiAmout) as KpiAmout,
	if(a.TrafficFee is null,0,a.TrafficFee) as TrafficFee ,
	if(a.FoodFee is null,0,a.FoodFee) as FoodFee ,
	if(a.AttFee is null,0,a.AttFee) as AttFee ,
	if(a.AttAwardFee is null,0,a.AttAwardFee) as AttAwardFee ,
	if(a.ChiDaoFee is null,0,a.ChiDaoFee) as ChiDaoFee ,
	if(a.ChiDaoCount is null,0,a.ChiDaoCount) as ChiDaoCount ,
	if(a.ZaotuiCount is null,0,a.ZaotuiCount) as ZaotuiCount ,
	if(a.ZaotuiFee is null,0,a.ZaotuiFee) as ZaotuiFee ,
	if(a.SickFee is null,0,a.SickFee) as SickFee ,
	if(a.SickNum is null,0,a.SickNum) as SickNum ,
	if(a.ShiJiaFee is null,0,a.ShiJiaFee) as ShiJiaFee ,
	if(a.ShiJiaNum is null,0,a.ShiJiaNum) as ShiJiaNum ,
	if(a.AnnualLeaveNum is null,0,a.AnnualLeaveNum) as AnnualLeaveNum ,
	if(a.AnnualLeaveFee is null,0,a.AnnualLeaveFee ) as AnnualLeaveFee ,
	if(a.ProfessionaFee is null,0,a.ProfessionaFee) as ProfessionaFee ,
	if(a.JobAllowance is null,0,a.JobAllowance) as JobAllowance ,
	if(a.AssignFee is null,0,a.AssignFee) as AssignFee ,
	if(a.MobileFee is null,0,a.MobileFee) as MobileFee ,
	if(a.SheBaoFee is null,0,a.SheBaoFee) as SheBaoFee ,
	if(a.HouseFee is null,0,a.HouseFee) as HouseFee ,
	if(a.SpecJobAllowance is null,0,a.SpecJobAllowance) as SpecJobAllowance ,
	if(a.SangJiaFee is null,0,a.SangJiaFee) as SangJiaFee ,
	if(a.HunJIaFee is null,0,a.HunJIaFee) as HunJIaFee ,
	if(a.BuRvJiaFee is null,0,a.BuRvJiaFee) as BuRvJiaFee ,
	if(a.ChanJiaFee is null,0,a.ChanJiaFee) as ChanJiaFee ,
	if(a.PeiChanJiaFee is null,0,a.PeiChanJiaFee) as PeiChanJiaFee ,
	if(a.KongBaiFee is null,0,a.KongBaiFee) as KongBaiFee ,
	if(a.QueQinFee is null,0,a.QueQinFee) as QueQinFee ,
	if(a.TiaoXiuFee is null,0,a.TiaoXiuFee) as TiaoXiuFee ,
	if(a.DaiXinJiaFee is null,0,a.DaiXinJiaFee)as DaiXinJiaFee ,
	if(a.WagesPayableAmount is null,0,a.WagesPayableAmount) as WagesPayableAmount ,
	if(a.ActualPayableAmount is null,0,a.ActualPayableAmount) as ActualPayableAmount ,
	if(a.AdjustFee is null,0,a.AdjustFee) as AdjustFee ,
	if(a.FreeTaxAmount is null,0,a.FreeTaxAmount) as FreeTaxAmount ,
	if(a.IncomeTaxAmount is null,0,a.IncomeTaxAmount) as IncomeTaxAmount ,
	if(a.CustAddFee1 is null,0,a.CustAddFee1) as CustAddFee1 ,
	if(a.CustAddFee2 is null,0,a.CustAddFee2) as CustAddFee2 ,
	if(a.CustSubFee1 is null,0,a.CustSubFee1) as CustSubFee1 ,
	if(a.CustSubFee2 is null,0,a.CustSubFee2) as CustSubFee2 ,
	a.ExtInfo,
	c.RealName,
	c.DingId,
	c.PersonId,
	c.JoinDate,
	if(b.BasicSallaryAmount is null,0,b.BasicSallaryAmount) as BasicSallaryAmount01,
	if(b.KpiAmout is null,0,b.KpiAmout) as KpiAmout01,
	if(b.TrafficFee is null,0,b.TrafficFee)  as TrafficFee01 ,
	if(b.FoodFee  is null,0,b.FoodFee) as FoodFee01,	                                   
	if(b.ProfessionaFee  is null,0,b.ProfessionaFee) as ProfessionaFee01,
	if(b.JobAllowance  is null,0,b.JobAllowance) as JobAllowance01,
	if(b.AssignFee  is null,0,b.AssignFee) as AssignFee01,
	if(b.MobileFee  is null,0,b.MobileFee) as MobileFee01,
	if(b.SheBaoFee  is null,0,b.SheBaoFee) as SheBaoFee01,
	if(b.HouseFee   is null,0,b.HouseFee) as HouseFee01,
	if(b.AttAwardFee  is null,0,b.AttAwardFee) as AttAwardFee01,
	if(b.SpecJobAllowance  is null,0,b.SpecJobAllowance) as SpecJobAllowance01,	
	if(b.CustAddFee1 is null,0,b.CustAddFee1) as CustAddFee101 ,
	if(b.CustAddFee2 is null,0,b.CustAddFee2) as CustAddFee201 ,
	if(b.CustSubFee1 is null,0,b.CustSubFee1) as CustSubFee101 ,
	if(b.CustSubFee2 is null,0,b.CustSubFee2) as CustSubFee201 ,
	if(d.WorkdayQty  is null,0,d.WorkdayQty) as WorkDay
from sa_monthpayroll a
inner join sa_personpayroll b on a.PersonId=b.PersonId and b.ValidDate<now() and b.InvalidDate>now()
inner join sa_personal c on c.PersonId=a.PersonId
inner join sa_basemonth d on d.MonthId=a.MonthId and d.TenantID=a.TenantId
where b.InvalidDate>d.StartDate   `

	personData, _ := DoQuery(sql)

	return personData
}
