package handlers

import (
	"payroll/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("g-menu", gMenuAll)
	groupApi.GET("g-menu/:id",  gMenuOne)
	groupApi.POST("g-menu",  gMenuCreate)
	groupApi.PATCH("g-menu",  gMenuUpdate)
	groupApi.DELETE("g-menu/:id",  gMenuDelete)
}
//All
func gMenuAll(c *gin.Context) {
	mdl := models.GMenu{}
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
func gMenuOne(c *gin.Context) {
	var mdl models.GMenu
	
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
func gMenuCreate(c *gin.Context) {
	var mdl models.GMenu
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
func gMenuUpdate(c *gin.Context) {
	var mdl models.GMenu
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
func gMenuDelete(c *gin.Context) {
	var mdl models.GMenu

	
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
