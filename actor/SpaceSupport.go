package actor

import (
	"github.com/GUMI-golang/gumi"
)

const Namespace = "actor"
func ActorSpace(space *gumi.Space) error {
	var err error
	err = space.Support(Namespace, &gumi.SupportTag{
		Name:"empty",
		New: func() gumi.GUMI {
			return Empty0()
		},
	})
	if err != nil {
		return err
	}
	err = space.Support(Namespace, &gumi.SupportTag{
		Name:"text",
		New: func() gumi.GUMI {
			return Text0()
		},
	})
	if err != nil {
		return err
	}
	err = space.Support(Namespace, &gumi.SupportTag{
		Name:"image",
		New: func() gumi.GUMI {
			return Image0()
		},
	})
	if err != nil {
		return err
	}
	return nil
}