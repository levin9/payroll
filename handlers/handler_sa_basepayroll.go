package handlers

import (
	"payroll/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("basepayroll", basepayrollAll)
	groupApi.GET("basepayroll/:id",  basepayrollOne)
	groupApi.POST("basepayroll",  basepayrollCreate)
	groupApi.PATCH("basepayroll",  basepayrollUpdate)
	groupApi.DELETE("basepayroll/:id",  basepayrollDelete)
}
//All
func basepayrollAll(c *gin.Context) {
	mdl := models.Basepayroll{}
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
func basepayrollOne(c *gin.Context) {
	var mdl models.Basepayroll
	
	Pid := c.Param("id")
		
	mdl.Pid = Pid
	data, err := mdl.One()
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}
//Create
func basepayrollCreate(c *gin.Context) {
	var mdl models.Basepayroll
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
func basepayrollUpdate(c *gin.Context) {
	var mdl models.Basepayroll
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
func basepayrollDelete(c *gin.Context) {
	var mdl models.Basepayroll

	
	Pid := c.Param("id")
		
	
	mdl.Pid = Pid
	errResult := mdl.Delete()
	if handleError(c, errResult) {
		return
	}
	jsonSuccess(c)
}
