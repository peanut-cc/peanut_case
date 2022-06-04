package plugin

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

const (
	DingXiang = "dingxiang"
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
func ReadExcel(p Plugin, fileName string) ([][]string, error) {
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	firstSheet := f.GetSheetName(0)
	rows, err := f.GetRows(firstSheet)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
