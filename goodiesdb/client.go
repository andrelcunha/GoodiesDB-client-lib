package goodiesdb

import (
	"bufio"
	"fmt"
	"net"
	"net/url"
	"strings"
)

type Client struct {
	Address  string
	Password string
	conn     net.Conn
	reader   *bufio.Reader
}

// NewClient creates a new Client with TCP connection
func NewClientByUrl(urlStr string) (*Client, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		// handle error
	}
	address := u.Host

	password := u.User.String()

	return NewClient(address, password)
}

// NewClient creates a new Client with TCP connection
func NewClient(address, password string) (*Client, error) {

	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	client := &Client{
		Address:  address,
		Password: password,
		conn:     conn,
		reader:   bufio.NewReader(conn),
	}

	if password != "" {
		err = client.Auth(password)
		if err != nil {
			return nil, err
		}
	}

	return client, nil
}

// Auth authenticates the client with the server
func (c *Client) Auth(password string) error {
	_, err := c.sendCommand(fmt.Sprintf("AUTH %s", password))
	return err
}

// Set sends the SET command to the server
func (c *Client) Set(key, value string) (string, error) {
	return c.sendCommand(fmt.Sprintf("SET %s %s", key, value))
}

// Get sends the GET command to the server
func (c *Client) Get(key string) (string, error) {
	return c.sendCommand(fmt.Sprintf("GET %s", key))
}

// sendCommand sends a command to the server and reads the response
func (c *Client) sendCommand(cmd string) (string, error) {
	_, err := fmt.Fprintln(c.conn, cmd)
	if err != nil {
		return "", err
	}

	response, err := c.reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(response), nil
}
