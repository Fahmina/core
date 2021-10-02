package tmcstepper

import (
	"context"
	"math"
	"time"

	"github.com/go-errors/errors"
	"github.com/mitchellh/mapstructure"

	"go.viam.com/utils"

	"go.viam.com/core/board"
	"go.viam.com/core/config"
	"go.viam.com/core/motor"
	"go.viam.com/core/registry"
	"go.viam.com/core/robot"

	pb "go.viam.com/core/proto/api/v1"

	"github.com/edaniels/golog"

	"go.uber.org/multierr"
)

// TMC5072Config extends motor.Config, mainly for RegisterComponentAttributeMapConverter
type TMC5072Config struct {
	motor.Config
	SPIBus     string  `json:"spi_bus"`
	ChipSelect string  `json:"chip_select"`
	Index      int     `json:"index"`
	SGThresh   int32   `json:"sg_thresh"`
	CalFactor  float64 `json:"cal_factor"`
}

const (
	modelname = "TMC5072"
)

func init() {
	registry.RegisterMotor(modelname, registry.Motor{Constructor: func(ctx context.Context, r robot.Robot, config config.Component, logger golog.Logger) (motor.Motor, error) {
		m, err := NewMotor(ctx, r, config.ConvertedAttributes.(*TMC5072Config), logger)
		if err != nil {
			return nil, err
		}
		return m, nil
	}})

	config.RegisterComponentAttributeMapConverter(config.ComponentTypeMotor, modelname, func(attributes config.AttributeMap) (interface{}, error) {
		var conf TMC5072Config
		decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{TagName: "json", Squash: true, Result: &conf})
		if err != nil {
			return nil, err
		}
		if err := decoder.Decode(attributes); err != nil {
			return nil, err
		}
		return &conf, nil
	})
}

// A Motor represents a brushless motor connected via a TMC controller chip (ex: TMC5072)
type Motor struct {
	board       board.Board
	bus         board.SPI
	csPin       string
	index       int
	enPin       string
	stepsPerRev int
	maxRPM      float64
	maxAcc      float64
	fClk        float64
	logger      golog.Logger
}

// TMC5072 Values
const (
	baseClk = 13200000 // Nominal 13.2mhz internal clock speed
	uSteps  = 256      // Microsteps per fullstep
)

// TMC5072 Register Addressses (for motor index 0)
// TODO full register set
const (
	// add 0x10 for motor 2
	chopConf  = 0x6C
	coolConf  = 0x6D
	drvStatus = 0x6F

	// add 0x20 for motor 2
	rampMode   = 0x20
	xActual    = 0x21
	vActual    = 0x22
	vStart     = 0x23
	a1         = 0x24
	v1         = 0x25
	aMax       = 0x26
	vMax       = 0x27
	dMax       = 0x28
	d1         = 0x2A
	vStop      = 0x2B
	xTarget    = 0x2D
	iHoldIRun  = 0x30
	vCoolThres = 0x31
	swMode     = 0x34
	rampStat   = 0x35
)

// TMC5072 ramp modes
const (
	modePosition = int32(0)
	modeVelPos   = int32(1)
	modeVelNeg   = int32(2)
	modeHold     = int32(3)
)

// NewMotor returns a TMC5072 driven motor
func NewMotor(ctx context.Context, r robot.Robot, c *TMC5072Config, logger golog.Logger) (*Motor, error) {
	board, ok := r.BoardByName(c.BoardName)
	if !ok {
		return nil, errors.Errorf("can't find Board (%s) requested by Motor", c.BoardName)
	}

	bus, ok := board.SPIByName(c.SPIBus)
	if !ok {
		return nil, errors.Errorf("can't find SPI bus (%s) requested by Motor", c.SPIBus)
	}

	if c.CalFactor == 0 {
		c.CalFactor = 1.0
	}

	m := &Motor{
		board:       board,
		bus:         bus,
		csPin:       c.ChipSelect,
		index:       c.Index,
		stepsPerRev: c.TicksPerRotation * uSteps,
		enPin:       c.Pins["en"],
		maxRPM:      c.MaxRPM,
		maxAcc:      c.MaxAcceleration,
		fClk:        baseClk / c.CalFactor,
		logger:      logger,
	}

	rawMaxAcc := m.rpmsToA(m.maxAcc)

	if c.SGThresh > 63 {
		c.SGThresh = 63
	} else if c.SGThresh < -64 {
		c.SGThresh = -64
	}
	// The register is a 6 bit signed int
	if c.SGThresh < 0 {
		c.SGThresh = int32(64 + math.Abs(float64(c.SGThresh)))
	}

	coolConfig := c.SGThresh << 16

	err := multierr.Combine(
		m.writeReg(ctx, chopConf, 0x000100C3),  // TOFF=3, HSTRT=4, HEND=1, TBL=2, CHM=0 (spreadCycle)
		m.writeReg(ctx, iHoldIRun, 0x00080F0A), // IHOLD=8 (half current), IRUN=15 (max current), IHOLDDELAY=6
		m.writeReg(ctx, coolConf, coolConfig),  // Sets just the SGThreshold (for now)

		// Set max acceleration and decceleration
		m.writeReg(ctx, a1, rawMaxAcc),
		m.writeReg(ctx, aMax, rawMaxAcc),
		m.writeReg(ctx, d1, rawMaxAcc),
		m.writeReg(ctx, dMax, rawMaxAcc),

		m.writeReg(ctx, vStart, 1),                         // Always start at min speed
		m.writeReg(ctx, vStop, 10),                         // Always count a stop as LOW speed, but where vStop > vStart
		m.writeReg(ctx, v1, m.rpmToV(m.maxRPM/4)),          // Transition ramp at 25% speed (if d1 and a1 are set different)
		m.writeReg(ctx, vCoolThres, m.rpmToV(m.maxRPM/20)), // Set minimum speed for stall detection and coolstep
		m.writeReg(ctx, vMax, m.rpmToV(0)),                 // Max velocity to zero, we don't want to move

		m.writeReg(ctx, rampMode, modeVelPos), // Lastly, set velocity mode to force a stop in case chip was left in moving state
		m.writeReg(ctx, xActual, 0),           // Zero the position
	)

	if err != nil {
		return nil, err
	}

	return m, nil
}

