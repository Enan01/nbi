package bill

import (
	"context"
	"log"

	"github.com/Enan01/nbi/api/notion"
)

func SyncToNotion(bill Bill) error {
	req, err := bill.ToCreatePageRequest()
	if err != nil {
		return err
	}

	napi := notion.NewApi(notion.NotionApiHost, notion.Token)
	resp, err := napi.CreatePage(context.Background(), req)
	if err != nil {
		return err
	}

	log.Printf("napi.CreatePage response = %+v", resp)
	return nil
}

func SyncWechatBillToNotion(filePath string) error {
	var billParser BillParser
	billParser = WechatBillParser{BillFilePath: filePath}

	bills, err := billParser.Parse()
	if err != nil {
		return err
	}

	for _, bill := range bills {
		if err := SyncToNotion(bill); err != nil {
			return err
		}
	}

	return nil
}

func SyncAlipayBillToNotion(filePath string) error {
	var billParser BillParser
	billParser = AlipayBillParser{BillFilePath: filePath}

	bills, err := billParser.Parse()
	if err != nil {
		return err
	}

	for _, bill := range bills {
		if err := SyncToNotion(bill); err != nil {
			return err
		}
	}

	return nil
}
