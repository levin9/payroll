package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/levin9/payroll/blls"
)

func init() {

	groupApi.POST("calculatepayroll", CalculatePayroll)

}

//CalculatePayroll
func CalculatePayroll(c *gin.Context) {
	var para CalcParameter
	err :=c.ShouldBind(&para)
	if c.Handler(c,err){
		return 
	}
	//serv := blls.ClaculateHandler
	//serv..CalculateMonthPayroll("","")
	//blls.CalculateMonthPayroll("","")

	fmt.Println(para)
	//foreach all data

	// var mdl models.GDemo
	// err := c.ShouldBind(&mdl)
	// if handleError(c, err) {
	// 	return
	// }
	// err = mdl.Create()
	// if handleError(c, err) {
	// 	return
	// }
	jsonData(c, nil)
}

type CalcParameter struct{
	MonthId		string 	'json:"month_id" comment:"月份编号"'
	TenantId	string 	'json:"tenant_id" comment:"租户编号"'
}