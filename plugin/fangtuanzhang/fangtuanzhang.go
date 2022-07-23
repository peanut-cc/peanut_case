package fangtuanzhang

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
	fangtuanzhang := &FangTuanZhang{Name: plugin.DingXiang, CustomName: "无痕-房团长"}
	plugin.PluginMap[plugin.FangTuanZhang] = fangtuanzhang
}

type FangTuanZhang struct {
	Name       string
	CustomName string
}

func (p *FangTuanZhang) GetPluginName() string {
	return p.Name
}

func (p *FangTuanZhang) HandleUploadFile(fileName string) error {
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
		province := row[5]
		city := row[6]
		county := row[7]
		address := row[8]
		salesChannelName := p.CustomName
		productName := row[10]
		numbers := row[13]
		unitPrice := ""
		barcode := ""
		if strings.Contains(productName, "规格一") {
			barcode = "6973601560997"
			unitPrice = "23.5"
		} else if strings.Contains(productName, "规格二") {
			barcode = "0010595"
			unitPrice = "54"
		} else if strings.Contains(productName, "规格三") {
			barcode = "0010596"
			unitPrice = "79"
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
	filename := fmt.Sprintf("./result/房团长/房团长%v.xlsx", uuid.MustString())
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

func (p *FangTuanZhang) DeleteUploadFile(fileName string) error {
	return os.Remove(fileName)
}
