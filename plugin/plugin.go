package plugin

import (
	"encoding/csv"
	"github.com/xuri/excelize/v2"
	"time"

	"log"
	"os"
)

const (
	DingXiang             = "丁香"
	MeiChu                = "美初"
	ShangHaiFanQi         = "上海梵迄"
	XiaoXiaoBaoMaMa       = "小小包麻麻"
	JinYouNingJiao        = "仅有凝胶"
	FuYuanDaJinYouRuanTan = "媛福达_仅有软糖"
	FuYuanDaFuYanJie      = "媛福达_妇炎洁"
	YunFan                = "云帆"
	JiWuKeJi              = "极物科技"
	HanMaiMai             = "涵卖卖"
	FangTuanZhang         = "房团长"
	HaoYueShangMao        = "皓跃商贸"
	AiKang                = "爱康"
)

var RowHeader = []string{"导入编号", "网店订单号", "下单时间", "付款时间", "承诺发货时间", "客户账号", "客户名称",
	"客户邮箱", "QQ", "收货人", "手机", "固定电话", "国家", "省份", "市（区）", "区（县）",
	"收货地址", "邮政编码", "发货仓库", "应收邮资", "平台佣金", "客付税额", "应收合计", "客服备注",
	"客户备注", "销售渠道名称", "结算方式", "货品名称", "条码", "货品编号", "规格", "批次号",
	"数量", "单价", "货品优惠", "金额", "网店子订单号", "定制码", "货品备注", "发票抬头",
	"发票类型", "证件类型", "证件号码", "证件使用姓名", "物流公司", "物流单号", "支付单号",
	"收款账户", "业务员", "跟单员", "标记"}

type Plugin interface {
	HandleUploadFile(fileName string) error
	DeleteUploadFile(fileName string) error
	GetPluginName() string
}

var PluginMap map[string]Plugin

func init() {
	PluginMap = make(map[string]Plugin, 10)
}

// ReadExcel 读取 excel 并返回所有行数据
func ReadExcel(fileName string) ([][]string, error) {
	var f *excelize.File
	var err error
	var num = 1
	// 每个文件尝试打开10次,10次之后打不开就放弃
	for num <= 10 {
		f, err = excelize.OpenFile(fileName)
		if err != nil {
			time.Sleep(1 * time.Second)
			log.Printf("打开文件错误:{%v} ,睡眠 1s 第{%v}次尝试打开文件{%v}", err, num, fileName)
			num += 1
			continue
		}
		break
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}()
	if err != nil {
		return nil, err
	}
	firstSheet := f.GetSheetName(0)
	rows, err := f.GetRows(firstSheet)
	if err != nil {
		return nil, err
	}
	log.Printf("读取{%v}内容成功\n", fileName)
	return rows, nil
}

// ReadCSV 解析csv数据
func ReadCSV(fileName string) ([][]string, error) {
	var f *os.File
	var err error
	var num = 1
	for num <= 10 {
		f, err = os.Open(fileName)
		if err != nil {
			time.Sleep(1 * time.Second)
			log.Printf("打开文件错误:{%v} ,睡眠 1s 第{%v}次尝试打开文件{%v}", err, num, fileName)
			num += 1
			continue
		}
		break
	}
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}()
	csvReader := csv.NewReader(f)
	csvReader.TrimLeadingSpace = true
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}
