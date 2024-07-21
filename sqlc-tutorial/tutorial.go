package main

import (
	"context"
	"log"
	"reflect"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"tutorial.dev/app/sqlc"
)

func run() error {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://demouser:demouser@localhost:5432/tutorial?sslmode=disable")
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := sqlc.New(conn)

	// list all authors
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		return err
	}
	log.Println(authors)

	// create an author
	insertedAuthor, err := queries.CreateAuthor(ctx, sqlc.CreateAuthorParams{
		//Name: "Brian Kernighan",
		Name: "Joe biden",
		Bio:  pgtype.Text{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	})
	if err != nil {
		return err
	}
	log.Println(insertedAuthor)

	// get the author we just inserted
	fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthor.ID)
	if err != nil {
		return err
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedAuthor, fetchedAuthor))
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
