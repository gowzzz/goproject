package handler

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/extrame/xls"
	"github.com/gin-gonic/gin"

	"crypto/tls"
	"encoding/base64"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	// "bufio"
	"github.com/nfnt/resize"
)

// 写入xlsx表的时候要先缩放大小，再填充数据。否则填充完数据再缩放容易让图片被二次缩放，导致意想不到的效果。

// 该方法只读取Sheet1
func ReadExcel(content []byte) (datas [][]string, err error) {
	// var datas [][]string

	if f, err := excelize.OpenReader(bytes.NewReader(content)); err == nil {
		sheets := f.GetSheetMap()
		if sheets == nil {
			return nil, errors.New("xlsx sheet is null")
		}
		// fmt.Printf("Sheet:%+v\n", sheets)
		// fmt.Println("ReadExcel:excelize")
		// 获取 Sheet1 上所有单元格
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
		if xlsFile, err := xls.OpenReader(bytes.NewReader(content), "utf-8"); err == nil {
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
			if datas, err := csv.NewReader(bytes.NewReader(content)).ReadAll(); err == nil {
				return datas, nil
			} else {
				fmt.Println("ReadExcel csv:err:", err)
				return nil, err
			}
		}
	}
}

type Exp struct{}

// 导出刷证记录
func ExportIdentityExs(datas []Exp, c *gin.Context) {
	filename := "全部刷证记录_" + time.Now().Format("20060102150405") + ".xlsx"

	// 创建excel表格
	xlsx := excelize.NewFile()
	defer func() {
		// c.Writer.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
		/*xlsx data:application/vnd.openxmlformats-officedocument.spreadsheetml.sheet;base64,*/
		/*doc data:application/vnd.openxmlformats-officedocument.wordprocessingml.document;base64,*/
		/*pdf data:application/pdf;base64,*/
		/*xls data:application/vnd.ms-excel;base64,*/
		/*rtf data:application/msword;base64,*/
		/*data:application/vnd.ms-excel;base64,*/
		v := url.Values{}
		v.Add("filename", filename)
		c.Writer.Header().Set("Content-Disposition", "attachment;"+v.Encode())
		c.Writer.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

		xlsx.WriteTo(c.Writer)
		c.Writer.Flush()
	}()
	// 生成标题 2019-07-03 10:14:23
	title := []string{"序号", "客户姓名", "身份证号", "性别", "刷证时间", "首次人脸抓拍时间", "身份证件照", "现场核验照", "刷证来源"}
	for i := 0; i < len(title); i++ {
		//65==A   A1~I1
		xlsx.SetCellStr("Sheet1", string(65+i)+"1", title[i])
	}
	// 先设置列宽 C-F设为2cm    每行高度2cm
	width, height := 2.54, 2.54 //cm
	sheet := "Sheet1"

	if err := SetCellWtdth(xlsx, sheet, "G", "H", width); err != nil {
		fmt.Println("SetCellWtdth err:", err)
	}

	for rownum := 2; rownum <= len(datas)+2; rownum++ {
		// for k, _ := range datas {
		// 每行高度2cm 第一行是标题，数据从第二行开始
		if err := SetCellHeight(xlsx, sheet, rownum, height); err != nil {
			fmt.Println("SetCellHeight err  :", err)
		}
	}

	for k, _ := range datas {
		rownum := k + 2
		// 每行高度2cm 第一行是标题，数据从第二行开始
		// if err := SetCellHeight(xlsx, sheet, rownum, height); err != nil {
		// 	fmt.Println("SetCellHeight err  :", err)
		// }
		// timeLayout := "2006-01-02 15:04:05"
		// A 		B 		C  			D 			E 		F 		 	    G 		   H		   I
		// 序号	  客户姓名	身份证号	性别	  刷证时间	 首次人脸抓拍时间	身份证件照	现场核验照	刷证来源
		xlsx.SetCellValue(sheet, "A"+strconv.Itoa(rownum), rownum-1) //从第二行开始
		// xlsx.SetCellStr(sheet, "B"+strconv.Itoa(rownum), p.Realname)
		// // SetColStyle
		// xlsx.SetCellStr(sheet, "C"+strconv.Itoa(rownum), p.CardId)
		// xlsx.SetCellStr(sheet, "D"+strconv.Itoa(rownum), p.Gender)
		// if p.CreatedAt.Unix() > 0 {
		// 	xlsx.SetCellStr(sheet, "E"+strconv.Itoa(rownum), p.CreatedAt.Format(timeLayout)) //time=>string
		// }
		// if p.FirstCapturedTime > 0 {
		// 	xlsx.SetCellStr(sheet, "F"+strconv.Itoa(rownum), time.Unix(p.FirstCapturedTime, 0).Format(timeLayout)) //unix=>string 设置时间戳 使用模板格式化为日期字符串
		// }
		// xlsx.SetCellStr(sheet, "I"+strconv.Itoa(rownum), p.Source)
		// // 证件照是png，抓拍照jpg
		// if err := AddImgBase64ToExcel(xlsx, sheet, "G"+strconv.Itoa(rownum), width, height, p.CardFace, ".png", p.Realname); err != nil {
		// 	fmt.Println("AddImgBase64ToExcel 1 err  :", err)
		// }
		// if err := AddImgBase64ToExcel(xlsx, sheet, "H"+strconv.Itoa(rownum), width, height, p.CheckFace, ".jpg", p.Realname); err != nil {
		// 	fmt.Println("AddImgBase64ToExcel 2 err  :", err)
		// }
		// for k, v := range i {
		// 	location := string(65+k+2) + strconv.Itoa(rownum)
		// 	if code := AddImgToExcel(xlsx, sheet, location, width, height, v.SavePath); code != 0 {
		// 		return "", code
		// 	}
		// }
	}
	return
}
func initDealDetailXLSX(f *excelize.File, sheetname string) {

	// style, err = f.NewStyle(`{"alignment":{"horizontal":"center","ident":1,"justify_last_line":true,"reading_order":0,"relative_indent":1,"shrink_to_fit":true,"text_rotation":45,"vertical":"top","wrap_text":true}}`)

	f.SetSheetName("Sheet1", sheetname)
	style, _ := f.NewStyle(`{"font":{"bold":true}}`)
	f.SetCellStyle(sheetname, "A1", "F1", style)
	f.SetCellStyle(sheetname, "A7", "F7", style)
	// f.SetCellStyle(sheetname, "A8", "E8", style)
	SetCellHeight(f, sheetname, 1, 1)
	SetCellHeight(f, sheetname, 7, 1)
	SetCellHeight(f, sheetname, 8, 1)
	f.SetCellStr(sheetname, "A1", "基本信息：")
	f.SetCellStr(sheetname, "A7", "流水证明:")
	f.SetColWidth(sheetname, "A", "F", 20)
	f.SetColWidth(sheetname, "F2", "F6", 12)
	// key1
	f.SetCellStr(sheetname, "B2", "客户姓名")
	f.SetCellStr(sheetname, "B3", "客户电话")
	f.SetCellStr(sheetname, "B4", "成交时间")
	f.SetCellStr(sheetname, "B5", "经纪人")
	f.SetCellStr(sheetname, "B6", "所属渠道")
	f.SetCellStr(sheetname, "D2", "性别")
	f.SetCellStr(sheetname, "D3", "客户证件号码")
	f.SetCellStr(sheetname, "D4", "项目")
	f.SetCellStr(sheetname, "D5", "房间")
	f.SetCellStr(sheetname, "D6", "当前置业顾问")
	// key2
	f.SetCellStr(sheetname, "A8", "时间")
	f.SetCellStr(sheetname, "B8", "留痕记录")
	f.SetCellStr(sheetname, "C8", "抓拍照")
	f.SetCellStr(sheetname, "D8", "人脸相似度")
	f.SetCellStr(sheetname, "E8", "是否本人")
	f.SetCellStr(sheetname, "F8", "风险指标")

	// 合并单元格
	f.MergeCell(sheetname, "A1", "F1")
	f.MergeCell(sheetname, "A7", "F7")
	f.MergeCell(sheetname, "A2", "A6")
	f.MergeCell(sheetname, "F2", "F6")
}

