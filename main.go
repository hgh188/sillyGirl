package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/cdle/sillyGirl/core"
)

func main() {
	go monitorGoroutine()
	core.Init123()
	sillyGirl := core.Bucket("sillyGirl")
	port := sillyGirl.Get("port", "8080")
	logs.Info("Http服务已运行(%s)。", sillyGirl.Get("port", "8080"))
	go core.Server.Run("0.0.0.0:" + port)
	logs.Info("关注频道 https://t.me/kczz2021 获取最新消息。")
	reader := bufio.NewReader(os.Stdin)
	d := false
	for _, arg := range os.Args {
		if arg == "-d" {
			d = true
		}
	}
	if !d {
		for _, arg := range os.Args {
			if arg == "-t" {
				logs.Info("终端交互已启用。")
				for {
					data, _, _ := reader.ReadLine()
					core.Senders <- &core.Faker{
						Type:    "terminal",
						Message: string(data),
					}

				}
			}
		}
		logs.Info("终端交互不可用，运行带-t参数即可启用。")
	}

	select {}
}

func monitorGoroutine() {
	ticker := time.NewTicker(time.Millisecond * 100)
	lastGNum := 0
	for {
		<-ticker.C
		if lastGNum != runtime.NumGoroutine() {
			fmt.Println("<========================", time.Now().Format("2006-01-02 15:04:05"), "Goroutine Number :", runtime.NumGoroutine(), "=========================>")
			lastGNum = runtime.NumGoroutine()
		}
	}
}
