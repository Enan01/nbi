package bill

import "time"

// 账单来源
type BillSource string
type BillCatecory string

const (
	Unknown BillSource = "unknown"
	Alipay  BillSource = "alipay"
	Wechat  BillSource = "wechat"
)

const (
	Diet BillCatecory = "饮食"
)

// 账单类型
type Bill struct {
	DealTime    time.Time
	DealType    string
	DealPartner string
	Product     string
	Inout       int
	Amount      float32
	PayType     string
	Source      BillSource
	Catecory    BillCatecory
}
