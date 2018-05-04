package layout

import (
	"fmt"
	"image"
	"github.com/GUMI-golang/gumi"
)

// Layout::Center
//
// Make all child center
type Center struct {
	gumi.ParentGUMI
	gumi.ParentBounder

}
func (s *Center) MaxChildrun() int {
	return 1
}
func (s *Center) String() string {
	return fmt.Sprintf("%s", "Center")
}
func (s *Center) RelayBound() image.Rectangle{
	rectangle := s.GetBound()
	chsz := s.Pipe().ProximateChildrunSize()
	if len(chsz) < 1{
		return image.ZR
	}
	sz := chsz[0]
	if sz == nil{
		return image.ZR
	}
	var vert, hori int
	if int(sz.Vertical.Max) < rectangle.Dy() {
		vert = int(sz.Vertical.Max)
	} else {
		vert = rectangle.Dy()
	}
	if int(sz.Horizontal.Max) < rectangle.Dx() {
		hori = int(sz.Horizontal.Max)
	} else {
		hori = rectangle.Dx()
	}
	left := (rectangle.Dx()-hori)/2 + rectangle.Min.X
	top := (rectangle.Dy()-vert)/2 + rectangle.Min.Y
	return image.Rect(left, top, left+hori, top+vert)
}
// Constructor 0
func Center0() *Center {
	return new(Center)
}