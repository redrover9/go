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
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	//f2 := excelize.NewFile()
	//index := f2.NewSheet("Sheet2")
	max := len(rows)
    locations := make(map[string]string)
    for i := 1; i <= max; i++ {
        cellNum := strconv.Itoa(i)
        cellNum = "A" + cellNum
        locations[cellNum] = "placeholder"
        for _, row := range rows {
            
        }
    }
    for key, element := range locations {
        fmt.Println("Key:", key, "=>", "Element:", element)
}
/*
for _, row := range rows {
		for _, colCell := range row {
			for num, _ := range nums {
				words = append(colCell, 
                strCell := strconv.Itoa(num)
				strCell = "A" + strCell
                
				fmt.Println(strCell)

				fmt.Println(colCell)
			}
		}
		f.SetActiveSheet(index)
		if err := f2.SaveAs("Book2.xlsx"); err != nil {
			fmt.Println(err)
	}

}
*/
}
