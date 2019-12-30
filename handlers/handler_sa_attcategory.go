package handlers

import (
	"payroll/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("attcategory", attcategoryAll)
	groupApi.GET("attcategory/:id",  attcategoryOne)
	groupApi.POST("attcategory",  attcategoryCreate)
	groupApi.PATCH("attcategory",  attcategoryUpdate)
	groupApi.DELETE("attcategory/:id",  attcategoryDelete)
}
//All
func attcategoryAll(c *gin.Context) {
	mdl := models.Attcategory{}
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
func attcategoryOne(c *gin.Context) {
	var mdl models.Attcategory
	
	Id := c.Param("id")
		
	mdl.Id = Id
	data, err := mdl.One()
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}
//Create
func attcategoryCreate(c *gin.Context) {
	var mdl models.Attcategory
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
func attcategoryUpdate(c *gin.Context) {
	var mdl models.Attcategory
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
func attcategoryDelete(c *gin.Context) {
	var mdl models.Attcategory

	
	Id := c.Param("id")
		
	
	mdl.Id = Id
	errResult := mdl.Delete()
	if handleError(c, errResult) {
		return
	}
	jsonSuccess(c)
}
