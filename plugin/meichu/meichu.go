package meichu

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
	"E211231-3": {
		"barcode":   "0010414",
		"unitPrice": "59",
	},
	"E211231-1": {
		"barcode":   "6973601560836",
		"unitPrice": "28",
	},
	"skuTCMR9CI9": {
		"barcode":   "6973601560997",
		"unitPrice": "30",
	},
	"skuP84OS9U4": {
		"barcode":   "0010595",
		"unitPrice": "68",
	},
	"sku0CD830PD": {
		"barcode":   "0010596",
		"unitPrice": "99",
	},
}

func init() {
	meichu := &Meichu{Name: plugin.MeiChu, CustomName: "无痕-美初"}
	plugin.PluginMap[plugin.MeiChu] = meichu
}

type Meichu struct {
	Name       string
	CustomName string
}

func (p *Meichu) GetPluginName() string {
	return p.Name
}

func (p *Meichu) HandleUploadFile(fileName string) error {
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
		receivePeople := row[3]
		phone := row[4]
		province := row[5]
		city := row[6]
		county := row[7]
		address := row[9]
		salesChannelName := p.CustomName
		productName := strings.TrimSpace(row[10])
		tmpbarcode := strings.TrimSpace(row[12])
		unitPrice := ""
		barcode := ""
		numbers := row[13]
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
	filename := fmt.Sprintf("./result/美初/美初%v.xlsx", uuid.MustString())
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

func (p *Meichu) DeleteUploadFile(fileName string) error {
	return os.Remove(fileName)
}
