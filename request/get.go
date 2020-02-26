package request

import (
	"net/http"
)

// Get struct
type Get struct{}

// DefaultGet DefaultGet
var DefaultGet = NewGet()

// NewGet init
func NewGet() *Get {
	return &Get{}
}

//Request get 请求
func (g *Get) Request(uri string) (*http.Response, error) {
	response, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	return response, nil
}
