package gobs

import (
	"context"
	"fmt"
	"reflect"

	"github.com/traphamxuan/gobs/common"
)

func GenerateSetupConfig(asyncMode map[common.ServiceStatus]bool, deps ...any) (*ServiceLifeCycle, error) {
	dependencies := make(Dependencies, len(deps))

	for i, ptr := range deps {
		if ptr == nil {
			return nil, fmt.Errorf("Cannot fill nil pointer. %w", common.ErrorInvalidType)
		}
		ptrType := reflect.TypeOf(ptr)
		// Make sure the prt is address to a pointer
		if ptrType.Kind() != reflect.Ptr {
			return nil, fmt.Errorf("Cannot fill non-pointer value. %w", common.ErrorInvalidType)
		}

		ptrTypeElem := ptrType.Elem()
		// Ensure the type implements IService before creating a new instance
		if !ptrTypeElem.Implements(reflect.TypeOf((*IService)(nil)).Elem()) {
			return nil, fmt.Errorf("%s does not implement IService interface %w",
				ptrTypeElem.String(), common.ErrorInvalidType)
		}

		// Create a new instance of the struct type that ptr points to
		newInstance := reflect.New(ptrTypeElem.Elem())

		// Initialize the new instance if it has an Init method
		if initable, ok := newInstance.Interface().(IService); ok {
			dependencies[i] = initable
		} else {
			return nil, fmt.Errorf("Cannot convert %s to IService instance %w",
				newInstance.Type().String(), common.ErrorInvalidType)
		}
	}
	onSetup := func(ctx context.Context, d Dependencies) error {
		return d.Assign(deps...)
	}

	return &ServiceLifeCycle{
		Deps:      dependencies,
		OnSetup:   onSetup,
		AsyncMode: asyncMode,
	}, nil
}
