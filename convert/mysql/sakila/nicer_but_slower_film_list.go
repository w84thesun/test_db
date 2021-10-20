package sakila

//go:generate reform

// NicerButSlowerFilmList represents a row in nicer_but_slower_film_list table.
//reform:nicer_but_slower_film_list
type NicerButSlowerFilmList struct {
	FID         *int16  `reform:"FID"`
	Title       *string `reform:"title"`
	Description *string `reform:"description"`
	Category    string  `reform:"category"`
	Price       *string `reform:"price"`
	Length      *int16  `reform:"length"`
	Rating      []byte  `reform:"rating"` // FIXME unhandled database type "enum"
	Actors      *string `reform:"actors"`
}
