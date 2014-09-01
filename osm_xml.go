package main

// http://wiki.openstreetmap.org/wiki/OSM_XML

// OSM ids
type Idtype uint64

// <osm version="0.6" generator="CGImap 0.0.2">
type Osm struct {
	Version   string     `xml:"version,attr"`
	Bounds    Bounds     `xml:"bounds"`
	Nodes     []Node     `xml:"node"`
	Ways      []Way      `xml:"way"`
	Relations []Relation `xml:"relation"`
}

// <bounds minlat="54.0889580" minlon="12.2487570" maxlat="54.0913900" maxlon="12.2524800"/>
type Bounds struct {
	MinLat float64 `xml:"minlat,attr"`
	MinLon float64 `xml:"minlon,attr"`
	MaxLat float64 `xml:"maxlat,attr"`
	MaxLon float64 `xml:"maxlon,attr"`
}

// <node id="298884269" lat="54.0901746" lon="12.2482632" user="SvenHRO" uid="46882" visible="true" version="1" changeset="676636" timestamp="2008-09-21T21:37:45Z"/>
type Node struct {
	Id   Idtype  `xml:"id,attr"`
	Lat  float64 `xml:"lat,attr"`
	Lon  float64 `xml:"lon,attr"`
	Tags []Tag   `xml:"tag"`
}

// <tag k="name" v="Neu Broderstorf"/>
type Tag struct {
	K string `xml:"k,attr"`
	V string `xml:"v,attr"`
}

// <way id="26659127" user="Masch" uid="55988" visible="true" version="5" changeset="4142606" timestamp="2010-03-16T11:47:08Z">
type Way struct {
	Id   Idtype `xml:"id,attr"`
	Nds  []Nd   `xml:"nd"`
	Tags []Tag  `xml:"tag"`
}

// <nd ref="292403538"/>
type Nd struct {
	Ref Idtype `xml:"ref,attr"`
}

// <relation id="56688" user="kmvar" uid="56190" visible="true" version="28" changeset="6947637" timestamp="2011-01-12T14:23:49Z">
type Relation struct {
	Id      Idtype   `xml:"id,attr"`
	Members []Member `xml:"member"`
	Tags    []Tag    `xml:"tag"`
}

// <member type="node" ref="294942404" role=""/>
type Member struct {
	Type string `xml:"type,attr"`
	Ref  Idtype `xml:"ref,attr"`
	Role string `xml:"role,attr"`
}
