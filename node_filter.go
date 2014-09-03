package main

func FilterNodes(data *Osm) error {
	for i := 0; i < len(data.Nodes); i++ {
		if !nodeInBounds(&data.Bounds, &data.Nodes[i]) {
			// Drop it
			data.Nodes = append(data.Nodes[:i], data.Nodes[i+1:]...)
			i--
		}
	}
	return nil
}
