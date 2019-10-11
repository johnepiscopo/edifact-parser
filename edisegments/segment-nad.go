package edisegments

func NewNADModel() *EdiSegment {
	m := new(EdiSegment)

	m.Name = "NAD"
	m.Elements = []*EdiElement{
		&EdiElement{
			Name: "PartyQualifier",
			Ref:  "NAD010",
			Tag:  "3035",
		},
		&EdiElement{
			Name: "PartyIdentificationDetails",
			Ref:  "NAD020",
			Tag:  "1001",
			SubElements: []*EdiElement{
				&EdiElement{
					Name: "PartyIdIdentification",
					Ref:  "NAD020-010",
					Tag:  "3039",
				},
				&EdiElement{
					Name: "CodeListAgency",
					Ref:  "NAD020-30",
					Tag:  "3055",
				},
			},
		},
	}

	// m.elements = []string{
	// 	"NameAndAddressIdentifier",   //This should be NAD
	// 	"PartyQualifier",             //Ref: NAD010 Tag: 3035
	// 	"PartyIdentificationDetails", //Ref: NAD020	Tag: 1001
	// 	"PartyIdIdentification",      //Ref: NAD020-010 Tag: 3039
	// 	"CodeListAgency",             //Ref: NAD020-30 Tag: 3055
	// }

	// m.separators = map[string]string{
	// 	"NameAndAddressIdentifier":   "+",
	// 	"PartyQualifier":             "+",
	// 	"PartyIdentificationDetails": ":",
	// 	"PartyIdIdentification":      ":",
	// 	"CodeListAgency":             "",
	// }

	return m
}
