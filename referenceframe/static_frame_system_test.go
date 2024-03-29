package referenceframe

import (
	"errors"
	"math"
	"testing"

	"github.com/golang/geo/r3"
	"go.viam.com/test"

	spatial "go.viam.com/rdk/spatialmath"
)

var blankPos map[string][]Input

func frameNames(frames []Frame) []string {
	names := make([]string, len(frames))
	for i, f := range frames {
		names[i] = f.Name()
	}
	return names
}

func TestSimpleFrameSystemFunctions(t *testing.T) {
	// build the system
	fs := NewEmptySimpleFrameSystem("test")
	frame3Pt := r3.Vector{0., 4., 0.} // location of frame3 with respect to world frame
	f3, err := FrameFromPoint("frame3", frame3Pt)
	test.That(t, err, test.ShouldBeNil)
	err = fs.AddFrame(f3, fs.World())
	test.That(t, err, test.ShouldBeNil)
	frame1Pt := r3.Vector{0., 3., 0.} // location of frame1 with respect to frame3
	f1, err := FrameFromPoint("frame1", frame1Pt)
	test.That(t, err, test.ShouldBeNil)
	err = fs.AddFrame(f1, fs.GetFrame("frame3"))
	test.That(t, err, test.ShouldBeNil)
	frame2Pt := r3.Vector{5., 1., 0.} // location of frame2 with respect to world frame
	f2, err := FrameFromPoint("frame2", frame2Pt)
	test.That(t, err, test.ShouldBeNil)
	err = fs.AddFrame(f2, fs.World())
	test.That(t, err, test.ShouldBeNil)

	frames := fs.FrameNames()
	test.That(t, len(frames), test.ShouldEqual, 3)

	f1Parents, err := fs.TracebackFrame(f1)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, len(f1Parents), test.ShouldEqual, 3)
	test.That(t, frameNames(f1Parents), test.ShouldResemble, []string{"frame1", "frame3", "world"})

	parent, err := fs.Parent(fs.GetFrame("frame1"))
	test.That(t, err, test.ShouldBeNil)
	test.That(t, parent, test.ShouldResemble, fs.GetFrame("frame3"))

	parent, err = fs.Parent(fs.GetFrame("frame3"))
	test.That(t, err, test.ShouldBeNil)
	test.That(t, parent, test.ShouldResemble, fs.World())
	// Pruning frame3 should also remove frame1
	fs.RemoveFrame(fs.GetFrame("frame3"))
	frames = fs.FrameNames()
	test.That(t, len(frames), test.ShouldEqual, 1)

	e2 := errors.New("parent frame with name \"foo\" not in frame system")
	e3 := errors.New("frame with name \"frame2\" already in frame system")
	e4 := errors.New("frame with name \"frame1\" not in frame system")
	if sfs, ok := fs.(*simpleFrameSystem); ok {
		err = sfs.AddFrame(f2, nil)
		test.That(t, err, test.ShouldBeError, NewParentFrameMissingError())

		fFoo := NewZeroStaticFrame("foo")
		err = sfs.checkName("bar", fFoo)
		test.That(t, err, test.ShouldBeError, e2)

		err = sfs.checkName("frame2", fs.World())
		test.That(t, err, test.ShouldBeError, e3)
	}

	_, err = fs.TracebackFrame(f1)
	test.That(t, err, test.ShouldBeError, e4)
}

// A simple Frame translation from the world frame to a frame right above it at (0, 3, 0)
// And then back to the world frame
// transforming a point at (1, 3, 0).
func TestSimpleFrameTranslation(t *testing.T) {
	// build the system
	fs := NewEmptySimpleFrameSystem("test")
	frame, err := FrameFromPoint("frame", r3.Vector{0., 3., 0.}) // location of frame with respect to world frame
	test.That(t, err, test.ShouldBeNil)
	err = fs.AddFrame(frame, fs.World())
	test.That(t, err, test.ShouldBeNil)

	// do the transformation
	pointWorld := r3.Vector{1., 3., 0.} // the point from PoV of world
	pointFrame := r3.Vector{1., 0., 0.} // the point from PoV of frame

	// transform point from world to frame
	transformPoint1, err := fs.TransformPoint(blankPos, pointWorld, World, "frame")
	test.That(t, err, test.ShouldBeNil)
	test.That(t, transformPoint1, test.ShouldResemble, pointFrame)

	// transform point from frame to world
	transformPoint2, err := fs.TransformPoint(blankPos, pointFrame, "frame", World)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, transformPoint2, test.ShouldResemble, pointWorld)
}

