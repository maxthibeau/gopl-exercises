package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type tTimeFunc func([]string)

// exercises 1.2 and 1.1
func echoIndex(input []string) {
	s, sep := "", ""
	for i, arg := range os.Args[0:] {
		s += sep + strconv.Itoa(i) + " " + arg
		sep = " "
	}
	fmt.Println(s)
}

// exercise 1.3 testing of strings.join() with exponentially scaled input
func doubleInput(iters int) []string {
	test := os.Args[1:]
	for i := 0; i < iters; i++ {
		test = append(test, test...)
	}
	return test
}

func echo(input []string) {
	s, sep := "", ""
	for _, arg := range input {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echoJoin(input []string) {
	fmt.Println(strings.Join(input, " "))
}

func timeFunc(f tTimeFunc, funcName string, iters int) {
	test := doubleInput(iters)
	start := time.Now()
	f(test)
	fmt.Printf(funcName+": %.10fs\n", time.Since(start).Seconds())
}

func main() {
	timeFunc(echo, "echo", 10)
	timeFunc(echoJoin, "echoJoin", 10)
}
