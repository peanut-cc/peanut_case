package shidiandushu

import (
	"fmt"
	"github.com/peanut-cc/peanut_case/plugin"
	"github.com/peanut-cc/peanut_case/utils/uuid"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
	"strings"
)

var BoardeToPrice = map[string]string{
	"0010036":       "27",
	"0010037":       "41",
	"6970869081288": "18",
	"6970869080946": "18",
	"6970869081295": "18",
	"200050039":     "36",
	"0010414":       "77",
	"0010301":       "36",
	"0010302":       "36",
	"0010303":       "36",
}

func init() {
	shidiandushu := &ShiDianDuShu{Name: plugin.DingXiang, CustomName: "无痕-厦门十点电子商务有限公司（十点读书)"}
	plugin.PluginMap[plugin.ShiDianDuShu] = shidiandushu
}

type ShiDianDuShu struct {
	Name       string
	CustomName string
}

func (p *ShiDianDuShu) GetPluginName() string {
	return p.Name
}

func (p *ShiDianDuShu) HandleUploadFile(fileName string) error {
	rows, err := plugin.ReadCSV(fileName)
	if err != nil {
		fmt.Printf("客户{%v} 打开csv 报错{%v}\n", p.GetPluginName(), err)
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
		logisticsNo := row[15]
		if strings.TrimSpace(logisticsNo) != "" {
			break
		}
		sn := row[0]
		shopSn := row[0]
		receivePeople := row[2]
		phone := row[3]
		province := ""
		city := ""
		county := ""
		address := strings.TrimSpace(row[4]) + strings.TrimSpace(row[7])
		salesChannelName := p.CustomName
		productName := row[9]
		barcode := row[10]
		numbers := row[11]
		unitPrice := ""
		realBarcode := strings.Split(barcode, "FS-")
		if len(realBarcode) != 2 {
			unitPrice = "数据错误"
		}
		barcode = realBarcode[1]
		_, ok := BoardeToPrice[barcode]
		if !ok {
			unitPrice = "数据错误"
		} else {
			unitPrice = BoardeToPrice[barcode]
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
	filename := fmt.Sprintf("./result/十点读书/十点读书%v.xlsx", uuid.MustString())
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

func (p *ShiDianDuShu) DeleteUploadFile(fileName string) error {
	return os.Remove(fileName)
}
