package main

import "time"

var nextNewId Idtype = 0

func getNode(data *Osm, id Idtype) *Node {
	for i := 0; i < len(data.Nodes); i++ {
		if data.Nodes[i].Id == id {
			return &data.Nodes[i]
		}
	}
	return nil
}

func appendNode(data *Osm, node Node) *Node {
	data.Nodes = append(data.Nodes, node)
	return &data.Nodes[len(data.Nodes)-1]
}

func nodeInBounds(bounds *Bounds, node *Node) bool {
	if node.Lat < bounds.MinLat || node.Lat > bounds.MaxLat {
		return false
	}
	if bounds.MaxLon >= bounds.MinLon && (node.Lon < bounds.MinLon || node.Lon > bounds.MaxLon) {
		return false
	}
	if bounds.MaxLon < bounds.MinLon && (node.Lon < bounds.MinLon && node.Lon > bounds.MaxLon) {
		return false
	}
	return true
}

// Create new empty node with Id < 0
// Lat, Lon = 0, 0
func createNewEmptyNode() Node {
	nextNewId--
	return Node{
		Id:        nextNewId,
		Visible:   true,
		Timestamp: time.Now().Format("2006-01-02T15:04:05Z"),
		Version:   "1",
	}
}

// Create corner nodes for bounds
//
//  0        1
//   *--------* <- max(lat)
//   |        |
//   |        |
//   |        | 2
// 3 *--------* <- min(lat)
//   ^        ^
//    \        \
//     min(lon) \
//               max(lon)
//
func createBoundsNodes(bounds *Bounds) [4]Node {
	nodes := [4]Node{createNewEmptyNode(), createNewEmptyNode(),
		createNewEmptyNode(), createNewEmptyNode()}

	nodes[0].Lat = bounds.MaxLat
	nodes[0].Lon = bounds.MinLon
	nodes[1].Lat = bounds.MaxLat
	nodes[1].Lon = bounds.MaxLon
	nodes[2].Lat = bounds.MinLat
	nodes[2].Lon = bounds.MaxLon
	nodes[3].Lat = bounds.MinLat
	nodes[3].Lon = bounds.MinLon

	return nodes
}
