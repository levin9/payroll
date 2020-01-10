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
	fmt.Println(input.FilePath)
	if input.FilePath == "" {
		return utils.Throw("对象为空或者路径为空.")
	}
	xlsx, err := excelize.OpenFile(input.FilePath)
	if err != nil {
		return utils.Throw(fmt.Sprintf("打开excel失败,error：%s", err.Error()))
	}
	style, _ := xlsx.NewStyle(`{"border":[{"type":"left","color":"000000","style":1},{"type":"top","color":"000000","style":1},{"type":"bottom","color":"000000","style":1},{"type":"right","color":"000000","style":1}],"fill":{"type":"pattern","color":["#ffeb00"],"pattern":1},"alignment":{"horizontal":"left","ident":1,"vertical":"center","wrap_text":true}}`)
	
	
	attDal := models.Attendance{}
	personDal := models.Personal{}
	attendanceList,_,_ :=attDal.All(models.CreateCondition(input.MonthId,input.TenantId))
	personList,_,_:=personDal.All(models.CreateCondition(input.MonthId,input.TenantId))

	sheetName := "Sheet1"
	rows := xlsx.GetRows(sheetName)
	if len(rows) <= 2 {
		return utils.Throw("excel格式错误或无有效数据.")
	}
	
	for rIndex, row := range rows {
		if rIndex == 0 || rIndex == 1 {
			continue
		}
		insertFlag :=true
		errorMsg := []string{}
		personAtt := models.Attendance{}
		person :=models.Personal{}
		for _,p := range *personList{
			if p.Realname ==utils.Trim(row[0]){
				person=p
				break
			}
		}
		if person.Realname==""{
			errorMsg = append(errorMsg, "用户名不存在.")
		}else{
			for _,att := range *attendanceList{
				if att.Personid== person.Personid{
					personAtt=att
					insertFlag=false
					break
				}
			}
			personAtt.Personid = utils.Trim(row[0])
			//personAtt.Id = m
			personAtt.Tenantid = input.TenantId
			personAtt.Personid = person.Personid
			personAtt.Monthid = input.MonthId

			personAtt.Chidaocount = o.toFloat32(row[5])
			personAtt.Chidaonum = o.toFloat32(row[5])
			personAtt.ZaoTuicount = o.toFloat32(row[5])
			personAtt.Tiaoxiunum = o.toFloat32(row[5])
			personAtt.ZaoTuinum = o.toFloat32(row[5])
			personAtt.ZaoTuinum = o.toFloat32(row[5])
			personAtt.ZaoTuinum = o.toFloat32(row[5])
			personAtt.ZaoTuinum = o.toFloat32(row[5])
			personAtt.ZaoTuinum = o.toFloat32(row[5])
			personAtt.ZaoTuinum = o.toFloat32(row[5])
			personAtt.ZaoTuinum = o.toFloat32(row[5])
			personAtt.ZaoTuinum = o.toFloat32(row[5])
		}
		
		//如果有错误，将背景设为警示颜色
		if len(errorMsg) > 0 {
			xlsx.SetCellStyle(sheetName, fmt.Sprintf("A%v", rIndex+1), fmt.Sprintf("AA%v", rIndex+1), style)
		}else{
			if insertFlag {
				personAtt.Id= utils.ToNewID()
				personAtt.Create()
			}else{
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
