package pointcloud

import (
	"testing"

	"github.com/golang/geo/r3"
	"go.viam.com/test"

	"go.viam.com/rdk/spatialmath"
)

func makeClouds(t *testing.T) []PointCloud {
	t.Helper()
	// create cloud 0
	cloud0 := New()
	p00 := NewVector(0, 0, 0)
	test.That(t, cloud0.Set(p00, nil), test.ShouldBeNil)
	p01 := NewVector(0, 0, 1)
	test.That(t, cloud0.Set(p01, nil), test.ShouldBeNil)
	p02 := NewVector(0, 1, 0)
	test.That(t, cloud0.Set(p02, nil), test.ShouldBeNil)
	p03 := NewVector(0, 1, 1)
	test.That(t, cloud0.Set(p03, nil), test.ShouldBeNil)
	// create cloud 1
	cloud1 := New()
	p10 := NewVector(30, 0, 0)
	test.That(t, cloud1.Set(p10, nil), test.ShouldBeNil)
	p11 := NewVector(30, 0, 1)
	test.That(t, cloud1.Set(p11, nil), test.ShouldBeNil)
	p12 := NewVector(30, 1, 0)
	test.That(t, cloud1.Set(p12, nil), test.ShouldBeNil)
	p13 := NewVector(30, 1, 1)
	test.That(t, cloud1.Set(p13, nil), test.ShouldBeNil)
	p14 := NewVector(28, 0.5, 0.5)
	test.That(t, cloud1.Set(p14, nil), test.ShouldBeNil)

	return []PointCloud{cloud0, cloud1}
}

func TestBoundingBoxFromPointCloud(t *testing.T) {
	clouds := makeClouds(t)
	cases := []struct {
		pc             PointCloud
		expectedCenter r3.Vector
		expectedDims   r3.Vector
	}{
		{clouds[0], r3.Vector{0, 0.5, 0.5}, r3.Vector{0, 1, 1}},
		{clouds[1], r3.Vector{29.6, 0.5, 0.5}, r3.Vector{2, 1, 1}},
	}

	for _, c := range cases {
		expectedBox, err := spatialmath.NewBox(spatialmath.NewPoseFromPoint(c.expectedCenter), c.expectedDims)
		test.That(t, err, test.ShouldBeNil)
		box, err := BoundingBoxFromPointCloud(c.pc)
		test.That(t, err, test.ShouldBeNil)
		test.That(t, box, test.ShouldNotBeNil)
		test.That(t, box.AlmostEqual(expectedBox), test.ShouldBeTrue)
	}
}

func TestMergePoints(t *testing.T) {
	clouds := makeClouds(t)
	mergedCloud, err := MergePointClouds(clouds)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, CloudContains(mergedCloud, 0, 0, 0), test.ShouldBeTrue)
	test.That(t, CloudContains(mergedCloud, 30, 0, 0), test.ShouldBeTrue)
}

func TestMergePointsWithColor(t *testing.T) {
	clouds := makeClouds(t)
	mergedCloud, err := MergePointCloudsWithColor(clouds)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, mergedCloud.Size(), test.ShouldResemble, 9)

	a, got := mergedCloud.At(0, 0, 0)
	test.That(t, got, test.ShouldBeTrue)

	b, got := mergedCloud.At(0, 0, 1)
	test.That(t, got, test.ShouldBeTrue)

	c, got := mergedCloud.At(30, 0, 0)
	test.That(t, got, test.ShouldBeTrue)

	test.That(t, a.Color(), test.ShouldResemble, b.Color())
	test.That(t, a.Color(), test.ShouldNotResemble, c.Color())
}

func TestPrune(t *testing.T) {
	clouds := makeClouds(t)
	// before prune
	test.That(t, len(clouds), test.ShouldEqual, 2)
	test.That(t, clouds[0].Size(), test.ShouldEqual, 4)
	test.That(t, clouds[1].Size(), test.ShouldEqual, 5)
	// prune
	clouds = PrunePointClouds(clouds, 5)
	test.That(t, len(clouds), test.ShouldEqual, 1)
	test.That(t, clouds[0].Size(), test.ShouldEqual, 5)
}
