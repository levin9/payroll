package models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

type PayrollMapItem struct {
	PayrollId         string
	TenantPayrollName string
	FieldName         string
	StandardFieldName string
	TenantFormula     string
}

type TenantFormulaService struct {
}

func (m *TenantFormulaService) GetTenantMap(TenantId string) (list *[]PayrollMapItem, err error) {
	var tx *gorm.DB
	list = &[]PayrollMapItem{}
	sql := "select  a.PayrollId,b.TenantPayrollName,a.FieldName,a.StandardFieldName,b.TenantFormula  from sa_payroll a inner join sa_tenantpayroll b on a.PayrollId=b.PayrollId where TenantId is not NULL and b.IsUsed=1 ORDER BY b.F_SortCode asc "
	tx.DB().QueryRow(sql).Scan(&list)

	return list, nil
}

func (m *TenantFormulaService) GetAllPayrollData(TenantId, MonthId string) []map[string]interface{} {
	//
	var tx *sql.DB
	sql := ""
	personData, _ := DoQuery(tx, sql)

	return personData
}
