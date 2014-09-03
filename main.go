package main

import (
	"encoding/xml"
	"flag"
	"log"
	"os"

	// "github.com/davecgh/go-spew/spew"
)

const APP_NAME = "osmboundscut"

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
	binData, err := xml.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	fout, err := os.OpenFile(fName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	if _, err := fout.WriteString(xml.Header); err != nil {
		return err
	}
	if _, err := fout.Write(binData); err != nil {
		return err
	}
	return nil
}

func main() {
	flag.Parse()

	var data = &Osm{}
	if err := readXml(*inputFile, data); err != nil {
		log.Fatalf("Failed to read xml: %v", err)
	}
	//spew.Printf("%#v\n", data)
	if err := CutWays(data); err != nil {
		log.Fatalf("CutWays error: %v", err)
	}
	data.Generator = APP_NAME
	if err := writeXml(*outputFile, data); err != nil {
		log.Fatalf("Failed to write xml: %v", err)
	}
}
