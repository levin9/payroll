package handlers

import (
	"fmt"
	"payroll/models"

	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("tenantadjustitem", tenantadjustitemAll)
	groupApi.GET("tenantadjustitem/:id", tenantadjustitemOne)
	groupApi.POST("tenantadjustitem", tenantadjustitemCreate)
	groupApi.PATCH("tenantadjustitem", tenantadjustitemUpdate)
	groupApi.DELETE("tenantadjustitem/:id", tenantadjustitemDelete)
}

//All
func tenantadjustitemAll(c *gin.Context) {
	mdl := models.Tenantadjustitem{}
	query := &models.PaginationQuery{}
	err := c.ShouldBindQuery(query)

	query.Where = WhereIf(c, query.Where, "MonthId", "MonthId='?'")
	query.Where = WhereIf(c, query.Where, "TenantId", "TenantId='?'")
	query.Where = WhereIf(c, query.Where, "keyword", "MonthId='?'")

	fmt.Println("查询条件是：", query.Where)
	if handleError(c, err) {
		return
	}
	list, total, err := mdl.FindAll(query)
	if handleError(c, err) {
		return
	}
	jsonPagination(c, list, total, query)
}

//One
func tenantadjustitemOne(c *gin.Context) {
	var mdl models.Tenantadjustitem

	Id := c.Param("id")

	mdl.Id = Id
	data, err := mdl.One()
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}

//Create
func tenantadjustitemCreate(c *gin.Context) {
	var mdl models.Tenantadjustitem
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
func tenantadjustitemUpdate(c *gin.Context) {
	var mdl models.Tenantadjustitem
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
func tenantadjustitemDelete(c *gin.Context) {
	var mdl models.Tenantadjustitem

	Id := c.Param("id")

	mdl.Id = Id
	errResult := mdl.Delete()
	if handleError(c, errResult) {
		return
	}
	jsonSuccess(c)
}
