package main

type svg struct {
	XMLName    struct{} `xml:"svg"`
	Width      int      `xml:"width,attr"`
	Height     int      `xml:"height,attr"`
	XMLNS      string   `xml:"xmlns,attr"`
	XMLNSXLink string   `xml:"xmlns:xlink,attr"`
	Items      []interface{}
}

type svgRect struct {
	XMLName     struct{} `xml:"rect"`
	X           int      `xml:"x,attr"`
	Y           int      `xml:"y,attr"`
	Width       int      `xml:"width,attr"`
	Height      int      `xml:"height,attr"`
	Fill        string   `xml:"fill,attr"`
	Stroke      string   `xml:"stroke,attr"`
	StrokeWidth int      `xml:"stroke-width,attr"`
}

type svgText struct {
	XMLName    struct{} `xml:"text"`
	Text       string   `xml:",chardata"`
	X          int      `xml:"x,attr"`
	Y          int      `xml:"y,attr"`
	TextAnchor string   `xml:"text-anchor,attr"`
}
