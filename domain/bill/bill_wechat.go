package bill

import (
	"strconv"
	"time"

	"github.com/Enan01/notion_bill/common"
)

var WechatBillStartLineNo int

type WechatBillParser struct {
	BillFilePath string
}

func (p WechatBillParser) Parse() ([]Bill, error) {
	lineNo := WechatBillStartLineNo
	lineCols, err := common.LineColumnsCsvFile(p.BillFilePath, lineNo, -1)
	if err != nil {
		return nil, err
	}

	var bills = make([]Bill, 0)
	for _, lineCol := range lineCols {
		var (
			dealTime time.Time
			inout    BillInout
			amount   float64
		)

		dealTime, err = time.Parse("2006-01-02 15:04:05", lineCol[0])
		if err != nil {
			return nil, err
		}

		if lineCol[4] == "收入" {
			inout = BillIn
		} else if lineCol[4] == "支出" {
			inout = BillOut
		} else {
			inout = BillInoutUnknown
		}

		amountRune := []rune(lineCol[5])
		amount, err = strconv.ParseFloat(string(amountRune[1:]), 64)
		if err != nil {
			return nil, err
		}

		bill := Bill{
			DealTime:    dealTime,
			DealType:    lineCol[1],
			DealPartner: lineCol[2],
			Product:     lineCol[3],
			Inout:       inout,
			Amount:      amount,
			Source:      BillSourceWechat,
			Catecory:    "",
		}
		bills = append(bills, bill)
	}

	return bills, nil
}
