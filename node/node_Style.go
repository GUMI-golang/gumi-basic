package node

import (
	"github.com/GUMI-golang/gumi"
	"fmt"
	"github.com/GUMI-golang/gumi/gime"
)

type Style struct {
	gumi.ParentGUMI
}
func (s *Style) ListValue() []string {
	return gumi.ListStringStyleData()
}
func (s *Style) GetValue(t string) (gime.Value) {
	if temp := gumi.FromStringStyleData(t); temp != gumi.STYLE_INVALID{
		value := s.Pipe().GetStyle(temp)
		// TODO : if value is gime.Value valid, return it direct
		res, err := temp.Type().Marshal(value)
		if err != nil {
			return err
		}
		return res
	}
	return gumi.ErrorNotControlable
}
func (s *Style) SetValue(t string, v gime.Value) error {
	// TODO, ReRendering Request
	if temp := gumi.FromStringStyleData(t); temp != gumi.STYLE_INVALID{
		// TODO : if v is gime.Value valid, return it direct
		if vstr, ok := v.(string); ok{
			res, err := temp.Type().Unmarshal(vstr)
			if err != nil {
				return err
			}
			s.Pipe().SetStyle(temp, res)
			return nil
		}

	}
	return gumi.ErrorNotControlable
}


// GUMI
func (s *Style) MaxChildrun() int {
	return 1
}
func (s *Style) String() string {
	temp := map[gumi.StyleData]interface{}{}
	s.Pipe().ListStyle(func(sdata gumi.StyleData, value interface{}) {
		temp[sdata] = value
	})
	return fmt.Sprint("Style ", temp)
}

func Style0() *Style {
	return &Style{}
}
