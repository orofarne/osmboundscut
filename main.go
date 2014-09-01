package main

import (
	"encoding/xml"
	"flag"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
)

var inputFile = flag.String("in", "", "Input file")

func readXml(fName string, data interface{}) error {
	fin, err := os.Open(fName)
	if err != nil {
		return err
	}
	dec := xml.NewDecoder(fin)
	return dec.Decode(data)
}

func main() {
	flag.Parse()

	var data = &Osm{}
	if err := readXml(*inputFile, data); err != nil {
		log.Fatalf("Failed to read xml: %v", err)
	}
	spew.Printf("%#v\n", data)
}
