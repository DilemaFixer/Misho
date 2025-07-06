package main 

import (
	"log"
	"time"
	"os"
	scr "github.com/DilemaFixer/Misho/src/screen"
	con "github.com/DilemaFixer/Misho/src/console"
)

const MinDelay = 10 * time.Millisecond
var HeightBefore , WidthBefore uint 
var IsChange bool = false

func resizeIfNeed(s *scr.Screen){
	h , w , err := con.GetConsoleSize()

	if err != nil {
		log.Println("error getting console window size :" , err)
	}

	if HeightBefore != h || WidthBefore != w {
		if HeightBefore != h {
			HeightBefore = h
		}

		if WidthBefore != w {
			WidthBefore = w
		}
		IsChange = true
		return
	}

	if IsChange {
		s.Resize(h , w)
		IsChange = false
	}
}

func main(){
	con.HideCursor()
	defer con.ShowCursor()
	h , w , err := con.GetConsoleSize()

	if err != nil {
		log.Println("error getting console window size :" , err)
	}
	screen := scr.NewScreen(h,w)
	
	for {
		resizeIfNeed(screen)
		screen.SetAll('*')
		screen.Display()
		os.Stdout.Sync()
		time.Sleep(MinDelay)
		con.Clear()
		os.Stdout.Sync()
	}
}
