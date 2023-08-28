package shapes

func Perimeter(rectange Rectangle) float64 {
	return 2 * (rectange.Width + rectange.Height)
}

func Area(rectange Rectangle) float64 {
	return (rectange.Width * rectange.Height)
}
