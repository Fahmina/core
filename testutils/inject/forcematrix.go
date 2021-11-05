package inject

import (
	"context"

	"go.viam.com/core/sensor"
	"go.viam.com/core/sensor/forcematrix"
)

// ForceMatrix is an injected ForceMatrix
type ForceMatrix struct {
	forcematrix.ForceMatrix
	MatrixFunc func(ctx context.Context) ([][]int, error)
}

// Matrix calls the injected MatrixFunc or the real variant
func (m *ForceMatrix) Matrix(ctx context.Context) ([][]int, error) {
	if m.MatrixFunc == nil {
		return m.ForceMatrix.Matrix(ctx)
	}
	return m.MatrixFunc(ctx)
}

// Desc returns that this is a force matrix.
func (m *ForceMatrix) Desc() sensor.Description {
	return sensor.Description{Type: forcematrix.Type}
}