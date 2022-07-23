package haoyueshangmao

import (
	"fmt"
	"github.com/peanut-cc/peanut_case/plugin"
	"github.com/peanut-cc/peanut_case/utils/uuid"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
	"strings"
)

func init() {
	haoyueshangmao := &HaoYueShangMao{Name: plugin.HaoYueShangMao, CustomName: "无痕-广州市皓跃商贸有限公司（皓跃商贸）"}
	plugin.PluginMap[plugin.HaoYueShangMao] = haoyueshangmao
}

type HaoYueShangMao struct {
	Name       string
	CustomName string
}

func (p *HaoYueShangMao) GetPluginName() string {
	return p.Name
}

func (p *HaoYueShangMao) HandleUploadFile(fileName string) error {
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
		receivePeople := row[2]
		phone := row[3]
		province := row[4]
		city := row[5]
		county := row[6]
		address := row[7]
		salesChannelName := p.CustomName
		productName := row[9]
		numbers := row[12]

		barcode := ""
		unitPrice := ""
		if strings.Contains(productName, "三种香味各1瓶") {
			barcode = "0010139"
			unitPrice = "33"
		} else if strings.Contains(productName, "樱花香三瓶装") {
			barcode = "3000830103"
			unitPrice = "33"
		} else if strings.Contains(productName, "植物香三瓶装") {
			barcode = "3000830104"
			unitPrice = "33"
		} else if strings.Contains(productName, "薰衣草三瓶装") {
			barcode = "3000830102"
			unitPrice = "33"
		} else {
			barcode = "错误"
			unitPrice = "错误"
			log.Printf("客户 {%v} excel 第 {%v} 商品信息错误\n", p.Name, index)
			return err
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
	filename := fmt.Sprintf("./result/皓跃商贸/皓跃商贸%v.xlsx", uuid.MustString())
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

func (p *HaoYueShangMao) DeleteUploadFile(fileName string) error {
	return os.Remove(fileName)
}
