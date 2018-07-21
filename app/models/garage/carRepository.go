package garage

import mgo "gopkg.in/mgo.v2"

// CarRepository contains all the interactions
// with the car collection stored in mongo.
type CarRepository struct {
	Session *mgo.Session
}

// collection returns the car collection.
func (repo *CarRepository) collection() *mgo.Collection {
	return repo.Session.DB("dingo_car_api").C("cars")
}

// FindAll returns all the cars stored in the database.
func (repo *CarRepository) FindAll() ([]*Car, error) {
	var cars []*Car
	err := repo.collection().Find(nil).All(&cars)
	return cars, err
}

// FindByID retrieves the car with the given id from the database.
func (repo *CarRepository) FindByID(id string) (*Car, error) {
	var car *Car
	err := repo.collection().FindId(id).One(&car)
	return car, err
}

// Insert inserts a car in the database.
func (repo *CarRepository) Insert(car *Car) error {
	return repo.collection().Insert(car)
}

// Update updates all the caracteristics of a car.
func (repo *CarRepository) Update(car *Car) error {
	return repo.collection().UpdateId(car.ID, car)
}

// Delete removes the car with the given id.
func (repo *CarRepository) Delete(id string) error {
	return repo.collection().RemoveId(id)
}

// IsNotFoundErr returns true if the error concerns a not found document.
func (repo *CarRepository) IsNotFoundErr(err error) bool {
	return err == mgo.ErrNotFound
}

// IsAlreadyExistErr returns true if the error is related
// to the insertion of an already existing document.
func (repo *CarRepository) IsAlreadyExistErr(err error) bool {
	return mgo.IsDup(err)
}
