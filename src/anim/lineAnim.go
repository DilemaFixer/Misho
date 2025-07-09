package anim

import (
	"time"
	"math"
	scr "github.com/DilemaFixer/Misho/src/screen"
)

type LineDrawer struct {
	BaseAnimProp
	startPoint Point
	endPoint Point
	speed float64
	char rune
}

func NewLineDrawer(char rune, start, end Point, speed float64, duration time.Duration) *LineDrawer {
	return &LineDrawer{
		BaseAnimProp: BaseAnimProp{
			startTime: time.Now(),
			duration:  duration,
		},
		startPoint: start,
		endPoint:   end,
		speed:      speed,
		char: 		char,
	}
}

func (drw *LineDrawer) Drow(screen *scr.Screen) bool {
	elapsed := time.Since(drw.startTime)
	
	if elapsed > drw.duration {
		return true
	}
	
	progress := float64(elapsed) / float64(drw.duration)
	progress = math.Min(progress, 1.0)
	
	drw.drawProgressiveLine(screen, progress)
	
	return false
}

func (drw *LineDrawer)OnEnd(screen *scr.Screen){
	return
}

func (drw *LineDrawer) drawProgressiveLine(screen *scr.Screen, progress float64) {
	dx := float64(drw.endPoint.X) - float64(drw.startPoint.X)
	dy := float64(drw.endPoint.Y) - float64(drw.startPoint.Y)
	distance := math.Sqrt(dx*dx + dy*dy)
	
	steps := int(distance * progress)
	if steps < 1 {
		steps = 1
	}
	
	for i := 0; i <= steps; i++ {
		t := float64(i) / distance
		if t > progress {
			break
		}
		
		x := float64(drw.startPoint.X) + dx*t
		y := float64(drw.startPoint.Y) + dy*t
		
		screen.Set(uint(math.Round(x)), uint(math.Round(y)), drw.char)
	}
}
