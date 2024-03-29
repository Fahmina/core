// Package register registers all relevant arms
package register

import (

	// register arms.
	_ "go.viam.com/rdk/component/arm/eva"
	_ "go.viam.com/rdk/component/arm/fake"
	_ "go.viam.com/rdk/component/arm/universalrobots"
	_ "go.viam.com/rdk/component/arm/varm"
	_ "go.viam.com/rdk/component/arm/vx300s"
	_ "go.viam.com/rdk/component/arm/wx250s"
	_ "go.viam.com/rdk/component/arm/xarm"
	_ "go.viam.com/rdk/component/arm/yahboom"
)
