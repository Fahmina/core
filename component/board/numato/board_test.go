package numato

import (
	"context"
	"errors"
	"testing"

	"github.com/edaniels/golog"
	"go.viam.com/test"
	"go.viam.com/utils"

	"go.viam.com/rdk/component/board"
	"go.viam.com/rdk/config"
)

func TestMask(t *testing.T) {
	m := newMask(32)
	test.That(t, len(m), test.ShouldEqual, 4)
	test.That(t, m.hex(), test.ShouldEqual, "00000000")

	m.set(0)
	test.That(t, m.hex(), test.ShouldEqual, "00000001")

	m.set(6)
	m.set(7)
	test.That(t, m.hex(), test.ShouldEqual, "000000c1")

	m.set(31)
	test.That(t, m.hex(), test.ShouldEqual, "800000c1")
}

func TestFixPins(t *testing.T) {
	test.That(t, fixPin(128, "0"), test.ShouldEqual, "000")
	test.That(t, fixPin(128, "00"), test.ShouldEqual, "000")
	test.That(t, fixPin(128, "000"), test.ShouldEqual, "000")

	test.That(t, fixPin(128, "1"), test.ShouldEqual, "001")
	test.That(t, fixPin(128, "01"), test.ShouldEqual, "001")
	test.That(t, fixPin(128, "001"), test.ShouldEqual, "001")
}

func TestNumato1(t *testing.T) {
	ctx := context.Background()
	logger := golog.NewTestLogger(t)
	b, err := connect(
		ctx,
		&board.Config{
			Attributes: config.AttributeMap{"pins": 128},
			Analogs:    []board.AnalogConfig{{Name: "foo", Pin: "01"}},
		},
		logger,
	)
	if errors.Is(err, errNoBoard) {
		t.Skip("no numato board connected")
	}
	test.That(t, err, test.ShouldBeNil)
	defer utils.TryClose(ctx, b)

	// For this to work 0 has be plugged into 1

	zeroPin, err := b.GPIOPinByName("0")
	test.That(t, err, test.ShouldBeNil)
	onePin, err := b.GPIOPinByName("1")
	test.That(t, err, test.ShouldBeNil)

	// set to low
	err = zeroPin.Set(context.Background(), false)
	test.That(t, err, test.ShouldBeNil)

	res, err := onePin.Get(ctx)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, res, test.ShouldEqual, false)

	// set to high
	err = zeroPin.Set(context.Background(), true)
	test.That(t, err, test.ShouldBeNil)

	res, err = onePin.Get(ctx)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, res, test.ShouldEqual, true)

	// set back to low
	err = zeroPin.Set(context.Background(), false)
	test.That(t, err, test.ShouldBeNil)

	res, err = onePin.Get(ctx)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, res, test.ShouldEqual, false)

	// test analog
	ar, ok := b.AnalogReaderByName("foo")
	test.That(t, ok, test.ShouldEqual, true)

	res2, err := ar.Read(ctx)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, res2, test.ShouldBeLessThan, 100)

	err = zeroPin.Set(context.Background(), true)
	test.That(t, err, test.ShouldBeNil)

	res2, err = ar.Read(ctx)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, res2, test.ShouldBeGreaterThan, 1000)

	err = zeroPin.Set(context.Background(), false)
	test.That(t, err, test.ShouldBeNil)

	res2, err = ar.Read(ctx)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, res2, test.ShouldBeLessThan, 100)
}
