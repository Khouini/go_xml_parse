package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"time"
)

func MxjStruct() {
	fmt.Println("Using struct-based XML Parsing")
	var startLoading = time.Now()

	// Specify the path to your XML file
	filePath := "./hotels_soap.xml"

	// Open the XML file
	xmlFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	// Read the entire XML file into a byte slice
	xmlData, err := io.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML file:", err)
		return
	}
	fmt.Println("Load Time: ", time.Since(startLoading))

	// Unmarshal the XML data into the Envelope struct
	startTime := time.Now()
	var envelope Envelope
	err = xml.Unmarshal(xmlData, &envelope)
	if err != nil {
		fmt.Println("Error unmarshaling XML data:", err)
		return
	}

	// Record the end time
	endTime := time.Now()

	// Calculate and log the parsing time
	parseDuration := endTime.Sub(startTime)
	fmt.Printf("Parsed XML in %v\n", parseDuration)

	// Extract the number of hotels parsed
	numHotels := len(envelope.Body.Hotels.Hotels)
	fmt.Printf("Number of hotels parsed: %d\n", numHotels)
}
