package storage

import (
	"slices"
	"uuid"

	"github.com/nawaphonOHM/warehouse/position"
)

type Storages []*Storage

func (s *Storages) PutProduct(new *Storage) {
	*s = append(*s, new)
}

func (s *Storages) PopProduct(productSearch *Storage) {

	index := slices.IndexFunc(*s, func(product *Storage) bool {
		return product.productId == productSearch.productId
	})

	firstPart := (*s)[:index]
	lastPart := (*s)[index+1:]

	*s = append(firstPart, lastPart...)
}

type Storage struct {
	productId uuid.UUID

	start *position.Position
	end   *position.Position
}

func NewStorage(end *position.Position, start *position.Position, productId uuid.UUID) *Storage {
	return &Storage{end: end, start: start, productId: productId}
}
