package main

func main() {

	shapesList := []Shape{
		Rectangle{Length: 5, Width: 3},
		Circle{Radius: 4},
		Square{Length: 6},
		Triangle{SideA: 3, SideB: 4, SideC: 5},
	}

	for _, shape := range shapesList {
		PrintShapeDetails(shape)
	}
}
