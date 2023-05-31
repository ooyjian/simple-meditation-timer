package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/aerth/playwav"
)

var minutes int
var testsound bool

func main() {
	flag.IntVar(&minutes, "t", 20, "The length of the meditation in minutes. Default to 20 minutes")
	flag.BoolVar(&testsound, "s", false, "Use this flag to test the sound. No timer will be started.")

	flag.Parse()

	if testsound {
		playwav.FromFile("sounds/bell.wav")
		return
	}

	fmt.Println("Meditation begins...")
	duration := time.Duration(minutes) * time.Second
	time.Sleep(duration)
	fmt.Println("Meditation ends...")

	playwav.FromFile("sounds/bell.wav")
}
