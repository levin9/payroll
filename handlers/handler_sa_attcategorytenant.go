package handlers

import (
	"payroll/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("attcategorytenant", attcategorytenantAll)
	groupApi.GET("attcategorytenant/:id",  attcategorytenantOne)
	groupApi.POST("attcategorytenant",  attcategorytenantCreate)
	groupApi.PATCH("attcategorytenant",  attcategorytenantUpdate)
	groupApi.DELETE("attcategorytenant/:id",  attcategorytenantDelete)
}
//All
func attcategorytenantAll(c *gin.Context) {
	mdl := models.Attcategorytenant{}
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
func attcategorytenantOne(c *gin.Context) {
	var mdl models.Attcategorytenant
	
	Holidaycategoryid := c.Param("id")
		
	mdl.Holidaycategoryid = Holidaycategoryid
	data, err := mdl.One()
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}
//Create
func attcategorytenantCreate(c *gin.Context) {
	var mdl models.Attcategorytenant
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
func attcategorytenantUpdate(c *gin.Context) {
	var mdl models.Attcategorytenant
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
func attcategorytenantDelete(c *gin.Context) {
	var mdl models.Attcategorytenant

	
	Holidaycategoryid := c.Param("id")
		
	
	mdl.Holidaycategoryid = Holidaycategoryid
	errResult := mdl.Delete()
	if handleError(c, errResult) {
		return
	}
	jsonSuccess(c)
}
