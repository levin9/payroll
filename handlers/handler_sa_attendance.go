package handlers

import (
	"fmt"
	"payroll/models"

	//"unsafe"

	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("attendance", attendanceAll)
	groupApi.GET("attendance/:id", attendanceOne)
	groupApi.POST("attendance", attendanceCreate)
	groupApi.POST("attendanceimport", attendanceImport)
	groupApi.PATCH("attendance", attendanceUpdate)
	groupApi.DELETE("attendance/:id", attendanceDelete)
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

//Update
func attendanceImport(c *gin.Context) {
	input := AttImportInput{}
	if err := c.ShouldBind(&input); handleError(c, err) {
		return
	}
	fmt.Println(input.TenantId)

	// respBytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// if err1 := json.Unmarshal(respBytes, &input); err1 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }
	// fmt.Println(input.TenantId)

	jsonSuccess(c)
}

type AttImportInput struct {
	BizType  string `json:"biz_type"`
	FilePath string `json:"file_path"`
	TenantId string `json:"tenant_id"`
}
