package bill

import (
	"fmt"
	"reflect"
	"time"

	"github.com/Enan01/notion_bill/api/notion"
)

// 账单来源
type BillSource string
type BillCatecory string
type BillInout int

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

// 账单类型
type Bill struct {
	DealTime    time.Time    `propertyName:"交易时间" propertyType:"date"`
	DealType    string       `propertyName:"交易类型" propertyType:"rich_text"`
	DealPartner string       `propertyName:"交易对方" propertyType:"rich_text"`
	Product     string       `propertyName:"商品名称" propertyType:"title"`
	Inout       BillInout    `propertyName:"收/支" propertyType:"select"`
	Amount      float32      `propertyName:"金额（元）" propertyType:"number"`
	Source      BillSource   `propertyName:"来源" propertyType:"select"`
	Catecory    BillCatecory `propertyName:"账单类型" propertyType:"select"`
}

func (b Bill) ToCreatePageRequest() (req notion.CreatePageRequest, err error) {
	type propTag struct {
		name string
		typ  string
	}

	var (
		pageProperties   = make(notion.PageProperties)
		fieldPropMapping = make(map[string]propTag)
	)

	val := reflect.ValueOf(b).Elem()

	for i := 0; i < val.NumField(); i++ {
		vf := val.Field(i)
		tf := val.Type().Field(i)
		tag := tf.Tag

		fmt.Println(tag.Get("propertyName"), ",", tag.Get("propertyType"))

		fn, fy := tf.Name, tf.Type
		pn, pt := tag.Get("propertyName"), tag.Get("propertyType")

		if len(pn) == 0 || len(pt) == 0 {
			continue
		}

		
	}

	// ----------------------------

	t := reflect.TypeOf(b)
	reflect.ValueOf(b).Elem()

	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		fmt.Println(sf.Tag.Get("propertyName"), ",", sf.Tag.Get("propertyType"))

		fn, fy := sf.Name, sf.Type
		pn, pt := sf.Tag.Get("propertyName"), sf.Tag.Get("propertyType")
		if len(pn) == 0 || len(pt) == 0 {
			continue
		}

		fieldPropMapping[fn] = propTag{
			name: pn,
			typ:  pt,
		}
	}

	v := reflect.ValueOf(b)

	// switch notion.PropertyType(pt) {
	// case notion.PropDate:
	// case notion.PropTitle:
	// case notion.PropNumber:
	// case notion.PropRichText:
	// case notion.PropSelect:
	// default:
	// }
	return
}
