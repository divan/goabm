package worlds

type XY interface {
	X() int
	Y() int
}

type XYZ interface {
	XY
	Z() int
}
