package handlers

import (
	"payroll/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("g-role-menu", gRoleMenuAll)
	groupApi.GET("g-role-menu/:id",  gRoleMenuOne)
	groupApi.POST("g-role-menu",  gRoleMenuCreate)
	groupApi.PATCH("g-role-menu",  gRoleMenuUpdate)
	groupApi.DELETE("g-role-menu/:id",  gRoleMenuDelete)
}
//All
func gRoleMenuAll(c *gin.Context) {
	mdl := models.GRoleMenu{}
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
func gRoleMenuOne(c *gin.Context) {
	var mdl models.GRoleMenu
	
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
func gRoleMenuCreate(c *gin.Context) {
	var mdl models.GRoleMenu
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
func gRoleMenuUpdate(c *gin.Context) {
	var mdl models.GRoleMenu
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
func gRoleMenuDelete(c *gin.Context) {
	var mdl models.GRoleMenu

	
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
