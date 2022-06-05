package shanghaifanqi

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
	"6973601560041":  "60",
	"300030001":      "109",
	"3000300040":     "158",
	"6973601560034":  "43",
	"300030005022":   "75",
	"0010073":        "97",
	"6973601560317":  "34.5",
	"200010050":      "54",
	"0010057":        "75",
	"6973601560836":  "33",
	"0010414":        "66",
	"6973601560829":  "9",
	"0010423":        "18.5",
	"9421900027570":  "219",
	"6973601560706":  "132",
	"0010222":        "220",
	"6973601560720":  "132",
	"0010221":        "220",
	"6973601560768":  "132",
	"6973601560751":  "132",
	"0010518":        "220",
	"6973601560485":  "65",
	"0010106":        "97",
	"0010125":        "108",
	"6973601560805":  "18",
	"0010432":        "30",
	"6973601560799":  "18",
	"0010315":        "30",
	"6973601560492":  "99",
	"20010109":       "148",
	"6973601560782":  "71",
	"0010362":        "130",
	"0010348":        "183",
	"6972138271079":  "18",
	"0010144":        "30",
	"0010162":        "40",
	"6973601560195":  "49",
	"0010355":        "76",
	"200010079":      "99",
	"0010369":        "160",
	"6973601560324":  "32",
	"0010052":        "57",
	"6973601560355":  "29",
	"0010053":        "56",
	"0010054":        "75",
	"6973601560867":  "27",
	"0010449":        "45",
	"0010450":        "58",
	"0010451":        "72",
	"6973601560652":  "30",
	"0010468":        "53",
	"0010469":        "59",
	"6973601560553":  "41",
	"0010193":        "77",
	"0010177":        "108",
	"6973601560362盒": "23",
	"0010357":        "41",
	"0010358":        "59",
	"6973601560379":  "23",
	"0010056":        "41",
	"0010148":        "59",
	"6973601560409":  "54",
	"10010071":       "98",
	"6973601560072":  "8",
	"300091001":      "24",
	"6973601560096":  "8",
	"300091002":      "24",
	"6973601560065":  "11",
	"300083002":      "24",
	"6973601560027":  "19",
	"300087004":      "34",
	"6973601560089":  "15",
	"300083001036":   "24",
	"6973601560119":  "19",
	"0010041":        "32",
	"6973601560003":  "14",
	"0010117":        "24",
	"6949805914188":  "11",
	"0010036":        "23",
	"0010037":        "35",
	"6970869081288":  "16",
	"6970869080946":  "16",
	"6970869081295":  "16",
	"200050039":      "31",
	"6939713003579":  "16",
	"6939713005085":  "16",
	"6939713005092":  "16",
	"0010139":        "36",
	"6923601907377":  "13",
	"0010163":        "21",
	"6923601900675":  "19",
	"0010291":        "30",
	"0010180":        "17",
	"0010178":        "21.5",
}

func init() {
	shanghaifanqi := &ShangHaiFanQi{Name: plugin.ShangHaiFanQi, CustomName: "无痕-上海东荟西通商贸有限公司（梵迄）"}
	plugin.PluginMap[plugin.ShangHaiFanQi] = shanghaifanqi
}

type ShangHaiFanQi struct {
	Name       string
	CustomName string
}

func (p *ShangHaiFanQi) GetPluginName() string {
	return p.Name
}

func (p *ShangHaiFanQi) HandleUploadFile(fileName string) error {
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
		sn := row[2]
		shopSn := row[2]
		receivePeople := row[8]
		phone := row[9]
		province := ""
		city := ""
		county := ""
		address := row[10]
		salesChannelName := p.CustomName
		productName := row[3]
		barcode := strings.TrimSpace(row[7])
		//unitPrice := ""
		unitPrice, ok := BoardeToPrice[barcode]
		if !ok {
			key := fmt.Sprintf("00%v", barcode)
			unitPrice2, ok2 := BoardeToPrice[key]
			if ok2 {
				unitPrice = unitPrice2
			} else {
				unitPrice = "条码错误"
			}
		}
		numbers := row[5]
		axis := fmt.Sprintf("A%d", index+1)
		err = f.SetSheetRow(sheetName, axis, &[]string{sn, shopSn, "", "", "", "", p.CustomName, "", "", receivePeople, phone, "", "",
			province, city, county, address, "", "", "", "", "", "", "", "", salesChannelName, "",
			productName, barcode, "", "", "", numbers, unitPrice, "", "", "", "", "", "", "", "", "",
			"", "", "", "", "", "", "", ""})
		if err != nil {
			log.Printf("客户 {%v} 写excel 第 {%v} 错误:{%v}", p.Name, index, err)
			return err
		}
	}
	filename := fmt.Sprintf("./result/上海梵迄/梵迄%v.xlsx", uuid.MustString())
	err = f.SaveAs(filename)
	if err != nil {
		log.Printf("保存{%v} 失败:{%v}\n", filename, err)
		return err
	}
	err = p.DeleteUploadFile(fileName)
	if err != nil {
		log.Printf("删除{%v} 失败:{%v}\n", fileName, err)
		return err
	}
	log.Printf("客户{%v} 订单{%v}处理完毕\n", p.Name, fileName)
	return nil
}

func (p *ShangHaiFanQi) DeleteUploadFile(fileName string) error {
	return os.Remove(fileName)
}
