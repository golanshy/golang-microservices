package shipping_box

type Product struct {
	Name   string
	Length int // Millimeters
	Width  int // Millimeters
	Height int // Millimeters
}

type Box struct {
	Length int // Millimeters
	Width  int // Millimeters
	Height int // Millimeters
}

func getBestBox(availableBoxes []Box, products []Product) Box {

	//TODO: Complete!

	return Box{
		Length: 0,
		Width:  0,
		Height: 0,
	}
}
