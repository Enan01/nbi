package main

import (
	"fmt"

	cli "github.com/urfave/cli/v2"
)

func main() {
	fmt.Println("hello notion_bill")

	app := &cli.App{
		Name:   "notion_bill",
		Usage:  "import your alipay and wechat bill into notion",
		Action: func(*cli.Context) error { panic("not implemented") },
	}
}
