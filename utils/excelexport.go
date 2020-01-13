package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/tealeg/xlsx"
)

type ExcelExportCfg struct {
	Title     string
	SheetName string
	Headers   []string
	Data      [][]string
}

func Export(cfg ExcelExportCfg) (string, error) {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet(cfg.SheetName)
	if err != nil {
		return "", err
	}
	row := sheet.AddRow()
	var cell *xlsx.Cell
	for _, title := range cfg.Headers {
		cell = row.AddCell()
		cell.Value = title
	}
	for _, r := range cfg.Data {
		row = sheet.AddRow()
		for _, c := range r {
			fmt.Printf("arr[%v][%v]=%v \t \n", "i", "", c)
			cell = row.AddCell()
			cell.Value = c
		}
	}
	
	time := strconv.Itoa(int(time.Now().Unix()))
	filename := cfg.Title + "-" + time + ".xlsx"

	fullPath := GetExcelFullPath() + filename
	err = file.Save(fullPath)
	if err != nil {
		return "", err
	}

	return filename, nil

}

func GetExcelFullPath() string {
	return ""
}
