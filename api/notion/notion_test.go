// This package implement notion rest api by go.
package notion

import (
	"context"
	"testing"
)

func Test_api_CreatePage(t *testing.T) {
	type fields struct {
		Host  string
		Token string
	}
	type args struct {
		ctx context.Context
		req CreatePageRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "create a page",
			fields: fields{
				Host:  NotionApiHost,
				Token: "", // please use your notion token
			},
			args: args{
				ctx: context.Background(),
				req: CreatePageRequest{
					Parent: struct {
						DatabaseID string `json:"database_id"`
					}{
						DatabaseID: "", // please use your database id
					},
					Properties: map[string]map[PropertyType]Property{
						"Name": {
							PropTypeTitle: TitleP{
								RichText{
									Type: RichTextTypeText,
									Text: &RichTextText{
										Content: "notion api test",
									},
								},
							},
						},
						"Tags": {
							PropTypeSelect: SelectP{
								Id:    "",
								Name:  "play",
								Color: "yellow",
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := NewApi(tt.fields.Host, tt.fields.Token)
			_, err := ac.CreatePage(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("api.CreatePage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
