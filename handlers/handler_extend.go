package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"payroll/blls"
	"payroll/models"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
)

func init() {

	groupApi.POST("calculatepayroll", CalculatePayroll)
	groupApi.GET("loadvariable/:id", LoadVariable)
	groupApi.POST("downloadAttendance", downloadAttendance)

}

//CalculatePayroll
func downloadAttendance(c *gin.Context) {
	var para blls.CalcParameter
	err := c.ShouldBind(&para)
	if handleError(c, err) {
		return
	}
	dao := models.Attendance{}
	attList, _, err := dao.All(models.CreateCondition(para.MonthId, para.TenantId))
	if err != nil {

	}
	xlsxNews := excelize.NewFile()

	xlsxNews.SetSheetRow("Sheet1", "A1", &[]interface{}{"部门", "姓名", "实际工作日"})
	for index, att := range *attList {
		lint := strconv.Itoa(index + 2)
		xlsxNews.SetSheetRow("Sheet1", "A"+lint, &[]interface{}{att.DeptName, att.RealName, att.Actualworkday})
	}
	//xlsxNews.SaveAs("360.xlsx")

	// c.Request.Response.Header.Set("Content-Type", "application/octet-stream")
	// c.Request.Response.Header.Set("Content-Disposition", "attachment; filename="+"360.xlsx")
	// c.Request.Response.Header.Set("Content-Transfer-Encoding", "binary")

	// //c.Request.Response.Header.Set("Content-Type", "application/vnd.ms-excel;charset=UTF-8")
	// //回写到web 流媒体 形成下载
	// _ = xlsxNews.Write(c.Writer)
	buf := new(bytes.Buffer)
	err = xlsxNews.Write(buf)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename=contacts.xlsx")
	c.Data(http.StatusOK, "text/xlsx", buf.Bytes())
	

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
