package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprintf(out, finalWord)
}

func main() {
	defaultSleeper := &DefaultSleeper{}
	Countdown(os.Stdout, defaultSleeper)
}
