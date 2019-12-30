package handlers

import (
	"payroll/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("g-demo", gDemoAll)
	groupApi.GET("g-demo/:id",  gDemoOne)
	groupApi.POST("g-demo",  gDemoCreate)
	groupApi.PATCH("g-demo",  gDemoUpdate)
	groupApi.DELETE("g-demo/:id",  gDemoDelete)
}
//All
func gDemoAll(c *gin.Context) {
	mdl := models.GDemo{}
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
func gDemoOne(c *gin.Context) {
	var mdl models.GDemo
	
	Id, err := parseParamID(c)
	if handleError(c, err) {
		return
	}
    	
	mdl.Id = Id
	data, err := mdl.One()
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}
//Create
func gDemoCreate(c *gin.Context) {
	var mdl models.GDemo
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
func gDemoUpdate(c *gin.Context) {
	var mdl models.GDemo
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
func gDemoDelete(c *gin.Context) {
	var mdl models.GDemo

	
	Id, err := parseParamID(c)
	if handleError(c, err) {
		return
	}
    	
	
	mdl.Id = Id
	errResult := mdl.Delete()
	if handleError(c, errResult) {
		return
	}
	jsonSuccess(c)
}
