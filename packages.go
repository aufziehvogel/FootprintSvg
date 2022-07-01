package main

import (
	"encoding/xml"
	"log"
	"os"
	"strconv"
)

const (
	pinSize      = 20
	packageY     = 30
	packageWidth = 90
	// Reserve a maximum space for labels
	labelWidth = 100
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

func (p *packageDip) height() int {
	return len(p.pins)/2*pinSize + (len(p.pins)/2-1)*p.pinSpacing + 2*p.pinSpacing
}

func (p *packageDip) draw() {
	var drawItems []interface{}

	packageX := labelWidth + pinSize

	r := &svgRect{
		X:           packageX,
		Y:           packageY,
		Width:       packageWidth,
		Height:      p.height(),
		Fill:        "transparent",
		Stroke:      "black",
		StrokeWidth: 2,
	}
	drawItems = append(drawItems, r)

	// TODO: We assume that len(p.pins) is always an even number
	for i := 0; i < len(p.pins)/2; i++ {
		rLeft := &svgRect{
			X:           packageX - pinSize,
			Y:           packageY + p.pinSpacing + i*(pinSize+p.pinSpacing),
			Width:       pinSize,
			Height:      pinSize,
			Fill:        "transparent",
			Stroke:      "black",
			StrokeWidth: 1,
		}
		pinNumberLeft := &svgText{
			Text: strconv.Itoa(i + 1),
			X:    packageX + 5,
			Y:    packageY + (i+1)*(pinSize+p.pinSpacing),
		}
		labelLeft := &svgText{
			Text:       p.pins[i],
			X:          packageX - pinSize - 5,
			Y:          packageY + (i+1)*(pinSize+p.pinSpacing),
			TextAnchor: "end",
		}
		rRight := &svgRect{
			X:           packageX + packageWidth,
			Y:           packageY + p.pinSpacing + i*(pinSize+p.pinSpacing),
			Width:       pinSize,
			Height:      pinSize,
			Fill:        "transparent",
			Stroke:      "black",
			StrokeWidth: 1,
		}
		pinNumberRight := &svgText{
			Text:       strconv.Itoa(len(p.pins) - i),
			X:          packageX + packageWidth - 5,
			Y:          packageY + (i+1)*(pinSize+p.pinSpacing),
			TextAnchor: "end",
		}
		labelRight := &svgText{
			Text: p.pins[len(p.pins)-i-1],
			X:    packageX + packageWidth + pinSize + 5,
			Y:    packageY + (i+1)*(pinSize+p.pinSpacing),
		}

		drawItems = append(
			drawItems,
			rLeft,
			pinNumberLeft,
			labelLeft,
			rRight,
			pinNumberRight,
			labelRight,
		)
	}

	svg := svg{
		Width:      2*labelWidth + 2*pinSize + packageWidth,
		Height:     p.height() + packageY + 1,
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
