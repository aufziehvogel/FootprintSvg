package main

func main() {
	p := packageSip{
		chipPackage: chipPackage{
			pins:       []string{"A", "B", "C", "D"},
			pinSpacing: 15,
		},
	}
	p.draw()
}
