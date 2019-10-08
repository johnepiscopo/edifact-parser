package edisegments

type UNHModel struct {
	elements   []string
	separators map[string]string
}

func NewUNHModel() *UNHModel {
	m := new(UNHModel)
	m.elements = []string{
		"MessageReferenceNumber",   //Ref: UNH010 	 Tag: 0062
		"MessageIdentifier",        //Ref: UNH020 	 Tag: S009
		"MessageTypeIdentifier",    //Ref: UNH020-010 Tag: 0065
		"MessageTypeVersionNumber", //Ref: UNH020-020 Tag: 0052
		"MessageTypeReleaseNumber", //Ref: UNH020-030 Tag: 0054
		"ControllingAgency",        //Ref: UNH020-040 Tag: 0051
		"AssociationAssignedCode",  //Ref: UNH020-050 Tag: 0057
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
