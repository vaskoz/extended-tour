// +build OMIT

package main

import (
	"fmt"

	"github.com/Netflix/chaosmonkey/clock"
)

func main() {
	c := clock.New()
	fmt.Printf("c is %v of type %[1]T\n", c)
}
