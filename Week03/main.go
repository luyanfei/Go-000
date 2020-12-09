package main

import (
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)

	g := new(errgroup.Group)
	done := make(chan struct{})
	g.Go(func() error {
		for {
			time.Sleep(1 * time.Second)
			fmt.Println("listen1 is working...")
			select {
			case <-done:
				return errors.New("listen1 canceled")
			default:
			}
		}
	})
	g.Go(func() error {
		for {
			time.Sleep(1 * time.Second)
			fmt.Println("listen2 is working...")
			select {
			case <-done:
				return errors.New("listen2 canceled")
			default:
			}
		}
	})
	if _, ok := <-c; ok {
		close(done)
	}
	if err := g.Wait(); err != nil {
		fmt.Printf("something is wrong: %v", err)
	}
}
