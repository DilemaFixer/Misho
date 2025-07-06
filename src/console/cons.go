package console 

import (
	"fmt"
	"os"
	"golang.org/x/term"
)

func GetConsoleSize() (uint, uint, error) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
    return uint(height), uint(width), err
}

func HideCursor() {
 	fmt.Print("\033[?25l")
}

func ShowCursor() {
 	fmt.Print("\033[?25h")
}

func Clear(){
	fmt.Print("\033[2J\033[H")
}