func (m *Motor) shiftAddr(addr uint8) uint8 {
	// Shift register address for motor 1 instead of motor zero
	if m.index == 1 {
		if addr >= 0x10 && addr <= 0x11 {
			addr += 0x08
		} else if addr >= 0x20 && addr <= 0x3C {
			addr += 0x20
		} else if addr >= 0x6A && addr <= 0x6F {
			addr += 0x10
		}
	}
	return addr
}

func (m *Motor) writeReg(ctx context.Context, addr uint8, value int32) error {

	addr = m.shiftAddr(addr)

	var buf [5]byte
	buf[0] = addr | 0x80
	buf[1] = 0xFF & byte(value>>24)
	buf[2] = 0xFF & byte(value>>16)
	buf[3] = 0xFF & byte(value>>8)
	buf[4] = 0xFF & byte(value)

	handle, err := m.bus.OpenHandle()
	if err != nil {
		return err
	}
	defer func() {
		if err := handle.Close(); err != nil {
			m.logger.Error(err)
		}
	}()

	//m.logger.Debug("Write: ", buf)

	_, err = handle.Xfer(ctx, 1000000, m.csPin, 3, buf[:]) // SPI Mode 3, 1mhz
	if err != nil {
		return err
	}

	return nil
}

func (m *Motor) readReg(ctx context.Context, addr uint8) (int32, error) {

	addr = m.shiftAddr(addr)

	var tbuf [5]byte
	tbuf[0] = addr

	handle, err := m.bus.OpenHandle()
	if err != nil {
		return 0, err
	}
	defer func() {
		if err := handle.Close(); err != nil {
			m.logger.Error(err)
		}
	}()

	//m.logger.Debug("ReadT: ", tbuf)

	// Read access returns data from the address sent in the PREVIOUS "packet," so we transmit, then read
	_, err = handle.Xfer(ctx, 1000000, m.csPin, 3, tbuf[:]) // SPI Mode 3, 1mhz
	if err != nil {
		return 0, err
	}

	rbuf, err := handle.Xfer(ctx, 1000000, m.csPin, 3, tbuf[:])
	if err != nil {
		return 0, err
	}

	var value int32
	value = int32(rbuf[1])
	value <<= 8
	value |= int32(rbuf[2])
	value <<= 8
	value |= int32(rbuf[3])
	value <<= 8
	value |= int32(rbuf[4])

	//m.logger.Debug("ReadR: ", rbuf)
	//m.logger.Debug("Read: ", value)

	return value, nil

}

// GetSG returns the current StallGuard reading (effectively an indication of motor load.)
func (m *Motor) GetSG(ctx context.Context) (int32, error) {
	rawRead, err := m.readReg(ctx, drvStatus)
	if err != nil {
		return 0, err
	}

	rawRead &= 1023
	return rawRead, nil
}

// Position gives the current motor position
func (m *Motor) Position(ctx context.Context) (float64, error) {
	rawPos, err := m.readReg(ctx, xActual)
	if err != nil {
		return 0, err
	}
	return float64(rawPos) / float64(m.stepsPerRev), nil
}

// PositionSupported returns true.
func (m *Motor) PositionSupported(ctx context.Context) (bool, error) {
	return true, nil
}

// Power TODO (Should it be amps, not throttle?)
func (m *Motor) Power(ctx context.Context, powerPct float32) error {
	return errors.New("power not supported for stepper motors")
}

// Go sets a velocity as a percentage of maximum
func (m *Motor) Go(ctx context.Context, d pb.DirectionRelative, powerPct float32) error {
	mode := modeVelPos
	if d == pb.DirectionRelative_DIRECTION_RELATIVE_BACKWARD {
		mode = modeVelNeg
	}
	speed := m.rpmToV(float64(powerPct) * m.maxRPM)

	return multierr.Combine(
		m.writeReg(ctx, rampMode, mode),
		m.writeReg(ctx, vMax, speed),
	)
}

