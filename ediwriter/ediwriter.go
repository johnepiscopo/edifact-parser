package ediwriter

import (
	"fmt"
	"strings"

	"github.com/johnepiscopo/edifactParser/common"
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

	segment := edisegments.Segments[segmentCode]
	if segment == nil {
		return edisegments.NewSegmentError(segmentCode)
	}

	var b strings.Builder

	for _, element := range segment.GetElements() {
		toIgnore := false
		for _, i := range elementsToIgnore {
			if i == element {
				toIgnore = true
				break
			}
		}

		if toIgnore {
			continue
		}

		b.WriteString(fmt.Sprintf("%s%s", data[element], segment.GetSeparator(element)))
	}

	s := possiblyTrimEnd(b.String())

	e.file = append(e.file, fmt.Sprintf("%s%s", s, common.TERMINATOR))

	return nil
}

func possiblyTrimEnd(str string) string {
	chars := []string{
		"+",
		":",
		"?",
	}

	for _, c := range chars {
		str = strings.TrimSuffix(str, c)
	}

	return str
}
