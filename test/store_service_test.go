package gobs_test

import (
	"context"

	"github.com/traphamxuan/gobs"
)

type A struct{}

func (a *A) Init(ctx context.Context) (*gobs.ServiceLifeCycle, error) {
	return nil, nil
}

var _ gobs.IService = (*B)(nil)

type B struct {
	A *A
}

func (b *B) Init(ctx context.Context) (*gobs.ServiceLifeCycle, error) {
	return gobs.GenerateSetupConfig(nil, &b.A)
}

var _ gobs.IService = (*C)(nil)

type C struct {
	A *A
	B *B
}

func (c *C) Init(ctx context.Context) (*gobs.ServiceLifeCycle, error) {
	return gobs.GenerateSetupConfig(nil, &c.A, &c.B)
}

var _ gobs.IService = (*D)(nil)

type D struct {
	B *B
	C *C
}

func (d *D) Init(ctx context.Context) (*gobs.ServiceLifeCycle, error) {
	return gobs.GenerateSetupConfig(nil, &d.B, &d.C)
}
