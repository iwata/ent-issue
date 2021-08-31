package main

import (
	"context"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/iwata/ent-issue/ent"
)

var _ mysql.MySQLDriver

func main() {
	client, err := ent.Open("mysql", "root:pass@(localhost:3306)/test")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	created, err := client.User.Create().SetName("hoge").Save(ctx)
	if err != nil {
		log.Fatalf("failed to create a user: %v", err)
	}

	_, err = client.User.Get(ctx, created.ID)
	if err != nil {
		log.Fatalf("failed to get a user: %v", err)
	}
	log.Println("select a user successfully")
}
