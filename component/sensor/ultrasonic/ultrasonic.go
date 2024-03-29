// Package ultrasonic implements an ultrasonic sensor based of the yahboom ultrasonic sensor
package ultrasonic

import (
	"context"
	"time"

	"github.com/edaniels/golog"
	"github.com/pkg/errors"
	rdkutils "go.viam.com/utils"

	"go.viam.com/rdk/component/board"
	"go.viam.com/rdk/component/sensor"
	"go.viam.com/rdk/config"
	"go.viam.com/rdk/registry"
	"go.viam.com/rdk/robot"
)

const (
	modelname = "ultrasonic"
)

// AttrConfig is used for converting config attributes.
type AttrConfig struct {
	TriggerPin string `json:"trigger_pin"`
	EchoInt    string `json:"echo_int"`
	Board      string `json:"board"`
}

func init() {
	registry.RegisterComponent(
		sensor.Subtype,
		modelname,
		registry.Component{Constructor: func(
			ctx context.Context,
			r robot.Robot,
			config config.Component,
			logger golog.Logger,
		) (interface{}, error) {
			return newSensor(ctx, r, config.Name, config.ConvertedAttributes.(*AttrConfig))
		}})

	config.RegisterComponentAttributeMapConverter(sensor.SubtypeName, modelname,
		func(attributes config.AttributeMap) (interface{}, error) {
			var conf AttrConfig
			return config.TransformAttributeMapToStruct(&conf, attributes)
		}, &AttrConfig{})
}

func newSensor(ctx context.Context, r robot.Robot, name string, config *AttrConfig) (sensor.Sensor, error) {
	r.Logger().Debug("building ultrasonic sensor")
	s := &Sensor{Name: name, config: config}
	b, err := board.FromRobot(r, config.Board)
	if err != nil {
		return nil, errors.Errorf("ultrasonic : cannot find board %q", config.Board)
	}
	i, ok := b.DigitalInterruptByName(config.EchoInt)
	if !ok {
		return nil, errors.Errorf("ultrasonic: cannot grab digital interrupt %q", config.EchoInt)
	}
	g, err := b.GPIOPinByName(config.TriggerPin)
	if err != nil {
		return nil, errors.Wrapf(err, "ultrasonic: cannot grab gpio %q", config.TriggerPin)
	}
	s.echoInt = i
	s.triggerPin = g
	if err := s.triggerPin.Set(ctx, false); err != nil {
		return nil, errors.Wrap(err, "ultrasonic: cannot set trigger pin to low")
	}
	s.intChan = make(chan bool)
	s.echoInt.AddCallback(s.intChan)
	return s, nil
}

// Sensor ultrasonic sensor.
type Sensor struct {
	Name       string
	config     *AttrConfig
	echoInt    board.DigitalInterrupt
	triggerPin board.GPIOPin
	intChan    chan bool
}

// GetReadings returns the calculated distance.
func (s *Sensor) GetReadings(ctx context.Context) ([]interface{}, error) {
	if err := s.triggerPin.Set(ctx, true); err != nil {
		return nil, errors.Wrap(err, "ultrasonic cannot set trigger pin to high")
	}

	rdkutils.SelectContextOrWait(ctx, time.Microsecond*20)
	if err := s.triggerPin.Set(ctx, false); err != nil {
		return nil, errors.Wrap(err, "ultrasonic cannot set trigger pin to low")
	}
	var timeA, timeB time.Time
	select {
	case <-s.intChan:
		timeB = time.Now()
	case <-ctx.Done():
		return nil, errors.New("ultrasonic: context canceled")
	case <-time.After(time.Second * 1):
		return nil, errors.New("ultrasonic timeout")
	}
	select {
	case <-s.intChan:
		timeA = time.Now()
	case <-ctx.Done():
		return nil, errors.New("ultrasonic: context canceled")
	case <-time.After(time.Second * 1):
		return nil, errors.New("ultrasonic timeout")
	}
	dist := timeA.Sub(timeB).Seconds() * 340 / 2
	return []interface{}{dist}, nil
}
