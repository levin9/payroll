package handlers

import (
	"payroll/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("bankaccount", bankaccountAll)
	groupApi.GET("bankaccount/:id",  bankaccountOne)
	groupApi.POST("bankaccount",  bankaccountCreate)
	groupApi.PATCH("bankaccount",  bankaccountUpdate)
	groupApi.DELETE("bankaccount/:id",  bankaccountDelete)
}
//All
func bankaccountAll(c *gin.Context) {
	mdl := models.Bankaccount{}
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
func bankaccountOne(c *gin.Context) {
	var mdl models.Bankaccount
	
	Id := c.Param("id")
		
	mdl.Id = Id
	data, err := mdl.One()
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}
//Create
func bankaccountCreate(c *gin.Context) {
	var mdl models.Bankaccount
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
func bankaccountUpdate(c *gin.Context) {
	var mdl models.Bankaccount
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
func bankaccountDelete(c *gin.Context) {
	var mdl models.Bankaccount

	
	Id := c.Param("id")
		
	
	mdl.Id = Id
	errResult := mdl.Delete()
	if handleError(c, errResult) {
		return
	}
	jsonSuccess(c)
}
