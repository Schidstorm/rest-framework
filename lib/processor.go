package lib

import "net/http"

type Processor interface {
	ProcessRequest(http.ResponseWriter, *http.Request)
}
