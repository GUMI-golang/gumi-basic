package drawing

import (
	"github.com/GUMI-golang/gumi"
	"github.com/GUMI-golang/gorat"
	"strconv"
	"github.com/GUMI-golang/gorat/textrat"
	"github.com/GUMI-golang/gumi/gcore"
)


func (Graduation) Vertical(pivot float32) Drawer {
	return DrawerFunc(func(pipe *gumi.Pipe, r gorat.VectorDrawer) {
		_, h := r.Size()
		for f := float32(0.0); f <= h; f += pivot {
			r.MoveTo(gorat.Vec2(0, f))
			r.LineTo(gorat.Vec2(Drawing.RulerWidth, f))
		}
		r.Stroke()
	})
}
func (Graduation) Horizontal(pivot float32) Drawer {
	return DrawerFunc(func(pipe *gumi.Pipe, r gorat.VectorDrawer) {
		w, _ := r.Size()
		for f := float32(0.0); f <= w; f += pivot {
			r.MoveTo(gorat.Vec2(f, 0 ))
			r.LineTo(gorat.Vec2(0, Drawing.RulerWidth))
		}
		r.Stroke()
	})
}

func (Grid) Vertical(pivot float32) Drawer {
	return DrawerFunc(func(pipe *gumi.Pipe, r gorat.VectorDrawer) {
		w, h := r.Size()
		for f := float32(0.0); f <= float32(w); f += pivot{
			r.MoveTo(gorat.Vec2(f, 0))
			r.MoveTo(gorat.Vec2(f, h))
		}
		r.Stroke()
	})
}
func (Grid) Horizontal(pivot float32) Drawer {
	return DrawerFunc(func(pipe *gumi.Pipe, r gorat.VectorDrawer) {
		w, h := r.Size()
		for f := float32(0.0); f <= float32(h); f += pivot{
			r.MoveTo(gorat.Vec2(0, f))
			r.MoveTo(gorat.Vec2(w, f))
		}
		r.Stroke()
	})
}


func (Hint) Vertical(pivot float32) Drawer {
	return DrawerFunc(func(pipe *gumi.Pipe, r gorat.VectorDrawer) {
		_, h := r.Size()
		f := pipe.GetStyle(gumi.STYLE_Font).(textrat.Font)
		for p := float32(pivot); p <= h; p += pivot {
			txt := strconv.FormatInt(int64(p), 10)
			f.PathText(r, txt, gorat.Vec2(0, p), gcore.AlignLeft | gcore.AlignTop)
		}
		r.Fill()
	})
}
func (Hint) Horizontal(pivot float32) Drawer {
	return DrawerFunc(func(pipe *gumi.Pipe, r gorat.VectorDrawer) {
		w, _ := r.Size()
		f := pipe.GetStyle(gumi.STYLE_Font).(textrat.Font)
		for p := float32(pivot); p <= w; p += pivot {
			txt := strconv.FormatInt(int64(p), 10)
			f.PathText(r, txt, gorat.Vec2(p, 0), gcore.AlignLeft | gcore.AlignTop)
		}
		r.Fill()
	})
}
