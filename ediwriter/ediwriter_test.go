package ediwriter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func Test_GGMSegmentCreated(t *testing.T) {
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
