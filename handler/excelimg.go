package handler

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// 读取包含图片的xlsx
func ReadXlsxWithImg(reader io.ReadSeeker) (datas [][]string, err error) {
	// var datas [][]string
	if f, err := excelize.OpenReader(reader); err == nil {
		sheets := f.GetSheetMap()
		sheetname := sheets[1]
		if sheets == nil {
			return nil, errors.New("xlsx sheet is null")
		}

		rows := f.GetRows(sheets[1])
		for k := 1; k < len(rows); k++ {
			// row := rows[k]
			imgcell := "B" + strconv.Itoa(k+1) //第一行数据是 B2
			contentname, content := f.GetPicture(sheetname, imgcell)
			fmt.Println("GetPicture contentname:", contentname, len(content))
		}
		return datas, nil
	}
	fmt.Println("ReadXlsxWithImg err:", err)
	return nil, err
}
