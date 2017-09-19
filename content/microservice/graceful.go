package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Application struct {
	notifyWhenDone chan struct{}
}

func (app *Application) Run() <-chan struct{} {
	go func() {
		time.Sleep(30 * time.Second)
		app.Stop()
	}()
	return app.notifyWhenDone
}

func (app *Application) Stop() {
	app.notifyWhenDone <- struct{}{}
}

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	application := &Application{make(chan struct{}, 1)}

	select {
	case <-application.Run():
		log.Println("I've completed my work")
	case <-sigChan:
		application.Stop()
		log.Println("I was asked to stop")
	}
}
