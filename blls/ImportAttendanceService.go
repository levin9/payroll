package blls

import (
	"fmt"
	"os"

	"payroll/models"
	"payroll/utils"

	//"unsafe"
	"github.com/360EntSecGroup-Skylar/excelize"
)

type ImportAttendanceService struct {
}

func (o *ImportAttendanceService) Import(input models.AttImportInput) error {
	fmt.Println(input.FilePath)
	if input.FilePath == "" {
		return utils.Throw("对象为空或者路径为空.")
	}
	root, _ := os.Getwd()
	xlsx, err := excelize.OpenFile(root + input.FilePath)
	if err != nil {
		return utils.Throw(fmt.Sprintf("打开excel失败,error：%s", err.Error()))
	}
	style, _ := xlsx.NewStyle(`{"border":[{"type":"left","color":"000000","style":1},{"type":"top","color":"000000","style":1},{"type":"bottom","color":"000000","style":1},{"type":"right","color":"000000","style":1}],"fill":{"type":"pattern","color":["#ffeb00"],"pattern":1},"alignment":{"horizontal":"left","ident":1,"vertical":"center","wrap_text":true}}`)

	attDal := models.Attendance{}
	personDal := models.Personal{}
	attendanceList, _, _ := attDal.All(models.CreateCondition(input.MonthId, input.TenantId))
	personList, _, _ := personDal.All(models.CreateCondition(input.MonthId, input.TenantId))

	sheetName := "Sheet1"
	rows := xlsx.GetRows(sheetName)
	if len(rows) <= 2 {
		return utils.Throw("excel格式错误或无有效数据.")
	}

	for rIndex, row := range rows {
		if rIndex == 0 || rIndex == 1 {
			continue
		}
		insertFlag := true
		errorMsg := []string{}
		personAtt := models.Attendance{}
		person := models.Personal{}
		for _, p := range *personList {
			if p.Realname == utils.Trim(row[1]) {
				person = p
				break
			}
		}
		if person.Realname == "" {
			errorMsg = append(errorMsg, "用户名不存在.")
		} else {
			for _, att := range *attendanceList {
				if att.Personid == person.Personid {
					personAtt = att
					insertFlag = false
					break
				}
			}

			//personAtt.Id = m
			personAtt.Tenantid = input.TenantId
			personAtt.Personid = person.Personid
			personAtt.Monthid = input.MonthId

			personAtt.Chidaocount = o.toFloat32(row[3])
			personAtt.Chidaonum = o.toFloat32(row[4])
			personAtt.ZaoTuicount = o.toFloat32(row[5])
			personAtt.ZaoTuinum = o.toFloat32(row[6])

			personAtt.Tiaoxiunum = o.toFloat32(row[7])
			personAtt.Shijianum = o.toFloat32(row[8])
			personAtt.Sicknum = o.toFloat32(row[9])
			personAtt.Daixinshijianum = o.toFloat32(row[10])
			personAtt.Annualleavenum = o.toFloat32(row[11])

			personAtt.Kuanggongnum = o.toFloat32(row[12])
			personAtt.Kongbainum = o.toFloat32(row[13])

			personAtt.Burvjianum = o.toFloat32(row[14])
			personAtt.Peichanjianum = o.toFloat32(row[15])
			personAtt.Chanjianum = o.toFloat32(row[16])
			personAtt.Hunjianum = o.toFloat32(row[17])
			personAtt.Sangjianum = o.toFloat32(row[18])

			personAtt.Actualworkday = o.toFloat32(row[2])

		}

		//如果有错误，将背景设为警示颜色
		if len(errorMsg) > 0 {
			xlsx.SetCellStyle(sheetName, fmt.Sprintf("A%v", rIndex+1), fmt.Sprintf("AA%v", rIndex+1), style)
		} else {
			if insertFlag {
				personAtt.Id = utils.ToNewID()
				personAtt.Create()
			} else {
				//personAtt.FMo
				personAtt.Update()
			}
		}
	}
	return nil
}

func (o *ImportAttendanceService) toFloat32(str string) float32 {
	return utils.ToFloat32(str)
}
