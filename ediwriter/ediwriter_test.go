package ediwriter

import (
	"fmt"
	"testing"

	"github.com/johnepiscopo/edifactParser/edisegments"
	"github.com/stretchr/testify/assert"
)

func Test_InvalidSegmentReturnsError(t *testing.T) {
	segmentCode := "D'OH"

	data := map[string]string{}
	ignore := []string{}

	e := new(EdiWriter)

	err := e.WriteSegment(segmentCode, data, ignore)

	assert.NotNil(t, err, "Err is nil")
	assert.IsType(t, &edisegments.SegmentError{}, err, fmt.Sprintf("Expected type edisegments.SegmentError Got %T", err))
	assert.Equal(t, "D'OH is not a valid segment", err.Error())
}

func Test_UNHSegmentCreated(t *testing.T) {
	segmentCode := "UNH"

	data := map[string]string{
		"MessageReferenceNumber":   "UNH",
		"MessageIdentifier":        "123456",
		"MessageTypeIdentifier":    "DESADV",
		"MessageTypeVersionNumber": "D",
		"MessageTypeReleaseNumber": "96A",
		"ControllingAgency":        "UN",
	}

	ignore := []string{
		"AssociationAssignedCode",
	}

	e := new(EdiWriter)

	e.WriteSegment(segmentCode, data, ignore)

	s := e.GetFile()
	fmt.Println(s)

	assert.Equal(t, "UNH+123456+DESADV:D:96A:UN'\n", s)

}

func Test_BGMSegmentCreated(t *testing.T) {
	segmentCode := "BGM"

	data := map[string]string{
		"MessageName":          "BGM",
		"MessageNameCoded":     "351",
		"MessageNumber":        "012345",
		"MessageFunctionCoded": "9",
	}

	ignore := []string{}

	e := new(EdiWriter)

	e.WriteSegment(segmentCode, data, ignore)

	s := e.GetFile()
	fmt.Println(s)

	assert.Equal(t, "BGM+351+012345+9'\n", s)
}

func Test_DTMSegmentCreated(t *testing.T) {
	segmentCode := "DTM"

	data := map[string]string{
		"DateTimePeriod":                "DTM",
		"DateTimePeriodQualifier":       "137",
		"DateTimePeriodValue":           "20081001",
		"DateTimePeriodFormatQualifier": "102",
	}

	ignore := []string{}

	e := new(EdiWriter)

	e.WriteSegment(segmentCode, data, ignore)

	s := e.GetFile()
	fmt.Println(s)

	assert.Equal(t, "DTM+137:20081001:102'\n", s)
}

func Test_NADSegmentCreated(t *testing.T) {
	segmentCode := "NAD"

	data := map[string]string{
		"NameAndAddressIdentifier":   "NAD",
		"PartyQualifier":             "BY",
		"PartyIdentificationDetails": "123456",
		"PartyIdIdentification":      "160",
		"CodeListAgency":             "9",
	}

	ignore := []string{}

	e := new(EdiWriter)

	e.WriteSegment(segmentCode, data, ignore)

	s := e.GetFile()
	fmt.Println(s)

	assert.Equal(t, "NAD+BY+123456:160:9'\n", s)
}

func Test_RFFSegmentCreated(t *testing.T) {
	segmentCode := "RFF"

	data := map[string]string{
		"Reference":          "RFF",
		"ReferenceQualifier": "PK",
		"ReferenceNumber":    "123456",
	}

	ignore := []string{}

	e := new(EdiWriter)

	e.WriteSegment(segmentCode, data, ignore)

	s := e.GetFile()
	fmt.Println(s)

	assert.Equal(t, "RFF+PK:123456'\n", s)
}

func Test_TDTSegmentCreated(t *testing.T) {
	segmentCode := "TDT"

	data := map[string]string{
		"Carrier":                       "TDT",
		"CarrierIdentification":         "UPSN",
		"CodeListIdentificationCode":    "172",
		"CodeListResponsibleAgencyCode": "",
		"CarrierName":                   "DPD",
	}

	ignore := []string{}

	e := new(EdiWriter)

	e.WriteSegment(segmentCode, data, ignore)

	s := e.GetFile()
	fmt.Println(s)

	assert.Equal(t, "TDT+++++UPSN:172::DPD'\n", s)
}

func Test_LOCSegmentCreated(t *testing.T) {
	segmentCode := "LOC"

	data := map[string]string{
		"LocationQualifier":       "LOC",
		"LocationIdentifier":      "103",
		"PlaceLocationIdentifier": "CENTRAL WAREHOUSE",
	}

	ignore := []string{}

	e := new(EdiWriter)

	e.WriteSegment(segmentCode, data, ignore)

	s := e.GetFile()
	fmt.Println(s)

	assert.Equal(t, "LOC+103+CENTRAL WAREHOUSE'\n", s)
}
