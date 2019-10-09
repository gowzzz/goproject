package handler

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ImportReq struct {
	Content string `form:"content" json:"content" xml:"content"  binding:"required"`
}

func OneFile(c *gin.Context) {
	// single file
	fileheader, err := c.FormFile("onefile")
	if err != nil {
		UseLog().Errorf("fileheader err:", err)
	}
	fmt.Printf("Contains:%+v\n", strings.Contains(fileheader.Header["Content-Type"][0], "123"))
	fmt.Printf("%+v\n", fileheader.Header["Content-Type"][0]) //application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
	UseLog().Info("fileheader.Filename:", fileheader.Filename)
	file, err := fileheader.Open()
	if err != nil {
		UseLog().Errorf("file err:", err)
	}
	datas, err := ReadExcelStr(file)
	if err != nil {
		UseLog().Errorf("ReadExcel err:", err)
	}
	fmt.Println("len(datas):", len(datas))
	// Upload the file to specific dst.
	// c.SaveUploadedFile(file, dst)

	// c.String(http.StatusOK, fmt.Sprintf("%s", datas))
	ExportStrCsv("aaa.csv", datas, c)
}
func MultFile(c *gin.Context) {
	// Multipart form
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]

	for _, file := range files {
		log.Println(file.Filename)

		// Upload the file to specific dst.
		// c.SaveUploadedFile(file, dst)
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}
func Import(c *gin.Context) {
	var imp ImportReq
	// caseid := c.Param("caseID")
	if err := c.ShouldBind(&imp); err != nil {
		c.SecureJSON(http.StatusOK, "imp err")
		return
	} else {
		// errDatas, err := ReadXlsFromBase64(caseid, imp.Content)
		// if err != nil {
		// 	if errDatas != nil {
		// 		// ExportImpErrorCSV(errDatas, c)
		// 		return
		// 	} else {
		// 		c.SecureJSON(http.StatusOK, "ok")
		// 		return
		// 	}
		// }
		// c.SecureJSON(http.StatusOK, "ok")
		// return
	}
}
