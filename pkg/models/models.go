package models

type Book struct {
	ID     int   
	Name   string
	Price  int    
	AuthorID int
	GenreID  int
}

type Author struct {
	ID   int
	Name string
}

type Genre struct {
	ID   int
	Ganre string
}