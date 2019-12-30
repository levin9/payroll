package handlers

import (
	"payroll/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("g-user-role", gUserRoleAll)
	groupApi.GET("g-user-role/:id",  gUserRoleOne)
	groupApi.POST("g-user-role",  gUserRoleCreate)
	groupApi.PATCH("g-user-role",  gUserRoleUpdate)
	groupApi.DELETE("g-user-role/:id",  gUserRoleDelete)
}
//All
func gUserRoleAll(c *gin.Context) {
	mdl := models.GUserRole{}
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
func gUserRoleOne(c *gin.Context) {
	var mdl models.GUserRole
	
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
func gUserRoleCreate(c *gin.Context) {
	var mdl models.GUserRole
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
func gUserRoleUpdate(c *gin.Context) {
	var mdl models.GUserRole
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
func gUserRoleDelete(c *gin.Context) {
	var mdl models.GUserRole

	
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
