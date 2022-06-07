package jinyouningjiao

import (
	"fmt"
	"github.com/peanut-cc/peanut_case/plugin"
	"github.com/peanut-cc/peanut_case/utils/uuid"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
)

var ProductTypePrice = map[string]map[string]string{
	"DF*1盒【妇炎洁妇用洁炎凝胶】": {
		"barcode":   "6973601560997",
		"unitPrice": "30",
	},
	"DF*3盒【妇炎洁妇用洁炎凝胶】": {
		"barcode":   "0010595",
		"unitPrice": "68",
	},
	"DF*5盒【妇炎洁妇用洁炎凝胶】": {
		"barcode":   "0010596",
		"unitPrice": "99",
	},
}

// 备注这里会有多个客户都是属于 “无痕-毛伟大团长”
func init() {
	jinYouNingJiao := &JinYouNingJiao{Name: plugin.JinYouNingJiao, CustomName: "无痕-毛伟大团长"}
	plugin.PluginMap[plugin.JinYouNingJiao] = jinYouNingJiao
}

type JinYouNingJiao struct {
	Name       string
	CustomName string
}

func (p *JinYouNingJiao) GetPluginName() string {
	return p.Name
}

func (p *JinYouNingJiao) HandleUploadFile(fileName string) error {
	rows, err := plugin.ReadExcel(fileName)
	if err != nil {
		fmt.Printf("客户{%v} 打开excel 报错{%v}", p.GetPluginName(), err)
		return err
	}
	f := excelize.NewFile()
	sheet := f.NewSheet("Sheet1")
	sheetName := f.GetSheetName(sheet)
	err = f.SetSheetRow(sheetName, "A1", &plugin.RowHeader)
	if err != nil {
		log.Printf("客户 {%v} 创建excel失败\n", p.Name)
		return err
	}

	for index, row := range rows {
		// 第一行的标题不做处理
		if index == 0 {
			continue
		}
		sn := row[1]
		shopSn := row[1]
		receivePeople := row[7]
		phone := row[8]
		province := ""
		city := ""
		county := ""
		address := row[9]
		salesChannelName := p.CustomName
		productName := row[4]
		numbers := row[6]
		bardcodeAndPriceMap, ok := ProductTypePrice[productName]
		barcode := ""
		unitPrice := ""
		if ok {
			barcode = bardcodeAndPriceMap["barcode"]
			unitPrice = bardcodeAndPriceMap["unitPrice"]
		} else {
			barcode = "数据错误"
			unitPrice = "数据错误"
		}
		axis := fmt.Sprintf("A%d", index+1)
		err = f.SetSheetRow(sheetName, axis, &[]string{sn, shopSn, "", "", "", "", p.CustomName, "", "", receivePeople, phone, "", "",
			province, city, county, address, "", "", "", "", "", "", "", "", salesChannelName, "",
			productName, barcode, "", "", "", numbers, unitPrice, "", "", "", "", "", "", "", "", "",
			"", "", "", "", "", "", "", ""})
		if err != nil {
			log.Printf("客户 {%v} 写excel 第 {%v} 错误\n", p.Name, index)
			return err
		}
	}
	filename := fmt.Sprintf("./result/跟团号/仅有凝胶/跟团号-仅有凝胶%v.xlsx", uuid.MustString())
	err = f.SaveAs(filename)
	if err != nil {
		log.Printf("保存{%v} 失败\n", filename)
		return err
	}
	if err := p.DeleteUploadFile(fileName); err != nil {
		log.Printf("删除{%v} 失败\n", fileName)
		return err
	}
	log.Printf("客户{%v} 订单{%v}处理完毕\n", p.Name, fileName)
	return nil
}

func (p *JinYouNingJiao) DeleteUploadFile(fileName string) error {
	return os.Remove(fileName)
}
