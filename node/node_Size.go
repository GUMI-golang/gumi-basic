package node

import (
	"github.com/GUMI-golang/gumi/gcore"
	"github.com/GUMI-golang/gumi"
	"fmt"
	"github.com/GUMI-golang/gumi/gime"
	"github.com/pkg/errors"
)

type Size struct {
	gumi.ParentGUMI
	gumi.ParentBounder
	sz gcore.Size
}

// ValueManager
const (
	size_size  = "size"
)
func (s *Size) ListValue() []string {
	return []string{size_size}
}
func (s *Size) GetValue(t string) (gime.Value) {
	switch t {
	case size_size:
		return gcore.MarshalSize(s.GetSize())
	}
	return gumi.ErrorNotControlable
}
func (s *Size) SetValue(t string, v gime.Value) error {
	switch t {
	case size_size:
		if vstr, ok := v.(string);ok{
			sz, err := gcore.UnmarshalSize(vstr)
			if err != nil {
				return err
			}
			s.SetSize(sz)
		}
		return errors.WithMessage(gumi.ErrorInvalidValue, "value must be string")
	}
	return gumi.ErrorNotControlable
}

func (s *Size) Size() (result gcore.Size ){
	chszs := s.Pipe().ProximateChildrunSize()
	if len(chszs) < 1{
		return s.sz
	}
	chsz := chszs[0]
	result = s.sz
	if result.Vertical == gcore.AUTOLENGTH {
		result.Vertical = chsz.Vertical
	} else if s.sz.Vertical == gcore.MINLENGTH {
		result.Vertical = chsz.Vertical
		result.Vertical.Max = chsz.Vertical.Min
	} else if s.sz.Vertical == gcore.MAXLENGTH {
		result.Vertical = chsz.Vertical
		result.Vertical.Min = chsz.Vertical.Max
	}
	if s.sz.Horizontal == gcore.AUTOLENGTH {
		result.Horizontal = chsz.Horizontal
	} else if s.sz.Horizontal == gcore.MINLENGTH {
		result.Horizontal = chsz.Horizontal
		result.Horizontal.Max = chsz.Horizontal.Min
	} else if s.sz.Horizontal == gcore.MAXLENGTH {
		result.Horizontal = chsz.Horizontal
		result.Horizontal.Min = chsz.Horizontal.Max
	}

	return
}
func (s *Size) MaxChildrun() int {
	return 1
}
func (s *Size) String() string {
	return fmt.Sprintf("%s(GUMISize:%v)", "Size", s.sz)
}

func Size0() *Size {
	return &Size{
		sz: gcore.AUTOSIZE,
	}
}
//

func (s *Size) GetSize() gcore.Size {
	return s.sz
}
func (s *Size) SetSize(sz gcore.Size) {
	s.sz = sz
	s.RequestResize()
}