package mock

import (
	"io"
	"net/http"
	"strings"
)

type Client struct{
	StatusCodeToReturn int
	BodyTextToReturn string
}

func (c Client) Do(req *http.Request) (*http.Response, error) {
	body := MockReadCloser{
		Reader: strings.NewReader(c.BodyTextToReturn),
		Closer: MockCloser{},
	}
	return &http.Response{
		StatusCode: c.StatusCodeToReturn,
		Body: body,
	}, nil
}

type MockReadCloser struct {
	io.Reader
	io.Closer
}

type MockCloser struct{}

func (m MockCloser) Close() error {
	return nil
}
