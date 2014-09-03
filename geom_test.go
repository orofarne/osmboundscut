package main

import (
	"math"

	. "gopkg.in/check.v1"
)

type GeomSuite struct{}

var _ = Suite(&GeomSuite{})

//
//          * (2,2)
//         /
// (-1,1) /
//  *----#-------* (4,1)
//      / (1,1)
//     /
//    * (0,0)
//
func (s *GeomSuite) TestCrossSimple1(c *C) {
	x, y := CrossLines(-1., 1., 4., 1., 0., 0., 2., 2.)
	c.Check(x, Near, 1., math.SmallestNonzeroFloat64)
	c.Check(y, Near, 1., math.SmallestNonzeroFloat64)
}

func (s *GeomSuite) TestCrossSimple1a(c *C) {
	x, y := CrossLines(4., 1., -1., 1., 0., 0., 2., 2.)
	c.Check(x, Near, 1., math.SmallestNonzeroFloat64)
	c.Check(y, Near, 1., math.SmallestNonzeroFloat64)
}

//
//        (2,2) *-----------* (4,2)
//
//  (1,1) *-----------------* (4,1)
//
func (s *GeomSuite) TestCrossParallellIsNaN(c *C) {
	x, y := CrossLines(1., 1., 4., 1., 2., 2., 4., 2.)
	c.Check(math.IsNaN(x), Equals, true)
	c.Check(math.IsNaN(y), Equals, true)
}

//
//         * (0,1)
// (-1,0)  |
//    *---------* (1,0)
//         |
//         * (0,-1)
//
func (s *GeomSuite) TestCrossSimple2(c *C) {
	x, y := CrossLines(-1., 0., 1., 0., 0., 1., 0., -1.)
	c.Check(x, Near, 0., math.SmallestNonzeroFloat64)
	c.Check(y, Near, 0., math.SmallestNonzeroFloat64)
}

//
//          * (2,2)
//         /
//        /   (3,1)
//   ----#---*---* (4,1)
//      / (1,1)
//     /
//    * (0,0)
//
func (s *GeomSuite) TestCrossSimple3(c *C) {
	x, y := CrossLines(3., 1., 4., 1., 0., 0., 2., 2.)
	c.Check(x, Near, 1., math.SmallestNonzeroFloat64)
	c.Check(y, Near, 1., math.SmallestNonzeroFloat64)
}

//
//          * (2,2)
//         /
// (-1,1) /
//  *----#-------* (4,1)
//      / (1,1)
//     /
//    * (0,0)
//
func (s *GeomSuite) TestCrossSegmentsIn(c *C) {
	x, y := CrossSegments(-1., 1., 4., 1., 0., 0., 2., 2.)
	c.Check(x, Near, 1., math.SmallestNonzeroFloat64)
	c.Check(y, Near, 1., math.SmallestNonzeroFloat64)
}

//
//          * (2,2)
//         /
//        /   (3,1)
//   ----#---*---* (4,1)
//      / (1,1)
//     /
//    * (0,0)
//
func (s *GeomSuite) TestCrossSegmentsOut(c *C) {
	x, y := CrossSegments(3., 1., 4., 1., 0., 0., 2., 2.)
	c.Check(math.IsNaN(x), Equals, true)
	c.Check(math.IsNaN(y), Equals, true)
}
