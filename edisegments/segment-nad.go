package edisegments

type NADModel struct {
	elements   []string
	separators map[string]string
}

func NewNADModel() *NADModel {
	m := new(NADModel)
	m.elements = []string{
		"NameAndAddressIdentifier",   //This should be NAD
		"PartyQualifier",             //Ref: NAD010 Tag: 3035
		"PartyIdentificationDetails", //Ref: NAD020	Tag: 1001
		"PartyIdIdentification",      //Ref: NAD020-010 Tag: 3039
		"CodeListAgency",             //Ref: NAD020-30 Tag: 3055
	}

	m.separators = map[string]string{
		"NameAndAddressIdentifier":   "+",
		"PartyQualifier":             "+",
		"PartyIdentificationDetails": ":",
		"PartyIdIdentification":      ":",
		"CodeListAgency":             "",
	}

	return m
}

func (u *NADModel) GetElements() []string {
	return u.elements
}

func (u *NADModel) GetSeparator(element string) string {
	return u.separators[element]
}
