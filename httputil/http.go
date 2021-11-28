package httputil

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"time"
)

var (
	_ Client   = (*DefaultClient)(nil)
	_ Request  = (*DefaultRequest)(nil)
	_ Response = (*DefaultResponse)(nil)
)

type (
	DefaultClient   http.Client
	DefaultRequest  http.Request
	DefaultResponse http.Response
	Client          interface {
		NewRequest(ctx context.Context, method, url string) (Request, error)
		Do(request Request) (Response, error)
	}
	Request interface {
		GetPath() string
		GetContext() context.Context
		AddHeader(header http.Header)
		AddQuery(values url.Values)
		WithBody(reader io.ReadCloser)
	}
	Response interface {
		GetStatus() int
		GetHeader() http.Header
		GetBody() io.ReadCloser
	}
)

func WrapClient(client http.Client) Client {
	return DefaultClient(client)
}

func NewClient() Client {
	return WrapClient(http.Client{
		Timeout: 10 * time.Second,
	})
}

func OptionalHeader(from Response) http.Header {
	if from == nil {
		return nil
	}

	return from.GetHeader()
}

func (c DefaultClient) NewRequest(ctx context.Context, method, url string) (Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return nil, err
	}

	return (*DefaultRequest)(req), nil
}

func (c DefaultClient) Do(request Request) (Response, error) {
	req := request.(*DefaultRequest)

	res, err := (*http.Client)(&c).Do((*http.Request)(req))
	if err != nil {
		return nil, err
	}

	return (*DefaultResponse)(res), nil
}

func (r *DefaultRequest) GetPath() string {
	return r.URL.Path
}

func (r *DefaultRequest) GetContext() context.Context {
	return (*http.Request)(r).Context()
}

func (r *DefaultRequest) AddQuery(values url.Values) {
	queries := r.URL.Query()
	for k, v := range values {
		queries[k] = append(queries[k], v...)
	}

	r.URL.RawQuery = queries.Encode()
}

func (r *DefaultRequest) AddHeader(header http.Header) {
	for k, v := range header {
		r.Header[k] = append(r.Header[k], v...)
	}
}

func (r *DefaultRequest) WithBody(body io.ReadCloser) {
	r.Body = body
}

func (d *DefaultResponse) GetStatus() int {
	return d.StatusCode
}

func (d *DefaultResponse) GetHeader() http.Header {
	return d.Header
}

func (d *DefaultResponse) GetBody() io.ReadCloser {
	return d.Body
}
