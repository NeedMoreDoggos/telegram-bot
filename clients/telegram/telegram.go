package telegram

import "net/http"

type Clien struct {
	host     string
	basePath string
	client   http.Client
}

func New(host string, token string) Clien {
	return Clien{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

func newBasePath(token string) string {
	return "bot" + token
}

func (c *Clien) Update() {

}

func (c *Clien) SendMessage() {

}
