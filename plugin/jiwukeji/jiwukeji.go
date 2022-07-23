package jiwukeji

import (
	"fmt"
	"github.com/peanut-cc/peanut_case/plugin"
	"github.com/peanut-cc/peanut_case/utils/uuid"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
	"strings"
)

var SkuToBarcodePrice = map[string]map[string]string{
	"skuAUA3FB67": {
		"barcode":   "6973601560317",
		"unitPrice": "37.95",
	},
	"skuS2R5AOQD": {
		"barcode":   "0010057",
		"unitPrice": "82.5",
	},
	"skuEFAF5RM0": {
		"barcode":   "300030001",
		"unitPrice": "99",
	},
	"skuUH5SP2UQ": {
		"barcode":   "200010050",
		"unitPrice": "59.4",
	},
}

func init() {
	jiwukeji := &JiWuKeJi{Name: plugin.MeiChu, CustomName: "无痕-极物科技（广州）有限公司（极物科技）"}
	plugin.PluginMap[plugin.JiWuKeJi] = jiwukeji
}

type JiWuKeJi struct {
	Name       string
	CustomName string
}

func (p *JiWuKeJi) GetPluginName() string {
	return p.Name
}

func (p *JiWuKeJi) HandleUploadFile(fileName string) error {
	rows, err := plugin.ReadExcel(fileName)
	if err != nil {
		fmt.Printf("客户{%v} 打开excel 报错{%v}\n", p.GetPluginName(), err)
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
		sn := row[0]
		shopSn := row[0]
		receivePeople := row[2]
		phone := row[3]
		province := row[4]
		city := row[5]
		county := row[6]
		address := row[7]
		salesChannelName := p.CustomName
		productName := strings.TrimSpace(row[1])
		tmpbarcode := strings.TrimSpace(row[12])
		unitPrice := ""
		barcode := ""
		numbers := row[14]
		bardcodeAndPriceMap, ok := SkuToBarcodePrice[tmpbarcode]
		if ok {
			barcode = bardcodeAndPriceMap["barcode"]
			unitPrice = bardcodeAndPriceMap["unitPrice"]
		} else {
			barcode = "条码错误"
			log.Printf("客户{%v} 处理excel {%v} 行 条码处理错误\n", p.Name, index)
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
	filename := fmt.Sprintf("./result/极物科技/极物科技%v.xlsx", uuid.MustString())
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

func (p *JiWuKeJi) DeleteUploadFile(fileName string) error {
	return os.Remove(fileName)
}
