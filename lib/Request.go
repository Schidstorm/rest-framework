package lib

import (
	"io"
	"net/http"
	"path"
	"strings"
)

type Request struct {
	queryArguments *QueryArguments
	name           string
	body           io.ReadCloser
	method         string
}

func NewRequest(req *http.Request) *Request {
	return &Request{
		queryArguments: NewQueryArguments(req.URL.Query()),
		name:           path.Base(req.URL.Path),
		body:           req.Body,
		method:         strings.ToUpper(req.Method),
	}
}

func (req *Request) GetQueryArguments() *QueryArguments {
	return req.queryArguments
}

func (req *Request) GetBody() io.ReadCloser {
	return req.body
}

func (req *Request) GetName() string {
	return req.name
}

func (req *Request) GetMethod() string {
	return req.method
}

func (req *Request) IsGet() bool {
	return req.name == "GET"
}

func (req *Request) IsPost() bool {
	return req.name == "POST"
}

func (req *Request) IsPut() bool {
	return req.name == "PUT"
}

func (req *Request) IsDelete() bool {
	return req.name == "DELETE"
}
