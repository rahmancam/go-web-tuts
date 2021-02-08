package models

// Person type
type Person struct {
	Name string
	Age  int
	ID   int
}

// Store interface
type Store interface {
	Save(n int, p *Person)
	Get(n int) (*Person, error)
	TotalRecords() int
}