// 列宽 12.56字符=120pixel
// 行号 72磅=120pixel
// 1yc=72bang=120pixel=2.54cm
func SetCellHeight(f *excelize.File, sheet string, row int, height float64) error {
	// 传入cm，写成像素
	// 设置单元格长宽
	//行高有个很奇怪的舍入：0.08~0.12都记作0.1   0.13~0.17都记作0.15 ，舍入单位为0.05。所有为了全入起见，原大小+0.05
	pixelH := height / 2.54 * 120 //像素
	ch := pixelH*0.6 - 0.05 + 0.05
	// fmt.Println("ch:", ch)
	f.SetRowHeight(sheet, row, ch)
	return nil
}
func SetCellWtdth(f *excelize.File, sheet, startcol, endcol string, width float64) error {
	// 传入cm，写成像素
	// 设置单元格长宽
	pixelW := width / 2.54 * 120 //像素
	cw := (pixelW-7)/9 + 7.0/9

	// fmt.Println("cw:", cw)
	f.SetColWidth(sheet, startcol, endcol, cw)

	return nil
}

func AddImgBase64ToExcel(xlsx *excelize.File, sheet, location string, width, height float64, imgbase64, imgext, realname string) error {
	// tmpfile := realname + "." + imgext
	// os.Remove(tmpfile)
	content, err := base64.StdEncoding.DecodeString(imgbase64)
	if err != nil {
		fmt.Println("err1:", err)
		return err
	}

	img, err := jpeg.Decode(bytes.NewReader(content))
	if err != nil {
		fmt.Println("jpeg Decode err  :", err)
		img, err = png.Decode(bytes.NewReader(content))
		if err != nil {
			fmt.Println("jpeg and png.Decode err  :", err)
			return err
		}
	}
	// pixelH := height / 2.54 * 90 //像素
	// pixelW := width / 2.54 * 90  //像素
	// pixelSize := pixelH
	// if pixelW<pixelH{
	// 	pixelSize=pixelW
	// }

	var m image.Image
	if img.Bounds().Dx() > img.Bounds().Dy() {
		m = resize.Resize(90, 0, img, resize.Lanczos3)
	} else {
		m = resize.Resize(0, 90, img, resize.Lanczos3)
	}
	buffer := bytes.NewBuffer(nil)
	jpeg.Encode(buffer, m, nil)

	format := `{"lock_aspect_ratio": true}`
	err = xlsx.AddPictureFromBytes(sheet, location, format, "xx", ".jpg", buffer.Bytes())
	if err != nil {
		fmt.Println("AddPicture err  :", err)
		return err
	}
	return nil
}
func AddImgToExcel(xlsx *excelize.File, sheet, location string, width, height float64, imgpath string) error {
	var f io.Reader
	if strings.HasPrefix(imgpath, "http") {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		resp, err := client.Get(imgpath)
		if err != nil {
			fmt.Println("client.Get err  :", err)
			return err
		}
		f = resp.Body
	} else {
		file, err := os.Open(imgpath)
		if err != nil {
			fmt.Println("os.Open err  :", err)
			return err
		}

		defer file.Close()
		f = file
	}
	var buffer bytes.Buffer
	buffer.ReadFrom(f)
	img, err := jpeg.Decode(bytes.NewReader(buffer.Bytes()))
	if err != nil {
		fmt.Println("jpeg .Decode err  :", err)
		img, err = png.Decode(bytes.NewReader(buffer.Bytes()))
		if err != nil {
			fmt.Println("jpeg and png.Decode err  :", err)
			return err
		}
	}

	// img, err := jpeg.Decode(bufio.NewReader(f))
	// if err != nil {
	// 	fmt.Println("jpeg .Decode err  :", err)
	// 	img, err = png.Decode(bufio.NewReader(f))
	// 	if err != nil {
	// 		fmt.Println("jpeg and png.Decode err  :", err)
	// 		return err
	// 	}
	// }
	// pixelH := height / 2.54 * 90 //像素
	// pixelW := width / 2.54 * 90  //像素
	// pixelSize := pixelH
	// if pixelW<pixelH{
	// 	pixelSize=pixelW
	// }

	var m image.Image
	if img.Bounds().Dx() > img.Bounds().Dy() {
		m = resize.Resize(90, 0, img, resize.Lanczos3)
	} else {
		m = resize.Resize(0, 90, img, resize.Lanczos3)
	}

	// write new image to file
	encodebuffer := bytes.NewBuffer(nil)
	jpeg.Encode(encodebuffer, m, nil)

	format := `{"lock_aspect_ratio": true, "locked": true, "positioning": "oneCell"}`
	// err = xlsx.AddPicture(sheet, location, outname, `{"lock_aspect_ratio": true, "locked": true, "positioning": "absolute"}`)//oneCell
	err = xlsx.AddPictureFromBytes(sheet, location, format, "xx", ".jpg", encodebuffer.Bytes())
	if err != nil {
		fmt.Println("AddPicture err  :", err)
		return err
	}
	return nil
}
