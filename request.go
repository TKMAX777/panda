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

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return p.client.Do(req)
}

func (p Handler) get(uri string, query url.Values) (*http.Response, error) {
	req, _ := http.NewRequest("GET", uri, new(bytes.Buffer))
	if query != nil {
		req.URL.RawQuery = query.Encode()
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36")

	return p.client.Do(req)
}
