package screen

import (
    "fmt"
	"os"
)

type Screen struct {
	Height , Width uint
	heightCapacity , widthCapacity uint
	buffer [][]rune
}

func NewScreen(height , width uint) *Screen {
	s := &Screen {
		buffer : makeRuneBuffer(height, width),
		Height : height,
		Width : width,
		heightCapacity : height,
		widthCapacity : width,
	}

	s.SetAll(' ')
	return s;
}

func makeRuneBuffer(height , width uint) [][]rune {
	rs := make([][]rune , height)

	for i := range height {
		rs[i] = make([]rune, width)
	}

	return rs
}

func (s *Screen) SetAll(char rune){
	for y := range s.Height {
		for x := range s.Width {
			s.buffer[y][x] = char
		}
	}
}

func (s *Screen) Set(x,y uint , char rune) bool {
	if x > 0 && x < s.Width && y > 0 && y < s.Height {
		s.buffer[y][x] = char
		return true
	}
	return false
}

func (s *Screen) Get(x,y uint) (rune , bool) {
	if x > 0 && x < s.Width && y > 0 && y < s.Height {
		return s.buffer[y][x] , true
	}
	return ' ' , false
}

func (s *Screen) Display() {
	os.Stdout.Sync()
	for y := range s.Height {
		for x := range s.Width {
			fmt.Printf("%c", s.buffer[y][x])
		}
		fmt.Printf("%c" ,'\n')
	}
	os.Stdout.Sync()
}

func (s *Screen) Resize(height , width uint) {
	if height == 0 || width == 0 {
		return
	}
	
	var needResizeHeight, needResizeWidth bool = false , false

	s.Height , needResizeHeight = clampValue(s.heightCapacity , height)
	s.Width , needResizeWidth = clampValue(s.widthCapacity , width)

	if !needResizeHeight && !needResizeWidth {
		return
	}

	var newHeight , newWidth uint = s.heightCapacity , s.widthCapacity
	
	if needResizeWidth {
		newWidth = width
	}

	if needResizeHeight {
		newHeight = height
	}

	if !needResizeWidth && needResizeHeight {
		needAdd := s.Height - height

		for range needAdd {
			newLine := make([]rune , s.widthCapacity)
			s.buffer = append(s.buffer , newLine)
		}
		return
	}

	if needResizeWidth {
		buf := makeRuneBuffer(newHeight , newWidth)
		copyRunes(s.buffer , buf)
		s.buffer = buf 
		s.widthCapacity = width
		s.Width = width
	}

	s.Width = newWidth
	s.widthCapacity = newWidth
	s.Height = newHeight
	s.heightCapacity = newHeight
}

func clampValue(capacity uint , newValue uint) (uint , bool){
	if capacity >= newValue {
		return newValue , false
	}
	return 0 , true
}

func copyRunes(oldBuf, newBuf [][]rune){
	validateRunesBeforeCopy(oldBuf , newBuf)
	height := len(oldBuf)
	width := len(oldBuf[0])

	for y := range height {
		for x := range width {
			newBuf[y][x] = oldBuf[y][x]
		}
	}
}

func validateRunesBeforeCopy(oldBuf, newBuf [][]rune){
	oldBufLen := len(oldBuf) 
	newBufLen := len(newBuf)
	
	if oldBufLen == 0 || newBufLen == 0 {
		panic("Try copy buffer staff from old to new , but one of them is empty")
	}

	if oldBufLen > newBufLen {
		panic("Old buffer bigger than new (height), can't copy staff")
	}

	if len(oldBuf[0]) > len(newBuf[0]) {
		panic("Old buffer bigger than new (wight) , can't copy staff")
	}
}
