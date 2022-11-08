package bill

import (
	"testing"
	"time"

	"github.com/Enan01/nbi/api/notion"
)

func TestSyncToNotion(t *testing.T) {
	type args struct {
		bill Bill
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"case 1",
			args{
				bill: Bill{
					DealTime:    time.Now(),
					DealType:    "微信红包",
					DealPartner: "红包",
					Product:     "红包",
					Inout:       BillIn,
					Amount:      56.8,
					Source:      BillSourceWechat,
					Catecory:    "",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		BillDatabaseId = "" // please use your database id
		notion.Token = ""   // please use your notion token
		t.Run(tt.name, func(t *testing.T) {
			if err := SyncToNotion(tt.args.bill); (err != nil) != tt.wantErr {
				t.Errorf("SyncToNotion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSyncWechatBillToNotion(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"case 1",
			args{filePath: ""},
			false,
		},
	}
	for _, tt := range tests {
		BillDatabaseId = "" // please use your database id
		notion.Token = ""   // please use your notion token
		t.Run(tt.name, func(t *testing.T) {
			if err := SyncWechatBillToNotion(tt.args.filePath); (err != nil) != tt.wantErr {
				t.Errorf("SyncWechatBillToNotion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSyncAlipayBillToNotion(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"case 1",
			args{filePath: ""},
			false,
		},
	}
	for _, tt := range tests {
		BillDatabaseId = "" // please use your database id
		notion.Token = ""   // please use your notion token
		t.Run(tt.name, func(t *testing.T) {
			if err := SyncAlipayBillToNotion(tt.args.filePath); (err != nil) != tt.wantErr {
				t.Errorf("SyncWechatBillToNotion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
