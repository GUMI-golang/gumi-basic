package drawing

import (
	"github.com/GUMI-golang/gumi"
	"github.com/GUMI-golang/gorat"
	"github.com/GUMI-golang/gorat/textrat"
	"strconv"
	"github.com/GUMI-golang/gumi/gcore"
	"fmt"
)

func (Ruler) Size() Drawer {
	return DrawerFunc(func(pipe *gumi.Pipe, r gorat.VectorDrawer) {
		w, h := r.Size()
		f := pipe.GetStyle(gumi.STYLE_Font).(textrat.Font)
		// horizontal line
		r.MoveTo(gorat.Vec2(0, Drawing.RulerWidth))
		r.LineTo(gorat.Vec2(w, Drawing.RulerWidth))
		//
		r.MoveTo(gorat.Vec2(Drawing.RulerWidth/2, Drawing.RulerWidth/2))
		r.LineTo(gorat.Vec2(0, Drawing.RulerWidth))
		r.LineTo(gorat.Vec2(Drawing.RulerWidth/2, Drawing.RulerWidth/2 * 3))
		r.MoveTo(gorat.Vec2(w - Drawing.RulerWidth/2, Drawing.RulerWidth/2))
		r.LineTo(gorat.Vec2(w, Drawing.RulerWidth))
		r.LineTo(gorat.Vec2(w - Drawing.RulerWidth/2, Drawing.RulerWidth/2 * 3))
		// Vertical Line
		r.MoveTo(gorat.Vec2(Drawing.RulerWidth, 0))
		r.LineTo(gorat.Vec2(Drawing.RulerWidth, h))
		//
		r.MoveTo(gorat.Vec2(Drawing.RulerWidth/2, Drawing.RulerWidth/2))
		r.LineTo(gorat.Vec2(Drawing.RulerWidth, 0))
		r.LineTo(gorat.Vec2(Drawing.RulerWidth/2 * 3, Drawing.RulerWidth/2))
		r.MoveTo(gorat.Vec2(Drawing.RulerWidth/2, h - Drawing.RulerWidth/2))
		r.LineTo(gorat.Vec2(Drawing.RulerWidth, h))
		r.LineTo(gorat.Vec2(Drawing.RulerWidth/2 * 3, h - Drawing.RulerWidth/2))
		//
		r.Stroke()
		// text
		f.PathText(r, strconv.FormatInt(int64(w), 10), gorat.Vec2(w / 2, Drawing.RulerWidth), gcore.AlignCenter)
		f.PathText(r, strconv.FormatInt(int64(h), 10), gorat.Vec2(Drawing.RulerWidth, h / 2), gcore.AlignCenter)
		r.Fill()
	})
}
func (Ruler) Proportion() Drawer {
	return DrawerFunc(func(pipe *gumi.Pipe, r gorat.VectorDrawer) {
		w, h := r.Size()
		f := pipe.GetStyle(gumi.STYLE_Font).(textrat.Font)
		gcd := float32(gcore.GCD(int64(w), int64(h)))
		txt := fmt.Sprintf("%d : %d", int(w/gcd), int(h/gcd))
		//
		r.MoveTo(gorat.Vec2(0,0))
		r.LineTo(gorat.Vec2(w,h))
		r.Stroke()
		f.PathText(r, txt, gorat.Vec2(w/2, h/2), gcore.AlignCenter)
		r.Fill()
		r.Stroke()
	})
}
//func (_Ruler) Screen() Drawer {
//	return FunctionDrawer{func(context *gg.Context, style *Style) image.Rectangle {
//		style.UseContext(context)
//		defer style.ReleaseContext(context)
//		context.SetColor(rulerColor)
//		for _, v := range gcore.DefinedResolutions.Smaller(context.Width(), context.Height()) {
//			context.DrawRectangle(0, 0, float64(v.Width), float64(v.Height))
//			w, _ := context.MeasureString(v.Name[0])
//			context.DrawString(v.Name[0], float64(v.Width)-w-5, float64(v.Height)-5)
//		}
//		context.Stroke()
//		return image.Rect(0, 0, context.Width(), context.Height())
//	}}
//}
//func (_Ruler) Dots(pivot float64) Drawer {
//	return FunctionDrawer{func(context *gg.Context, style *Style) image.Rectangle {
//		style.UseContext(context)
//		defer style.ReleaseContext(context)
//		context.SetColor(rulerColor)
//		for x := 0.; x <= float64(context.Width()); x += float64(pivot) {
//			for y := 0.; y <= float64(context.Height()); y += float64(pivot) {
//				context.DrawPoint(x, y, style.Default.LineWidth)
//			}
//		}
//		context.Stroke()
//		return image.Rect(0, 0, context.Width(), context.Height())
//	}}
//}
