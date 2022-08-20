package main

import (
	"fmt"
	"github.com/peanut-cc/peanut_case/plugin"
	_ "github.com/peanut-cc/peanut_case/plugin/aikang"
	_ "github.com/peanut-cc/peanut_case/plugin/dingxiang"
	_ "github.com/peanut-cc/peanut_case/plugin/fangtuanzhang"
	_ "github.com/peanut-cc/peanut_case/plugin/hanmaimai"
	_ "github.com/peanut-cc/peanut_case/plugin/haoyueshangmao"
	_ "github.com/peanut-cc/peanut_case/plugin/jinyouningjiao"
	_ "github.com/peanut-cc/peanut_case/plugin/jiwukeji"
	_ "github.com/peanut-cc/peanut_case/plugin/meichu"
	_ "github.com/peanut-cc/peanut_case/plugin/shanghaifanqi"
	_ "github.com/peanut-cc/peanut_case/plugin/xiaoxiaobaomama"
	_ "github.com/peanut-cc/peanut_case/plugin/yuanfuda_fuyanjie"
	_ "github.com/peanut-cc/peanut_case/plugin/yuanfuda_jinyouruantang"
	_ "github.com/peanut-cc/peanut_case/plugin/yunfan"
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
	os.MkdirAll("./upload/媛福达/媛福达_仅有软糖", os.ModePerm)
	os.MkdirAll("./upload/媛福达/媛福达_妇炎洁", os.ModePerm)
	os.MkdirAll("./result/媛福达", os.ModePerm)
	os.MkdirAll("./upload/云帆", os.ModePerm)
	os.MkdirAll("./result/云帆", os.ModePerm)
	os.MkdirAll("./upload/极物科技", os.ModePerm)
	os.MkdirAll("./result/极物科技", os.ModePerm)
	os.MkdirAll("./upload/涵卖卖", os.ModePerm)
	os.MkdirAll("./result/涵卖卖", os.ModePerm)
	os.MkdirAll("./upload/房团长", os.ModePerm)
	os.MkdirAll("./result/房团长", os.ModePerm)
	os.MkdirAll("./upload/皓跃商贸", os.ModePerm)
	os.MkdirAll("./result/皓跃商贸", os.ModePerm)
	os.MkdirAll("./upload/爱康", os.ModePerm)
	os.MkdirAll("./result/爱康", os.ModePerm)
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
