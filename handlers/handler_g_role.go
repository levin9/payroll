package handlers

import (
	"payroll/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("g-role", gRoleAll)
	groupApi.GET("g-role/:id",  gRoleOne)
	groupApi.POST("g-role",  gRoleCreate)
	groupApi.PATCH("g-role",  gRoleUpdate)
	groupApi.DELETE("g-role/:id",  gRoleDelete)
}
//All
func gRoleAll(c *gin.Context) {
	mdl := models.GRole{}
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
func gRoleOne(c *gin.Context) {
	var mdl models.GRole
	
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
func gRoleCreate(c *gin.Context) {
	var mdl models.GRole
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
func gRoleUpdate(c *gin.Context) {
	var mdl models.GRole
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
func gRoleDelete(c *gin.Context) {
	var mdl models.GRole

	
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
