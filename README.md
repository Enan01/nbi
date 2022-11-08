# notion_bill

nbi is a tool to import bill into notion database.

bill support list: 

- [x] alipay
- [x] wechat

> Official notion api link: https://developers.notion.com/reference/intro

## install

```
$ go install github.com/Enan01/nbi@latest
```

## how to use

```
$ nbi -h
NAME:
   nbi - import your alipay and wechat bill into notion

USAGE:
   nbi [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --billFile value, --bf value            bill file path
   --billType value, --bt value            bill type, alipay is 1, wechat is 2 (default: 1)
   --notionDatabaseId value, --ndid value  notion database id
   --notionToken value, --nt value         notion token
   --help, -h                              show help (default: false)
```