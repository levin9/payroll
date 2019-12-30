package handlers

import (
	"payroll/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("g-menu-action", gMenuActionAll)
	groupApi.GET("g-menu-action/:id",  gMenuActionOne)
	groupApi.POST("g-menu-action",  gMenuActionCreate)
	groupApi.PATCH("g-menu-action",  gMenuActionUpdate)
	groupApi.DELETE("g-menu-action/:id",  gMenuActionDelete)
}
//All
func gMenuActionAll(c *gin.Context) {
	mdl := models.GMenuAction{}
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
func gMenuActionOne(c *gin.Context) {
	var mdl models.GMenuAction
	
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
func gMenuActionCreate(c *gin.Context) {
	var mdl models.GMenuAction
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
func gMenuActionUpdate(c *gin.Context) {
	var mdl models.GMenuAction
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
func gMenuActionDelete(c *gin.Context) {
	var mdl models.GMenuAction

	
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
