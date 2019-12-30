package handlers

import (
	"payroll/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("personpayroll", personpayrollAll)
	groupApi.GET("personpayroll/:id",  personpayrollOne)
	groupApi.POST("personpayroll",  personpayrollCreate)
	groupApi.PATCH("personpayroll",  personpayrollUpdate)
	groupApi.DELETE("personpayroll/:id",  personpayrollDelete)
}
//All
func personpayrollAll(c *gin.Context) {
	mdl := models.Personpayroll{}
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
func personpayrollOne(c *gin.Context) {
	var mdl models.Personpayroll
	
	Id := c.Param("id")
		
	mdl.Id = Id
	data, err := mdl.One()
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}
//Create
func personpayrollCreate(c *gin.Context) {
	var mdl models.Personpayroll
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
func personpayrollUpdate(c *gin.Context) {
	var mdl models.Personpayroll
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
func personpayrollDelete(c *gin.Context) {
	var mdl models.Personpayroll

	
	Id := c.Param("id")
		
	
	mdl.Id = Id
	errResult := mdl.Delete()
	if handleError(c, errResult) {
		return
	}
	jsonSuccess(c)
}
