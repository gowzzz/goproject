package handler

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"net/url"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/extrame/xls"
	"github.com/gin-gonic/gin"
)

// 读取xlsx xls csv都可以，只读取第一个表单
func ReadExcelStr(reader io.ReadSeeker) (datas [][]string, err error) {
	// var datas [][]string

	if f, err := excelize.OpenReader(reader); err == nil {
		sheets := f.GetSheetMap()
		if sheets == nil {
			return nil, errors.New("xlsx sheet is null")
		}
		// fmt.Printf("Sheet:%+v\n", sheets)
		// fmt.Println("ReadExcel:excelize")
		datas = f.GetRows(sheets[1])
		// for _, row := range rows {
		// 	for _, colCell := range row {
		// 		fmt.Print(colCell, "\t")
		// 	}
		// 	fmt.Println()
		// }
		return datas, nil
	} else {
		fmt.Println("ReadExcel excelize:err:", err)
		if xlsFile, err := xls.OpenReader(reader, "utf-8"); err == nil {
			defer func() {
				if err := recover(); err != nil {
					fmt.Printf("datas:%+v\n", datas)
				}
			}()
			fmt.Println("ReadExcel:xls")
			sheet1 := xlsFile.GetSheet(0)
			for i := 0; i <= int(sheet1.MaxRow); i++ {
				var rowdata []string
				row := sheet1.Row(i)
				for index := row.FirstCol(); index < row.LastCol(); index++ {
					rowdata = append(rowdata, row.Col(index))
				}
				datas = append(datas, rowdata)
			}
			return datas, nil
		} else {
			fmt.Println("ReadExcel xls:err:", err)
			fmt.Println("ReadExcel:csv")
			if datas, err := csv.NewReader(reader).ReadAll(); err == nil {
				return datas, nil
			} else {
				fmt.Println("ReadExcel csv:err:", err)
				return nil, err
			}
		}
	}
}

func ExportStrCsv(filename string, data [][]string, c *gin.Context) {
	// filename := "导入失败的渠道成交记录_" + time.Now().Format("20060102150405") + ".csv"
	v := url.Values{}
	v.Add("filename", filename)
	c.Writer.Header().Set("Content-Disposition", "attachment;"+v.Encode())
	// c.Writer.Header().Add("Content-Type", "application/vnd.ms-excel")
	c.Writer.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	w := csv.NewWriter(c.Writer)         //创建一个新的写入文件流
	// 10个
	// 客户姓名 客户身份证号(15/18位) 置业顾问 经纪人 所属渠道公司 报备时间(2017-11-07 11:11:11) 成交时间(2017-11-07 11:11:11) 房号
	// data := [][]string{
	// 	{
	// 		"客户姓名",
	// 		"客户身份证号",
	// 		"置业顾问",
	// 		"经纪人",
	// 		"所属渠道公司",
	// 		"报备时间",
	// 		"成交时间",
	// 		"房号",
	// 		"错误信息",
	// 	},
	// }
	// data = append(data, datas...)
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			data[i][j] = "	" + data[i][j]
		}
	}
	w.WriteAll(data) //写入数据
	w.Flush()
	c.Writer.Flush()
}
