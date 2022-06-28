package yuanfuda_fuyanjie

import (
	"fmt"
	"github.com/peanut-cc/peanut_case/plugin"
	"github.com/peanut-cc/peanut_case/utils/uuid"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
)

var SkuToBarcodePrice = map[string]map[string]string{
	"8171687192": {
		"barcode":   "300091001",
		"unitPrice": "24",
	},
	"8821982978": {
		"barcode":   "0010163",
		"unitPrice": "23.5",
	},
	"8418717878": {
		"barcode":   "0010291",
		"unitPrice": "20.7",
	},
	"8919894225": {
		"barcode":   "3000830015",
		"unitPrice": "39",
	},
	"8761455045": {
		"barcode":   "0010180",
		"unitPrice": "20",
	},
	"8170740142": {
		"barcode":   "6973601560782",
		"unitPrice": "71",
	},
	"8919894625": {
		"barcode":   "6939713005085",
		"unitPrice": "19",
	},
	"8170740954": {
		"barcode":   "0010348",
		"unitPrice": "183",
	},
	"8170106377": {
		"barcode":   "6973601560836",
		"unitPrice": "33",
	},
	"8172249756": {
		"barcode":   "300091002",
		"unitPrice": "24",
	},
	"8170740069": {
		"barcode":   "0010362",
		"unitPrice": "130",
	},
	"8170106513": {
		"barcode":   "0010414",
		"unitPrice": "66",
	},
	"8919894939": {
		"barcode":   "6939713003579",
		"unitPrice": "19",
	},
	"8171687556": {
		"barcode":   "6973601560072",
		"unitPrice": "8",
	},
}

func init() {
	fuYuanDaYuYanJie := &YuanFudaFuYanJie{Name: plugin.FuYuanDaFuYanJie, CustomName: "无痕-媛福达"}
	plugin.PluginMap[plugin.FuYuanDaFuYanJie] = fuYuanDaYuYanJie
}

type YuanFudaFuYanJie struct {
	Name       string
	CustomName string
}

func (p *YuanFudaFuYanJie) GetPluginName() string {
	return p.Name
}

func (p *YuanFudaFuYanJie) HandleUploadFile(fileName string) error {
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
		sn := row[0]
		shopSn := row[0]
		receivePeople := row[3]
		phone := row[4]
		province := ""
		city := ""
		county := ""
		address := row[5]
		salesChannelName := p.CustomName
		productName := row[7]
		numbers := row[10]
		sku := row[9]
		bardcodeAndPriceMap, ok := SkuToBarcodePrice[sku]
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
	filename := fmt.Sprintf("./result/媛福达/媛福达_妇炎洁%v.xlsx", uuid.MustString())
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

func (p *YuanFudaFuYanJie) DeleteUploadFile(fileName string) error {
	return os.Remove(fileName)
}
