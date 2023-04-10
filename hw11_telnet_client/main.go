package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "")
	flag.Parse()

	host, port := flag.Arg(0), flag.Arg(1)
	conn := NewTelnetClient(net.JoinHostPort(host, port), *timeout, os.Stdin, os.Stdout)
	if err := conn.Connect(); err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		SenderR(ctx, conn, wg)
	}()
	go func() {
		RecieverR(ctx, conn, wg)
	}()
	wg.Wait()
}

func SenderR(ctx context.Context, conn TelnetClient, wg *sync.WaitGroup) {
SENDER:
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			break SENDER
		default:
			err := conn.Send()
			if err != nil {
				wg.Done()
				log.Fatal(err)
			}
		}
	}
}

func RecieverR(ctx context.Context, conn TelnetClient, wg *sync.WaitGroup) {
RECEIVER:
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			break RECEIVER
		default:
			err := conn.Receive()
			if err != nil {
				wg.Done()
				log.Fatal(err)
			}
		}
	}
}
