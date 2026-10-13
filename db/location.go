package db

import "github.com/nawaphonOHM/warehouse/position"

type Location struct {
	start *position.Position
	end   *position.Position
}

func (l Location) Start() *position.Position {
	return l.start
}

func (l Location) End() *position.Position {
	return l.end
}

func NewLocation(start *position.Position, end *position.Position) *Location {
	return &Location{start: start, end: end}
}
