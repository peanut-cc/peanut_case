package niuyisite

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
	niuyisite := &NiuYiSiTe{Name: plugin.NiuYiSiTe, CustomName: "无痕-北京纽伊斯特科贸有限公司（纽伊斯特）"}
	plugin.PluginMap[plugin.NiuYiSiTe] = niuyisite
}

type NiuYiSiTe struct {
	Name       string
	CustomName string
}

func (p *NiuYiSiTe) GetPluginName() string {
	return p.Name
}

func (p *NiuYiSiTe) HandleUploadFile(fileName string) error {
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
		province := ""
		city := ""
		county := ""
		address := row[4]
		salesChannelName := p.CustomName
		productName := row[5]
		numbers := row[8]
		barcode := ""
		unitPrice := ""
		if strings.Contains(productName, "罗伊氏") {
			barcode = "0010702"
			unitPrice = "288"
		} else if strings.Contains(productName, "钙DK") {
			barcode = "0010703"
			unitPrice = "152"
		} else if strings.Contains(productName, "黑枸杞") {
			barcode = "0010706"
			unitPrice = "398"
		} else if strings.Contains(productName, "玫瑰") {
			barcode = "0010707"
			unitPrice = "398"
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
	filename := fmt.Sprintf("./result/纽伊斯特/纽伊斯特%v.xlsx", uuid.MustString())
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

func (p *NiuYiSiTe) DeleteUploadFile(fileName string) error {
	return os.Remove(fileName)
}
