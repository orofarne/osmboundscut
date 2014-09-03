package main

import (
	"fmt"
	"math"
)

func CutWays(data *Osm) error {
	// Bound nodes
	var boundNodesRef [4]*Node
	boundNodes := createBoundsNodes(&data.Bounds)
	for i := 0; i < len(boundNodes); i++ {
		boundNodesRef[i] = appendNode(data, boundNodes[i])
	}

	// Ways
	for i := 0; i < len(data.Ways); i++ {
		if err := cutWay(data, &data.Ways[i], boundNodesRef); err != nil {
			return err
		}
	}

	return nil
}

func findPointsOnBBox(bounds Bounds, p1 Node, p2 Node) ([]Node, error) {
	if bounds.MaxLon < bounds.MinLon {
		// Convert to continual plane
		if p1.Lon < bounds.MaxLon {
			p1.Lon = 180.0 + p1.Lon
		}
		if p2.Lon < bounds.MaxLon {
			p2.Lon = 180.0 + p2.Lon
		}
		bounds.MaxLon = 180.0 + bounds.MaxLon
	}

	//                   b
	// max(lat) *-----------------*
	//          |                 |
	//          |                 |
	//          |                 |
	//        a |                 | c
	//          |                 |
	//          |                 |
	//          |                 |
	// min(lat) *-----------------*
	//          ^        d        ^
	//           \                 \
	//            min(lon)          max(lon)
	// x <-> lon
	// y <-> lat
	//
	lonA, latA := CrossSegments(p1.Lon, p1.Lat, p2.Lon, p2.Lat,
		bounds.MinLon, bounds.MinLat, bounds.MinLon, bounds.MaxLat)
	lonB, latB := CrossSegments(p1.Lon, p1.Lat, p2.Lon, p2.Lat,
		bounds.MinLon, bounds.MaxLat, bounds.MaxLon, bounds.MaxLat)
	lonC, latC := CrossSegments(p1.Lon, p1.Lat, p2.Lon, p2.Lat,
		bounds.MaxLon, bounds.MaxLat, bounds.MaxLon, bounds.MinLat)
	lonD, latD := CrossSegments(p1.Lon, p1.Lat, p2.Lon, p2.Lat,
		bounds.MaxLon, bounds.MinLat, bounds.MinLon, bounds.MinLat)
	// Append results
	var result []Node
	var appendResult = func(lon, lat float64) {
		p := createNewEmptyNode()
		if lon > 180.0 {
			lon -= math.Mod(lon, 180.0)
		}
		p.Lon = lon
		p.Lat = lat
		result = append(result, p)
	}
	if !math.IsNaN(lonA) && !math.IsNaN(latA) {
		appendResult(lonA, latA)
	}
	if !math.IsNaN(lonB) && !math.IsNaN(latB) {
		appendResult(lonB, latB)
	}
	if !math.IsNaN(lonC) && !math.IsNaN(latC) {
		appendResult(lonC, latC)
	}
	if !math.IsNaN(lonD) && !math.IsNaN(latD) {
		appendResult(lonD, latD)
	}
	return result, nil
}

func dropWay(data *Osm, way *Way) error {
	for i := 0; i < len(data.Ways); i++ {
		if data.Ways[i].Id == way.Id {
			data.Ways = append(data.Ways[:i], data.Ways[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Way not found")
}

func validateTail(data *Osm, nds []Nd, boundNodes [4]*Node) ([]Nd, error) {
	if len(nds) < 2 {
		return nds, nil
	}
	p1 := getNode(data, nds[len(nds)-2].Ref)
	if p1 == nil {
		return nds, fmt.Errorf("Node %v not found", nds[len(nds)-2].Ref)
	}
	p2 := getNode(data, nds[len(nds)-1].Ref)
	if p2 == nil {
		return nds, fmt.Errorf("Node %v not found", nds[len(nds)-1].Ref)
	}
	if ne(p1.Lon, p2.Lon) && ne(p1.Lat, p2.Lat) {
		for _, bNode := range boundNodes {
			if (eq(bNode.Lon, p1.Lon) && eq(bNode.Lat, p2.Lat)) ||
				(eq(bNode.Lon, p2.Lon) && eq(bNode.Lat, p1.Lat)) {
				bNd := Nd{
					Ref: bNode.Id,
				}
				return append(nds[:len(nds)-1], bNd, nds[len(nds)-1]), nil
			}
		}
	}
	return nds, nil
}

func cutWay(data *Osm, way *Way, boundNodes [4]*Node) error {
	if len(way.Nds) == 0 {
		dropWay(data, way)
		return nil
	}

	prevNode := getNode(data, way.Nds[0].Ref)
	if prevNode == nil {
		return fmt.Errorf("Node %v not found", way.Nds[0].Ref)
	}
	var newWay []Nd
	inBBox := nodeInBounds(&data.Bounds, prevNode)
	if inBBox {
		newWay = append(newWay, Nd{Ref: prevNode.Id})
	}

	closedWay := way.Nds[0].Ref == way.Nds[len(way.Nds)-1].Ref

	for i := 1; i < len(way.Nds); i++ {
		curNode := getNode(data, way.Nds[i].Ref)
		if curNode == nil {
			return fmt.Errorf("Node %v not found", way.Nds[i].Ref)
		}

		xNodes, err := findPointsOnBBox(data.Bounds, *prevNode, *curNode)
		if err != nil {
			return err
		}
		if len(xNodes) == 0 {
			// segment fully in bbox or out of it
			if inBBox {
				newWay = append(newWay, Nd{Ref: curNode.Id})
			}
		} else {
			for _, n := range xNodes {
				newWay = append(newWay, Nd{Ref: appendNode(data, n).Id})
			}
			// Validate tail
			if closedWay {
				newWay, err = validateTail(data, newWay, boundNodes)
				if err != nil {
					return err
				}
			}
			// Check bounds
			if len(xNodes) == 1 {
				inBBox = !inBBox
			}
		}
		prevNode = curNode
	}
	// Empty result?
	if len(newWay) == 0 {
		dropWay(data, way)
		return nil
	}
	// Closed way
	if closedWay {
		newWay = append(newWay, newWay[0])
		var err error
		newWay, err = validateTail(data, newWay, boundNodes)
		if err != nil {
			return err
		}
	}
	way.Nds = newWay
	return nil
}
