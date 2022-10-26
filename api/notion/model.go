package notion

type (
	PageProperties map[string]map[PropertyType]Property

	CreatePageRequest struct {
		Parent struct {
			DatabaseID string `json:"database_id"`
		} `json:"parent"`
		Properties PageProperties `json:"properties"`
	}

	CreatePageResponse struct {
		Id string `json:"id"`
	}
)
