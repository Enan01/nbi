package notion

type PropertyType string
type RichTextType string
type Color string

type Property interface {
	GetType() PropertyType
}

const (
	PropDate     PropertyType = "date"
	PropNumber   PropertyType = "number"
	PropSelect   PropertyType = "select"
	PropTitle    PropertyType = "title"
	PropRichText PropertyType = "rich_text"
)

const (
	RichTextTypeText     RichTextType = "text"
	RichTextTypeMention  RichTextType = "mention"
	RichTextTypeEquation RichTextType = "equation"
)

const (
	// TODO
	Default Color = "default"
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
		PlainText   string               `json:"plain_text,omitempty"`
		Herf        string               `json:"herf,omitempty"`
		Annotations *RichTextAnnotations `json:"annotations,omitempty"`
		Type        RichTextType         `json:"type,omitempty"`
		Text        *RichTextText        `json:"text,omitempty"`
	}

	RichTextAnnotations struct {
		Bold          bool  `json:"bold,omitempty"`
		Italic        bool  `json:"italic,omitempty"`
		Strikethrough bool  `json:"strikethrough,omitempty"`
		Underline     bool  `json:"underline,omitempty"`
		Code          bool  `json:"code,omitempty"`
		Color         Color `json:"color,omitempty"`
	}

	RichTextText struct {
		Content string            `json:"content,omitempty"`
		Link    *RichTextTextLink `json:"link,omitempty"`
	}

	RichTextTextLink struct {
		Url string `json:"url,omitempty"`
	}

	NumberP struct {
		Number float64 `json:"number,omitempty"`
	}

	TitleP []RichTextP
)

func (p DateP) GetType() PropertyType {
	return PropertyType("date")
}
func (p SelectP) GetType() PropertyType {
	return PropertyType("select")
}
func (p RichTextP) GetType() PropertyType {
	return PropertyType("rich_text")
}
func (p NumberP) GetType() PropertyType {
	return PropertyType("number")
}
func (p TitleP) GetType() PropertyType {
	return PropertyType("title")
}
