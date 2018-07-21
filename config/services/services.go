package services

import (
	"os"
	"time"

	"github.com/sarulabs/di"
	"github.com/sarulabs/di-example/app/models/garage"
	"github.com/sarulabs/di-example/config/logging"
	"go.uber.org/zap"
	mgo "gopkg.in/mgo.v2"
)

// Services contains the definitions of the application services.
var Services = []di.Def{
	{
		Name:  "logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return logging.Logger, nil
		},
	},
	{
		Name:  "mongo-pool",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return mgo.DialWithTimeout(os.Getenv("MONGO_URL"), 5*time.Second)
		},
		Close: func(obj interface{}) error {
			obj.(*mgo.Session).Close()
			return nil
		},
	},
	{
		Name:  "mongo",
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			return ctn.Get("mongo-pool").(*mgo.Session).Copy(), nil
		},
		Close: func(obj interface{}) error {
			obj.(*mgo.Session).Close()
			return nil
		},
	},
	{
		Name:  "car-repository",
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			return &garage.CarRepository{
				Session: ctn.Get("mongo").(*mgo.Session),
			}, nil
		},
	},
	{
		Name:  "car-manager",
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			return &garage.CarManager{
				Repo:   ctn.Get("car-repository").(*garage.CarRepository),
				Logger: ctn.Get("logger").(*zap.Logger),
			}, nil
		},
	},
}
