package handlers

import (
	"payroll/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("tenantpayroll", tenantpayrollAll)
	groupApi.GET("tenantpayroll/:id",  tenantpayrollOne)
	groupApi.POST("tenantpayroll",  tenantpayrollCreate)
	groupApi.PATCH("tenantpayroll",  tenantpayrollUpdate)
	groupApi.DELETE("tenantpayroll/:id",  tenantpayrollDelete)
}
//All
func tenantpayrollAll(c *gin.Context) {
	mdl := models.Tenantpayroll{}
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
func tenantpayrollOne(c *gin.Context) {
	var mdl models.Tenantpayroll
	
	Tpid := c.Param("id")
		
	mdl.Tpid = Tpid
	data, err := mdl.One()
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}
//Create
func tenantpayrollCreate(c *gin.Context) {
	var mdl models.Tenantpayroll
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
func tenantpayrollUpdate(c *gin.Context) {
	var mdl models.Tenantpayroll
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
func tenantpayrollDelete(c *gin.Context) {
	var mdl models.Tenantpayroll

	
	Tpid := c.Param("id")
		
	
	mdl.Tpid = Tpid
	errResult := mdl.Delete()
	if handleError(c, errResult) {
		return
	}
	jsonSuccess(c)
}
