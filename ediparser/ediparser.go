package ediparser

import (
	"strings"

	"github.com/johnepiscopo/edifactParser/edisegments"
	"github.com/xiatechs/sdv/framework/middleware/logging"
)

type EdiParser struct{}

func (e *EdiParser) ParseEdifact(input string, ignoreList []string) []*edisegments.EdiSegment {

	//Split into array by the terminator delimiter
	segments := strings.Split(input, "'")

	segmentModels := []*edisegments.EdiSegment{}

	for _, s := range segments {
		m := processSegment(s, ignoreList)
		segmentModels = append(segmentModels, m)
	}

	return segmentModels
}

func processSegment(segment string, ignoreList []string) *edisegments.EdiSegment {
	//Split into the elements
	elements := strings.Split(segment, "+")
	//Get the relevant Segment model
	model := edisegments.Segments[elements[0]]
	model.ElementsToIgnore = ignoreList

	for i, element := range elements {
		if i == 0 {
			logging.Info.Printf("Processing segment %s", model.Name)
			//Will be the segment identifier so can ignore e.g. UNH
			continue
		}

		processElement(element, model, ignoreList)
	}

	return model
}

func processElement(element string, model *edisegments.EdiSegment, ignoreList []string) {
	subElements := strings.Split(element, ":")

	for _, element := range model.Elements {

		if len(subElements) == 0 {
			break
		}

		ignore := false

		for _, ig := range ignoreList {
			if element.Name == ig {
				ignore = true
				logging.Info.Printf("Ignoring element %s", element.Name)
				break
			}
		}

		if ignore {
			continue
		}

		//For 123456
		if len(element.SubElements) == 0 && len(element.Value) == 0 {
			element.Value = subElements[0]
			subElements = subElements[1:]
			logging.Info.Printf("Setting element %s to %s", element.Name, element.Value)
			continue
		}

		//For DESADV:D:96A:UN
		for _, subelement := range element.SubElements {
			for _, ig := range ignoreList {
				if subelement.Name == ig {
					logging.Info.Printf("Ignoring element %s", subelement.Name)
					ignore = true
					break
				}
			}

			if ignore || len(subelement.Value) > 0 {
				continue
			}

			if len(subElements) == 0 {
				break
			}

			subelement.Value = subElements[0]
			logging.Info.Printf("Setting element %s to %s", subelement.Name, subelement.Value)
			subElements = subElements[1:]

		}
	}
}
