// This package implement notion rest api by go.
package notion

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
)

var (
	ErrHttpStatusNotOK = errors.New("http status not 200")
)

const (
	NotionApiHost = "https://api.notion.com"
)

var (
	CreatePageUri = "/v1/pages"
)

var Token = ""
var DebugMode bool = true

type Api interface {
	CreatePage(ctx context.Context, req CreatePageRequest) (resp CreatePageResponse, err error)
}

type api struct {
	Resty *resty.Client
	Host  string
}

func NewApi(host string, token string) Api {
	r := resty.New()
	r.SetAuthToken(token)
	return &api{
		Resty: r,
		Host:  host,
	}
}

func (ac *api) CreatePage(ctx context.Context, req CreatePageRequest) (resp CreatePageResponse, err error) {
	if DebugMode {
		json, _ := json.Marshal(req)
		log.Printf("CreatePage req = %s", json)
	}

	url := ac.Host + CreatePageUri

	res, err := ac.Resty.R().
		SetHeaders(map[string]string{
			"Content-Type":   "application/json",
			"Notion-Version": "2022-06-28",
		}).
		SetBody(req).
		SetResult(&resp).
		Post(url)

	log.Printf("ac.CreatePage response %s", res)

	if err != nil {
		return
	}

	if res.StatusCode() != http.StatusOK {
		err = ErrHttpStatusNotOK
		return
	}

	return
}