// A simple Frame translation from the world frame to a frame right above it at (0, 3, 0) rotated 180 around Z
// And then back to the world frame
// transforming a point at (1, 3, 0).
func TestSimpleFrameTranslationWithRotation(t *testing.T) {
	// build the system
	fs := NewEmptySimpleFrameSystem("test")
	framePose := spatial.NewPoseFromAxisAngle(r3.Vector{0., 3., 0.}, r3.Vector{0., 0., 1.}, math.Pi)
	f1, err := NewStaticFrame("frame", framePose)
	test.That(t, err, test.ShouldBeNil)
	err = fs.AddFrame(f1, fs.World())
	test.That(t, err, test.ShouldBeNil)

	// define the point coordinates
	pointWorld := r3.Vector{1., 3., 0.}  // the point from PoV of world
	pointFrame := r3.Vector{-1., 0., 0.} // the point from PoV of frame

	// transform point from world to frame
	transformPoint1, err := fs.TransformPoint(blankPos, pointWorld, World, "frame")
	test.That(t, err, test.ShouldBeNil)
	test.That(t, transformPoint1.X, test.ShouldAlmostEqual, pointFrame.X)
	test.That(t, transformPoint1.Y, test.ShouldAlmostEqual, pointFrame.Y)
	test.That(t, transformPoint1.Z, test.ShouldAlmostEqual, pointFrame.Z)

	// transform point from frame to world
	transformPoint2, err := fs.TransformPoint(blankPos, pointFrame, "frame", World)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, transformPoint2, test.ShouldResemble, pointWorld)
}

/*
Transforms the pose of *object from *frame1 into *frame2. The Orientation of *frame1 and *frame2
are the same, so the final composed transformation is only made up of one translation.

|              |
|*frame1       |*object
|              |
|
|*frame3
|
|
|              *frame2
|________________
world

transform the point that is in the world frame is at (5, 7, 0) from frame1 to frame2.
frame1 has its origin at (0, 7, 0) in the world referenceframe. and frame2 has its origin at (5, 1, 0).
frame3 is an intermediate frame at (0, 4, 0) in the world referenceframe.
All 4 frames have the same orientation.
*/
func TestFrameTranslation(t *testing.T) {
	// build the system
	fs := NewEmptySimpleFrameSystem("test")
	frame3Pt := r3.Vector{0., 4., 0.} // location of frame3 with respect to world frame
	f3, err := FrameFromPoint("frame3", frame3Pt)
	test.That(t, err, test.ShouldBeNil)
	err = fs.AddFrame(f3, fs.World())
	test.That(t, err, test.ShouldBeNil)
	frame1Pt := r3.Vector{0., 3., 0.} // location of frame1 with respect to frame3
	f1, err := FrameFromPoint("frame1", frame1Pt)
	test.That(t, err, test.ShouldBeNil)
	err = fs.AddFrame(f1, fs.GetFrame("frame3"))
	test.That(t, err, test.ShouldBeNil)
	frame2Pt := r3.Vector{5., 1., 0.} // location of frame2 with respect to world frame
	f2, err := FrameFromPoint("frame2", frame2Pt)
	test.That(t, err, test.ShouldBeNil)
	err = fs.AddFrame(f2, fs.World())
	test.That(t, err, test.ShouldBeNil)

	// do the transformation
	pointFrame1 := r3.Vector{5., 0., 0.} // the point from PoV of frame 1
	pointFrame2 := r3.Vector{0., 6., 0.} // the point from PoV of frame 2
	transformPoint, err := fs.TransformPoint(blankPos, pointFrame1, "frame1", "frame2")
	test.That(t, err, test.ShouldBeNil)
	test.That(t, transformPoint, test.ShouldResemble, pointFrame2)
}

/*
Very similar to test above, but this time *frame2 is oriented 90 degrees counterclockwise around the z-axis
(+z is pointing out of the page), which means the orientation of the object will also change.

|              |
|*frame1       |*object
|              |
|
|*frame3
|              |
|              |
|          ____|*frame2
|________________
world
*/

