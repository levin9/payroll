package handlers

import (
	"payroll/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("messagerequest", messagerequestAll)
	groupApi.GET("messagerequest/:id",  messagerequestOne)
	groupApi.POST("messagerequest",  messagerequestCreate)
	groupApi.PATCH("messagerequest",  messagerequestUpdate)
	groupApi.DELETE("messagerequest/:id",  messagerequestDelete)
}
//All
func messagerequestAll(c *gin.Context) {
	mdl := models.Messagerequest{}
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
func messagerequestOne(c *gin.Context) {
	var mdl models.Messagerequest
	
	Id := c.Param("id")
		
	mdl.Id = Id
	data, err := mdl.One()
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}
//Create
func messagerequestCreate(c *gin.Context) {
	var mdl models.Messagerequest
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
func messagerequestUpdate(c *gin.Context) {
	var mdl models.Messagerequest
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
func messagerequestDelete(c *gin.Context) {
	var mdl models.Messagerequest

	
	Id := c.Param("id")
		
	
	mdl.Id = Id
	errResult := mdl.Delete()
	if handleError(c, errResult) {
		return
	}
	jsonSuccess(c)
}
