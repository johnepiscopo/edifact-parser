package edisegments

func NewMEAModel() *EdiSegment {
	m := new(EdiSegment)

	m.Name = "MEA"
	m.Elements = []*EdiElement{
		&EdiElement{
			Name: "Measurement Application Qualifier",
			Ref:  "MEA010",
			Tag:  "6311",
		},
		&EdiElement{
			Name: "Measurement Details",
			Ref:  "MEA020",
			Tag:  "C502",
			SubElements: []*EdiElement{
				&EdiElement{
					Name: "Measurement Dimension, Coded",
					Tag:  "6313",
					Ref:  "MEA020-010",
				},
			},
		},
		&EdiElement{
			Name: "ValueRange",
			Ref:  "MEA030",
			Tag:  "C174",
			SubElements: []*EdiElement{
				&EdiElement{
					Name: "Measure Unit Qualifier",
					Ref:  "MEA030-010",
					Tag:  "6411",
				},
				&EdiElement{
					Name: "Measurement Value",
					Ref:  "MEA030-020",
					Tag:  "6314",
				},
			},
		},
	}

	// m.elements = []string{
	// 	"SegmentName",
	// 	"MeasurementAttributeQualifier", //Ref: MEA010 	 Tag: 6311
	// 	"MeasurementDetails",            //Ref: MEA020 	 Tag: C502
	// 	"MeasurementDimensionsCoded",    //Ref: MEA020-010 Tag: 6313
	// 	"ValueRange",                    //Ref: MEA030 Tag: C174
	// 	"MeasureUnitQualifier",          //Ref: MEA030-010 Tag: 6411
	// 	"MeasurementValue",              //Ref: MEA030-020 Tag: 6314
	// }

	// m.separators = map[string]string{
	// 	"SegmentName":                   "+",
	// 	"MeasurementAttributeQualifier": "+",
	// 	"MeasurementDetails":            "+",
	// 	"MeasurementDimensionsCoded":    ":",
	// 	"ValueRange":                    "+",
	// 	"MeasureUnitQualifier":          "+",
	// 	"MeasurementValue":              "",
	// }

	return m
}
