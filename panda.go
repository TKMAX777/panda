package panda

import (
	"net/http"
	"net/http/cookiejar"
)

const BaseURI = "https://panda.ecs.kyoto-u.ac.jp"
const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36"

type Handler struct {
	client *http.Client
	jar    *cookiejar.Jar
	auth   struct {
		ID   string
		Pass string
	}
	authOK bool
}

type unixTime struct {
	EpochSecond int64 `json:"epochSecond"`
	Nano        int64 `json:"nano"`
}

type Data struct {
	EntityPrefix string `json:"entityPrefix"`

	AssignmentCollection []Assignment `json:"assignment_collection"`
	ContentCollection    []Content    `json:"content_collection"`
}

// NewClient make new PandA Handler
func NewClient() *Handler {
	jar, _ := cookiejar.New(nil)

	var p = &Handler{
		client: &http.Client{Jar: jar},
		jar:    jar,
		authOK: false,
	}

	return p
}
