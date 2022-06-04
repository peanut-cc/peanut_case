package main

import (
	"fmt"
	"github.com/peanut-cc/peanut_case/plugin"
	_ "github.com/peanut-cc/peanut_case/plugin/dingxiang"
)

func main() {
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
