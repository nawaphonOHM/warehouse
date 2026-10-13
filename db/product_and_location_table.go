package db

import (
	"fmt"
	"math/rand"
	"uuid"

	"github.com/nawaphonOHM/warehouse/position"
)

type ProductAndLocationTable []*ProductAndLocationRow

func (p *ProductAndLocationTable) CanFindConsecutiveSpace(width float64, length float64, heigh float64) (yes bool) {

	fmt.Printf("width: %f, length: %f, heigh: %f", width, length, heigh)

	return rand.Int31n(10) > 5
}

type ProductAndLocationRow struct {
	productId uuid.UUID

	location *Location
}

func (p *ProductAndLocationRow) ProductId() uuid.UUID {
	return p.productId
}

func (p *ProductAndLocationRow) GetLocation() *Location {
	return p.location
}

func NewProductAndLocationRow(productId uuid.UUID, start *position.Position, end *position.Position) *ProductAndLocationRow {
	return &ProductAndLocationRow{productId: productId, location: NewLocation(start, end)}
}
