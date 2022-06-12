package yunfan

import (
	"fmt"
	"github.com/peanut-cc/peanut_case/plugin"
	"github.com/peanut-cc/peanut_case/utils/uuid"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
	"strings"
)

var (
	JiGuo           = "无痕-天津极果优品电子商务有限公司（极果）"
	JunJun          = "无痕-上海卿禾科技工作室（君君辅食记）"
	LeTuiZu         = "无痕-北京银铃文化传播有限公司（乐退族）"
	ABaoYanXuan     = "无痕-湖南态度日升网络科技有限公司（态度心选）"
	TaoMeiWu        = "无痕-桃美物"
	WeiYangYouPin   = "无痕-未央优品"
	YiFeiFan        = "无痕-上海艺梵电子商务有限公司（艺非凡）"
	ShiDianDuShu    = "无痕-厦门十点电子商务有限公司（十点读书)"
	LiXiangShengHuo = "无痕-浙江和跃天明贸易有限公司（小羽私厨）"
)

var SaleChannelBarcodePrice = map[string]map[string]string{
	JiGuo: {
		"0010414":       "59",
		"6973601560836": "28",
	},
	JunJun: {
		"6973601560485": "59",
		"0010106":       "89",
		"0010125":       "99",
		"0010368":       "99",
		"6973601560379": "26",
		"0010056":       "46",
		"0010148":       "65.5",
		"300083002":     "26",
	},
	ABaoYanXuan: {
		"300087005033":  "33.64",
		"6973601560027": "19.07",
	},
	TaoMeiWu: {
		"0010414":       "59",
		"6973601560836": "28",
	},
	WeiYangYouPin: {
		"0010414":       "59",
		"6973601560836": "28",
		"6973601560867": "24",
		"0010449":       "39",
		"0010450":       "49",
		"0010451":       "60",
	},
	ShiDianDuShu: {
		"0010036":       "27",
		"0010037":       "41",
		"6970869081288": "18",
		"6970869080946": "18",
		"6970869081295": "18",
		"200050039":     "36",
	},
	LiXiangShengHuo: {
		"300083001036": "30.88",
	},
}

var NoBarcode = map[string]map[string]map[string]string{
	LeTuiZu: {
		"3g*45粒/盒": {
			"barcode":   "6973601560362盒",
			"unitPrice": "23",
		},
		"3g*45粒/盒*2": {
			"barcode":   "0010357",
			"unitPrice": "41",
		},
		"买2送1": {
			"barcode":   "0010358",
			"unitPrice": "60",
		},
	},
	YiFeiFan: {
		"50g": {
			"barcode":   "6973601560119",
			"unitPrice": "19.9",
		},
	},
}

func init() {
	// 云帆比较特殊，会包含多个不同 CustomName
	yunfan := &YunFan{Name: plugin.YunFan, CustomName: ""}
	plugin.PluginMap[plugin.YunFan] = yunfan
}

type YunFan struct {
	Name       string
	CustomName string
}

func (p *YunFan) GetPluginName() string {
	return p.Name
}

