package ediwriter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func Test_InvalidSegmentReturnsError(t *testing.T) {
// 	segmentCode := "D'OH"

// 	data := map[string]string{}
// 	ignore := []string{}

// 	e := new(EdiWriter)

// 	err := e.WriteSegment(segmentCode, data, ignore)

// 	assert.NotNil(t, err, "Err is nil")
// 	assert.IsType(t, &edisegments.SegmentError{}, err, fmt.Sprintf("Expected type edisegments.SegmentError Got %T", err))
// 	assert.Equal(t, "D'OH is not a valid segment", err.Error())
// }

func Test_UNHSegmentCreated(t *testing.T) {
	segmentCode := "UNH"

	data := map[string]string{
		"MessageReferenceNumber":   "123456",
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
		"LocationQualifier":       "103",
		"PlaceLocationIdentifier": "CENTRAL WAREHOUSE",
	}

	ignore := []string{}

	e := new(EdiWriter)

	e.WriteSegment(segmentCode, data, ignore)

	s := e.GetFile()
	fmt.Println(s)

	assert.Equal(t, "LOC+103+CENTRAL WAREHOUSE'\n", s)
}

func Test_CPSSegmentCreated(t *testing.T) {
	segmentCode := "CPS"

	data := map[string]string{
		"SegmentName":         "CPS",
		"HierachicalIDNumber": "1",
		"HierachicalParentId": "",
		"PackagingLevelCoded": "4",
	}

	ignore := []string{}

	e := new(EdiWriter)

	e.WriteSegment(segmentCode, data, ignore)

	s := e.GetFile()
	fmt.Println(s)

	assert.Equal(t, "CPS+1++4'\n", s)
}

func Test_PACSegmentCreated(t *testing.T) {
	segmentCode := "PAC"

	type testStruct struct {
		data     map[string]string
		ignore   []string
		expected string
	}

	t1 := testStruct{
		data: map[string]string{
			"SegmentName":                     "PAC",
			"NoOfPackages":                    "",
			"PackagingDetails":                "",
			"PackagingLevelCoded":             "3",
			"PackagingRelatedDescriptionCode": "52",
			"PackageType":                     "CARTON",
			"PackagesIdentification":          "CAR10N",
		},
		ignore: []string{
			//"PackagingDetails",
			"PackagesIdentification",
		},
		expected: "PAC++3:52+CARTON'\n",
	}

	t2 := testStruct{
		data: map[string]string{
			"SegmentName":  "PAC",
			"NoOfPackages": "1",
		},
		ignore: []string{
			"PackagingDetails",
			"PackagingLevelCoded",
			"PackagingRelatedDescriptionCode",
			"PackageType",
			"PackagesIdentification",
		},
		expected: "PAC+1'\n",
	}

	testcases := []testStruct{
		t1, t2,
	}

	for _, testcase := range testcases {
		e := new(EdiWriter)

		e.WriteSegment(segmentCode, testcase.data, testcase.ignore)

		s := e.GetFile()
		fmt.Println(s)

		assert.Equal(t, testcase.expected, s, fmt.Sprintf("Got %s Expected %s", testcase.expected, s))
	}

}
