package main

import (
	"fmt"
)

/*
43 44 45 46 47 48 49
42 21 22 23 24 25 26
41 20 7  8  9  10 27
40 19 6  1  2  11 28
39 18 5  4  3  12 29
38 17 16 15 14 13 30
37 36 35 34 33 32 31
*/
// 找到每一圈的起始位置,围绕转一圈即可.,内往外是(2n-1)^2
/*
外往内是
1.首先，读者可以先自行画出边长为3，4，5的螺旋矩阵观察。不难得出可以将每四个指向箭头定义成一个圈。
每个圈都是从左到右–从上到下–从右到左–从下到上递增1。因此，找出每个圈的起始位置，依次往后赋值即可完成此算法。

2.第一个圈的第一个数必为1。因此可以用它作为计算其它圈的起始数字。总结规律可以得到：
后一个圈的起始数字 = 前一个圈的起始数字 + 4（前一个圈的边长 -1）；*

*/
// 奇数是完整圈数，偶数少了左上,补上
func PrintArr(x, y int) [][]int {
	if x%2 == 0 {
		x = x + 1
	}
	if y%2 == 0 {
		y = y + 1
	}
	arrlen := x
	if y > x {
		arrlen = y
	}
	fmt.Println("arrlen:", arrlen)
	quanshu := (arrlen + 1) / 2
	var arr = make([][]int, arrlen, arrlen)
	for i := 0; i < arrlen; i++ {
		arr[i] = make([]int, arrlen)
	}
	// 找到原点
	X, Y := arrlen/2, arrlen/2
	arr[Y][X] = 1
	// 找到XY坐标系，以原点为中心，
	for i := 1; i < quanshu; i++ {
		// 计算轴数值
		// 右 	(2*n-1)^2+n  上 0~i-1 下 0~i
		arr[Y][X+i] = (2*i-1)*(2*i-1) + i
		tmp := arr[Y][X+i]
		for j := 1; j < i; j++ {
			arr[Y-j][X+i] = tmp - j
			arr[Y+j][X+i] = tmp + j
		}
		arr[Y+i][X+i] = tmp + i
		// 下	(2*n-1)^2+3n 左 0~i 右 0~i-1
		arr[Y+i][X] = (2*i-1)*(2*i-1) + 3*i
		tmp = arr[Y+i][X]
		for j := 1; j < i; j++ {
			arr[Y+i][X+j] = tmp - j
			arr[Y+i][X-j] = tmp + j
		}
		arr[Y+i][X-i] = tmp + i
		// 左	(2*n-1)^2+5n 上 0~i 下 0~i-1
		arr[Y][X-i] = (2*i-1)*(2*i-1) + 5*i
		tmp = arr[Y][X-i]
		for j := 1; j < i; j++ {
			arr[Y-j][X-i] = tmp + j
			arr[Y+j][X-i] = tmp - j
		}
		arr[Y-i][X-i] = tmp + i
		// 上	(2*n-1)^2+7n 左 0~i-1 右 0~i
		arr[Y-i][X] = (2*i-1)*(2*i-1) + 7*i
		tmp = arr[Y-i][X]
		for j := 1; j < i; j++ {
			arr[Y-i][X+j] = tmp + j
			arr[Y-i][X-j] = tmp - j
		}
		arr[Y-i][X+i] = tmp + i
	}
	for _, v := range arr {
		fmt.Printf("%+v\n", v)
	}
	return arr
}
func main() {
	aa := make(map[string]string)
	aa["aa"] = "AA"
	if a, ok := aa["BB"]; !ok {
		fmt.Println("!ok")
	} else {
		fmt.Println("aa:", a)

	}
	PrintArr(3, 4)
}
