package bill

import "time"

// 账单来源
type BillSource string

const (
	Unknown BillSource = "unknown"
	Alipay  BillSource = "alipay"
	Wechat  BillSource = "wechat"
)

// 账单类型
type Bill struct {
	DealTime     time.Time
	DealType     string
	DealPartner  string
	Product      string
	Inout        int
	Amount       float32
	PayType      string
	Source       BillSource
	BillCatecory string
}
