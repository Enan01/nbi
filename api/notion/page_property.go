package notion

type PropertyType string
type RichTextType string
type AnnotationColor string

const (
	PropDate   PropertyType = "date"
	PropNumber PropertyType = "number"
)

const (
	RichTextText     RichTextType = "text"
	RichTextMention  RichTextType = "mention"
	RichTextEquation RichTextType = "equation"
)

const (
	// TODO
	Default string = "default"
)

type (
	DateP struct {
		Start    string `json:"start,omitempty"`
		End      string `json:"end,omitempty"`
		TimeZone string `json:"time_zone,omitempty"`
	}

	SelectP struct {
		Id    string `json:"id,omitempty"`
		Name  string `json:"name,omitempty"`
		Color string `json:"color,omitempty"`
	}

	RichTextP struct {
		PlainText   string `json:"plain_text,omitempty"`
		Herf        string `json:"herf,omitempty"`
		Annotations struct {
			Bold          bool            `json:"bold,omitempty"`
			Italic        bool            `json:"italic,omitempty"`
			Strikethrough bool            `json:"strikethrough,omitempty"`
			Underline     bool            `json:"underline,omitempty"`
			Code          bool            `json:"code,omitempty"`
			Color         AnnotationColor `json:"color,omitempty"`
		} `json:"annotations,omitempty"`
		Type RichTextType `json:"type,omitempty"`
		Text struct {
			Content string `json:"content,omitempty"`
			Link    struct {
				Url string `json:"url,omitempty"`
			} `json:"link,omitempty"`
		} `json:"text,omitempty"`
	}

	NumberP struct {
		Number float64 `json:"number,omitempty"`
	}
)
