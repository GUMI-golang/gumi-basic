package actor

import (
	"fmt"
	"image/color"

	"github.com/GUMI-golang/gumi"
	"github.com/GUMI-golang/gumi/gcore"
	"github.com/GUMI-golang/gorat"
	"github.com/GUMI-golang/gorat/textrat"
	"github.com/pkg/errors"
	"github.com/GUMI-golang/gumi/gime"
)

// Actor::Text
//
// Text use for render text
type Text struct {
	gumi.ParentGUMI
	gumi.ParentDoRender
	//
	align gcore.Align
	text  string
	//
}

// GUMI
func (s *Text) MaxChildrun() int {
	return 0
}
func (s *Text) Size() gcore.Size {
	//f := s.Pipe().GetStyle(gumi.STYLE_Font).(*gcore.Font)
	f := textrat.Default
	size := f.MeasureText(s.text)
	temp := gcore.Size{
		Horizontal: gcore.MinLength(uint16(size[0])),
		Vertical:   gcore.MinLength(uint16(size[1])),
	}
	return temp
}
func (s *Text) String() string {
	return fmt.Sprintf("%s(text:%s, align:%s)", "Text", s.text, s.align)
}

// DoRender
func (s *Text) DoRender(rst gorat.SubRasterizer) {
	bd := rst.Bound()
	bd = bd.Sub(bd.Min)
	rst.SetFiller(gorat.NewColorFiller(s.Pipe().GetStyle(gumi.STYLE_TextColor).(color.Color)))
	textrat.Default.TextInRect(rst, s.text, bd, s.align)
}

// Constructor 0
func Text0() *Text {
	return &Text{}
}

// ValueManager
const (
	text_text  = "text"
	text_align = "align"
)

func (s *Text) ListValue() []string {
	return []string{
		text_text,
		text_align,
	}
}
func (s *Text) GetValue(t string) (gime.Value) {
	switch t {
	case text_text:
		return s.GetText()
	case text_align:
		return gcore.MarshalAlign(s.GetAlign())

	}
	return  gumi.ErrorNotControlable
}
func (s *Text) SetValue(t string, v gime.Value) error {
	switch t {
	case text_text:
		if vstr, ok := v.(string); ok{
			s.SetText(vstr)
			return nil
		}
		return errors.WithMessage(gumi.ErrorInvalidValue, "Invalid value type")
	case text_align:
		if vstr, ok := v.(string); ok{
			if temp, err := gcore.UnmarshalAlign(vstr); err != nil {
				return errors.WithMessage(gumi.ErrorParsingFail, err.Error())
			} else {
				s.SetAlign(temp)
				return nil
			}
		}
		return errors.WithMessage(gumi.ErrorInvalidValue, "Invalid value type")
	}
	return gumi.ErrorNotControlable
}

// DefaultManager
func (s *Text) GetDefault() (gime.Value) {
	return s.GetText()
}
func (s *Text) SetDefault(v gime.Value) error {
	return s.SetValue(text_text, v)
}

// Method Text
func (s *Text) SetText(text string) {
	if s.text != text {
		s.text = text
		s.Throw()
	}
}
func (s *Text) GetText() string {
	return s.text
}

// Method Align
func (s *Text) SetAlign(align gcore.Align) {
	if s.align != align {
		s.align = align
		s.Throw()
	}
}
func (s *Text) GetAlign() gcore.Align {
	return s.align
}
