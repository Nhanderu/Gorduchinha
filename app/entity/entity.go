package entity

import (
	"time"
)

type entity struct {
	ID        uint32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Team struct {
	entity
	Abbr     string
	Name     string
	FullName string
	Trophies []Trophy
}

type Champ struct {
	entity
	Slug string
	Name string
}

type Trophy struct {
	entity
	UUID  string
	Year  int
	Champ Champ
}
