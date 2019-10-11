package ediwriter

import (
	"strings"

	"github.com/johnepiscopo/edifactParser/edisegments"
)

//EdiWriter creates an EDIFact file
type EdiWriter struct {
	file []string
}

//GetFile returns the EDIFact file
func (e *EdiWriter) GetFile() string {
	var sb strings.Builder
	for _, f := range e.file {
		sb.WriteString(f)
		sb.WriteString("\n")
	}

	return sb.String()
}

func (e *EdiWriter) WriteSegment(segmentCode string, data map[string]string, elementsToIgnore []string) error {

	segment := edisegments.Segments[strings.ToUpper(segmentCode)]

	for key, value := range data {
		segment.SetValueByName(key, value)
	}

	e.file = append(e.file, segment.String())

	return nil
}
