package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
)

func main() {
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	f2 := excelize.NewFile()
	rows, err := f.GetRows("Sheet1")
	i := 0
	for _, row := range rows {
		for _, colCell := range row {
			i++
			cell := strconv.Itoa(i)
			cell = "A" + cell
			f2.SetCellValue("Sheet1", cell, colCell)
		}
	}
	if err := f2.SaveAs("Book2.xlsx"); err != nil {
		fmt.Println(err)

	}
}
