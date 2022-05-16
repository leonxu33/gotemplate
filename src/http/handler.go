package http

import (
	"encoding/json"
	"io"
	"net/http"
)

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (hdl *TestHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		hdl.handleGet(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		hdl.handlePost(rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (hdl *TestHandler) handleGet(rw http.ResponseWriter, r *http.Request) {

}

func (hdl *TestHandler) handlePost(rw http.ResponseWriter, r *http.Request) {

}

type TestRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type TestResponse struct {
	ErrCode int    `json:"errCode"`
	ErrMsg  string `json:"errMesg"`
}

func (p *TestRequest) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

func (p *TestResponse) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	return encoder.Encode(p)
}
