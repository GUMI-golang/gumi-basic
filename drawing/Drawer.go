package drawing

import (
	"github.com/GUMI-golang/gumi"
	"github.com/GUMI-golang/gorat"
)

type Drawer interface {
	Draw(pipe *gumi.Pipe, r gorat.VectorDrawer)
}
type ImfomaticDrawer interface {
	Drawer
	Informate (tick gumi.EventTick) bool
}
type DrawerFunc func(pipe *gumi.Pipe, r gorat.VectorDrawer)
func (s DrawerFunc) Draw(pipe *gumi.Pipe, r gorat.VectorDrawer){
	s(pipe, r)
}
