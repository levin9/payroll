package handlers

import (
	"payroll/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("personfreetax", personfreetaxAll)
	groupApi.GET("personfreetax/:id",  personfreetaxOne)
	groupApi.POST("personfreetax",  personfreetaxCreate)
	groupApi.PATCH("personfreetax",  personfreetaxUpdate)
	groupApi.DELETE("personfreetax/:id",  personfreetaxDelete)
}
//All
func personfreetaxAll(c *gin.Context) {
	mdl := models.Personfreetax{}
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
func personfreetaxOne(c *gin.Context) {
	var mdl models.Personfreetax
	
	Id := c.Param("id")
		
	mdl.Id = Id
	data, err := mdl.One()
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}
//Create
func personfreetaxCreate(c *gin.Context) {
	var mdl models.Personfreetax
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
func personfreetaxUpdate(c *gin.Context) {
	var mdl models.Personfreetax
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
func personfreetaxDelete(c *gin.Context) {
	var mdl models.Personfreetax

	
	Id := c.Param("id")
		
	
	mdl.Id = Id
	errResult := mdl.Delete()
	if handleError(c, errResult) {
		return
	}
	jsonSuccess(c)
}