// transform the point that is in the world frame is at (5, 7, 0) from frame1 to frame2.
// frame1 has its origin at (0, 7, 0) in the world referenceframe. and frame2 has its origin
// at (5, 1, 0), and orientation 90 degrees around z.
// frame3 is an intermediate frame at (0, 4, 0) in the world referenceframe.
func TestFrameTransform(t *testing.T) {
	// build the system
	fs := NewEmptySimpleFrameSystem("test")
	// location of frame3 with respect to world frame
	frame3Pt := r3.Vector{0., 4., 0.} // location of frame3 with respect to world frame
	f3, err := FrameFromPoint("frame3", frame3Pt)
	test.That(t, err, test.ShouldBeNil)
	err = fs.AddFrame(f3, fs.World())
	test.That(t, err, test.ShouldBeNil)
	frame1Pt := r3.Vector{0., 3., 0.} // location of frame1 with respect to frame3
	f1, err := FrameFromPoint("frame1", frame1Pt)
	test.That(t, err, test.ShouldBeNil)
	err = fs.AddFrame(f1, fs.GetFrame("frame3"))
	test.That(t, err, test.ShouldBeNil)
	frame2Pose := spatial.NewPoseFromAxisAngle(r3.Vector{5., 1., 0.}, r3.Vector{0., 0., 1.}, math.Pi/2)
	f2, err := NewStaticFrame("frame2", frame2Pose)
	test.That(t, err, test.ShouldBeNil)
	err = fs.AddFrame(f2, fs.World())
	test.That(t, err, test.ShouldBeNil)

	// do the transformation
	pointFrame1 := r3.Vector{5., 0., 0.} // the point from PoV of frame 1
	pointFrame2 := r3.Vector{6., 0., 0.} // the point from PoV of frame 2
	transformPoint, err := fs.TransformPoint(blankPos, pointFrame1, "frame1", "frame2")
	test.That(t, err, test.ShouldBeNil)
	test.That(t, transformPoint.X, test.ShouldAlmostEqual, pointFrame2.X)
	test.That(t, transformPoint.Y, test.ShouldAlmostEqual, pointFrame2.Y)
	test.That(t, transformPoint.Z, test.ShouldAlmostEqual, pointFrame2.Z)
}

/*
This test uses the same setup as the above test, but this time is concede with representing a geometry in a difference reference frame

|              |
|*frame1       |*object
|              |
|
|*frame3
|              |
|              |
|          ____|*frame2
|________________
world
*/

// transform the object geometry that is in the world frame is at (5, 7, 0) from frame1 to frame2.
// frame1 has its origin at (0, 7, 0) in the world referenceframe. and frame2 has its origin
// at (5, 1, 0), and orientation 90 degrees around z.
// frame3 is an intermediate frame at (0, 4, 0) in the world referenceframe.
func TestGeometriesOfFrame(t *testing.T) {
	// build the system
	fs := NewEmptySimpleFrameSystem("test")
	// location of frame3 with respect to world frame
	frame3Pt := r3.Vector{0., 4., 0.} // location of frame3 with respect to world frame
	f3, err := FrameFromPoint("frame3", frame3Pt)
	test.That(t, err, test.ShouldBeNil)
	err = fs.AddFrame(f3, fs.World())
	test.That(t, err, test.ShouldBeNil)
	frame1Pt := r3.Vector{0., 3., 0.} // location of frame1 with respect to frame3
	f1, err := FrameFromPoint("frame1", frame1Pt)
	test.That(t, err, test.ShouldBeNil)
	err = fs.AddFrame(f1, fs.GetFrame("frame3"))
	test.That(t, err, test.ShouldBeNil)
	frame2Pose := spatial.NewPoseFromAxisAngle(r3.Vector{5., 1., 0.}, r3.Vector{0., 0., 1.}, math.Pi/2)
	f2, err := NewStaticFrame("frame2", frame2Pose)
	test.That(t, err, test.ShouldBeNil)
	err = fs.AddFrame(f2, fs.World())
	test.That(t, err, test.ShouldBeNil)
	objectFromFrame1 := r3.Vector{5., 0., 0.}
	gc, err := spatial.NewBoxCreator(r3.Vector{2, 2, 2}, spatial.NewZeroPose())
	test.That(t, err, test.ShouldBeNil)
	object, err := NewStaticFrameWithGeometry("object", spatial.NewPoseFromPoint(objectFromFrame1), gc)

	test.That(t, err, test.ShouldBeNil)
	err = fs.AddFrame(object, f1)
	test.That(t, err, test.ShouldBeNil)

	objectFromFrame2 := r3.Vector{6., 0., 0.} // the point from PoV of frame 2
	geometries, _ := fs.GeometriesOfFrame(blankPos, "object", "frame2")
	test.That(t, geometries, test.ShouldNotBeNil)
	test.That(t, spatial.R3VectorAlmostEqual(geometries["object"].Pose().Point(), objectFromFrame2, 1e-8), test.ShouldBeTrue)
}

