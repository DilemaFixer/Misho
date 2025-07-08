package eng 

import (
	"log"
	"time"
	"os"
	scr "github.com/DilemaFixer/Misho/src/screen"
	con "github.com/DilemaFixer/Misho/src/console"
)

const MinDelay = 10 * time.Millisecond
var heightBefore , widthBefore uint 
var isChange bool = false

type Eng struct {
	screen *scr.Screen
	isStopt bool
}

func NewEng() *Eng{
	height , width , err := con.GetConsoleSize()

	if err != nil {
		log.Println("error getting console window size :" , err)
	}
	screen := scr.NewScreen(height,width)
	return &Eng{ 
		screen:screen, 
		isStopt:false,
	}
}

func (eng *Eng)StartWorkLoop(){
	con.HideCursor()
	defer con.ShowCursor()

	for {
		resizeScreenIfNeed(eng.screen)
		eng.screen.SetAll('*')
		eng.screen.Display()
		os.Stdout.Sync()
		time.Sleep(MinDelay)
		con.Clear()
		os.Stdout.Sync()
	}
}


func resizeScreenIfNeed(s *scr.Screen){
	height , width  , err := con.GetConsoleSize()

	if err != nil {
		log.Println("Error getting console window size :" , err)
	}

	if heightBefore != height || widthBefore != width{
		if heightBefore != height{
			heightBefore = height
		}

		if widthBefore != width {
			widthBefore = width
		}
		isChange = true
		return
	}

	if isChange {
		s.Resize(height , width)
		isChange = false
	}
}
