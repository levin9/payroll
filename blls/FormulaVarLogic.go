package blls

type FormulaVarDto struct {
}

//select  a.PayrollId,b.TenantPayrollName,a.FieldName,a.StandardFieldName,b.TenantFormula  from sa_payroll a inner join sa_tenantpayroll b on a.PayrollId=b.PayrollId where TenantId is not NULL and b.IsUsed=1 ORDER BY b.F_SortCode asc
type PayrollMapEntity struct {
	PayrollId         string
	TenantPayrollName string
	FieldName         string
	StandardFieldName string
	TenantFormula     string
}

type FormulaVarLogic struct {
}

func (o *FormulaVarLogic) GetAllVar(TenantId string) (list *[]FormulaVarDto, err error) {
	return nil, nil
}

func (o *FormulaVarLogic) GetTenantMap(TenantId string) (list *[]FormulaVarDto, err error) {
	return nil, nil
}
