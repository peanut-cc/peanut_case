package aikang

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
	"6970869081288": "17.00",
}

func init() {
	aikang := &AiKang{Name: plugin.AiKang, CustomName: "无痕-天津爱康互联网信息服务有限公司"}
	plugin.PluginMap[plugin.AiKang] = aikang
}

type AiKang struct {
	Name       string
	CustomName string
}

func (p *AiKang) GetPluginName() string {
	return p.Name
}

func (p *AiKang) HandleUploadFile(fileName string) error {
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
		sn := row[3]
		shopSn := row[3]
		receivePeople := row[4]
		phone := row[5]
		province := ""
		city := ""
		county := ""
		address := row[6]
		salesChannelName := p.CustomName
		productName := row[7]
		barcode := strings.TrimSpace(row[9])
		numbers := row[10]
		unitPrice, ok := BoardeToPrice[barcode]
		if !ok {
			unitPrice = "条码错误"
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
	filename := fmt.Sprintf("./result/爱康/爱康%v.xlsx", uuid.MustString())
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

func (p *AiKang) DeleteUploadFile(fileName string) error {
	return os.Remove(fileName)
}
