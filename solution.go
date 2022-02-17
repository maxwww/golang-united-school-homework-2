package square

import (
	"math"
)

type Side int

const (
	SidesTriangle Side = 3
	SidesSquare   Side = 4
	SidesCircle   Side = 0
)

func CalcSquare(sideLen float64, sidesNum Side) float64 {
	switch sidesNum {
	case SidesSquare:
		return math.Pow(sideLen, 2)
	case SidesTriangle:
		return math.Pow(sideLen, 2) * math.Sqrt(3) / 4
	case SidesCircle:
		return math.Pi * math.Pow(sideLen, 2)
	}

	return 0
}
