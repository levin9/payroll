package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/levin9/payroll/blls"
	"payroll/models"
)

func init() {

	groupApi.POST("calculatepayroll", CalculatePayroll)
	groupApi.GET("loadvariable/:id", LoadVariable)

}

//CalculatePayroll
func CalculatePayroll(c *gin.Context) {
	var para blls.CalcParameter
	err := c.ShouldBind(&para)
	if handleError(c, err) {
		return
	}

	payrollHandler := blls.OperatorPayrollHandler{}
	result := payrollHandler.CalculateMonthPayroll(para)
	fmt.Println(result)
	jsonData(c, result)
}

//LoadVariable
func LoadVariable(c *gin.Context) {
	provider := models.VariableProvider{}
	TenantId := c.Param("id")
	list := provider.Load(TenantId)
	jsonData(c, list)
}
