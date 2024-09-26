package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go"
	countDownStart = 3
)

type Sleeper interface {
	Sleep()
}

func CountDown(out io.Writer, s Sleeper) {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		s.Sleep()
	}

	fmt.Fprint(out, finalWord)
}

type DefaultSleeper struct{}

func (s DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	sleeper := DefaultSleeper{}
	CountDown(os.Stdout, sleeper)
}
