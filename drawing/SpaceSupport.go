package drawing

import (
	"github.com/GUMI-golang/gumi"
)

const Namespace = "drawing"
func DrawingSpace(space *gumi.Space) error {
	var err error
	err = space.Support(Namespace, &gumi.SupportTag{
		Name:"draw",
		New: func() gumi.GUMI {
			return new(Draw)
		},
	})
	if err != nil {
		return err
	}
	return nil
}