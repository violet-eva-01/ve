// Package utils @author: Violet-Eva @date  : 2025/9/19 @notes :
package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

func WriteExcelForStruct[t any](fileName, path string, sheetName, title []string, data ...[]t) (err error) {

	var (
		titles     []string
		sheetNames []string
		isStruct   bool
	)

	if len(data) == 0 {
		return fmt.Errorf("没有数据传入,无法写Excel文档")
	}

	if sheetName == nil || len(sheetName) < 1 {
		for i := 1; i <= len(data); i++ {
			sheetNames = append(sheetNames, fmt.Sprintf("第%d页", i))
		}
	} else {
		sheetNames = sheetName
	}

	tmpData := data[0][0]
	tmpRowType := reflect.TypeOf(tmpData)
	tmpRowTypeStr := tmpRowType.Kind().String()

	if tmpRowTypeStr == "struct" {
		isStruct = true
	}

	if title == nil || (len(title) < 1 && isStruct) {
		for index := 0; index < tmpRowType.NumField(); index++ {
			titles = append(titles, tmpRowType.Field(index).Name)
		}
	} else {
		titles = title
	}

	excelFile := excelize.NewFile()

	switch {
	case isStruct:

		for sheetIndex, sheetData := range data {
			sn := sheetNames[sheetIndex]
			_, err = excelFile.NewSheet(sn)
			if err != nil {
				return
			}

			for titleIndex, colTitle := range titles {
				cell, _ := excelize.CoordinatesToCellName(titleIndex+1, 1)
				err = excelFile.SetCellValue(sn, cell, colTitle)
				if err != nil {
					return
				}
			}

			for rowIndex, row := range sheetData {
				rowType := reflect.TypeOf(row)
				rowValue := reflect.ValueOf(row)
				for columnIndex := 0; columnIndex < rowType.NumField(); columnIndex++ {
					switch rowValue.Field(columnIndex).String() {
					case "int", "int8", "int16", "int32", "int64":
						cell, _ := excelize.CoordinatesToCellName(columnIndex+1, rowIndex+2)
						content := rowValue.Field(columnIndex).Int()
						_ = excelFile.SetCellInt(sn, cell, content)
						style, _ := excelFile.NewStyle(&excelize.Style{NumFmt: 1})
						_ = excelFile.SetCellStyle(sn, cell, cell, style)
					case "float32", "float64":
						cell, _ := excelize.CoordinatesToCellName(columnIndex+1, rowIndex+2)
						content := rowValue.Field(columnIndex).Float()
						_ = excelFile.SetCellFloat(sn, cell, content, -1, 64)
						style, _ := excelFile.NewStyle(&excelize.Style{NumFmt: 2})
						_ = excelFile.SetCellStyle(sn, cell, cell, style)
					default:
						content := rowValue.Field(columnIndex).Interface()
						cell, _ := excelize.CoordinatesToCellName(columnIndex+1, rowIndex+2)
						_ = excelFile.SetCellValue(sn, cell, content)
					}
				}
			}
		}

		_ = excelFile.DeleteSheet("Sheet1")

	case !isStruct:

		for sheetIndex, sheetData := range data {
			sn := "Sheet1"
			if sheetIndex != 0 {
				sn = fmt.Sprintf("Sheet%d", sheetIndex+1)
				_, err = excelFile.NewSheet(sn)
				if err != nil {
					return
				}
			}
			cell, _ := excelize.CoordinatesToCellName(1, 1)
			err = excelFile.SetCellValue(sn, cell, tmpRowType.Kind().String())
			if err != nil {
				return
			}
			for rowIndex, row := range sheetData {
				rowStr := fmt.Sprintf("%v", row)
				switch tmpRowTypeStr {
				case "int", "int8", "int16", "int32", "int64":
					rowCell, _ := excelize.CoordinatesToCellName(1, rowIndex+2)
					content, _ := strconv.ParseInt(rowStr, 10, 64)
					_ = excelFile.SetCellInt(sn, rowCell, content)
					style, _ := excelFile.NewStyle(&excelize.Style{NumFmt: 1})
					_ = excelFile.SetCellStyle(sn, rowCell, rowCell, style)
				case "float32", "float64":
					rowCell, _ := excelize.CoordinatesToCellName(1, rowIndex+2)
					content, _ := strconv.ParseFloat(rowStr, 64)
					_ = excelFile.SetCellFloat(sn, rowCell, content, -1, 64)
					style, _ := excelFile.NewStyle(&excelize.Style{NumFmt: 2})
					_ = excelFile.SetCellStyle(sn, rowCell, rowCell, style)
				default:
					rowCell, _ := excelize.CoordinatesToCellName(1, rowIndex+2)
					_ = excelFile.SetCellValue(sn, rowCell, rowStr)
				}
			}
		}
	}

	excelFileName := fileName + ".xlsx"

	err = excelFile.SaveAs(path + "/" + excelFileName)
	if err != nil {
		return
	}
	return
}

