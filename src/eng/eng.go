package eng 

import (
	"log"
	"time"
	"github.com/DilemaFixer/Misho/src/anim"
	scr "github.com/DilemaFixer/Misho/src/screen"
	con "github.com/DilemaFixer/Misho/src/console"
)

const MinDelay = 50 * time.Millisecond
var heightBefore , widthBefore uint 
var isChange bool = false

type Eng struct {
	screen *scr.Screen
	drowers []anim.Drower
}

func NewEng() *Eng{
	height , width , err := con.GetConsoleSize()

	if err != nil {
		log.Println("error getting console window size :" , err)
	}
	screen := scr.NewScreen(height,width)
	return &Eng{ 
		screen:screen, 
		drowers:make([]anim.Drower , 0),
	}
}

func (eng *Eng)StartWorkCycle(){
	con.HideCursor()
	defer con.ShowCursor()

	for !eng.isAnimationsEnd() {
		resizeScreenIfNeed(eng.screen)

		for i := len(eng.drowers) - 1; i >= 0; i-- {
			isDone := eng.drowers[i].Drow(eng.screen)

			if isDone {
				eng.drowers = append(eng.drowers[:i] , eng.drowers[i+1:]...)
			}
		}

		eng.drowFrame()
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

func (eng *Eng)isAnimationsEnd() bool{
	if len(eng.drowers) == 0 {
		return true	
	}
	return false
}

func (eng *Eng)drowFrame(){
	eng.screen.Display()
	time.Sleep(MinDelay)
	con.Clear()
}

func (eng *Eng)AddDrower(drw anim.Drower) {
	eng.drowers = append(eng.drowers , drw)
}
