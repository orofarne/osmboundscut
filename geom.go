package main

import "math"

const EPS = 0.000000001

// a > b
func gt(a, b float64) bool {
	return a-b > EPS
}

// a < b
func lt(a, b float64) bool {
	return b-a > EPS
}

// a == b
func eq(a, b float64) bool {
	return math.Abs(a-b) < EPS
}

// a != b
func ne(a, b float64) bool {
	return math.Abs(a-b) > EPS
}

// Search crosspoint of two lines
// First line: points (x11, y11), (x12, y12)
// Second line: points (x21, y21), (x22, y22)
func CrossLines(x11, y11, x12, y12, x21, y21, x22, y22 float64) (x, y float64) {
	// Line coeffs
	a1 := y11 - y12
	b1 := x12 - x11
	c1 := x11*y12 - x12*y11
	a2 := y21 - y22
	b2 := x22 - x21
	c2 := x21*y22 - x22*y21
	// Determinant
	D := a1*b2 - a2*b1
	if math.Abs(D) < EPS {
		x, y = math.NaN(), math.NaN()
		return
	}
	k := -1.0 / D
	x = k * (b2*c1 - b1*c2)
	y = k * (-a2*c1 + a1*c2)
	return
}

// Search crosspoint of two segments
// First segment: points (x11, y11), (x12, y12)
// Second segment: points (x21, y21), (x22, y22)
func CrossSegments(x11, y11, x12, y12, x21, y21, x22, y22 float64) (x, y float64) {
	minX1 := math.Min(x11, x12)
	maxX1 := math.Max(x11, x12)
	minY1 := math.Min(y11, y12)
	maxY1 := math.Max(y11, y12)
	minX2 := math.Min(x21, x22)
	maxX2 := math.Max(x21, x22)
	minY2 := math.Min(y21, y22)
	maxY2 := math.Max(y21, y22)
	if gt(minX1, maxX2) || gt(minX2, maxX1) || gt(minY1, maxY2) || gt(minY2, maxY1) {
		x, y = math.NaN(), math.NaN()
		return
	}

	x, y = CrossLines(x11, y11, x12, y12, x21, y21, x22, y22)
	if math.IsNaN(x) || math.IsNaN(y) {
		return
	}
	if lt(x, minX1) || lt(x, minX2) || gt(x, maxX1) || gt(x, maxX2) ||
		lt(y, minY1) || lt(y, minY2) || gt(y, maxY1) || gt(y, maxY2) {
		x, y = math.NaN(), math.NaN()
		return
	}
	return
}
