package main

import (
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"time"
)

func listen1(done <-chan interface{}) error {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("listen1 is working...")
		select {
		case <-done:
			return errors.New("listen1 canceled")
		default:
		}
	}
}

func listen2(done <-chan interface{}) error {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("listen2 is working...")
		select {
		case <-done:
			return errors.New("listen2 canceled")
		default:
		}
	}
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	g := new(errgroup.Group)
	done := make(chan struct{})
	g.Go(listen1)
	g.Go(listen2)
	if _, ok := <-c; ok {
		close(done)
	}
	if err := g.Wait(); err != nil {
		fmt.Println("something is wrong: %v", err)
	}
}
