package main

import (
	"fmt"
	"second_simple_server/pkg/repository"

	"log"
)
const connStr = "postgres://mylt1c:Zexecmdirjkt8@localhost:5432/pushka"

func main() {
	db, err := repository.New(connStr)
	if err != nil {
		log.Fatal(err.Error())
	}
	item, err := db.GetBooksByID(7)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("ID: %d, Name: %s, AuthorID: %d, GenreID: %d, Price: %d\n", item.ID, item.Name, item.AuthorID, item.GenreID, item.Price)
}

