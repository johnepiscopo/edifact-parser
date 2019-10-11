package edisegments

func NewRFFModel() *EdiSegment {
	m := new(EdiSegment)

	m.Name = "RFF"
	m.Elements = []*EdiElement{
		&EdiElement{
			Name: "Reference",
			Ref:  "RFF010",
			Tag:  "C506",
			SubElements: []*EdiElement{
				&EdiElement{
					Name: "ReferenceQualifier",
					Ref:  "RFF010-010",
					Tag:  "1153",
				},
				&EdiElement{
					Name: "ReferenceNumber",
					Ref:  "RFF010-020",
					Tag:  "1154",
				},
			},
		},
	}

	// m.elements = []string{
	// 	"Reference",          //Ref: RFF010 	 Tag: C506
	// 	"ReferenceQualifier", //Ref: RFF010-010 	 Tag: 1153
	// 	"ReferenceNumber",    //Ref: RFF010-020 Tag: 1154
	// }

	// m.separators = map[string]string{
	// 	"Reference":          "+",
	// 	"ReferenceQualifier": ":",
	// 	"ReferenceNumber":    "",
	// }

	return m
}
