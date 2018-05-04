package layout

import "github.com/GUMI-golang/gumi"

const Namespace = "layout"
func LayoutSpace(space *gumi.Space) error {
	var err error
	err = space.Support(Namespace, &gumi.SupportTag{
		Name:"center",
		New: func() gumi.GUMI {
			return Center0()
		},
	})
	if err != nil {
		return err
	}
	return nil
}