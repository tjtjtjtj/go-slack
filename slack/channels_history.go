package slack

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

type Message struct {
	Type string `json:"type"`
	Ts   string `json:ts"`
	User string `json:"user"`
}

type History struct {
	Ok       bool   `json:"ok"`
	Latest   string `json:"latest"`
	Messages []struct {
		Message
	} `json:"messages"`
}

func (c *Client) GetChannlesHistory(ctx context.Context, channel string, count string) (*History, error) {
	spath := fmt.Sprintf("/channels.history")
	values := url.Values{}
	values.Add("channel", channel)
	values.Add("count", count)
	req, err := c.newRequest(ctx, "post", spath, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, errors.Errorf("history NotFound", spath)
	}
	fmt.Printf("status:%s", res.Status)

	var history History
	if err := decodeBody(res, &history); err != nil {
		return nil, err
	}

	return &history, nil
}
