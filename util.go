package main

import (
	"os"
	"fmt"
	"runtime"
	"os/exec"
)

func CLS() {
	command := ""
	switch runtime.GOOS {
	case "windows":
		command = "cls"
	case "linux":
		command = "clear"
	default:
		panic(fmt.Sprintf("can't clear the screen on %s", runtime.GOOS))
	}
	runCmd(command)
}

func runCmd(cmd string) {
	c := exec.Command(cmd)
	c.Stdout = os.Stdout
	c.Run()
}

func printSymbols(length int, symbol string, newLine bool) {
	for i := 0; i < length; i++ {
		print(symbol)
	}
	if newLine {
		println()
	}
}

func printTitle(width, height, generation int) {
	printSymbols(width, "  ", true)
	printSymbols((width-21/2)/2, "  ", false)
	print(" Conway's Game of Life ")
	printSymbols(width-21/2, "  ", true)
	universeString := fmt.Sprintf(" Universe %dx%d", width, height)
	generationString := fmt.Sprintf("Generation %d ", generation)
	print(universeString)
	printSymbols(width-len(universeString)/2-len(generationString)/2, "  ", false)
	println(generationString)
	printSymbols(width, "  ", true)
}
