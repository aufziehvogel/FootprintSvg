package main

func main() {
	p := packageDip{
		chipPackage: chipPackage{
			pins: []string{
				"RST",
				"P3",
				"P4",
				"GND",
				"P0",
				"P1",
				"P2",
				"VCC",
			},
			pinSpacing: 15,
		},
	}
	p.draw()
}
