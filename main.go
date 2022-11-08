package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Enan01/nbi/api/notion"
	"github.com/Enan01/nbi/domain/bill"
	cli "github.com/urfave/cli/v2"
)

func main() {
	var (
		billFile         string
		billType         int
		notionDatabaseId string
		notionToken      string
	)

	app := &cli.App{
		Name:  "nbi",
		Usage: "import your alipay and wechat bill into notion",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "billFile",
				Aliases:     []string{"bf"},
				Value:       "",
				Usage:       "bill file path",
				Destination: &billFile,
			},
			&cli.IntFlag{
				Name:        "billType",
				Aliases:     []string{"bt"},
				Value:       1,
				Usage:       "bill type, alipay is 1, wechat is 2",
				Destination: &billType,
			},
			&cli.StringFlag{
				Name:        "notionDatabaseId",
				Aliases:     []string{"ndid"},
				Value:       "",
				Usage:       "notion database id",
				Destination: &notionDatabaseId,
			},
			&cli.StringFlag{
				Name:        "notionToken",
				Aliases:     []string{"nt"},
				Value:       "",
				Usage:       "notion token",
				Destination: &notionToken,
			},
		},
		Action: func(ctx *cli.Context) error {
			notion.Token = notionToken
			bill.BillDatabaseId = notionDatabaseId

			if bill.BillType(billType) == bill.BillTypeAlipay {
				bill.AlipayBillStartLineNo = 5
				if err := bill.SyncAlipayBillToNotion(billFile); err != nil {
					return cli.NewExitError(fmt.Sprintf("cause error %s", err), 1)
				}
			} else if bill.BillType(billType) == bill.BillTypeWechat {
				bill.WechatBillStartLineNo = 17
				if err := bill.SyncWechatBillToNotion(billFile); err != nil {
					return cli.NewExitError(fmt.Sprintf("cause error %s", err), 1)
				}
			} else {
				return cli.NewExitError("BillType error, please check the billType param", 1)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
