package main

import (
	"bufio"
	"fmt"
	"github.com/clbanning/mxj"
	"io"
	"os"
	"time"
)

func Mxj() {
	fmt.Println("Mxj package Parse")
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

	// Use buffered I/O for reading the file
	reader := bufio.NewReader(xmlFile)
	xmlData, err := reader.ReadBytes(0)
	if err != nil && err != io.EOF {
		fmt.Println("Error reading XML file:", err)
		return
	}
	fmt.Println("Load Time: ", time.Since(startLoading))

	// Unmarshal the XML data into a map
	startTime := time.Now()
	mv, err := mxj.NewMapXml(xmlData)
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
	hotels, err := mv.ValuesForPath("Envelope.Body.Hotels.Hotel")
	if err != nil {
		fmt.Println("Error extracting hotels:", err)
		return
	}
	numHotels := len(hotels)
	fmt.Printf("Number of hotels parsed: %d\n", numHotels)
}
