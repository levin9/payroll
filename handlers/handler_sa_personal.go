package handlers

import (
	"payroll/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("personal", personalAll)
	groupApi.GET("personal/:id",  personalOne)
	groupApi.POST("personal",  personalCreate)
	groupApi.PATCH("personal",  personalUpdate)
	groupApi.DELETE("personal/:id",  personalDelete)
}
//All
func personalAll(c *gin.Context) {
	mdl := models.Personal{}
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
func personalOne(c *gin.Context) {
	var mdl models.Personal
	
	Personid := c.Param("id")
		
	mdl.Personid = Personid
	data, err := mdl.One()
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}
//Create
func personalCreate(c *gin.Context) {
	var mdl models.Personal
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
func personalUpdate(c *gin.Context) {
	var mdl models.Personal
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
func personalDelete(c *gin.Context) {
	var mdl models.Personal

	
	Personid := c.Param("id")
		
	
	mdl.Personid = Personid
	errResult := mdl.Delete()
	if handleError(c, errResult) {
		return
	}
	jsonSuccess(c)
}
