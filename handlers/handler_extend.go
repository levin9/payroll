package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/levin9/payroll/blls"
)

func init() {

	groupApi.POST("calculatepayroll", CalculatePayroll)

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
