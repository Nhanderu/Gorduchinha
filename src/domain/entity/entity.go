package entity

type Champ struct {
	ID   int
	Slug string
	Name string
}

type Team struct {
	ID       int
	Abbr     string
	Name     string
	FullName string
	Trophies []Trophy
}

type Trophy struct {
	ID    int
	Year  int
	Champ Champ
}
