package entity

type Team struct {
	ID       int
	Abbr     string
	Name     string
	FullName string
	Trophies []Trophy
}

type Champ struct {
	ID   int
	Slug string
	Name string
}

type Trophy struct {
	ID    int
	Year  int
	Champ Champ
}