func TestComplicatedFrameTransform(t *testing.T) {
	// build the system
	fs := NewEmptySimpleFrameSystem("test")

	// frame 1 rotate by 45 degrees around z axis and translate
	frame1, err := NewStaticFrame("frame1", spatial.NewPoseFromAxisAngle(r3.Vector{-1., 2., 0.}, r3.Vector{0., 0., 1.}, math.Pi/4))
	test.That(t, err, test.ShouldBeNil)
	err = fs.AddFrame(frame1, fs.World())
	test.That(t, err, test.ShouldBeNil)
	// frame 2 rotate by 45 degree (relative to frame 1) around z axis and translate
	frame2, err := NewStaticFrame("frame2",
		spatial.NewPoseFromAxisAngle(r3.Vector{2. * math.Sqrt(2), 0., 0.}, r3.Vector{0., 0., 1.}, math.Pi/4))
	test.That(t, err, test.ShouldBeNil)
	err = fs.AddFrame(frame2, fs.GetFrame("frame1"))
	test.That(t, err, test.ShouldBeNil)

	// test out a transform from world to frame
	pointStart := r3.Vector{1., 7., 0.} // the point from PoV of world
	pointEnd := r3.Vector{3., 0., 0.}   // the point from PoV of frame 2
	transformPoint, err := fs.TransformPoint(blankPos, pointStart, World, "frame2")
	test.That(t, err, test.ShouldBeNil)
	test.That(t, transformPoint.X, test.ShouldAlmostEqual, pointEnd.X)
	test.That(t, transformPoint.Y, test.ShouldAlmostEqual, pointEnd.Y)
	test.That(t, transformPoint.Z, test.ShouldAlmostEqual, pointEnd.Z)

	// test out transform between frames
	// frame3 - pure rotation around y 90 degrees
	frame3, err := NewStaticFrame("frame3", spatial.NewPoseFromAxisAngle(r3.Vector{}, r3.Vector{0., 1., 0.}, math.Pi/2))
	test.That(t, err, test.ShouldBeNil)
	err = fs.AddFrame(frame3, fs.World())
	test.That(t, err, test.ShouldBeNil)

	// frame4 - pure translation
	frame4, err := NewStaticFrame("frame4", spatial.NewPoseFromPoint(r3.Vector{-2., 7., 1.}))
	test.That(t, err, test.ShouldBeNil)
	err = fs.AddFrame(frame4, fs.GetFrame("frame3"))
	test.That(t, err, test.ShouldBeNil)

	pointStart = r3.Vector{3., 0., 0.} // the point from PoV of frame 2
	pointEnd = r3.Vector{2., 0., 0.}   // the point from PoV of frame 4
	transformPoint, err = fs.TransformPoint(blankPos, pointStart, "frame2", "frame4")
	test.That(t, err, test.ShouldBeNil)
	test.That(t, transformPoint.X, test.ShouldAlmostEqual, pointEnd.X)
	test.That(t, transformPoint.Y, test.ShouldAlmostEqual, pointEnd.Y)
	test.That(t, transformPoint.Z, test.ShouldAlmostEqual, pointEnd.Z)

	// back to world frame
	pointStart = r3.Vector{2., 0., 0.} // the point from PoV of frame 4
	pointEnd = r3.Vector{1., 7., 0.}   // the point from PoV of world
	transformPoint, err = fs.TransformPoint(blankPos, pointStart, "frame4", World)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, transformPoint.X, test.ShouldAlmostEqual, pointEnd.X)
	test.That(t, transformPoint.Y, test.ShouldAlmostEqual, pointEnd.Y)
	test.That(t, transformPoint.Z, test.ShouldAlmostEqual, pointEnd.Z)
}

