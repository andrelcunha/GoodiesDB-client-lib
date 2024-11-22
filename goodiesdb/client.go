package goodiesdb

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Address string
}

func NewClient(address string) *Client {
	return &Client{Address: address}
}

func (c *Client) Set(key, value string) error {
	url := fmt.Sprintf("%s/set", c.Address)
	data := fmt.Sprintf("key=%s&value=%s", key, value)
	resp, err := http.Post(url, "application/x-www-form-urlencoded", bytes.NewBufferString(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	return err
}

func (c *Client) Get(key string) (string, error) {
	url := fmt.Sprintf("%s/get/%s", c.Address, key)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
