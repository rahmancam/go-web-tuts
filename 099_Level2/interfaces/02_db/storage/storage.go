package storage

import "github.com/rahmancam/go-web-tuts/099_Level2/interfaces/02_db/models"

// MapStore db
type MapStore struct {
	store models.Store
}

// Save record
func (m MapStore) Save(n int, p *models.Person) {
	m.store.Save(n, p)
}

// Get record
func (m MapStore) Get(n int) (*models.Person, error) {
	return m.store.Get(n)
}

// TotalRecords count
func (m MapStore) TotalRecords() int {
	return m.store.TotalRecords()
}

// NewStore constructor
func NewStore(s models.Store) *MapStore {
	return &MapStore{
		store: s,
	}
}
