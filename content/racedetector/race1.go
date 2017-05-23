package main

import "fmt"

func main() {
	var a int
	go func() {
		a = 2
	}()
	a = 3
}
