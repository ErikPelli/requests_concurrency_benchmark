package client

import (
	"errors"
	"net/http"
	"strconv"
	"time"
)

type client struct {
	client  http.Client
	address string
}

func newClient(host string, port int) *client {
	return &client{
		client: http.Client{
			Timeout: 10 * time.Second,
		},
		address: "http://" + host + ":" + strconv.Itoa(port),
	}
}

func (c *client) EchoRequest() error {
	req, err := http.NewRequest(http.MethodGet, c.address, nil)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check if body isn't empty
	tmp := make([]byte, 1)
	n, err := resp.Body.Read(tmp)
	if err != nil || n != len(tmp) {
		return errors.New("failed to read response body")
	}

	return nil
}
