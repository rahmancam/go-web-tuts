package storage_test

import (
	"testing"

	"github.com/rahmancam/go-web-tuts/099_Level2/interfaces/02_db/models"
	"github.com/rahmancam/go-web-tuts/099_Level2/interfaces/02_db/storage"
	"github.com/rahmancam/go-web-tuts/099_Level2/interfaces/02_db/storage/mongo"
)

func TestGet(t *testing.T) {
	store := storage.NewStore(mongo.Mongo{})
	p1 := &models.Person{
		Name: "Abdul",
	}
	store.Save(10001, p1)
	got, _ := store.Get(10001)
	if got.ID != 10001 {
		t.Errorf("ID mismatch in mongo store")
	}

	if p1 != got {
		t.Fatalf("Want %v, got %v", p1, got)
	}
}
