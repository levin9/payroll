package handlers

import (
	"payroll/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("tenant", tenantAll)
	groupApi.GET("tenant/:id",  tenantOne)
	groupApi.POST("tenant",  tenantCreate)
	groupApi.PATCH("tenant",  tenantUpdate)
	groupApi.DELETE("tenant/:id",  tenantDelete)
}
//All
func tenantAll(c *gin.Context) {
	mdl := models.Tenant{}
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
func tenantOne(c *gin.Context) {
	var mdl models.Tenant
	
	Tenantid := c.Param("id")
		
	mdl.Tenantid = Tenantid
	data, err := mdl.One()
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}
//Create
func tenantCreate(c *gin.Context) {
	var mdl models.Tenant
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
func tenantUpdate(c *gin.Context) {
	var mdl models.Tenant
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
func tenantDelete(c *gin.Context) {
	var mdl models.Tenant

	
	Tenantid := c.Param("id")
		
	
	mdl.Tenantid = Tenantid
	errResult := mdl.Delete()
	if handleError(c, errResult) {
		return
	}
	jsonSuccess(c)
}
