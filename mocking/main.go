package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

// type sleeper struct{}
type Sleeper interface {
	Sleep()
}

// func (s sleeper) Sleep() {
// 	time.Sleep(1 * time.Second)
// }

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (config ConfigurableSleeper) Sleep() {
	config.sleep(config.duration)
}

func Countdown(out io.Writer, t Sleeper) {
	for i := 3; i > 0; i-- {
		t.Sleep()
		fmt.Fprintf(out, strconv.Itoa(i)+"\n")
	}
	t.Sleep()
	fmt.Fprintf(out, "Go!")
}
