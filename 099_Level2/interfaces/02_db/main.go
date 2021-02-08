package main

import (
	"errors"
	"fmt"
)

// Person type
type Person struct {
	Name string
	age  int
	ID   int
}

// Mongo Map
type Mongo map[int]*Person

// Postgres Map
type Postgres map[int]*Person

func (m Mongo) save(n int, p *Person) {
	p.ID = n
	m[n] = p
}

func (m Mongo) get(n int) (p *Person, err error) {
	if p, ok := m[n]; ok {
		return p, nil
	}
	return nil, errors.New("Person not found : " + fmt.Sprint(n))
}

func (m Mongo) totalRecords() int {
	return len(m)
}

func (pg Postgres) save(n int, p *Person) {
	p.ID = n
	pg[n] = p
}

func (pg Postgres) get(n int) (p *Person, err error) {
	if p, ok := pg[n]; ok {
		return p, nil
	}
	return nil, errors.New("Person not found : " + fmt.Sprint(n))
}

func (pg Postgres) totalRecords() int {
	return len(pg)
}

// Store interface
type Store interface {
	save(n int, p *Person)
	get(n int) (*Person, error)
	totalRecords() int
}

// MapStore db
type MapStore struct {
	store Store
}

func (m MapStore) save(n int, p *Person) {
	m.store.save(n, p)
}

func (m MapStore) get(n int) (*Person, error) {
	return m.store.get(n)
}

func (m MapStore) totalRecords() int {
	return m.store.totalRecords()
}

// NewStore constructor
func NewStore(s Store) *MapStore {
	return &MapStore{
		store: s,
	}
}

func main() {
	store := NewStore(make(Mongo))
	store.save(101, &Person{
		Name: "Abdul Rahman",
		age:  100,
	})
	store.save(102, &Person{
		Name: "Peter Parker",
		age:  98,
	})
	if p, err := store.get(101); err == nil {
		fmt.Printf("MongoStore : %d, %s\n", p.ID, p.Name)
	}
	if p, err := store.get(102); err == nil {
		fmt.Printf("MongoStore : %d, %s\n", p.ID, p.Name)
	}
	fmt.Printf("Total records in MongoStore : %d\n", store.totalRecords())

	store = NewStore(make(Postgres))
	store.save(201, &Person{
		Name: "Josh",
		age:  100,
	})
	store.save(202, &Person{
		Name: "John Smith",
		age:  98,
	})
	store.save(203, &Person{
		Name: "Aruba shaw",
		age:  99,
	})
	if p, err := store.get(201); err == nil {
		fmt.Printf("PostgresStore : %d, %s\n", p.ID, p.Name)
	}
	if p, err := store.get(202); err == nil {
		fmt.Printf("PostgresStore : %d, %s\n", p.ID, p.Name)
	}
	fmt.Printf("Total records in PostgresStore : %d\n", store.totalRecords())
}
