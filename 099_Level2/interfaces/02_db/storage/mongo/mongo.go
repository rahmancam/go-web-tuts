package mongo

import (
	"errors"
	"fmt"

	"github.com/rahmancam/go-web-tuts/099_Level2/interfaces/02_db/models"
)

// Mongo Map
type Mongo map[int]*models.Person

// Save records
func (m Mongo) Save(n int, p *models.Person) {
	p.ID = n
	m[n] = p
}

// Get records
func (m Mongo) Get(n int) (p *models.Person, err error) {
	if p, ok := m[n]; ok {
		return p, nil
	}
	return nil, errors.New("Person not found : " + fmt.Sprint(n))
}

// TotalRecords count
func (m Mongo) TotalRecords() int {
	return len(m)
}
