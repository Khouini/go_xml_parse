package main

import (
	"fmt"
	"time"

	"github.com/beevik/etree"
)

func Entree() {
	fmt.Println("Entree package Parse")
	var startLoading = time.Now()
	// Specify the path to your XML file
	filePath := "./hotels_soap.xml"

	// Read the XML file
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(filePath); err != nil {
		fmt.Println("Error reading XML file:", err)
		return
	}
	fmt.Println("Load Time: ", time.Since(startLoading))

	// Unmarshal the XML data into the Envelope struct
	startTime := time.Now()
	var envelope Envelope
	root := doc.SelectElement("Envelope")
	if root == nil {
		fmt.Println("Error: No Envelope element found")
		return
	}

	// Parse the Envelope element
	if err := parseEnvelope(root, &envelope); err != nil {
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

func parseEnvelope(element *etree.Element, envelope *Envelope) error {
	// Parse the Body element
	bodyElement := element.SelectElement("Body")
	if bodyElement == nil {
		return fmt.Errorf("no Body element found")
	}

	// Parse the Hotels element
	hotelsElement := bodyElement.SelectElement("Hotels")
	if hotelsElement == nil {
		return fmt.Errorf("no Hotels element found")
	}

	// Parse each Hotel element
	for _, hotelElement := range hotelsElement.SelectElements("Hotel") {
		var hotel Hotel
		if err := parseHotel(hotelElement, &hotel); err != nil {
			return err
		}
		envelope.Body.Hotels.Hotels = append(envelope.Body.Hotels.Hotels, hotel)
	}

	return nil
}

func parseHotel(element *etree.Element, hotel *Hotel) error {
	// Parse each field of the Hotel struct
	hotel.HotelId = element.SelectElement("HotelId").Text()
	hotel.Name = element.SelectElement("Name").Text()
	// Add parsing for other fields as necessary
	return nil
}
