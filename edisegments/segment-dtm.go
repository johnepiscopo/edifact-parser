package edisegments

func NewDTMModel() *EdiSegment {
	m := new(EdiSegment)

	m.Name = "DTM"
	m.Elements = []*EdiElement{
		&EdiElement{
			Name: "DateTimePeriod",
			Ref:  "DTM010",
			Tag:  "C507",
			SubElements: []*EdiElement{
				&EdiElement{
					Name: "DateTimePeriodQualifier",
					Ref:  "DTM010-010",
					Tag:  "2005",
				},
				&EdiElement{
					Name: "DateTimePeriodValue",
					Ref:  "DTM010-020",
					Tag:  "2380",
				},
				&EdiElement{
					Name: "DateTimePeriodFormatQualifier",
					Ref:  "DTM010-030",
					Tag:  "2379",
				},
			},
		},
	}

	// m.elements = []string{
	// 	"DateTimePeriod",                //Ref: DTM010 	 Tag: C507
	// 	"DateTimePeriodQualifier",       //Ref: DTM010-010 	 Tag: 2005
	// 	"DateTimePeriodValue",           //Ref: DTM010-020 Tag: 2380
	// 	"DateTimePeriodFormatQualifier", //Ref: DTM010-030 Tag: 2379
	// }

	// m.separators = map[string]string{
	// 	"DateTimePeriod":                "+",
	// 	"DateTimePeriodQualifier":       ":",
	// 	"DateTimePeriodValue":           ":",
	// 	"DateTimePeriodFormatQualifier": "",
	// }

	return m
}
