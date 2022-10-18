package notion

type PropertyType string

const (
	Date   PropertyType = "date"
	Number PropertyType = "number"
)

type (
	DateP struct {
		Start    string `json:"start,omitempty"`
		End      string `json:"end,omitempty"`
		TimeZone string `json:"time_zone,omitempty"`
	}
)
