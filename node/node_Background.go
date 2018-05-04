package node


import (
	"image"
	"github.com/GUMI-golang/gumi"
	"github.com/GUMI-golang/gumi/gcore"
	"math"
	"github.com/GUMI-golang/gorat"
	"image/draw"
	"github.com/GUMI-golang/gumi/gime"
)

type Background struct {
	gumi.ParentGUMI
	gumi.ParentDoRender

	img *image.RGBA
}

func Background0() *Background {
	return new(Background)
}
const (
	background_image = "image"
)
func (s *Background) ListValue() []string {
	return []string{
		background_image,
	}
}
func (s *Background) GetValue(t string) (gime.Value) {
	switch t {
	case background_image:
		temp := image.NewRGBA(s.img.Rect)
		draw.Draw(temp, temp.Rect, s.img, s.img.Rect.Min, draw.Src)
		return temp
	}
	return gumi.ErrorNotControlable
}
func (s *Background) SetValue(t string, v gime.Value) error {
	switch t {
	case background_image:
		if vimg, ok := v.(image.Image);ok{
			s.SetImage(vimg)
			return nil
		}
		return gumi.ErrorInvalidValue
	}
	return gumi.ErrorNotControlable
}

func (s *Background) MaxChildrun() int {
	return 1
}
func (s *Background) String() string {
	return "Background"
}
func (s *Background) DoRender(rst gorat.SubRasterizer) {
	rst.SetFiller(gorat.NewImageFiller(s.img, gorat.ImageFillerGausian))
	w, h := rst.Size()
	rst.MoveTo(gorat.Vec2(0,0))
	rst.LineTo(gorat.Vec2(0,h))
	rst.LineTo(gorat.Vec2(w,h))
	rst.LineTo(gorat.Vec2(w,0))
	rst.Fill()
}
func (s *Background) Size() gcore.Size {
	sz := s.img.Bounds().Size()
	return gcore.Size{
		Horizontal:gcore.Length{
			Min:uint16(sz.X),
			Max:math.MaxUint16,
		},
		Vertical:gcore.Length{
			Min:uint16(sz.Y),
			Max:math.MaxUint16,
		},
	}
}

func (s *Background) SetImage(i image.Image) {
	s.img = image.NewRGBA(i.Bounds())
	draw.Draw(s.img, s.img.Rect, i, i.Bounds().Min, draw.Src)
}
func (s *Background) GetImage() image.Image {
	return s.img
}