package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FullSegmentCreated(t *testing.T) {
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
