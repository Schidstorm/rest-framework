package lib

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

type HttpProcessor struct {
	application *Application
}

func NewHttpProcessor(application *Application) *HttpProcessor {
	return &HttpProcessor{application: application}
}

func (parser HttpProcessor) ProcessRequest(writer http.ResponseWriter, request *http.Request) {
	logrus.Infoln(fmt.Sprintf("Request %s", request.URL.Path))

	response := Response{Success: false}
	modelResponse, err := parser.application.ProcessController(NewRequest(request))
	if err != nil {
		response.ErrorMessage = err.Error()
		logrus.Error(err)
	} else {
		response.Success = true
		response.Data = modelResponse
	}

	buffer, err := json.Marshal(response)
	if err != nil {
		response.Data = nil
		response.Success = false
		response.ErrorMessage = err.Error()
	}

	writer.Header().Set("Content-Type", "application/json")
	if response.Success {
		writer.WriteHeader(200)
	} else {
		writer.WriteHeader(500)
	}
	_, err = writer.Write(buffer)

	if err != nil {
		logrus.Error(err)
	}
}
