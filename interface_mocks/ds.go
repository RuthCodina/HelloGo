package interface_mocks

import (
	"fmt"
	"log"
)

// Datastore defines and interface for storing retrievable data
// Any type that implements this interface (has these two methods) is also of type ´datastore´.
// This means any value of Type `MockDatastore` is also of type `Datastore`
// This means we could have a value of type `*sql.DB` and it can also be of type `Datastore`
// This means we can write functions to take Type `Datastore` and use either of these:
// -- `MockDatastore`
// -- `*sql.DB`
type Datastore interface {
	GetUser(id int) (User, error)
	SaveUser(u User) error
}

// service represents a service that stores retrievable data
// we will attach methods to `Service` so that cna use either of these
// --- `MockDatastore`
// --- `*sql.DB`

type Service struct {
	ds Datastore
}

func (s Service) GetUser(id int) (User, error) {
	return s.ds.GetUser(id)
}
func (s Service) SaveUser(u User) error {
	return s.ds.SaveUser(u)
}

func DB() {
	db := MockDatastore{
		Users: map[int]User{},
	}
	srvc := Service{
		ds: db,
	}

	u1 := User{
		ID:    1,
		First: "James",
	}

	err := srvc.SaveUser(u1)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	u1Returned, err := srvc.GetUser(u1.ID)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	fmt.Println(u1)
	fmt.Println(u1Returned)
}
