package sic

type Database struct {
	Notes   []string  `xml:"notes"`
	LabelID []string  `xml:"label_id"`
	File    [][]File  `xml:"file"`
	Ghost   []Ghost   `xml:"ghost"`
	Label   []Label   `xml:"label"`
	Card    []Card    `xml:"card"`
	Field   [][]Field `xml:"field"`
}

type Ghost struct {
	ID        string `xml:"id,attr"`
	TimeStamp string `xml:"time_stamp,attr"`
}

type Label struct {
	Type      string `xml:"type,attr"`
	TimeStamp string `xml:"time_stamp,attr"`
	ID        string `xml:"id,attr"`
	Name      string `xml:"name,attr"`
}

type Card struct {
	ID          string `xml:"id,attr"`
	Symbol      string `xml:"symbol,attr"`
	Template    string `xml:"template,attr"`
	Type        string `xml:"type,attr"`
	WebsiteIcon string `xml:"website_icon,attr"`
	TimeStamp   string `xml:"time_stamp,attr"`
	Deleted     string `xml:"deleted,attr"`
	Title       string `xml:"title,attr"`
	Color       string `xml:"color,attr"`
	Star        string `xml:"star,attr"`
}

type Field struct {
	Hash    string `xml:"hash,attr"`
	History string `xml:"history,attr"`
	Name    string `xml:"name,attr"`
	Type    string `xml:"type,attr"`
	Text    string `xml:",chardata"`
	Score   string `xml:"score,attr"`
}

type File struct {
	Name string `xml:"name,attr"`
	Text string `xml:",chardata"`
}
