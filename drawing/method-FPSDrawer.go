package drawing

import (
	"fmt"
	"github.com/GUMI-golang/gorat"
	"github.com/GUMI-golang/gumi"
	"math"
	"github.com/GUMI-golang/gorat/textrat"
	"github.com/GUMI-golang/gumi/gcore"
)

const (
	fpsDrawerHistory = 32
	fpsPos           = 12
)

func (_Drawing) FPS(prefix string) Drawer {
	return &fpsDrawer{
		prefix: prefix,
	}
}

type fpsDrawer struct {
	prefix string
	dts    [fpsDrawerHistory]float32
	i      int
	c      int
}

func (s *fpsDrawer) Draw(pipe *gumi.Pipe, r gorat.VectorDrawer) {
	var avg = clamp(average(s.dts[:s.c]...), 0.0001, math.MaxFloat32)
	txt := fmt.Sprintf("%s %.2f - %2.5f", s.prefix, 1000/float64(avg), avg)
	w, _ := r.Size()
	// TODO
	//f := pipe.GetStyle(gumi.STYLE_Font).(textrat.Font)
	//f.Text(r, txt, gorat.Vec2(w-fpsPos, fpsPos), gcore.AlignRight|gcore.AlignTop)
	textrat.Default.Text(r, txt, gorat.Vec2(w-fpsPos, fpsPos), gcore.AlignRight|gcore.AlignTop)
}
func (s *fpsDrawer) Informate(tick gumi.EventTick) bool {
	s.dts[s.i] = float32(tick.DeltaT.Seconds())
	s.i = (s.i + 1) % fpsDrawerHistory
	if s.c < fpsDrawerHistory {
		s.c += 1
	}
	if s.i == 0 {
		return true
	}
	return false
}

func clamp(i, min, max float32) float32 {
	if i < min {
		return min
	}
	if i > max {
		return max
	}
	return i
}
func sum(v ...float32) (res float32) {
	for _, f := range v {
		res += f
	}
	return
}
func average(v ...float32) float32 {
	return sum(v...) / float32(len(v))
}