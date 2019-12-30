package handlers

import (
	"payroll/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("operatelog", operatelogAll)
	groupApi.GET("operatelog/:id",  operatelogOne)
	groupApi.POST("operatelog",  operatelogCreate)
	groupApi.PATCH("operatelog",  operatelogUpdate)
	groupApi.DELETE("operatelog/:id",  operatelogDelete)
}
//All
func operatelogAll(c *gin.Context) {
	mdl := models.Operatelog{}
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
func operatelogOne(c *gin.Context) {
	var mdl models.Operatelog
	
	Id := c.Param("id")
		
	mdl.Id = Id
	data, err := mdl.One()
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}
//Create
func operatelogCreate(c *gin.Context) {
	var mdl models.Operatelog
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
func operatelogUpdate(c *gin.Context) {
	var mdl models.Operatelog
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
func operatelogDelete(c *gin.Context) {
	var mdl models.Operatelog

	
	Id := c.Param("id")
		
	
	mdl.Id = Id
	errResult := mdl.Delete()
	if handleError(c, errResult) {
		return
	}
	jsonSuccess(c)
}
