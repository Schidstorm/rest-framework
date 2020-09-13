package restframework

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type HttpServer struct {
	handler Processor
}

func NewHttpServer(handler Processor) *HttpServer {
	return &HttpServer{handler: handler}
}

func (httpServer *HttpServer) Listen(endpoint string) {
	s := &http.Server{
		Addr:           fmt.Sprintf(endpoint),
		Handler:        httpServer,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	logrus.Fatalln(s.ListenAndServe())
}

func (httpServer *HttpServer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	httpServer.handler.ProcessRequest(rw, r)
}
