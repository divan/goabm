package ui

// UI defines the minimal user interface type.
type UI interface {
	Stop()
	Loop()
}

type Charts interface {
	AddChart(name string, values <-chan float64)
}

type Grid interface {
	AddGrid(<-chan [][]interface{})
}

type Grid3D interface {
	Add3DGrid(<-chan [][][]interface{})
}
