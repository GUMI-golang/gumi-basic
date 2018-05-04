package actor

import (
	"image"
	"github.com/GUMI-golang/gumi"
	"github.com/GUMI-golang/gumi/gcore"
	"math"
	"github.com/GUMI-golang/gorat"
	"image/draw"
	"github.com/GUMI-golang/gumi/gime"
)

type Image struct {
	gumi.ParentGUMI
	gumi.ParentDoRender

	img *image.RGBA
}

func Image0() *Image {
	return new(Image)
}
const (
	image_image  = "image"
)
func (s *Image) ListValue() []string {
	return []string{
		image_image,
	}
}
func (s *Image) GetValue(t string) (gime.Value) {
	switch t {
	case image_image:
		temp := image.NewRGBA(s.img.Rect)
		draw.Draw(temp, temp.Rect, s.img, s.img.Rect.Min, draw.Src)
		return temp
	}
	return gumi.ErrorNotControlable
}
func (s *Image) SetValue(t string, v gime.Value) error {
	switch t {
	case image_image:
		if vimg, ok := v.(image.Image);ok{
			s.SetImage(vimg)
			return nil
		}
		return gumi.ErrorInvalidValue
	}
	return gumi.ErrorNotControlable
}

func (s *Image) GetDefault() (gime.Value) {
	return s.GetValue(image_image)
}
func (s *Image) SetDefault(v gime.Value) error {
	return s.SetValue(image_image, v)
}

func (s *Image) MaxChildrun() int {
	return 0
}
func (s *Image) String() string {
	return "Image"
}
func (s *Image) DoRender(rst gorat.SubRasterizer) {
	rst.SetFiller(gorat.NewImageFiller(s.img, gorat.ImageFillerGausian))
	w, h := rst.Size()
	rst.MoveTo(gorat.Vec2(0,0))
	rst.LineTo(gorat.Vec2(0,h))
	rst.LineTo(gorat.Vec2(w,h))
	rst.LineTo(gorat.Vec2(w,0))
	rst.Fill()
}
func (s *Image) Size() gcore.Size {
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

func (s *Image ) SetImage(i image.Image) {
	s.img = image.NewRGBA(i.Bounds())
	draw.Draw(s.img, s.img.Rect, i, i.Bounds().Min, draw.Src)
}
func (s *Image ) GetImage() image.Image {
	return s.img
}