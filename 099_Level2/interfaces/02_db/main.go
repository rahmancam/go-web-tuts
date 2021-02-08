package main

import (
	"fmt"

	"github.com/rahmancam/go-web-tuts/099_Level2/interfaces/02_db/models"
	"github.com/rahmancam/go-web-tuts/099_Level2/interfaces/02_db/storage"
	"github.com/rahmancam/go-web-tuts/099_Level2/interfaces/02_db/storage/mongo"
	"github.com/rahmancam/go-web-tuts/099_Level2/interfaces/02_db/storage/postgres"
)

func main() {
	store := storage.NewStore(mongo.Mongo{})
	store.Save(101, &models.Person{
		Name: "Abdul Rahman",
		Age:  100,
	})
	store.Save(102, &models.Person{
		Name: "Peter Parker",
		Age:  98,
	})
	if p, err := store.Get(101); err == nil {
		fmt.Printf("MongoStore : %d, %s\n", p.ID, p.Name)
	}
	if p, err := store.Get(102); err == nil {
		fmt.Printf("MongoStore : %d, %s\n", p.ID, p.Name)
	}
	fmt.Printf("Total records in MongoStore : %d\n", store.TotalRecords())

	store = storage.NewStore(postgres.Postgres{})
	store.Save(201, &models.Person{
		Name: "Josh",
		Age:  100,
	})
	store.Save(202, &models.Person{
		Name: "John Smith",
		Age:  98,
	})
	store.Save(203, &models.Person{
		Name: "Aruba shaw",
		Age:  99,
	})
	if p, err := store.Get(201); err == nil {
		fmt.Printf("PostgresStore : %d, %s\n", p.ID, p.Name)
	}
	if p, err := store.Get(202); err == nil {
		fmt.Printf("PostgresStore : %d, %s\n", p.ID, p.Name)
	}
	fmt.Printf("Total records in PostgresStore : %d\n", store.TotalRecords())
}
