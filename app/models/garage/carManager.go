package garage

import (
	"github.com/sarulabs/di-example/app/models/helpers"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2/bson"
)

// CarManager handles the creation, modification and deletion of cars.
// It uses a CarRepository to communicate with the database.
type CarManager struct {
	Repo   *CarRepository
	Logger *zap.Logger
}

// GetAll returns the list of cars.
func (m *CarManager) GetAll() ([]*Car, error) {
	cars, err := m.Repo.FindAll()

	if cars == nil {
		cars = []*Car{}
	}

	if err != nil {
		m.Logger.Error(err.Error())
	}

	return cars, err
}

// Get returns the car with the given id.
// If the car does not exist an helpers.ErrNotFound is returned.
func (m *CarManager) Get(id string) (*Car, error) {
	car, err := m.Repo.FindByID(id)

	if m.Repo.IsNotFoundErr(err) {
		return nil, helpers.NewErrNotFound("Car `" + id + "` does not exist.")
	}

	if err != nil {
		m.Logger.Error(err.Error())
	}

	return car, err
}

// Create inserts the given car in the database.
// It returns the inserted car.
func (m *CarManager) Create(car *Car) (*Car, error) {
	if err := ValidateCar(car); err != nil {
		return nil, err
	}

	car.ID = bson.NewObjectId().Hex()

	err := m.Repo.Insert(car)

	if m.Repo.IsAlreadyExistErr(err) {
		return m.Create(car)
	}

	if err != nil {
		m.Logger.Error(err.Error())
		return nil, err
	}

	return car, nil
}

// Update updates the car with the given id.
// It uses the values contained in the given car fields.
// It returns the updated car.
func (m *CarManager) Update(id string, car *Car) (*Car, error) {
	if err := ValidateCar(car); err != nil {
		return nil, err
	}

	car.ID = id

	err := m.Repo.Update(car)

	if m.Repo.IsNotFoundErr(err) {
		return nil, helpers.NewErrNotFound("Car `" + id + "` does not exist.")
	}

	if err != nil {
		m.Logger.Error(err.Error())
		return nil, err
	}

	return car, err
}

// Delete removes the car with the given id.
func (m *CarManager) Delete(id string) error {
	err := m.Repo.Delete(id)

	if m.Repo.IsNotFoundErr(err) {
		return nil
	}

	if err != nil {
		m.Logger.Error(err.Error())
	}

	return err
}
