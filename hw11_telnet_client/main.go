package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var timeout string

func init() {
	flag.StringVar(&timeout, "timeout", "10s", "connection timeout")
	flag.Parse()
}

const MinArgsCount = 3

func Worker(ctx context.Context, handler func() error, cancelFunc context.CancelFunc) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			err := handler()
			if err != nil {
				cancelFunc()
				return
			}
		}
	}
}
func main() {
	argsCount := len(os.Args)
	if argsCount < MinArgsCount {
		log.Fatalf("Expected to have at least 3 arguments, but got %d", argsCount)
	}

	host, port := os.Args[2], os.Args[3]
	parsedTimeout, err := time.ParseDuration(timeout)
	if err != nil {
		log.Fatalln(err)
	}

	client := NewTelnetClient(net.JoinHostPort(host, port), parsedTimeout, os.Stdin, os.Stdout)
	err = client.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err = client.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	ctx, cancelFunc := context.WithCancel(context.Background())
	go Worker(ctx, client.Receive, cancelFunc)
	go Worker(ctx, client.Send, cancelFunc)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-sigCh:
			cancelFunc()
			close(sigCh)
			return

		case <-ctx.Done():
			close(sigCh)
			return
		}
	}
}
