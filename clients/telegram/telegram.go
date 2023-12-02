package telegram

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

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

func (c *Clien) SendMessage(chatID int, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)

	_, err := c.doRequest("sendMessage", q)
	if err != nil {
		fmt.Errorf("can't send message: %w", err)
	}
	return nil
}

func (c *Clien) Update(offset int, limit int) ([]Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("offset", strconv.Itoa(limit))

	data, err := c.doRequest("getUpdates", q)
	if err != nil {
		return nil, fmt.Errorf("can't get updates: %w", err)
	}

	var res UpdatesResponse

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, fmt.Errorf("can't unmarshal response: %w", err)
	}
	if !res.Ok {
		return nil, errors.New("not ok")
	}
	return res.Result, nil
}

func (c *Clien) doRequest(method string, query url.Values) ([]byte, error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)

	if err != nil {
		return nil, fmt.Errorf("can't create request: %w", err)
	}

	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("can't do request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("can't read body: %w", err)
	}
	return body, nil
}

func newBasePath(token string) string {
	return "bot" + token
}
