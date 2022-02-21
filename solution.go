package square

import "math"

type Sides int

const (
	SidesCircle   Sides = 0
	SidesTriangle Sides = 3
	SidesSquare   Sides = 4
)

func CalcSquare(sideLen float64, sidesNum Sides) float64 {
	switch sidesNum {
	case SidesCircle:
		return math.Pi * math.Pow(sideLen, 2)
	case SidesTriangle:
		return math.Pow(sideLen, 2) * math.Sqrt(3) / 4
	case SidesSquare:
		return math.Pow(sideLen, 2)
	default:
		return 0
	}
}
