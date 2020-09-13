package rest_framework

import (
	"encoding/json"
	"io/ioutil"
)

type Controller struct {
	application *Application
}

func NewController() *Controller {
	return &Controller{application: nil}
}

func (controller *Controller) Get(request *Request, id int64) (interface{}, error) {

	if id < 0 {
		model, err := controller.application.FindModelArray(request.GetName())
		if err != nil {
			return nil, err
		}

		err = controller.application.Db.Find(model).Error
		if err != nil {
			return nil, err
		}
		return model, err
	} else {
		model, err := controller.application.FindModel(request.GetName())
		if err != nil {
			return nil, err
		}

		err = controller.application.Db.First(model, id).Error
		if err != nil {
			return nil, err
		}
		return model, err
	}

}

func (controller *Controller) Post(request *Request) (interface{}, error) {
	model, err := controller.application.FindModel(request.GetName())
	if err != nil {
		return nil, err
	}

	buffer, err := ioutil.ReadAll(request.GetBody())
	if err != nil {
		return nil, err
	}
	_ = request.GetBody().Close()

	err = json.Unmarshal(buffer, model)
	if err != nil {
		return nil, err
	}

	err = controller.application.Db.Create(model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (controller *Controller) Put(request *Request, id int64) (interface{}, error) {
	model, err := controller.application.FindModel(request.GetName())
	if err != nil {
		return nil, err
	}
	controller.application.Db.First(model, id)

	reader := request.GetBody()
	buffer, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	_ = reader.Close()

	err = json.Unmarshal(buffer, model)
	if err != nil {
		return nil, err
	}

	err = controller.application.Db.Updates(model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (controller *Controller) Delete(request *Request, id int64) error {
	model, err := controller.application.FindModel(request.GetName())

	if err != nil {
		return err
	}

	err = controller.application.Db.Delete(model, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (controller *Controller) SetApplication(application *Application) {
	controller.application = application
}
