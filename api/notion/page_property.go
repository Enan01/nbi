package notion

type PropertyType string

func ToPropertyType(t string) PropertyType {
	switch PropertyType(t) {
	case PropTypeDate:
		return PropTypeDate
	case PropTypeNumber:
		return PropTypeNumber
	case PropTypeSelect:
		return PropTypeSelect
	case PropTypeTitle:
		return PropTypeTitle
	case PropTypeRichText:
		return PropTypeRichText
	default:
		return PropTypeUnknown
	}
}

type RichTextType string
type Color string

type Property interface {
	GetType() PropertyType
}

const (
	PropTypeUnknown  PropertyType = "unknown"
	PropTypeDate     PropertyType = "date"
	PropTypeNumber   PropertyType = "number"
	PropTypeSelect   PropertyType = "select"
	PropTypeTitle    PropertyType = "title"
	PropTypeRichText PropertyType = "rich_text"
)

const (
	RichTextTypeText     RichTextType = "text"
	RichTextTypeMention  RichTextType = "mention"
	RichTextTypeEquation RichTextType = "equation"
)

const (
	ColorDefault Color = "default"
	ColorGray    Color = "gray"
	ColorBrown   Color = "brown"
	ColorRed     Color = "red"
	ColorOrange  Color = "orange"
	ColorYellow  Color = "yellow"
	ColorGreen   Color = "green"
	ColorBlue    Color = "blue"
	ColorPurple  Color = "purple"
	ColorPink    Color = "pink"
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
		Color Color  `json:"color,omitempty"`
	}

	RichTextP []RichText

	RichTextAnnotations struct {
		Bold          bool  `json:"bold,omitempty"`
		Italic        bool  `json:"italic,omitempty"`
		Strikethrough bool  `json:"strikethrough,omitempty"`
		Underline     bool  `json:"underline,omitempty"`
		Code          bool  `json:"code,omitempty"`
		Color         Color `json:"color,omitempty"`
	}

	RichText struct {
		PlainText   string               `json:"plain_text,omitempty"`
		Herf        string               `json:"herf,omitempty"`
		Annotations *RichTextAnnotations `json:"annotations,omitempty"`
		Type        RichTextType         `json:"type,omitempty"`
		Text        *RichTextText        `json:"text,omitempty"`
	}

	RichTextText struct {
		Content string            `json:"content,omitempty"`
		Link    *RichTextTextLink `json:"link,omitempty"`
	}

	RichTextTextLink struct {
		Url string `json:"url,omitempty"`
	}

	NumberP float64

	TitleP []RichText
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
