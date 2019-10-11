package edisegments

import (
	"fmt"
	"strings"
)

//Segmenter is the interface defining the contract
//for each of the EDI Segments
type Segmenter interface {
	GetElements() []string
	GetSeparator(element string) string
}

//Segments is the map that maps the Segment identifier to the EDI Segment model
var Segments = map[string]*EdiSegment{
	"UNH": NewUNHModel(),
	"BGM": NewBGMModel(),
	"DTM": NewDTMModel(),
	"RFF": NewRFFModel(),
	"NAD": NewNADModel(),
	"TDT": NewTDTModel(),
	"LOC": NewLOCModel(),
	"CPS": NewCPSModel(),
	"PAC": NewPACModel(),
}

//SegmentError is a placeholder comment
type SegmentError struct {
	segmentCode string
}

func (s *SegmentError) Error() string {
	return fmt.Sprintf("%s is not a valid segment", s.segmentCode)
}

//NewSegmentError returns a new segment error object
func NewSegmentError(segmentCode string) error {
	return &SegmentError{segmentCode: segmentCode}
}

var t = EdiSegment{}

type EdiSegment struct {
	Name             string
	Elements         []*EdiElement
	ElementsToIgnore []string
}

func (e *EdiSegment) GetElements() []string {
	elements := []string{}

	for _, el := range e.Elements {
		elements = append(elements, el.Name)
		if len(el.SubElements) == 0 {
			continue
		}

		for _, subel := range el.SubElements {
			elements = append(elements, subel.Name)
		}
	}

	return elements
}

func (e *EdiSegment) SetValueByTag(tag, value string) error {
	for _, ele := range e.Elements {
		if ele.Tag == tag {
			ele.Value = value
			return nil
		}

		for _, sub := range ele.SubElements {
			if sub.Tag == tag {
				sub.Value = value
				return nil
			}
		}
	}

	return fmt.Errorf("%s is not a valid tag for this segment", tag)
}

func (e *EdiSegment) SetValueByName(name, value string) error {
	for _, ele := range e.Elements {
		if ele.Name == name {
			ele.Value = value
			return nil
		}

		for _, sub := range ele.SubElements {
			if sub.Name == name {
				sub.Value = value
				return nil
			}
		}
	}

	return fmt.Errorf("%s is not a valid name for this segment", name)
}

func (e *EdiSegment) SetValueByRef(ref, value string) error {
	for _, ele := range e.Elements {
		if ele.Ref == ref {
			ele.Value = value
			return nil
		}

		for _, sub := range ele.SubElements {
			if sub.Ref == ref {
				sub.Value = value
				return nil
			}
		}
	}

	return fmt.Errorf("%s is not a valid ref for this segment", ref)
}

func (e *EdiSegment) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("%s%s", e.Name, "+"))

	for _, el := range e.Elements {
		if e.ignoreElement(el.Name) {
			continue
		}
		if len(el.Value) > 0 {
			sb.WriteString(el.Value)
			if len(el.SubElements) > 0 {
				sb.WriteString(":")
			}
		}

		for i, sel := range el.SubElements {
			if e.ignoreElement(sel.Name) {
				continue
			}
			sb.WriteString(sel.Value)
			if i != len(el.SubElements)-1 {
				sb.WriteString(":")
			}
		}
		sb.WriteString("+")
	}

	str := sb.String()

	s := trimEnd(str)

	return fmt.Sprintf("%s'", s)
}

func (e *EdiSegment) ignoreElement(name string) bool {
	for _, ignore := range e.ElementsToIgnore {
		if ignore == name {
			return true
		}
	}

	return false
}

func trimEnd(s string) string {
	index := 0
	chars := []string{
		"+",
		":",
		"?",
	}
	for i, c := range s {
		isSpecial := false

		for _, special := range chars {
			if string(c) == special {
				isSpecial = true
				continue
			}
		}

		if !isSpecial {
			index = i
		}
	}

	s = s[:index+1]

	return s

}

type EdiElement struct {
	Ref         string
	Tag         string
	Name        string
	Value       string
	Required    bool
	SubElements []*EdiElement
}
