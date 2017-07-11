package handlers

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

func NewRequestBodyMock(value string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(value)))
}

type HandlerResponseMock struct {
	StatusCode   int
	ResponseBody string
}

func (mock *HandlerResponseMock) Header() http.Header {
	return make(http.Header)
}

func (mock *HandlerResponseMock) Write(value []byte) (int, error) {
	mock.ResponseBody = string(value)
	return len(value), nil
}

func (mock *HandlerResponseMock) WriteHeader(statusCode int) {
	mock.StatusCode = statusCode
}
