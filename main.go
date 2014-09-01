package main

import (
	"encoding/xml"
	"flag"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
)

var inputFile = flag.String("in", "", "Input file")
var outputFile = flag.String("out", "", "Output file")

func readXml(fName string, data interface{}) error {
	fin, err := os.Open(fName)
	if err != nil {
		return err
	}
	dec := xml.NewDecoder(fin)
	return dec.Decode(data)
}

func writeXml(fName string, data interface{}) error {
	fout, err := os.OpenFile(fName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	_, err = fout.WriteString(xml.Header)
	if err != nil {
		return err
	}
	dec := xml.NewEncoder(fout)
	return dec.Encode(data)
}

func main() {
	flag.Parse()

	var data = &Osm{}
	if err := readXml(*inputFile, data); err != nil {
		log.Fatalf("Failed to read xml: %v", err)
	}
	spew.Printf("%#v\n", data)
	if err := writeXml(*outputFile, data); err != nil {
		log.Fatalf("Failed to write xml: %v", err)
	}
}
