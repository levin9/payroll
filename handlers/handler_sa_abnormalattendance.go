package handlers

import (
	"payroll/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("abnormalattendance", abnormalattendanceAll)
	groupApi.GET("abnormalattendance/:id",  abnormalattendanceOne)
	groupApi.POST("abnormalattendance",  abnormalattendanceCreate)
	groupApi.PATCH("abnormalattendance",  abnormalattendanceUpdate)
	groupApi.DELETE("abnormalattendance/:id",  abnormalattendanceDelete)
}
//All
func abnormalattendanceAll(c *gin.Context) {
	mdl := models.Abnormalattendance{}
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
func abnormalattendanceOne(c *gin.Context) {
	var mdl models.Abnormalattendance
	
	Id := c.Param("id")
		
	mdl.Id = Id
	data, err := mdl.One()
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}
//Create
func abnormalattendanceCreate(c *gin.Context) {
	var mdl models.Abnormalattendance
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
func abnormalattendanceUpdate(c *gin.Context) {
	var mdl models.Abnormalattendance
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
func abnormalattendanceDelete(c *gin.Context) {
	var mdl models.Abnormalattendance

	
	Id := c.Param("id")
		
	
	mdl.Id = Id
	errResult := mdl.Delete()
	if handleError(c, errResult) {
		return
	}
	jsonSuccess(c)
}
