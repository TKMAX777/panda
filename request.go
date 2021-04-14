package panda

import (
	"bytes"
	"net/http"
	"net/url"
	"strings"
)

func (p Handler) post(uri string, query, values url.Values) (*http.Response, error) {
	req, _ := http.NewRequest("POST", uri, strings.NewReader(values.Encode()))
	if query != nil {
		req.URL.RawQuery = query.Encode()
	}

	req.Header.Add("User-Agent", UserAgent)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return p.client.Do(req)
}

func (p Handler) get(uri string, query url.Values) (*http.Response, error) {
	req, _ := http.NewRequest("GET", uri, new(bytes.Buffer))
	if query != nil {
		req.URL.RawQuery = query.Encode()
	}

	req.Header.Add("User-Agent", UserAgent)

	return p.client.Do(req)
}
