package commands

// Enum of all commands
const (
	Document      = "_doc"
	Strong        = "bf"
	Italic        = "it"
	InlineCode    = "`"
	InlineMath    = "$"
	Link          = "href"
	Reference     = "ref"
	Dereference   = "deref"
	Paragraph     = "para"
	Quote         = "quote"
	Math          = "$$"
	Code          = "```"
	Figure        = "figure"
	Box           = "box"
	Title         = "title"
	Subtitle      = "subtitle"
	Heading       = "heading"
	OrderedList   = "ol"
	UnorderedList = "ul"
	ListItem      = "li"
)

// GetDisplayName Get the display name of the command
func GetDisplayName(command string) string {
	return map[string]string{
		Document:      "Document",
		Strong:        "Strong",
		Italic:        "Italic",
		InlineCode:    "Inline Code",
		InlineMath:    "Inline Math",
		Link:          "Link",
		Reference:     "Reference",
		Dereference:   "Dereference",
		Paragraph:     "Paragraph",
		Quote:         "Quote",
		Math:          "Math",
		Code:          "Code",
		Figure:        "Figure",
		Box:           "Box",
		Title:         "Title",
		Subtitle:      "Subtitle",
		Heading:       "Heading",
		OrderedList:   "Ordered List",
		UnorderedList: "Unordered List",
		ListItem:      "ListItem",
	}[command]
}
