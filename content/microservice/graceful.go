package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Application struct{}

func (app *Application) Run(ctx context.Context) <-chan struct{} {
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		for i := 0; i < 30; i++ {
			time.Sleep(1 * time.Second) // do some simulated work
			select {
			case <-ctx.Done():
				break
			default:
				log.Println("Completed iteration", i)
			}
		}
		cancel()
	}()
	return ctx.Done()
}

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	application := &Application{}
	ctx, cancel := context.WithCancel(context.Background())

	select {
	case <-application.Run(ctx):
		log.Println("I've completed my work")
	case <-sigChan:
		cancel()
		<-ctx.Done()
		log.Println("I was asked to stop")
	}
}
