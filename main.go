package main

import (
	"fmt"
	"github.com/peanut-cc/peanut_case/plugin"
	_ "github.com/peanut-cc/peanut_case/plugin/dingxiang"
	_ "github.com/peanut-cc/peanut_case/plugin/jinyouningjiao"
	_ "github.com/peanut-cc/peanut_case/plugin/meichu"
	_ "github.com/peanut-cc/peanut_case/plugin/shanghaifanqi"
	_ "github.com/peanut-cc/peanut_case/plugin/xiaoxiaobaomama"
	"os"
)

func initResultDir() {
	os.MkdirAll("./result/上海梵迄", os.ModePerm)
	os.MkdirAll("./upload/上海梵迄", os.ModePerm)
	os.MkdirAll("./result/丁香", os.ModePerm)
	os.MkdirAll("./upload/丁香", os.ModePerm)
	os.MkdirAll("./result/小小包麻麻", os.ModePerm)
	os.MkdirAll("./upload/小小包麻麻", os.ModePerm)
	os.MkdirAll("./upload/美初", os.ModePerm)
	os.MkdirAll("./result/美初", os.ModePerm)
	os.MkdirAll("./upload/跟团号/仅有凝胶", os.ModePerm)
	os.MkdirAll("./result/跟团号/仅有凝胶", os.ModePerm)
}

func main() {
	initResultDir()
	initMonitor()
	for {
		select {
		case message := <-fileChan:
			plugin, ok := plugin.PluginMap[message.custom]
			if ok {
				plugin.HandleUploadFile(message.filename)
			} else {
				fmt.Println("未找到插件")
			}

		default:
		}
	}
}
