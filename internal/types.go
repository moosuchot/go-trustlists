package internal

import "encoding/xml"

type TrustServiceStatusList struct {
	XMLName           xml.Name `xml:"TrustServiceStatusList"`
	SchemeInformation SchemeInformation
}

type SchemeInformation struct {
	XMLName                     xml.Name `xml:"SchemeInformation"`
	TSLVersionIdentifier        int      `xml:"TSLVersionIdentifier"`
	HistoricalInformationPeriod int
	PointersToOtherTSL          PointersToOtherTSL
}

type PointersToOtherTSL struct {
	XMLName         xml.Name `xml:"PointersToOtherTSL"`
	OtherTSLPointer OtherTSLPointer
}

type OtherTSLPointer struct {
	XMLName                  xml.Name `xml:"OtherTSLPointer"`
	ServiceDigitalIdentities []ServiceDigitalIdentities
}

type ServiceDigitalIdentities struct {
	XMLName                xml.Name `xml:"ServiceDigitalIdentities"`
	ServiceDigitalIdentity ServiceDigitalIdentity
}

type ServiceDigitalIdentity struct {
	XMLName   xml.Name `xml:"ServiceDigitalIdentity"`
	DigitalId DigitalId
}

type DigitalId struct {
	XMLName         xml.Name `xml:"DigitalId"`
	X509Certificate string   `xml:"X509Certificate"`
}
