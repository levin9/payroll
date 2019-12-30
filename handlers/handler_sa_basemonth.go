package handlers

import (
	"payroll/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("basemonth", basemonthAll)
	groupApi.GET("basemonth/:id",  basemonthOne)
	groupApi.POST("basemonth",  basemonthCreate)
	groupApi.PATCH("basemonth",  basemonthUpdate)
	groupApi.DELETE("basemonth/:id",  basemonthDelete)
}
//All
func basemonthAll(c *gin.Context) {
	mdl := models.Basemonth{}
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
func basemonthOne(c *gin.Context) {
	var mdl models.Basemonth
	
	Id := c.Param("id")
		
	mdl.Id = Id
	data, err := mdl.One()
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}
//Create
func basemonthCreate(c *gin.Context) {
	var mdl models.Basemonth
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
func basemonthUpdate(c *gin.Context) {
	var mdl models.Basemonth
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
func basemonthDelete(c *gin.Context) {
	var mdl models.Basemonth

	
	Id := c.Param("id")
		
	
	mdl.Id = Id
	errResult := mdl.Delete()
	if handleError(c, errResult) {
		return
	}
	jsonSuccess(c)
}
