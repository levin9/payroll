package handlers

import (
	"payroll/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("attendance", attendanceAll)
	groupApi.GET("attendance/:id",  attendanceOne)
	groupApi.POST("attendance",  attendanceCreate)
	groupApi.PATCH("attendance",  attendanceUpdate)
	groupApi.DELETE("attendance/:id",  attendanceDelete)
}
//All
func attendanceAll(c *gin.Context) {
	mdl := models.Attendance{}
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
func attendanceOne(c *gin.Context) {
	var mdl models.Attendance
	
	Id := c.Param("id")
		
	mdl.Id = Id
	data, err := mdl.One()
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}
//Create
func attendanceCreate(c *gin.Context) {
	var mdl models.Attendance
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
func attendanceUpdate(c *gin.Context) {
	var mdl models.Attendance
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
func attendanceDelete(c *gin.Context) {
	var mdl models.Attendance

	
	Id := c.Param("id")
		
	
	mdl.Id = Id
	errResult := mdl.Delete()
	if handleError(c, errResult) {
		return
	}
	jsonSuccess(c)
}
