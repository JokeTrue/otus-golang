package main

import (
	"bufio"
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
	Address    string
	Timeout    time.Duration
	In         io.ReadCloser
	Out        io.Writer
	Connection net.Conn
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &Client{
		Address: address,
		Timeout: timeout,
		In:      in,
		Out:     out,
	}
}

func (c *Client) Connect() (err error) {
	c.Connection, err = net.DialTimeout("tcp", c.Address, c.Timeout)
	return err
}

func (c *Client) Close() error {
	err := c.In.Close()
	if err != nil {
		return err
	}
	err = c.Connection.Close()
	return err
}

func (c *Client) execute(in io.Reader, out io.Writer) error {
	writer := bufio.NewWriter(out)
	scanner := bufio.NewScanner(in)
	for {
		scanned := scanner.Scan()
		if !scanned {
			if err := scanner.Err(); err != nil {
				return err
			}
			break
		}
		_, err := writer.WriteString(scanner.Text() + "\n")
		if err != nil {
			return err
		}
		err = writer.Flush()
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) Send() error {
	err := c.execute(c.In, c.Connection)
	return err
}

func (c *Client) Receive() error {
	err := c.execute(c.Connection, c.Out)
	return err
}
