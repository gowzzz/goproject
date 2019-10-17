package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func main() {
	var array_or_slice = []string{"aaa", "bbb", "ccc"}
	a := fmt.Sprint(array_or_slice)
	b := strings.Trim(a, "[]")
	c := strings.Replace(b, " ", ",", -1)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	return
	//我们另启一个goroutine来处理监控对象的事件
	go func() {
		//创建一个监控对象
		watch, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}
		defer watch.Close()
		//添加要监控的对象，文件或文件夹
		err = watch.Add("./")
		if err != nil {
			log.Fatal(err)
		}
		for {
			select {
			case ev := <-watch.Events:
				{
					//判断事件发生的类型，如下5种
					// Create 创建
					// Write 写入
					// Remove 删除
					// Rename 重命名
					// Chmod 修改权限
					if ev.Op&fsnotify.Create == fsnotify.Create {
						log.Println("创建文件 : ", ev.Name)
					}
					if ev.Op&fsnotify.Write == fsnotify.Write {
						switch ev.Name {
						case "homey.yml":
							fmt.Println("edit homey")
						case "feidan.yml":
							fmt.Println("edit feidan")
						default:
							fmt.Println("edit other:", ev.Name)
						}
					}
					if ev.Op&fsnotify.Remove == fsnotify.Remove {
						log.Println("删除文件 : ", ev.Name)
					}
					if ev.Op&fsnotify.Rename == fsnotify.Rename {
						log.Println("重命名文件 : ", ev.Name)
					}
					if ev.Op&fsnotify.Chmod == fsnotify.Chmod {
						log.Println("修改权限 : ", ev.Name)
					}
				}
			case err := <-watch.Errors:
				{
					log.Println("error : ", err)
					return
				}
			}
		}
	}()

	//循环
	select {}
}
