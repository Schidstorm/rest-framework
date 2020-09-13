package rest_framework

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
	"reflect"
)

type Application struct {
	Db              *gorm.DB
	modelTypes      []reflect.Type
	controllerTypes map[string]reflect.Type
	httpServer      *HttpServer
	httpProcessor   *HttpProcessor
	modelRegistry   *TypeRegistry
}

func LoadApplication(dialector gorm.Dialector) *Application {
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		logrus.Errorln("Database connection failed", err)
		os.Exit(1)
	}

	application := &Application{
		Db:              db,
		modelTypes:      []reflect.Type{},
		controllerTypes: map[string]reflect.Type{},
		modelRegistry:   NewTypeRegistry(),
	}

	application.httpProcessor = NewHttpProcessor(application)
	application.httpServer = &HttpServer{handler: application.httpProcessor}
	return application
}

func (app *Application) ProcessController(request *Request) (interface{}, error) {
	controller := NewController()
	controller.SetApplication(app)

	switch request.GetMethod() {
	case "GET":
		id, err := request.GetQueryArguments().GetInt64("id", 10)
		if err != nil {
			return controller.Get(request, -1)
		}
		return controller.Get(request, id)
	case "POST":
		return controller.Post(request)
	case "PUT":
		id, err := request.GetQueryArguments().GetInt64("id", 10)
		if err != nil {
			return nil, err
		}
		return controller.Put(request, id)
	case "DELETE":
		id, err := request.GetQueryArguments().GetInt64("id", 10)
		if err != nil {
			return nil, err
		}
		err = controller.Delete(request, id)
		return nil, err
	default:
		return nil, errors.New("method not implemented")
	}
}

func (app *Application) FindModel(name string) (interface{}, error) {
	if !app.modelRegistry.Has(name) {
		return nil, errors.New(fmt.Sprintf("Model %s not found", name))
	}

	return reflect.New(app.modelRegistry.Get(name)).Interface(), nil
}

func (app *Application) FindModelArray(name string) (interface{}, error) {
	if !app.modelRegistry.Has(name) {
		return nil, errors.New(fmt.Sprintf("Model %s not found", name))
	}

	return reflect.New(reflect.SliceOf(app.modelRegistry.Get(name))).Interface(), nil
}

func (app *Application) RegisterModel(model interface{}) {
	app.modelRegistry.Register(reflect.TypeOf(model))
	err := app.Db.AutoMigrate(model)

	if err != nil {
		logrus.Errorln(fmt.Sprintf("Auto Migration failed for model %s", reflect.TypeOf(model).Name()), err)
		os.Exit(1)
	}
}

func (app *Application) Start(httpEndpoint string) {
	app.httpServer.Listen(httpEndpoint)
}