func GetMaxLenMapList(input []map[string]interface{}) map[string]interface{} {
	var js = struct {
		Index  int
		Length int
	}{}
	for index, value := range input {
		length := len(value)
		if index == 0 {
			js.Index = index
			js.Length = length
		} else {
			if length > js.Length {
				js.Index = index
				js.Length = length
			}
		}
	}
	return input[js.Index]
}

func WriteExcelForMapList(fileName, path string, sheetName, title []string, data ...[]map[string]interface{}) (err error) {

	var (
		titles     []string
		sheetNames []string
	)

	if len(data) < 1 {
		return fmt.Errorf("没有数据传入,无法写Excel文档")
	}

	if sheetName == nil || len(sheetName) < 1 {
		for i := 1; i <= len(data); i++ {
			sheetNames = append(sheetNames, fmt.Sprintf("第%d页", i))
		}
	} else {
		sheetNames = sheetName
	}

	excelFile := excelize.NewFile()

	for sheetIndex, sheetData := range data {
		sn := "Sheet1"
		if sheetIndex != 0 {
			sn = fmt.Sprintf("Sheet%d", sheetIndex+1)
			_, err = excelFile.NewSheet(sn)
			if err != nil {
				return
			}
		}
		if title == nil || (len(title) < 1) {
			list := GetMaxLenMapList(sheetData)
			for k := range list {
				titles = append(titles, k)
			}
		} else {
			titles = title
		}
		for titleIndex, colTitle := range titles {
			cell, _ := excelize.CoordinatesToCellName(titleIndex+1, 1)
			err = excelFile.SetCellValue(sn, cell, colTitle)
			if err != nil {
				return
			}
		}
		for rowIndex, row := range sheetData {
			for columnIndex, columnName := range titles {
				rowType := reflect.TypeOf(row[columnName]).Kind().String()
				switch rowType {
				case "int", "int8", "int16", "int32", "int64":
					cell, _ := excelize.CoordinatesToCellName(columnIndex+1, rowIndex+2)
					content, _ := strconv.ParseInt(fmt.Sprintf("%s", row[columnName]), 10, 64)
					_ = excelFile.SetCellInt(sn, cell, content)
					style, _ := excelFile.NewStyle(&excelize.Style{NumFmt: 1})
					_ = excelFile.SetCellStyle(sn, cell, cell, style)
				case "float32", "float64":
					cell, _ := excelize.CoordinatesToCellName(columnIndex+1, rowIndex+2)
					content := row[columnName].(float64)
					precision := len(strings.Split(fmt.Sprintf("%s", row[columnName]), ".")[1])
					_ = excelFile.SetCellFloat(sn, cell, content, precision, 64)
					style, _ := excelFile.NewStyle(&excelize.Style{NumFmt: 2})
					_ = excelFile.SetCellStyle(sn, cell, cell, style)
				default:
					content := row[columnName]
					cell, _ := excelize.CoordinatesToCellName(columnIndex+1, rowIndex+2)
					_ = excelFile.SetCellValue(sn, cell, content)
				}
			}
		}
	}

	excelFileName := fileName + ".xlsx"

	err = excelFile.SaveAs(path + "/" + excelFileName)
	if err != nil {
		return
	}
	return
}
