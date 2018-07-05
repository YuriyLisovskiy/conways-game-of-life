package main

import (
	"fmt"
	"time"
	"runtime"
)

const (
	WIDTH       = 60
	HEIGHT      = 30
	GENERATIONS = 300
)

func main() {
	life := NewLife(WIDTH, HEIGHT)
	for i := 0; i < GENERATIONS; i++ {
		life.Step()
		CLS()
		printTitle(WIDTH, HEIGHT, i+1)
		fmt.Println(life)
		time.Sleep(time.Second / 10)
	}
	if runtime.GOOS == "windows" {
		runCmd("pause")
	}
}
