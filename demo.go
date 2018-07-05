package main

import (
	"fmt"
	"time"
	"runtime"
	"github.com/YuriyLisovskiy/TestRepo/src"
)

const (
	WIDTH       = 60
	HEIGHT      = 30
	GENERATIONS = 300
)

func main() {
	life := src.NewLife(WIDTH, HEIGHT)
	for i := 0; i < GENERATIONS; i++ {
		life.Step()
		src.CLS()
		src.PrintTitle(WIDTH, HEIGHT, i+1)
		fmt.Println(life)
		time.Sleep(time.Second / 10)
	}
	if runtime.GOOS == "windows" {
		src.RunCmd("pause")
	}
}