func (p *YunFan) HandleUploadFile(fileName string) error {
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
		tmpChannelName := row[0]
		sn := row[4]
		shopSn := row[4]
		receivePeople := row[9]
		phone := row[8]
		province := row[11]
		city := row[12]
		county := row[13]
		address := row[14]
		productName := row[19]
		numbers := row[28]

		// 不一定有barcode
		barcode := row[27]
		unitPrice := ""
		productType := row[23]

		if strings.Contains(tmpChannelName, "极果") {
			p.CustomName = JiGuo
			barcodePrice := SaleChannelBarcodePrice[JiGuo]
			_, ok := barcodePrice[barcode]
			if !ok {
				unitPrice = "数据错误"
			} else {
				unitPrice = barcodePrice[barcode]
			}
		} else if strings.Contains(tmpChannelName, "君君") {
			p.CustomName = JunJun
			barcodePrice := SaleChannelBarcodePrice[JunJun]
			_, ok := barcodePrice[barcode]
			if !ok {
				unitPrice = "数据错误"
			} else {
				unitPrice = barcodePrice[barcode]
			}
		} else if strings.Contains(tmpChannelName, "乐退族") {
			p.CustomName = LeTuiZu
			noBarcodePriceMap := NoBarcode[LeTuiZu]
			if strings.Contains(productType, "3g*45粒/盒") {
				productType = "3g*45粒/盒"
			} else if strings.Contains(productType, "3g*45粒/盒*2") {
				productType = "3g*45粒/盒*2"
			} else if strings.Contains(productType, "买2送1") {
				productType = "买2送1"
			} else {
				productType = ""
			}
			BarcodePrice, ok := noBarcodePriceMap[productType]
			if ok {
				barcode = BarcodePrice["barcode"]
				unitPrice = BarcodePrice["unitPrice"]
			} else {
				barcode = "数据错误"
				unitPrice = "数据错误"
			}
		} else if strings.Contains(tmpChannelName, "阿宝严选") {
			p.CustomName = ABaoYanXuan
			barcodePrice := SaleChannelBarcodePrice[ABaoYanXuan]
			_, ok := barcodePrice[barcode]
			if !ok {
				unitPrice = "数据错误"
			} else {
				unitPrice = barcodePrice[barcode]
			}
		} else if strings.Contains(tmpChannelName, "桃美物") {
			p.CustomName = TaoMeiWu
			barcodePrice := SaleChannelBarcodePrice[TaoMeiWu]
			_, ok := barcodePrice[barcode]
			if !ok {
				unitPrice = "数据错误"
			} else {
				unitPrice = barcodePrice[barcode]
			}
		} else if strings.Contains(tmpChannelName, "未央优品") {
			p.CustomName = WeiYangYouPin
			barcodePrice := SaleChannelBarcodePrice[WeiYangYouPin]
			_, ok := barcodePrice[barcode]
			if !ok {
				unitPrice = "数据错误"
			} else {
				unitPrice = barcodePrice[barcode]
			}
		} else if strings.Contains(tmpChannelName, "艺非凡") {
			p.CustomName = YiFeiFan
			noBarcodePriceMap := NoBarcode[YiFeiFan]
			if strings.Contains(productType, "50g") {
				productType = "50g"
			} else {
				productType = ""
			}
			BarcodePrice, ok := noBarcodePriceMap[productType]
			if ok {
				barcode = BarcodePrice["barcode"]
				unitPrice = BarcodePrice["unitPrice"]
			} else {
				barcode = "数据错误"
				unitPrice = "数据错误"
			}
		} else if strings.Contains(tmpChannelName, "十点读书") {
			p.CustomName = ShiDianDuShu
			barcodePrice := SaleChannelBarcodePrice[ShiDianDuShu]
			realBarcode := strings.Split(barcode, "FS-")
			if len(realBarcode) != 2 {
				unitPrice = "数据错误"
			}
			barcode = realBarcode[1]
			_, ok := barcodePrice[barcode]
			if !ok {
				unitPrice = "数据错误"
			} else {
				unitPrice = barcodePrice[barcode]
			}
		} else if strings.Contains(tmpChannelName, "理想生活") {
			p.CustomName = LiXiangShengHuo
			barcodePrice := SaleChannelBarcodePrice[LiXiangShengHuo]
			_, ok := barcodePrice[barcode]
			if !ok {
				unitPrice = "数据错误"
			} else {
				unitPrice = barcodePrice[barcode]
			}
		} else {
			p.CustomName = "未知"
			barcode = "数据错误"
			unitPrice = "数据错误"
		}

		salesChannelName := p.CustomName

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
	filename := fmt.Sprintf("./result/云帆/云帆%v.xlsx", uuid.MustString())
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

func (p *YunFan) DeleteUploadFile(fileName string) error {
	return os.Remove(fileName)
}
