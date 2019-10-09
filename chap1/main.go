package main

import (
	"fmt"
	"time"
)

func main() {
	var data = 0
	go func() {
		data++
	}()
	go func() {
		data2 := data
		fmt.Println("data2:", data2)
	}()
	if data == 0 {
		time.Sleep(1 * time.Second)
		fmt.Println("data==0:", data)
	}
}
