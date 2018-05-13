package main

import (
	//"math/big"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
)

type fileStruct struct {
	number         int
	opnumber       string
	nameproject    string
	nameapp        string
	namedocker     string
	organize       string
	responseperson string
	section        string
	portnumber     int64
	status         string
	doubleIDC      string
	startime       string
	operator       string
	letgo          string
	remarks        string
}

func main() {
	var docker []fileStruct

	xlsx, err := excelize.OpenFile("/Users/yp-tc-m-2642/Desktop/test.xlsx") //"../../../../../Desktop/test.xlsx"
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get value from cell by given worksheet name and axis.
	//cell := xlsx.GetCellValue("工作表1", "B2")
	//fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows := xlsx.GetRows("工作表1")
	docker = make([]fileStruct, len(rows))
	//for idx, row := range rows {
	for i := 1; i < len(rows); i++ {
		docker[i].number, _ = strconv.Atoi(rows[i][0])
		docker[i].opnumber = rows[i][1]
		docker[i].nameproject = rows[i][2]
		docker[i].nameapp = rows[i][3]
		docker[i].namedocker = rows[i][4]
		docker[i].organize = rows[i][5]
		docker[i].responseperson = rows[i][6]
		docker[i].section = rows[i][7]
		docker[i].portnumber, _ = strconv.ParseInt(rows[i][8], 10, 64)
		docker[i].status = rows[i][9]
		docker[i].doubleIDC = rows[i][10]
		docker[i].startime = rows[i][11]
		docker[i].operator = rows[i][12]
		docker[i].letgo = rows[i][13]
		docker[i].remarks = rows[i][14]

		// fmt.Printf("row: %v\n", row)
		fmt.Printf("docker: %v\n", docker[i])
	}

}
