package yuanfuda_jinyouruantang

import (
	"fmt"
	"github.com/peanut-cc/peanut_case/plugin"
	"github.com/peanut-cc/peanut_case/utils/uuid"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
)

var ProductTypePrice = map[string]map[string]string{
	"3g*30粒/盒*2": {
		"barcode":   "0010449",
		"unitPrice": "45",
	},
}

func init() {
	fuYuanDa := &JinYouRuanTang{Name: plugin.DingXiang, CustomName: "无痕-媛福达"}
	plugin.PluginMap[plugin.FuYuanDa] = fuYuanDa
}

type JinYouRuanTang struct {
	Name       string
	CustomName string
}

func (p *JinYouRuanTang) GetPluginName() string {
	return p.Name
}

func (p *JinYouRuanTang) HandleUploadFile(fileName string) error {
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
		receivePeople := row[5]
		phone := row[6]
		province := row[8]
		city := row[9]
		county := row[10]
		address := row[12]
		salesChannelName := p.CustomName
		productName := row[19]
		numbers := row[20]
		productType := row[15]
		bardcodeAndPriceMap, ok := ProductTypePrice[productType]
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
	filename := fmt.Sprintf("./result/媛福达/媛福达_仅有软糖%v.xlsx", uuid.MustString())
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

func (p *JinYouRuanTang) DeleteUploadFile(fileName string) error {
	return os.Remove(fileName)
}
