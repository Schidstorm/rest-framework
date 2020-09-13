package rest_framework

import "net/http"

type Processor interface {
	ProcessRequest(http.ResponseWriter, *http.Request)
}
