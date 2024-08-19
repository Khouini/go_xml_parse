package main

import (
	"fmt"
	"os"
	"time"

	"github.com/antchfx/xmlquery"
)

func Xmlquery() {
	fmt.Println("Xmlquery package Parse")
	// Specify the path to your XML file
	filePath := "./hotels_soap.xml"

	// Open the XML file
	xmlFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	var startLoading = time.Now()
	// Parse the XML file
	doc, err := xmlquery.Parse(xmlFile)
	if err != nil {
		fmt.Println("Error parsing XML file:", err)
		return
	}
	fmt.Println("Load Time: ", time.Since(startLoading))

	// Unmarshal the XML data into the Envelope struct
	startTime := time.Now()
	var envelope Envelope
	root := xmlquery.FindOne(doc, "//*[local-name()='Envelope']")
	if root == nil {
		fmt.Println("Error: No Envelope element found")
		return
	}

	// Parse the Envelope element
	if err := parseEnvelope2(root, &envelope); err != nil {
		fmt.Println("Error parsing XML data:", err)
		return
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

func parseEnvelope2(element *xmlquery.Node, envelope *Envelope) error {
	// Parse the Body element
	bodyElement := xmlquery.FindOne(element, "//*[local-name()='Body']")
	if bodyElement == nil {
		return fmt.Errorf("no Body element found")
	}

	// Parse the Hotels element
	hotelsElement := xmlquery.FindOne(bodyElement, "//*[local-name()='Hotels']")
	if hotelsElement == nil {
		return fmt.Errorf("no Hotels element found")
	}

	// Parse each Hotel element
	for _, hotelElement := range xmlquery.Find(hotelsElement, "//*[local-name()='Hotel']") {
		var hotel Hotel
		if err := parseHotel2(hotelElement, &hotel); err != nil {
			return err
		}
		envelope.Body.Hotels.Hotels = append(envelope.Body.Hotels.Hotels, hotel)
	}

	return nil
}

func parseHotel2(element *xmlquery.Node, hotel *Hotel) error {
	// Parse each field of the Hotel struct
	hotel.HotelId = xmlquery.FindOne(element, "//*[local-name()='HotelId']").InnerText()
	hotel.Name = xmlquery.FindOne(element, "//*[local-name()='Name']").InnerText()
	// Add parsing for other fields as necessary
	return nil
}
