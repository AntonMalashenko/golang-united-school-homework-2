package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

const SidesCircle int = 0
const SidesTriangle int = 3
const SidesSquare int = 4

func CalcSquare(sideLen float64, sidesNum int) float64 {
	switch sidesNum {
	case SidesCircle:
		return (sideLen * sideLen * math.Pi) / 2
	case SidesSquare:
		return sideLen * sideLen
	case SidesTriangle:
		return sideLen * math.Sqrt(3) / 4
	default:
		return 0
	}
}
