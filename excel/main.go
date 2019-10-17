package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main(c *gin.Context) {
	filename := "人脸风控系统系统账户_导入模版_" + time.Now().Format("20060102150405") + ".xlsx"
	f, err := excelize.OpenFile("feidanadmin.xlsx")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	sheets := f.GetSheetMap()
	if sheets == nil {
		return
	}
	f.SetCellStr(sheets[1], "B3", "wwww")
	v := url.Values{}
	v.Add("filename", filename)
	c.Writer.Header().Set("Content-Disposition", "attachment;"+v.Encode())
	c.Writer.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	f.WriteTo(c.Writer)
	c.Writer.Flush()
}
