package main

import "encoding/xml"

type Envelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  string   `xml:"Header"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	Hotels Hotels `xml:"http://example.com/hotel Hotels"`
}

type Hotels struct {
	XMLName xml.Name `xml:"http://example.com/hotel Hotels"`
	Hotels  []Hotel  `xml:"Hotel"`
}

type Hotel struct {
	HotelId       string   `xml:"HotelId"`
	Name          string   `xml:"Name"`
	Rating        int      `xml:"Rating"`
	Address       string   `xml:"Address"`
	Score         int      `xml:"Score"`
	HotelChainId  int      `xml:"HotelChainId"`
	AccTypeId     int      `xml:"AccTypeId"`
	City          string   `xml:"City"`
	CityId        int      `xml:"CityId"`
	ZoneId        int      `xml:"ZoneId"`
	Zone          string   `xml:"Zone"`
	Country       string   `xml:"Country"`
	CountryId     int      `xml:"CountryId"`
	Latitude      float64  `xml:"Latitude"`
	Longitude     float64  `xml:"Longitude"`
	MarketingText string   `xml:"MarketingText"`
	MinRate       float64  `xml:"MinRate"`
	MaxRate       float64  `xml:"MaxRate"`
	Currency      string   `xml:"Currency"`
	Rooms         []Room   `xml:"Rooms>Room"`
	Photos        []string `xml:"Photos>Photo"`
}

type Room struct {
	Code  int    `xml:"Code"`
	Name  string `xml:"Name"`
	Rates []Rate `xml:"Rates>Rate"`
}

type Rate struct {
	RateKey                string               `xml:"RateKey"`
	AmountWithoutPromotion float64              `xml:"AmountWithoutPromotion"`
	RateClass              string               `xml:"RateClass"`
	ContractId             string               `xml:"ContractId"`
	RateType               string               `xml:"RateType"`
	PaymentType            string               `xml:"PaymentType"`
	Allotment              int                  `xml:"Allotment"`
	Availability           string               `xml:"Availability"`
	Amount                 float64              `xml:"Amount"`
	BoardCode              string               `xml:"BoardCode"`
	BoardName              string               `xml:"BoardName"`
	CancellationPolicies   []CancellationPolicy `xml:"CancellationPolicies>CancellationPolicy"`
	Offers                 []Offer              `xml:"Offers>Offer"`
	Promotions             []Promotion          `xml:"Promotions>Promotion"`
	Supplements            []Supplement         `xml:"Supplements>Supplement"`
	Taxes                  []Tax                `xml:"Taxes>Tax"`
	Rooms                  int                  `xml:"Rooms"`
	Adults                 int                  `xml:"Adults"`
	Children               int                  `xml:"Children"`
	Infant                 int                  `xml:"Infant"`
	ChildrenAges           string               `xml:"ChildrenAges"`
}

type CancellationPolicy struct {
	Amount float64 `xml:"Amount"`
	From   string  `xml:"From"`
}

type Offer struct {
	// Add fields as necessary
}

type Promotion struct {
	// Add fields as necessary
}

type Supplement struct {
	// Add fields as necessary
}

type Tax struct {
	Name     string  `xml:"Name"`
	Amount   float64 `xml:"Amount"`
	Currency string  `xml:"Currency"`
	Included bool    `xml:"Included"`
}
