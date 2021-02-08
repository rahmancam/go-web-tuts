package postgres

import (
	"errors"
	"fmt"

	"github.com/rahmancam/go-web-tuts/099_Level2/interfaces/02_db/models"
)

// Postgres Map
type Postgres map[int]*models.Person

// Save count
func (pg Postgres) Save(n int, p *models.Person) {
	p.ID = n
	pg[n] = p
}

// Get count
func (pg Postgres) Get(n int) (p *models.Person, err error) {
	if p, ok := pg[n]; ok {
		return p, nil
	}
	return nil, errors.New("Person not found : " + fmt.Sprint(n))
}

// TotalRecords count
func (pg Postgres) TotalRecords() int {
	return len(pg)
}