// GoFor turns in the given direction the given number of times at the given speed. Does not block.
func (m *Motor) GoFor(ctx context.Context, d pb.DirectionRelative, rpm float64, rotations float64) error {
	curPos, err := m.Position(ctx)
	if err != nil {
		return err
	}

	if d == pb.DirectionRelative_DIRECTION_RELATIVE_BACKWARD {
		rotations *= -1
	}
	target := curPos + rotations
	return m.GoTo(ctx, rpm, target)
}

// Convert rpm to TMC5072 steps/s
func (m *Motor) rpmToV(rpm float64) int32 {
	if rpm > m.maxRPM {
		rpm = m.maxRPM
	}
	// Time constant for velocities in TMC5072
	tConst := m.fClk / math.Pow(2, 24)
	speed := rpm / 60 * float64(m.stepsPerRev) / tConst
	return int32(speed)
}

// Convert rpm/s to TMC5072 steps/taConst^2
func (m *Motor) rpmsToA(acc float64) int32 {
	// Time constant for accelerations in TMC5072
	taConst := math.Pow(2, 41) / math.Pow(m.fClk, 2)
	rawMaxAcc := acc / 60 * float64(m.stepsPerRev) * taConst
	return int32(rawMaxAcc)
}

// GoTo moves to the specified position in terms of rotations.
func (m *Motor) GoTo(ctx context.Context, rpm float64, position float64) error {
	position *= float64(m.stepsPerRev)
	return multierr.Combine(
		m.writeReg(ctx, rampMode, modePosition),
		m.writeReg(ctx, vMax, m.rpmToV(rpm)),
		m.writeReg(ctx, xTarget, int32(position)),
	)
}

// IsOn returns true if the motor is currently moving.
func (m *Motor) IsOn(ctx context.Context) (bool, error) {
	vel, err := m.readReg(ctx, vActual)
	on := vel != 0
	return on, err
}

// Enable pulls down the hardware enable pin, activating the power stage of the chip
func (m *Motor) Enable(ctx context.Context, turnOn bool) error {
	return m.board.GPIOSet(ctx, m.enPin, !turnOn)
}

// Off stops the motor.
func (m *Motor) Off(ctx context.Context) error {
	return m.Go(ctx, pb.DirectionRelative_DIRECTION_RELATIVE_FORWARD, 0)
}

// GoTillStop enables StallGuard detection, then moves in the direction/speed given until resistance (endstop) is detected.
func (m *Motor) GoTillStop(ctx context.Context, d pb.DirectionRelative, rpm float64, stopFunc func(ctx context.Context) bool) error {
	if err := m.GoFor(ctx, d, rpm, 1000); err != nil {
		return err
	}

	// Disable stallguard and turn off if we fail homing
	defer func() {
		if err := multierr.Combine(
			m.writeReg(ctx, swMode, 0x000),
			m.Off(ctx),
		); err != nil {
			m.logger.Error(err)
		}
	}()

	// Get up to speed
	var fails int
	for {
		// sg, _ := m.GetSG(ctx)
		// m.logger.Debugf("SGValueSpeed: %d", sg)

		if !utils.SelectContextOrWait(ctx, 10*time.Millisecond) {
			return errors.New("context cancelled during GoTillStop")
		}

		if stopFunc != nil && stopFunc(ctx) {
			return nil
		}

		stat, err := m.readReg(ctx, rampStat)
		if err != nil {
			return err
		}
		// Look for velocity_reached flag
		if stat&0x100 == 0x100 {
			break
		}

		if fails >= 500 {
			return errors.New("timed out during GoTillStop acceleration")
		}
		fails++
	}

	// Now enable stallguard
	if err := m.writeReg(ctx, swMode, 0x400); err != nil {
		return err
	}

	// Wait for motion to stop at endstop
	fails = 0
	for {
		// sg, _ := m.GetSG(ctx)
		// m.logger.Debugf("SGValueReady: %d", sg)

		if !utils.SelectContextOrWait(ctx, 10*time.Millisecond) {
			return errors.New("context cancelled during GoTillStop")
		}

		if stopFunc != nil && stopFunc(ctx) {
			return nil
		}

		stat, err := m.readReg(ctx, rampStat)
		if err != nil {
			return err
		}
		// Look for vzero flag
		if stat&0x400 == 0x400 {
			break
		}

		if fails >= 1000 {
			return errors.New("timed out during GoTillStop")
		}
		fails++
	}

	return nil
}

// Zero resets the current position to the offset given.
func (m *Motor) Zero(ctx context.Context, offset float64) error {
	on, err := m.IsOn(ctx)
	if err != nil {
		return err
	} else if on {
		return errors.New("can't zero while moving")
	}
	return multierr.Combine(
		m.writeReg(ctx, rampMode, modeHold),
		m.writeReg(ctx, xTarget, int32(offset*float64(m.stepsPerRev))),
		m.writeReg(ctx, xActual, int32(offset*float64(m.stepsPerRev))),
	)
}