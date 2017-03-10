package sic

// Database is the root container of the xml
type Database struct {
	Ghosts   []Ghost `xml:"ghost"`
	Labels   []Label `xml:"label"`
	Cards    []Card  `xml:"card"`
	Callback func()  `xml:"-"`
}

// File is contains one file
type File struct {
	Name string `xml:"name,attr"`
	Text string `xml:",chardata"`
}

// Ghost TODO
type Ghost struct {
	ID        int    `xml:"id,attr"`
	Timestamp string `xml:"time_stamp,attr"`
}

// Label is a label template
type Label struct {
	ID        int    `xml:"id,attr"`
	Name      string `xml:"name,attr"`
	Type      string `xml:"type,attr,omitempty"`
	Timestamp string `xml:"time_stamp,attr,omitempty"`
}

// Card contains the information for a item
type Card struct {
	ID          int     `xml:"id,attr"`
	Symbol      string  `xml:"symbol,attr"`
	WebsiteIcon string  `xml:"website_icon,attr,omitempty"`
	Title       string  `xml:"title,attr"`
	Color       string  `xml:"color,attr"`
	Timestamp   string  `xml:"time_stamp,attr,omitempty"`
	Type        string  `xml:"type,attr,omitempty"`
	Deleted     bool    `xml:"deleted,attr,omitempty"`
	Star        string  `xml:"star,attr,omitempty"`
	Template    bool    `xml:"template,attr"`
	Files       []File  `xml:"file"`
	Fields      []Field `xml:"field"`
	Notes       string  `xml:"notes,omitempty"`
	LabelID     []int   `xml:"label_id"`
}

// Field contains a field which is in a card
type Field struct {
	Name    string `xml:"name,attr"`
	Type    string `xml:"type,attr"` // login, password or website
	Value   string `xml:",chardata"`
	Score   string `xml:"score,attr,omitempty"`
	Hash    string `xml:"hash,attr,omitempty"` // MD5 of the value
	History string `xml:"history,attr,omitempty"`
}
