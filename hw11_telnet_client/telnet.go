package main

import (
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	Close() error
	Send() error
	Receive() error
}

type Client struct {
	address    string
	timeout    time.Duration
	in         io.ReadCloser
	out        io.Writer
	connection net.Conn
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &Client{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

func (c *Client) Connect() (err error) {
	c.connection, err = net.DialTimeout("tcp", c.address, c.timeout)
	return err
}

func (c *Client) Close() error {
	return c.connection.Close()
}

func (c *Client) Send() error {
	_, err := io.Copy(c.connection, c.in)
	return err
}

func (c *Client) Receive() error {
	_, err := io.Copy(c.out, c.connection)
	return err
}
