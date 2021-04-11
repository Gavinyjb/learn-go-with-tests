package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}
type ConfigurableSleeper struct {
	duration time.Duration
}

func (c *ConfigurableSleeper) Sleep() {
	time.Sleep(c.duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		//time.Sleep(1*time.Second)
		sleeper.Sleep()
		fmt.Fprintln(out, i)

	}
	//time.Sleep(1*time.Second)
	sleeper.Sleep()
	fmt.Fprintf(out, "Go!")
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second}
	Countdown(os.Stdout, sleeper)
}
