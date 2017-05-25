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

func BenchmarkItoa1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			Itoa1(tc.in)
		}
	}
}

func BenchmarkItoa2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			Itoa2(tc.in)
		}
	}
}

func main() {
	log.Println(runtime.Version())
	flag.Set("test.bench", "anything")
	flag.Set("test.v", "true")

	testing.Main(func(pat, str string) (bool, error) { return true, nil },
		[]testing.InternalTest{},
		[]testing.InternalBenchmark{{"BenchmarkItoa1", BenchmarkItoa1}, {"BenchmarkItoa2", BenchmarkItoa2}},
		[]testing.InternalExample{})
}
