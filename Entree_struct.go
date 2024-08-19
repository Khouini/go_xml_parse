package main

import (
	"fmt"
	"time"

	"github.com/beevik/etree"
)

func EntreeStruct() {
	fmt.Println("Entree package struct Parse")
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
	if err := parseEnvelope3(root, &envelope); err != nil {
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

func parseEnvelope3(element *etree.Element, envelope *Envelope) error {
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
		if err := parseHotel3(hotelElement, &hotel); err != nil {
			return err
		}
		envelope.Body.Hotels.Hotels = append(envelope.Body.Hotels.Hotels, hotel)
	}

	return nil
}

func parseHotel3(element *etree.Element, hotel *Hotel) error {
	hotel.HotelId = getElementText(element, "HotelId")
	hotel.Name = getElementText(element, "Name")
	hotel.Rating = getElementInt(element, "Rating")
	hotel.Address = getElementText(element, "Address")
	hotel.Score = getElementInt(element, "Score")
	hotel.HotelChainId = getElementInt(element, "HotelChainId")
	hotel.AccTypeId = getElementInt(element, "AccTypeId")
	hotel.City = getElementText(element, "City")
	hotel.CityId = getElementInt(element, "CityId")
	hotel.ZoneId = getElementInt(element, "ZoneId")
	hotel.Zone = getElementText(element, "Zone")
	hotel.Country = getElementText(element, "Country")
	hotel.CountryId = getElementInt(element, "CountryId")
	hotel.Latitude = getElementFloat(element, "Latitude")
	hotel.Longitude = getElementFloat(element, "Longitude")
	hotel.MarketingText = getElementText(element, "MarketingText")
	hotel.MinRate = getElementFloat(element, "MinRate")
	hotel.MaxRate = getElementFloat(element, "MaxRate")
	hotel.Currency = getElementText(element, "Currency")

	// Parse the Rooms
	for _, roomElement := range element.SelectElement("Rooms").SelectElements("Room") {
		var room Room
		room.Code = getElementInt(roomElement, "Code")
		room.Name = getElementText(roomElement, "Name")

		// Parse the Rates
		for _, rateElement := range roomElement.SelectElement("Rates").SelectElements("Rate") {
			var rate Rate
			rate.RateKey = getElementText(rateElement, "RateKey")
			rate.AmountWithoutPromotion = getElementFloat(rateElement, "AmountWithoutPromotion")
			rate.RateClass = getElementText(rateElement, "RateClass")
			rate.ContractId = getElementText(rateElement, "ContractId")
			rate.RateType = getElementText(rateElement, "RateType")
			rate.PaymentType = getElementText(rateElement, "PaymentType")
			rate.Allotment = getElementInt(rateElement, "Allotment")
			rate.Availability = getElementText(rateElement, "Availability")
			rate.Amount = getElementFloat(rateElement, "Amount")
			rate.BoardCode = getElementText(rateElement, "BoardCode")
			rate.BoardName = getElementText(rateElement, "BoardName")

			// Parse the CancellationPolicies
			for _, cpElement := range rateElement.SelectElement("CancellationPolicies").SelectElements("CancellationPolicy") {
				var cp CancellationPolicy
				cp.Amount = getElementFloat(cpElement, "Amount")
				cp.From = getElementText(cpElement, "From")
				rate.CancellationPolicies = append(rate.CancellationPolicies, cp)
			}

			// Parse other nested elements (Offers, Promotions, Supplements, Taxes) similarly
			// ...

			room.Rates = append(room.Rates, rate)
		}

		hotel.Rooms = append(hotel.Rooms, room)
	}

	// Parse the Photos
	for _, photoElement := range element.SelectElement("Photos").SelectElements("Photo") {
		hotel.Photos = append(hotel.Photos, photoElement.Text())
	}

	return nil
}

func getElementText(element *etree.Element, tag string) string {
	el := element.SelectElement(tag)
	if el != nil {
		return el.Text()
	}
	return ""
}

func getElementInt(element *etree.Element, tag string) int {
	el := element.SelectElement(tag)
	if el != nil {
		// If your XML values are actually integers
		var value int
		fmt.Sscanf(el.Text(), "%d", &value)
		return value
	}
	return 0
}

func getElementFloat(element *etree.Element, tag string) float64 {
	el := element.SelectElement(tag)
	if el != nil {
		// If your XML values are actually floats
		var value float64
		fmt.Sscanf(el.Text(), "%f", &value)
		return value
	}
	return 0.0
}
