package handlers

import (
	"payroll/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("g-user", gUserAll)
	groupApi.GET("g-user/:id",  gUserOne)
	groupApi.POST("g-user",  gUserCreate)
	groupApi.PATCH("g-user",  gUserUpdate)
	groupApi.DELETE("g-user/:id",  gUserDelete)
}
//All
func gUserAll(c *gin.Context) {
	mdl := models.GUser{}
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
func gUserOne(c *gin.Context) {
	var mdl models.GUser
	
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
func gUserCreate(c *gin.Context) {
	var mdl models.GUser
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
func gUserUpdate(c *gin.Context) {
	var mdl models.GUser
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
func gUserDelete(c *gin.Context) {
	var mdl models.GUser

	
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
