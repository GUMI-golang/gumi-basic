package drawing

import (
	"github.com/GUMI-golang/gumi"
	"github.com/GUMI-golang/gorat"
	"github.com/GUMI-golang/gumi/gime"
	"strings"
	"fmt"
)

//
// support 'fn'
//		- graduation.vertical			(pivot float32)
//		- graduation.horizontal			(pivot float32)
//		- grid.vertical					(pivot float32)
//		- grid.horizontal				(pivot float32)
//		- hint.vertical					(pivot float32)
//		- hint.horizontal				(pivot float32)
//		- fps							(prefix string)
type Draw struct {
	gumi.ParentGUMI
	drawfuncs []Drawer
}

const (
	draw_fn = "fn"
)
func (s *Draw) ListValue() []string {
	return []string{draw_fn}
}
func (s *Draw) GetValue(t string) (gime.Value) {
	switch t {
	case draw_fn:
		// TODO : better than nil?
		return nil
	}
	return gumi.ErrorNotControlable
}
func (s *Draw) SetValue(t string, v gime.Value) error {
	switch t {
	case draw_fn:
		if vstr, ok := v.(string); ok{
			var temp []Drawer = nil
			for _, fnstr := range strings.Fields(vstr) {
				fns, err := interpreter.Call(fnstr)
				fmt.Println(fns, err)
				if err != nil {
					return err
				}
				for _, f := range fns {
					temp = append(temp, f.(Drawer))
				}
			}
			s.drawfuncs =temp
			return nil
		}
		return gumi.ErrorInvalidValue
	}
	return gumi.ErrorNotControlable
}
func (s *Draw) PostRender(rasterzier gorat.SubRasterizer) {
	for _, fn := range s.drawfuncs {
		fn.Draw(s.Pipe(), rasterzier)
	}
}
func (s *Draw) MaxChildrun() int {
	return 1
}
func (s *Draw) String() string {
	return "Draw"
}
//func (s *NDrawing) BaseRender(subimg *image.RGBA) {
//}
//
//func (s *NDrawing) DecalRender(fullimg *image.RGBA) (updated image.Rectangle) {
//	sub := fullimg.SubImage(s.rnode.GetAllocation()).(*image.RGBA)
//	ctx := createContext(sub)
//	var res image.Rectangle
//	for _, v := range s.drawfuncs {
//		ctx.Push()
//		res = res.Union(v.Draw(ctx, s.style).Add(s.rnode.GetAllocation().Min))
//		ctx.Pop()
//	}
//	return res
//}
//
//func (s *NDrawing) GUMIInfomation(info Information) {
//	var changed bool
//	for _, v := range s.drawfuncs {
//		if v2, ok := v.(DrawerWithInformation); ok {
//			changed = changed || v2.Inform(info)
//		}
//	}
//	s.child.GUMIInfomation(info)
//}
//func (s *NDrawing) GUMIStyle(style *Style) {
//	s.style = style
//	s.child.GUMIStyle(style)
//}
//
//func (s *NDrawing) GUMIRenderSetup(man *renderline.Manager, parent renderline.Node) {
//	s.rmana = man
//	s.rnode = man.New(parent, nil)
//	s.rnode.SetJob(s)
//	s.child.GUMIRenderSetup(man, s.rnode)
//}
//
//func (s *NDrawing) GUMIHappen(event Event) {
//	s.child.GUMIHappen(event)
//}
//func (s *NDrawing) GUMISize() gcore.Size {
//	return s.child.GUMISize()
//}
//func (s *NDrawing) String() string {
//	return fmt.Sprintf("%s(drawing:%d GUMIRender)", "NDrawing", len(s.drawfuncs))
//}
//
////
//func NDrawing0(drawFuncs ...Drawer) *NDrawing {
//	return &NDrawing{
//		drawfuncs: drawFuncs,
//	}
//}

