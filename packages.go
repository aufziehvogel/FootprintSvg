package main

import (
	"encoding/xml"
	"log"
	"os"
)

const (
	pinSize  = 20
	packageX = 30
	packageY = 30
)

type chipPackage struct {
	pins       []string
	pinSpacing int
}

type packageSip struct {
	chipPackage
}

type packageDip struct {
	chipPackage
}

type packageQfp struct {
	chipPackage
}

func (p *packageSip) height() int {
	return len(p.pins)*pinSize + (len(p.pins)-1)*p.pinSpacing + 2*p.pinSpacing
}

func (p *packageSip) draw() {
	var drawItems []interface{}

	r := &svgRect{
		X:           packageX,
		Y:           packageY,
		Width:       90,
		Height:      p.height(),
		Fill:        "transparent",
		Stroke:      "black",
		StrokeWidth: 2,
	}
	drawItems = append(drawItems, r)

	for i, _ := range p.pins {
		r := &svgRect{
			X:           packageX - pinSize,
			Y:           packageY + p.pinSpacing + i*(pinSize+p.pinSpacing),
			Width:       pinSize,
			Height:      pinSize,
			Fill:        "transparent",
			Stroke:      "black",
			StrokeWidth: 1,
		}
		drawItems = append(drawItems, r)
	}

	svg := svg{
		Width:      300,
		Height:     p.height() + 50,
		XMLNS:      "http://www.w3.org/2000/svg",
		XMLNSXLink: "http://www.w3.org/1999/xlink",
		Items:      drawItems,
	}

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(svg); err != nil {
		log.Panic("could not marshal")
	}
}
