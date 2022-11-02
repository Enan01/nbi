package bill

import (
	"errors"
	"reflect"
	"time"

	"github.com/Enan01/notion_bill/api/notion"
)

var (
	ErrPropertyTypeInvalid = errors.New("property type is invalid")
)

var BillDatabaseId string

// 账单来源
type BillSource string

// 账单类型
type BillCatecory string

// 收支类型
type BillInout int

func (s BillSource) String() string {
	switch s {
	case BillSourceAlipay:
		return "支付宝"
	case BillSourceWechat:
		return "微信"
	default:
		return "未知"
	}
}

func (s BillSource) Color() notion.Color {
	switch s {
	case BillSourceAlipay:
		return notion.ColorBlue
	case BillSourceWechat:
		return notion.ColorGreen
	default:
		return notion.ColorDefault
	}
}

func (io BillInout) String() string {
	switch io {
	case BillIn:
		return "收入"
	case BillOut:
		return "支出"
	default:
		return "未知"
	}
}

func (io BillInout) Color() notion.Color {
	switch io {
	case BillIn:
		return notion.ColorGreen
	case BillOut:
		return notion.ColorRed
	default:
		return notion.ColorDefault
	}
}

const (
	BillSourceUnknown BillSource = "unknown"
	BillSourceAlipay  BillSource = "alipay"
	BillSourceWechat  BillSource = "wechat"
)

const (
	BillCatecoryDiet BillCatecory = "饮食"
)

const (
	BillIn  BillInout = 1
	BillOut BillInout = 2
)

const (
	BillTagPropName = "propertyName"
	BillTagPropType = "propertyType"
)

// 账单类型
type Bill struct {
	DealTime    time.Time    `propertyName:"交易时间" propertyType:"date"`
	DealType    string       `propertyName:"交易类型" propertyType:"rich_text"`
	DealPartner string       `propertyName:"交易对方" propertyType:"rich_text"`
	Product     string       `propertyName:"商品名称" propertyType:"title"`
	Inout       BillInout    `propertyName:"收/支" propertyType:"select"`
	Amount      float64      `propertyName:"金额（元）" propertyType:"number"`
	Source      BillSource   `propertyName:"来源" propertyType:"select"`
	Catecory    BillCatecory `propertyName:"-" propertyType:"-"`
}

func (b Bill) ToCreatePageRequest() (req notion.CreatePageRequest, err error) {
	var (
		pageProperties = make(notion.PageProperties)
	)

	val := reflect.ValueOf(b)

	for i := 0; i < val.NumField(); i++ {
		tf := val.Type().Field(i)
		pn, pt := tf.Tag.Get(BillTagPropName), tf.Tag.Get(BillTagPropType)

		if len(pn) == 0 || len(pt) == 0 || pn == "-" || pt == "-" {
			continue
		}

		fv := val.Field(i).Interface()

		var prop notion.Property
		prop, err = parseFieldToNotionProperty(notion.ToPropertyType(pt), fv)
		if err != nil {
			return
		}

		pageProperties.Append(pn, prop)
	}

	req = notion.CreatePageRequest{
		Parent: struct {
			DatabaseID string "json:\"database_id\""
		}{
			DatabaseID: BillDatabaseId,
		},
		Properties: pageProperties,
	}

	return
}

func parseFieldToNotionProperty(pt notion.PropertyType, fv interface{}) (prop notion.Property, err error) {
	switch pt {
	case notion.PropTypeDate:
		if d, ok := fv.(time.Time); !ok {
			err = ErrPropertyTypeInvalid
			return
		} else {
			prop = notion.DateP{
				Start: d.Format(time.RFC3339),
			}
			return
		}
	case notion.PropTypeNumber:
		if n, ok := fv.(float64); !ok {
			err = ErrPropertyTypeInvalid
			return
		} else {
			prop = notion.NumberP(n)
			return
		}
	case notion.PropTypeSelect:
		switch v := fv.(type) {
		case BillInout:
			prop = notion.SelectP{
				// Id:    v.SelectId(),
				Name: v.String(),
				// Color: v.Color(),
			}
			return
		case BillSource:
			prop = notion.SelectP{
				// Id:    v.SelectId(),
				Name: v.String(),
				// Color: v.Color(),
			}
			return
		default:
			err = ErrPropertyTypeInvalid
			return
		}
	case notion.PropTypeTitle:
		if s, ok := fv.(string); !ok {
			err = ErrPropertyTypeInvalid
			return
		} else {
			prop = notion.TitleP{
				notion.RichText{
					Type: notion.RichTextTypeText,
					Text: &notion.RichTextText{
						Content: s,
					},
				}}
			return
		}
	case notion.PropTypeRichText:
		if s, ok := fv.(string); !ok {
			err = ErrPropertyTypeInvalid
			return
		} else {
			prop = notion.RichTextP{
				notion.RichText{
					Type: notion.RichTextTypeText,
					Text: &notion.RichTextText{
						Content: s,
					},
				},
			}
			return
		}
	default:
		return
	}
}
