package position

import "math/rand/v2"

type Position struct {
	x float64
	y float64
	z float64
}

func (p *Position) X() float64 {
	return p.x
}

func (p *Position) Y() float64 {
	return p.y
}

func (p *Position) Z() float64 {
	return p.z
}

func FindBestLocation(width float64, length float64, heigh float64) (*Position, *Position) {

	xStart := rand.Float64() * (1000 - width)
	xEnd := xStart + width

	yStart := rand.Float64() * (1000 - length)
	yEnd := yStart + length

	zStart := rand.Float64() * (100 - heigh)
	zEnd := zStart + heigh

	return &Position{xStart, yStart, zStart}, &Position{xEnd, yEnd, zEnd}
}
