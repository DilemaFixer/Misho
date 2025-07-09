package anim

import (
	"time"
	scr "github.com/DilemaFixer/Misho/src/screen"
)

type Drower interface {
	Drow(screen *scr.Screen) bool
	OnEnd(screen *scr.Screen)
}

type Point struct {
	X,Y uint 
}

type BaseAnimProp struct {
	startTime time.Time
	duration time.Duration
}
