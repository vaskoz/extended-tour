package main

import "log"

func main() {
	var a int
	go func() {
		a = 2
	}()
	a = 3
	log.Println("a has value", a)
}
