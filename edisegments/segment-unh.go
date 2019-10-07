package edisegments

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
