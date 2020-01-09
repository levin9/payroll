package blls

import (
	"fmt"
	"payroll/models"
	"payroll/utils"
	//"unsafe"
	"github.com/360EntSecGroup-Skylar/excelize"
)

type ImportAttendanceService struct {
}

func (o *ImportAttendanceService) Import(input models.AttImportInput) error {
	if input.FilePath == "" {
		return throwException("对象为空或者路径为空.")
	}
	xlsx, err := excelize.OpenFile(input.FilePath)
	if err != nil {
		return throwException(fmt.Sprintf("打开excel失败,error：%s", err.Error()))
	}
	style, _ := xlsx.NewStyle(`{"border":[{"type":"left","color":"000000","style":1},{"type":"top","color":"000000","style":1},{"type":"bottom","color":"000000","style":1},{"type":"right","color":"000000","style":1}],"fill":{"type":"pattern","color":["#ffeb00"],"pattern":1},"alignment":{"horizontal":"left","ident":1,"vertical":"center","wrap_text":true}}`)
	sheetName := "Sheet1"
	rows := xlsx.GetRows(sheetName)
	if len(rows) <= 2 {
		return throwException("excel格式错误或无有效数据.")
	}
	for rIndex, row := range rows {
		if rIndex == 0 || rIndex == 1 {
			continue
		}
		errorMsg = []string{}
		attendance := new(model.Attendance)
		
		account.FullName = utils.Trim(row[0])
	}
	fmt.Println(xlsx)
	return nil
}
