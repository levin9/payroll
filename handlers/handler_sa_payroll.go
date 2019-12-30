package handlers

import (
	"payroll/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("payroll", payrollAll)
	groupApi.GET("payroll/:id",  payrollOne)
	groupApi.POST("payroll",  payrollCreate)
	groupApi.PATCH("payroll",  payrollUpdate)
	groupApi.DELETE("payroll/:id",  payrollDelete)
}
//All
func payrollAll(c *gin.Context) {
	mdl := models.Payroll{}
	query := &models.PaginationQuery{}
	err := c.ShouldBindQuery(query)
	if handleError(c, err) {
		return
	}
	list, total, err := mdl.All(query)
	if handleError(c, err) {
		return
	}
	jsonPagination(c, list, total, query)
}
//One
func payrollOne(c *gin.Context) {
	var mdl models.Payroll
	
	Payrollid := c.Param("id")
		
	mdl.Payrollid = Payrollid
	data, err := mdl.One()
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}
//Create
func payrollCreate(c *gin.Context) {
	var mdl models.Payroll
	err := c.ShouldBind(&mdl)
	if handleError(c, err) {
		return
	}
	err = mdl.Create()
	if handleError(c, err) {
		return
	}
	jsonData(c, mdl)
}
//Update
func payrollUpdate(c *gin.Context) {
	var mdl models.Payroll
	err := c.ShouldBind(&mdl)
	if handleError(c, err) {
		return
	}
	err = mdl.Update()
	if handleError(c, err) {
		return
	}
	jsonSuccess(c)
}
//Delete
func payrollDelete(c *gin.Context) {
	var mdl models.Payroll

	
	Payrollid := c.Param("id")
		
	
	mdl.Payrollid = Payrollid
	errResult := mdl.Delete()
	if handleError(c, errResult) {
		return
	}
	jsonSuccess(c)
}
