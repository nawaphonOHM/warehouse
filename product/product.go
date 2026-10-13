package product

import "uuid"

type Product struct {
	width  float64
	length float64
	height float64

	productId uuid.UUID
}

func (p *Product) Width() float64 {
	return p.width
}

func (p *Product) Length() float64 {
	return p.length
}

func (p *Product) Height() float64 {
	return p.height
}

func (p *Product) ProductId() uuid.UUID {
	return p.productId
}
