package bill

import (
	"strconv"
	"time"

	"github.com/Enan01/nbi/common"
)

var AlipayBillStartLineNo int

type AlipayBillParser struct {
	BillFilePath string
}

func (p AlipayBillParser) Parse() ([]Bill, error) {
	startLineNo := AlipayBillStartLineNo
	lineCols, err := common.LineColumnsCsvFile(p.BillFilePath, startLineNo, -1)
	if err != nil {
		return nil, err
	}

	var bills = make([]Bill, 0)
	for _, lineCol := range lineCols {
		if lineCol[0] == "------------------------------------------------------------------------------------" {
			break
		}

		var (
			dealTime time.Time
			inout    BillInout
			amount   float64
		)

		dealTime, err = time.Parse("2006-01-02 15:04:05", lineCol[2])
		if err != nil {
			return nil, err
		}

		if lineCol[10] == "收入" {
			inout = BillIn
		} else if lineCol[10] == "支出" {
			inout = BillOut
		} else {
			inout = BillInoutUnknown
		}

		amountRune := []rune(lineCol[9])
		amount, err = strconv.ParseFloat(string(amountRune[:]), 64)
		if err != nil {
			return nil, err
		}

		bill := Bill{
			DealTime:    dealTime,
			DealType:    lineCol[6],
			DealPartner: lineCol[7],
			Product:     lineCol[8],
			Inout:       inout,
			Amount:      amount,
			Source:      BillSourceAlipay,
			Catecory:    "",
		}
		bills = append(bills, bill)
	}

	return bills, nil
}
