const fs = require('fs').promises;
const { parseStringPromise } = require('xml2js');
const { XMLParser } = require('fast-xml-parser');

const filePath = '../hotels_soap.xml'; // Replace with your XML file path

// Function to parse XML using xml2js
async function parseWithXml2js(xmlData) {
    try {
        console.time('xml2js');
        const result = await parseStringPromise(xmlData, { explicitArray: false });
        console.timeEnd('xml2js');
        return result;
    } catch (error) {
        console.error('Error parsing XML with xml2js:', error);
        throw error;
    }
}

// Function to parse XML using fast-xml-parser
function parseWithFastXmlParser(xmlData) {
    try {
        console.time('fast-xml-parser');
        const parser = new XMLParser({
            ignoreAttributes: false,
            attributeNamePrefix: "@_",
            allowBooleanAttributes: true,
        });
        const result = parser.parse(xmlData);
        console.timeEnd('fast-xml-parser');
        return result;
    } catch (error) {
        console.error('Error parsing XML with fast-xml-parser:', error);
        throw error;
    }
}

// Run both parsers and compare
async function runComparison() {
    try {
        // Load the XML data asynchronously
        console.time("loading")
        const xmlData = await fs.readFile(filePath, 'utf-8');
        console.timeEnd("loading")
        // Parse using xml2js
        const xml2jsData = await parseWithXml2js(xmlData);
        console.log('Parsed XML with xml2js:', xml2jsData);

        // Parse using fast-xml-parser
        const fastXmlData = parseWithFastXmlParser(xmlData);
        console.log('Parsed XML with fast-xml-parser:', fastXmlData);

    } catch (error) {
        console.error('Error:', error);
    }
}

// Execute the comparison
runComparison();
