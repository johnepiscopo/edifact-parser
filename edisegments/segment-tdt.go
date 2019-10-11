package edisegments

func NewTDTModel() *EdiSegment {
	m := new(EdiSegment)

	m.Name = "TDT"
	m.Elements = []*EdiElement{
		&EdiElement{
			Name: "StageCodeQualifier",
			Ref:  "TDT010",
			Tag:  "8051",
		},
		&EdiElement{
			Name: "MeansOfTransport",
			Ref:  "TDT020",
			Tag:  "8028",
		},
		&EdiElement{
			Name: "ModeOfTransport",
			Ref:  "TDT030",
			Tag:  "C220",
		},
		&EdiElement{
			Name: "TransportMeans",
			Ref:  "TDT040",
			Tag:  "C040",
		},
		&EdiElement{
			Name: "Carrier",
			Ref:  "TDT050",
			Tag:  "C040",
			SubElements: []*EdiElement{
				&EdiElement{
					Name: "CarrierIdentification",
					Ref:  "TDT050-010",
					Tag:  "3127",
				},
				&EdiElement{
					Name: "CodeListIdentificationCode",
					Ref:  "TDT050-020",
					Tag:  "1131",
				},
				&EdiElement{
					Name: "CodeListResponsibleAgencyCode",
					Ref:  "TDT050-030",
					Tag:  "3055",
				},
				&EdiElement{
					Name: "CarrierName",
					Ref:  "TDT050-040",
					Tag:  "3128",
				},
			},
		},
	}

	// m.elements = []string{
	// 	"Carrier",                       //Ref: TDT050 	 Tag: C040
	// 	"CarrierIdentification",         //Ref: TDT050-010 	 Tag: 3127
	// 	"CodeListIdentificationCode",    //Ref: TDT050-020 Tag: 1131
	// 	"CodeListResponsibleAgencyCode", //Ref: TDT050-030 Tag: 3055
	// 	"CarrierName",                   //Ref: TDT050-040 Tag: 3128
	// }

	// m.separators = map[string]string{
	// 	//This has 5 separators as there are a number of elements
	// 	//missing http://www.unece.org/trade/untdid/d05b/trsd/trsdtdt.htm
	// 	"Carrier":                       "+++++",
	// 	"CarrierIdentification":         ":",
	// 	"CodeListIdentificationCode":    ":",
	// 	"CodeListResponsibleAgencyCode": ":",
	// 	"CarrierName":                   "",
	// }

	return m
}
