package main

const (
	//SEPARATOR is the data element separator for the segment
	SEPARATOR string = "+"
	//SUB_ELEMENT_SEPARATOR is the component data element separator
	SUB_ELEMENT_SEPARATOR string = ":"
	//TERMINATOR marks the end of the segment. I need your coat, your boots and your motorcycle.
	TERMINATOR string = "'"
	//ESCAPE_CHARACTER escapes special characters
	ESCAPE_CHARACTER string = "?"
)

var Segments = map[string]Segmenter{
	"UNH": NewUNHModel(),
}

//TODO: Convert this to use a map instead of a slice
var SegmentCodes = map[string]map[string]string{
	"UNH": UNHElements,
}

//TODO: Convert this to be a map of element name: separator
var UNHElements = map[string]string{
	"MessageReferenceNumber":   "+",
	"MessageIdentifier":        "+",
	"MessageTypeIdentifier":    ":",
	"MessageTypeVersionNumber": ":",
	"MessageTypeReleaseNumber": ":",
	"ControllingAgency":        ":",
	"AssociationAssignedCode":  ":",
}

type Segmenter interface {
	GetElements() []string
	GetSeparator(element string) string
}

type UNHModel struct {
	elements   []string
	separators map[string]string
}

func NewUNHModel() *UNHModel {
	m := new(UNHModel)
	m.elements = []string{
		"MessageReferenceNumber",
		"MessageIdentifier",
		"MessageTypeIdentifier",
		"MessageTypeVersionNumber",
		"MessageTypeReleaseNumber",
		"ControllingAgency",
		"AssociationAssignedCode",
	}

	m.separators = map[string]string{
		"MessageReferenceNumber":   "+",
		"MessageIdentifier":        "+",
		"MessageTypeIdentifier":    ":",
		"MessageTypeVersionNumber": ":",
		"MessageTypeReleaseNumber": ":",
		"ControllingAgency":        ":",
		"AssociationAssignedCode":  "",
	}

	return m
}

func (u *UNHModel) GetElements() []string {
	return u.elements
}

func (u *UNHModel) GetSeparator(element string) string {
	return u.separators[element]
}
