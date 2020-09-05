package main

import (
	"fmt"
	"os"
	"time"
)

//go:generate go get github.com/rakyll/statik
//go:generate statik

func main() {
	c, err := NewCmd()
	if err != nil {
		Panic(c, err)
	}

	year, month, err := c.ParseArgs(os.Args[1:], time.Now())
	if err != nil {
		Panic(c, err)
	}

	for _, line := range c.CreateCalendar(year, month) {
		fmt.Println(line)
	}
}

func Panic(c *Cmd, err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	for _, v := range c.CreateUsage() {
		fmt.Fprintln(os.Stderr, v)
	}
	os.Exit(1)
}
