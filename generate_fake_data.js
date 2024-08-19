const fs = require("fs")

// Function to convert a single hotel object to XML format
function hotelToXML(hotel) {
    return `
        <Hotel>
            <HotelId>${hotel.hotelId}</HotelId>
            <Name>${hotel.name}</Name>
            <Rating>${hotel.rating}</Rating>
            <Address>${hotel.address}</Address>
            <Score>${hotel.score}</Score>
            <HotelChainId>${hotel.hotelChainId}</HotelChainId>
            <AccTypeId>${hotel.accTypeId}</AccTypeId>
            <City>${hotel.city}</City>
            <CityId>${hotel.cityId}</CityId>
            <ZoneId>${hotel.zoneId}</ZoneId>
            <Zone>${hotel.zone}</Zone>
            <Country>${hotel.country}</Country>
            <CountryId>${hotel.countryId}</CountryId>
            <Latitude>${hotel.lat}</Latitude>
            <Longitude>${hotel.long}</Longitude>
            <MarketingText>${hotel.marketingText}</MarketingText>
            <MinRate>${hotel.minRate}</MinRate>
            <MaxRate>${hotel.maxRate}</MaxRate>
            <Currency>${hotel.currency}</Currency>
            <Rooms>
                ${hotel.rooms.map(room => `
                    <Room>
                        <Code>${room.code}</Code>
                        <Name>${room.name}</Name>
                        <Rates>
                            ${room.rates.map(rate => `
                                <Rate>
                                    <RateKey>${rate.rateKey}</RateKey>
                                    <AmountWithoutPromotion>${rate.amountWithoutPromotion}</AmountWithoutPromotion>
                                    <RateClass>${rate.rateClass}</RateClass>
                                    <ContractId>${rate.contractId}</ContractId>
                                    <RateType>${rate.rateType}</RateType>
                                    <PaymentType>${rate.paymentType}</PaymentType>
                                    <Allotment>${rate.allotment}</Allotment>
                                    <Availability>${rate.availability}</Availability>
                                    <Amount>${rate.amount}</Amount>
                                    <BoardCode>${rate.boardCode}</BoardCode>
                                    <BoardName>${rate.boardName}</BoardName>
                                    <CancellationPolicies>
                                        ${rate.cancellationPolicies.map(policy => `
                                            <CancellationPolicy>
                                                <Amount>${policy.amount}</Amount>
                                                <From>${policy.from}</From>
                                            </CancellationPolicy>
                                        `).join('').trim()}
                                    </CancellationPolicies>
                                    <Offers>
                                        ${rate.offers.map(offer => `
                                            <Offer>
                                                <!-- Add Offer Fields -->
                                            </Offer>
                                        `).join('').trim()}
                                    </Offers>
                                    <Promotions>
                                        ${rate.promotions.map(promotion => `
                                            <Promotion>
                                                <!-- Add Promotion Fields -->
                                            </Promotion>
                                        `).join('').trim()}
                                    </Promotions>
                                    <Supplements>
                                        ${rate.supplements.map(supplement => `
                                            <Supplement>
                                                <!-- Add Supplement Fields -->
                                            </Supplement>
                                        `).join('').trim()}
                                    </Supplements>
                                    <Taxes>
                                        ${rate.taxes.map(tax => `
                                            <Tax>
                                                <Name>${tax.name}</Name>
                                                <Amount>${tax.amount}</Amount>
                                                <Currency>${tax.currency}</Currency>
                                                <Included>${tax.included}</Included>
                                            </Tax>
                                        `).join('').trim()}
                                    </Taxes>
                                    <Rooms>${rate.rooms}</Rooms>
                                    <Adults>${rate.adults}</Adults>
                                    <Children>${rate.children}</Children>
                                    <Infant>${rate.infant}</Infant>
                                    <ChildrenAges>${rate.childrenAges}</ChildrenAges>
                                </Rate>
                            `).join('').trim()}
                        </Rates>
                    </Room>
                `).join('').trim()}
            </Rooms>
            <Photos>
                ${hotel.photos.map(photo => `
                    <Photo>${photo}</Photo>
                `).join('').trim()}
            </Photos>
        </Hotel>
    `.trim();
}

// Function to generate SOAP XML for the entire dataset
function generateSOAPXML(hotels) {
    return `
        <soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:hot="http://example.com/hotel">
            <soapenv:Header/>
            <soapenv:Body>
                <hot:Hotels>
                    ${hotels.map(hotel => hotelToXML(hotel)).join('').trim()}
                </hot:Hotels>
            </soapenv:Body>
        </soapenv:Envelope>
    `.trim();
}

// Generate and write the SOAP XML to a file
const data = JSON.parse(fs.readFileSync('./hotels_100_mb.json', 'utf-8'));
const soapXML = generateSOAPXML(data);
fs.writeFileSync('./hotels_soap.xml', soapXML);
console.log('SOAP XML generated and saved to hotels_soap.xml');
