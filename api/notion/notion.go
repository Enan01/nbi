// This package implement notion rest api by go.
package notion

import (
	"log"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	NotionApiHost = "https://api.notion.com"

	CreatePageUri = "/pages"

	DatabaseId = ""
	Token      = ""
)

type NotionApiClient struct {
	Resty *resty.Client
	Host  string
}

func NewNotionApiClient(host string, token string) NotionApiClient {
	r := resty.New()
	r.SetAuthToken(token)
	return NotionApiClient{
		Resty: r,
		Host:  host,
	}
}

type (
	CreatePageRequest struct {
		Parent struct {
			DatabaseID string `json:"database_id"`
		} `json:"parent"`
		Properties struct {
			Type struct {
				Select struct {
					ID    string `json:"id"`
					Name  string `json:"name"`
					Color string `json:"color"`
				} `json:"select"`
			} `json:"Type"`
			Score5 struct {
				Select struct {
					ID    string `json:"id"`
					Name  string `json:"name"`
					Color string `json:"color"`
				} `json:"select"`
			} `json:"Score /5"`
			Name struct {
				Title []struct {
					Text struct {
						Content string `json:"content"`
					} `json:"text"`
				} `json:"title"`
			} `json:"Name"`
			Status struct {
				Select struct {
					ID    string `json:"id"`
					Name  string `json:"name"`
					Color string `json:"color"`
				} `json:"select"`
			} `json:"Status"`
			Publisher struct {
				Select struct {
					ID    string `json:"id"`
					Name  string `json:"name"`
					Color string `json:"color"`
				} `json:"select"`
			} `json:"Publisher"`
			PublishingReleaseDate struct {
				Date struct {
					Start time.Time   `json:"start"`
					End   interface{} `json:"end"`
				} `json:"date"`
			} `json:"Publishing/Release Date"`
			Link struct {
				URL string `json:"url"`
			} `json:"Link"`
			Summary struct {
				RichText []struct {
					Type string `json:"type"`
					Text struct {
						Content string      `json:"content"`
						Link    interface{} `json:"link"`
					} `json:"text"`
					Annotations struct {
						Bold          bool   `json:"bold"`
						Italic        bool   `json:"italic"`
						Strikethrough bool   `json:"strikethrough"`
						Underline     bool   `json:"underline"`
						Code          bool   `json:"code"`
						Color         string `json:"color"`
					} `json:"annotations"`
					PlainText string      `json:"plain_text"`
					Href      interface{} `json:"href"`
				} `json:"rich_text"`
			} `json:"Summary"`
			Read struct {
				Checkbox bool `json:"checkbox"`
			} `json:"Read"`
		} `json:"properties"`
	}

	CreatePageResponse struct {
	}
)

func (ac NotionApiClient) CreatePage(req CreatePageRequest) (resp CreatePageResponse, err error) {
	url := ac.Host + CreatePageUri

	res, err := ac.Resty.R().
		SetHeaders(map[string]string{
			"Content-Type":   "application/json",
			"Notion-Version": "2021-08-16",
		}).
		SetBody(req).
		SetResult(&resp).
		Post(url)

	log.Printf("CreatePage response %s", res)

	if err != nil {
		return
	}

	return
}
