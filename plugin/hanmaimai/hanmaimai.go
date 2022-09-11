package hanmaimai

import (
	"fmt"
	"github.com/peanut-cc/peanut_case/plugin"
	"github.com/peanut-cc/peanut_case/utils/uuid"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
	"strings"
)

var ProductTypePrice = map[string]map[string]string{
	"闪亮蒸汽眼罩缓解眼疲劳热敷学生发热眼罩睡眠舒缓疲劳遮光透气【茶语】": {
		"barcode":   "6970869081288",
		"unitPrice": "22.4",
	},
	"闪亮蒸汽眼罩缓解眼疲劳热敷学生发热眼罩睡眠舒缓疲劳遮光透气【甜柚】": {
		"barcode":   "6970869081295",
		"unitPrice": "22.4",
	},
	"闪亮蒸汽眼罩缓解眼疲劳热敷学生发热眼罩睡眠舒缓疲劳遮光透气【榴恋】": {
		"barcode":   "6970869080946",
		"unitPrice": "22.4",
	},
	"闪亮蒸汽眼罩缓解眼疲劳热敷学生发热眼罩睡眠舒缓疲劳遮光透气【茶语+榴恋+甜柚】": {
		"barcode":   "200050039",
		"unitPrice": "44.8",
	},
	"仁和艾草帖颈椎贴多种中药草本精粹古法制成纯天然草本成分【12贴/盒*5】": {
		"barcode":   "0010037",
		"unitPrice": "47.2",
	},
	"仁和艾草帖颈椎贴多种中药草本精粹古法制成纯天然草本成分【12贴/盒*3】": {
		"barcode":   "0010036",
		"unitPrice": "31.2",
	},
	"仁和艾草帖颈椎贴多种中药草本精粹古法制成纯天然草本成分【12贴/盒】": {
		"barcode":   "6949805914188",
		"unitPrice": "15.2",
	},
	"【实发4瓶】叮当优品天然维VC软糖复合多种维生素0糖0脂45粒/瓶【买3赠1】": {
		"barcode":   "0010729",
		"unitPrice": "72.3",
	},
}

func init() {
	hanMaiMai := &HanMaiMai{Name: plugin.HanMaiMai, CustomName: "无痕-涵卖卖"}
	plugin.PluginMap[plugin.HanMaiMai] = hanMaiMai
}

type HanMaiMai struct {
	Name       string
	CustomName string
}

func (p *HanMaiMai) GetPluginName() string {
	return p.Name
}

func (p *HanMaiMai) HandleUploadFile(fileName string) error {
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
		if index <= 1 {
			continue
		}
		sn := row[1]
		shopSn := row[1]
		receivePeople := row[6]
		phone := row[7]
		province := ""
		city := ""
		county := ""
		address := row[5]
		salesChannelName := p.CustomName
		productInfo := row[8]
		productInfoList := strings.Split(productInfo, " ")
		if len(productInfoList) != 3 {
			log.Printf("客户 {%v} excel 第 {%v} 商品信息错误\n", p.Name, index)
			return err
		}
		productName := strings.TrimSpace(strings.Join(productInfoList[:2], ""))
		strings.TrimRight(strings.TrimLeft(strings.TrimSpace(productInfoList[2]), "【"), "】")
		numberRune := []rune(productInfoList[2])
		if len(numberRune) != 3 {
			log.Printf("客户 {%v} excel 第 {%v} 商品信息错误\n", p.Name, index)
			return err
		}
		numbers := string(numberRune[1])
		bardcodeAndPriceMap, ok := ProductTypePrice[productName]
		barcode := ""
		unitPrice := ""
		if ok {
			barcode = bardcodeAndPriceMap["barcode"]
			unitPrice = bardcodeAndPriceMap["unitPrice"]
		} else {
			barcode = "数据错误"
			unitPrice = "数据错误"
		}
		// 因为这家的表格是从第三行开始所以不用index+1
		axis := fmt.Sprintf("A%d", index)
		err = f.SetSheetRow(sheetName, axis, &[]string{sn, shopSn, "", "", "", "", p.CustomName, "", "", receivePeople, phone, "", "",
			province, city, county, address, "", "", "", "", "", "", "", "", salesChannelName, "",
			productName, barcode, "", "", "", numbers, unitPrice, "", "", "", "", "", "", "", "", "",
			"", "", "", "", "", "", "", ""})
		if err != nil {
			log.Printf("客户 {%v} 写excel 第 {%v} 错误\n", p.Name, index)
			return err
		}
	}
	filename := fmt.Sprintf("./result/涵卖卖/涵卖卖%v.xlsx", uuid.MustString())
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

func (p *HanMaiMai) DeleteUploadFile(fileName string) error {
	return os.Remove(fileName)
}
