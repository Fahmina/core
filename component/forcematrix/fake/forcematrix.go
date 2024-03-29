// Package fake implements a fake ForceMatrix
package fake

import (
	"context"

	"github.com/edaniels/golog"
	"github.com/pkg/errors"

	"go.viam.com/rdk/component/forcematrix"
	"go.viam.com/rdk/config"
	"go.viam.com/rdk/registry"
	"go.viam.com/rdk/robot"
)

func init() {
	registry.RegisterComponent(forcematrix.Subtype, "fake", registry.Component{
		Constructor: func(
			ctx context.Context,
			r robot.Robot,
			config config.Component,
			logger golog.Logger,
		) (interface{}, error) {
			if config.Attributes.Bool("fail_new", false) {
				return nil, errors.New("whoops")
			}
			return NewForceMatrix(config)
		},
	})
}

// NewForceMatrix returns a new fake ForceMatrix.
func NewForceMatrix(cfg config.Component) (forcematrix.ForceMatrix, error) {
	name := cfg.Name
	return &ForceMatrix{
		Name: name,
	}, nil
}

// ForceMatrix is a fake ForceMatrix that always returns the same matrix of values.
type ForceMatrix struct {
	Name string
}

// ReadMatrix always returns the same matrix.
func (fsm *ForceMatrix) ReadMatrix(ctx context.Context) ([][]int, error) {
	result := make([][]int, 4)
	for i := 0; i < len(result); i++ {
		result[i] = []int{1, 1, 1, 1}
	}
	return result, nil
}

// DetectSlip always return false.
func (fsm *ForceMatrix) DetectSlip(ctx context.Context) (bool, error) {
	return false, nil
}
