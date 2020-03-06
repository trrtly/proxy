package request

import (
	"net/http"
	"net/url"
	"time"
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

//RequestByProxy get 请求
func (g *Get) RequestByProxy(uri, ipproxy string, timeout time.Duration) (*http.Response, error) {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	transport := &http.Transport{
		Proxy: func(_ *http.Request) (*url.URL, error) {
			return url.Parse("http://" + ipproxy)
		},
	}
	response, err := (&http.Client{Transport: transport, Timeout: timeout * time.Second}).Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return response, nil
}
