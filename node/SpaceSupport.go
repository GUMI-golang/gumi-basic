package node

import "github.com/GUMI-golang/gumi"

const Namespace = "node"
func NodeSpace(space *gumi.Space) error {
	var err error
	err = space.Support(Namespace, &gumi.SupportTag{
		Name:"style",
		New: func() gumi.GUMI {
			return Style0()
		},
	})
	if err != nil {
		return err
	}
	err = space.Support(Namespace, &gumi.SupportTag{
		Name:"size",
		New: func() gumi.GUMI {
			return Size0()
		},
	})
	if err != nil {
		return err
	}
	err = space.Support(Namespace, &gumi.SupportTag{
		Name:"background",
		New: func() gumi.GUMI {
			return Background0()
		},
	})
	if err != nil {
		return err
	}
	return nil
}