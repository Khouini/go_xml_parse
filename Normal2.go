package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"time"
)

func Normal2() {
	fmt.Println("Normal 2 Parse")
	startLoading := time.Now()
	filePath := "./hotels_soap.xml"

	xmlFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	fmt.Println("Load Time: ", time.Since(startLoading))

	startTime := time.Now()
	decoder := xml.NewDecoder(xmlFile)
	var envelope Envelope

	// Decode envelope using streaming approach
	if err := decodeEnvelope(decoder, &envelope); err != nil {
		fmt.Println("Error unmarshaling XML data:", err)
		return
	}

	endTime := time.Now()
	parseDuration := endTime.Sub(startTime)
	fmt.Printf("Parsed XML in %v\n", parseDuration)

	numHotels := len(envelope.Body.Hotels.Hotels)
	fmt.Printf("Number of hotels parsed: %d\n", numHotels)
}

func decodeEnvelope(d *xml.Decoder, env *Envelope) error {
	for {
		t, err := d.Token()
		if err != nil {
			return err
		}
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "Envelope" {
				return d.DecodeElement(env, &se)
			}
		}
	}
}
