package main 

import (
	"time"
	eng "github.com/DilemaFixer/Misho/src/eng"
	anim "github.com/DilemaFixer/Misho/src/anim"
)

func main(){
	animEng := eng.NewEng()
	lineDrawer := anim.NewLineDrawer(
		'*',
		anim.Point{X: 5, Y: 5},   
		anim.Point{X: 75, Y: 20}, 
		15.0,                     		
		4*time.Second)
	lineDrawer1 := anim.NewLineDrawer(
		'*',
		anim.Point{X: 3, Y: 10},   
		anim.Point{X: 50, Y: 30}, 
		15.0,                     		
		8*time.Second)
	animEng.AddDrower(lineDrawer)
	animEng.AddDrower(lineDrawer1)

	animEng.StartWorkCycle()
}
