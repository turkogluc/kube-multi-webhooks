package v1

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
)

type CustomDefaulter interface {
	Default(ctx context.Context, obj runtime.Object) error
}

type customDef struct {
	name string
}

func (c *customDef) Default(_ context.Context, obj runtime.Object) error {
	fmt.Println("### DEF", obj, c.name)
	return nil
}
