package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"time"
)

func StreamingParse() {
	fmt.Println("Streaming Parse")
	// Specify the path to your XML file
	filePath := "./hotels_soap.xml"

	// Open the XML file
	xmlFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	// Create a new XML decoder
	decoder := xml.NewDecoder(xmlFile)

	// Unmarshal the XML data into the Envelope struct
	startTime := time.Now()
	var envelope Envelope

	for {
		token, err := decoder.Token()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error decoding XML:", err)
			return
		}

		switch t := token.(type) {
		case xml.StartElement:
			if t.Name.Local == "Hotel" {
				var hotel Hotel
				if err := decoder.DecodeElement(&hotel, &t); err != nil {
					fmt.Println("Error decoding element:", err)
					return
				}
				envelope.Body.Hotels.Hotels = append(envelope.Body.Hotels.Hotels, hotel)
			}
		}
	}

	// Record the end time
	endTime := time.Now()

	// Calculate and log the parsing time
	parseDuration := endTime.Sub(startTime)
	fmt.Printf("Parsed XML in %v\n", parseDuration)

	// Log the number of hotels parsed
	numHotels := len(envelope.Body.Hotels.Hotels)
	fmt.Printf("Number of hotels parsed: %d\n", numHotels)
}
