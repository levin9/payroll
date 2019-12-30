package handlers

import (
	"payroll/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("g-menu-resource", gMenuResourceAll)
	groupApi.GET("g-menu-resource/:id",  gMenuResourceOne)
	groupApi.POST("g-menu-resource",  gMenuResourceCreate)
	groupApi.PATCH("g-menu-resource",  gMenuResourceUpdate)
	groupApi.DELETE("g-menu-resource/:id",  gMenuResourceDelete)
}
//All
func gMenuResourceAll(c *gin.Context) {
	mdl := models.GMenuResource{}
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
func gMenuResourceOne(c *gin.Context) {
	var mdl models.GMenuResource
	
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
func gMenuResourceCreate(c *gin.Context) {
	var mdl models.GMenuResource
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
func gMenuResourceUpdate(c *gin.Context) {
	var mdl models.GMenuResource
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
func gMenuResourceDelete(c *gin.Context) {
	var mdl models.GMenuResource

	
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
