package handlers

import (
	"github.com/gin-gonic/gin"
	"payroll/models"
	"strconv"
)

func init() {
	groupApi.GET("element", elementAll)
	groupApi.GET("element/:id", elementOne)
	groupApi.POST("element", elementCreate)
	groupApi.PATCH("element", elementUpdate)
	groupApi.DELETE("element/:id", elementDelete)
}

//All
func elementAll(c *gin.Context) {
	mdl := models.Element{}
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
func elementOne(c *gin.Context) {
	var mdl models.Element

	Id := c.Param("id")

	mdl.Id, _ = strconv.Atoi(Id)
	data, err := mdl.One()
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}

//Create
func elementCreate(c *gin.Context) {
	var mdl models.Element
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
func elementUpdate(c *gin.Context) {
	var mdl models.Element
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
func elementDelete(c *gin.Context) {
	var mdl models.Element

	Id := c.Param("id")

	mdl.Id, _ = strconv.Atoi(Id)
	errResult := mdl.Delete()
	if handleError(c, errResult) {
		return
	}
	jsonSuccess(c)
}
