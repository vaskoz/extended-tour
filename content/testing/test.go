// +build OMIT

package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"testing"
)

func Itoa1(x int) string {
	return fmt.Sprintf("%d", x)
}

func Itoa2(x int) string {
	return strconv.Itoa(x)
}

var testcases = []struct {
	in  int
	out string
}{
	{5, "5"},
	{3003, "3003"},
}

func TestItoa1(t *testing.T) {
	for _, tc := range testcases {
		if out := Itoa1(tc.in); out != tc.out {
			t.Errorf("Expected %v, but got %v", tc.out, out)
		}
	}
}

func TestItoa2(t *testing.T) {
	for _, tc := range testcases {
		if out := Itoa2(tc.in); out != tc.out {
			t.Errorf("Expected %v, but got %v", tc.out, out)
		}
	}
}

func main() {
	log.Println(runtime.Version())
	flag.Set("test.bench", "anything")
	flag.Set("test.v", "true")

	testing.Main(func(pat, str string) (bool, error) { return true, nil },
		[]testing.InternalTest{{"TestItoa1", TestItoa1}, {"TestItoa2", TestItoa2}},
		[]testing.InternalBenchmark{},
		[]testing.InternalExample{})
}