func TestSystemSplitAndRejoin(t *testing.T) {
	// build the system
	fs := NewEmptySimpleFrameSystem("test")

	// frame 1 rotate by 45 degrees around z axis and translate
	frame1, err := NewStaticFrame("frame1", spatial.NewPoseFromAxisAngle(r3.Vector{-1., 2., 0.}, r3.Vector{0., 0., 1.}, math.Pi/4))
	test.That(t, err, test.ShouldBeNil)
	err = fs.AddFrame(frame1, fs.World())
	test.That(t, err, test.ShouldBeNil)
	// frame 2 rotate by 45 degree (relative to frame 1) around z axis and translate
	frame2, err := NewStaticFrame("frame2",
		spatial.NewPoseFromAxisAngle(r3.Vector{2. * math.Sqrt(2), 0., 0.}, r3.Vector{0., 0., 1.}, math.Pi/4))
	test.That(t, err, test.ShouldBeNil)
	err = fs.AddFrame(frame2, fs.GetFrame("frame1"))
	test.That(t, err, test.ShouldBeNil)

	// frame3 - pure rotation around y 90 degrees
	frame3, err := NewStaticFrame("frame3", spatial.NewPoseFromAxisAngle(r3.Vector{}, r3.Vector{0., 1., 0.}, math.Pi/2))
	test.That(t, err, test.ShouldBeNil)
	err = fs.AddFrame(frame3, fs.World())
	test.That(t, err, test.ShouldBeNil)

	// frame4 - pure translation
	frame4, err := NewStaticFrame("frame4", spatial.NewPoseFromPoint(r3.Vector{-2., 7., 1.}))
	test.That(t, err, test.ShouldBeNil)
	err = fs.AddFrame(frame4, fs.GetFrame("frame3"))
	test.That(t, err, test.ShouldBeNil)

	// complete fs
	t.Logf("frames in fs: %v", fs.FrameNames())

	// This should remove frames 3 and 4 from fs
	fs2, err := fs.DivideFrameSystem(frame3)
	test.That(t, err, test.ShouldBeNil)
	t.Logf("frames in fs after divide: %v", fs.FrameNames())
	t.Logf("frames in fs2 after divide: %v", fs2.FrameNames())

	f4 := fs.GetFrame("frame4")
	test.That(t, f4, test.ShouldBeNil)
	f1 := fs.GetFrame("frame1")
	test.That(t, f1, test.ShouldNotBeNil)

	f4 = fs2.GetFrame("frame4")
	test.That(t, f4, test.ShouldNotBeNil)
	f1 = fs2.GetFrame("frame1")
	test.That(t, f1, test.ShouldBeNil)

	pointStart := r3.Vector{2., 0., 0.} // the point from PoV of frame 4
	pointEnd := r3.Vector{0., 7., 1.}   // the point from PoV of world (frame3)
	transformPoint, err := fs2.TransformPoint(blankPos, pointStart, "frame4", World)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, transformPoint.X, test.ShouldAlmostEqual, pointEnd.X)
	test.That(t, transformPoint.Y, test.ShouldAlmostEqual, pointEnd.Y)
	test.That(t, transformPoint.Z, test.ShouldAlmostEqual, pointEnd.Z)

	_, err = fs2.TransformPoint(blankPos, pointStart, "frame4", "frame2")
	test.That(t, err, test.ShouldNotBeNil)

	// Put frame3 back where it was
	err = fs.AddFrame(frame3, fs.World())
	test.That(t, err, test.ShouldBeNil)
	err = fs.MergeFrameSystem(fs2, frame3)
	test.That(t, err, test.ShouldBeNil)

	// Comfirm that fs2 is empty now
	t.Logf("frames in fs after merge: %v", fs.FrameNames())
	t.Logf("frames in fs2 after merge: %v", fs2.FrameNames())

	// Confirm new combined frame system now works as it did before
	pointStart = r3.Vector{3., 0., 0.} // the point from PoV of frame 2
	pointEnd = r3.Vector{2., 0., 0.}   // the point from PoV of frame 4
	transformPoint, err = fs.TransformPoint(blankPos, pointStart, "frame2", "frame4")
	test.That(t, err, test.ShouldBeNil)
	test.That(t, transformPoint.X, test.ShouldAlmostEqual, pointEnd.X)
	test.That(t, transformPoint.Y, test.ShouldAlmostEqual, pointEnd.Y)
	test.That(t, transformPoint.Z, test.ShouldAlmostEqual, pointEnd.Z)
}
