package handlers

import (
	"github.com/gin-gonic/gin"
	"payroll/models"
)

func init() {
	groupApi.GET("monthpayroll", monthpayrollAll)
	groupApi.GET("monthpayroll/:id", monthpayrollOne)
	groupApi.POST("monthpayroll", monthpayrollCreate)
	groupApi.PATCH("monthpayroll", monthpayrollUpdate)
	groupApi.DELETE("monthpayroll/:id", monthpayrollDelete)
}

//All
func monthpayrollAll(c *gin.Context) {
	mdl := models.Monthpayroll{}
	query := &models.PaginationQuery{}
	err := c.ShouldBindQuery(query)
	if handleError(c, err) {
		return
	}
	query.Where = WhereIf(c, query.Where, "MonthId", "MonthId='?'")
	query.Where = WhereIf(c, query.Where, "TenantId", "TenantId='?'")
	//query.Where = WhereIf(c, query.Where, "keyword", "MonthId='?'")

	list, total, err := mdl.FindAll(query)
	if handleError(c, err) {
		return
	}
	jsonPagination(c, list, total, query)
}

//One
func monthpayrollOne(c *gin.Context) {
	var mdl models.Monthpayroll

	Id := c.Param("id")

	mdl.Id = Id
	data, err := mdl.One()
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}

//Create
func monthpayrollCreate(c *gin.Context) {
	var mdl models.Monthpayroll
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
func monthpayrollUpdate(c *gin.Context) {
	var mdl models.Monthpayroll
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
func monthpayrollDelete(c *gin.Context) {
	var mdl models.Monthpayroll

	Id := c.Param("id")

	mdl.Id = Id
	errResult := mdl.Delete()
	if handleError(c, errResult) {
		return
	}
	jsonSuccess(c)
}
