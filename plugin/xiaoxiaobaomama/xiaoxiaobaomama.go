package xiaoxiaobaomama

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
	CustomName     = "无痕-宿迁市亿宝文化传播有限公司（钱儿）"
	CustomName2    = "无痕-宿迁市百宝信息科技有限公司（小小包麻麻）"
	BarcodeToPrice = map[string]string{
		"6970869081288": "17",
		"6970869080946": "17",
		"6970869081295": "17",
		"300091001":     "27",
		"300091002":     "26.5",
		"300083005021":  "26",
		"6973601560157": "31.7",
		"200010038":     "54.9",
		"6949805914188": "13",
		"0010036":       "27",
		"0010037":       "41",
		"6939713003579": "21.73",
		"6939713005085": "21.73",
		"6939713005092": "21.73",
		"3000830015":    "43.46",
		"6973601560119": "36",
		"0010041":       "66.5",
		"6973601560041": "75.89",
		"300030001":     "151.78",
		"3000300040":    "222.67",
		"6973601560409": "58.18",
		"10010071":      "103.86",
		"6973601560317": "50.7",
		"200010050":     "78.9",
		"0010057":       "109.1",
		"6973601560348": "25.28",
		"0010068":       "46.06",
		"6973601560485": "72",
		"0010106":       "108",
		"0010125":       "120",
		"6973601560836": "47.6",
		"0010414":       "95.2",
		"0010509":       "42",
	}
)

func init() {
	mama := &XiaoXiaoBaoMaMa{Name: plugin.XiaoXiaoBaoMaMa, CustomName: "无痕-宿迁市亿宝文化传播有限公司（钱儿）"}
	plugin.PluginMap[plugin.XiaoXiaoBaoMaMa] = mama
}

type XiaoXiaoBaoMaMa struct {
	Name       string
	CustomName string
}

func (p *XiaoXiaoBaoMaMa) GetPluginName() string {
	return p.Name
}

func (p *XiaoXiaoBaoMaMa) HandleUploadFile(fileName string) error {
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
		sn := row[2]
		shopSn := row[2]
		receivePeople := row[5]
		phone := row[6]
		province := row[7]
		city := row[8]
		county := row[9]
		address := row[11]
		SalesChannelName := ""
		if strings.Contains(strings.TrimSpace(row[0]), "钱儿") {
			SalesChannelName = CustomName
		} else {
			SalesChannelName = CustomName2
		}
		productName := row[14]
		barcode := strings.TrimSpace(row[16])
		unitPrice, ok := BarcodeToPrice[barcode]
		if !ok {
			unitPrice = "条码错误"
		}
		numbers := row[18]
		axis := fmt.Sprintf("A%d", index+1)
		err = f.SetSheetRow(sheetName, axis, &[]string{sn, shopSn, "", "", "", "", SalesChannelName, "", "", receivePeople, phone, "", "",
			province, city, county, address, "", "", "", "", "", "", "", "", SalesChannelName, "",
			productName, barcode, "", "", "", numbers, unitPrice, "", "", "", "", "", "", "", "", "",
			"", "", "", "", "", "", "", ""})
		if err != nil {
			log.Printf("客户 {%v} 写excel 第 {%v} 错误\n", p.Name, index)
			return err
		}
	}
	filename := fmt.Sprintf("./result/小小包麻麻/小小包麻麻%v.xlsx", uuid.MustString())
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

func (p *XiaoXiaoBaoMaMa) DeleteUploadFile(fileName string) error {
	return os.Remove(fileName)
}
