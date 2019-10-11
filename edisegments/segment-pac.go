package edisegments

func NewPACModel() *EdiSegment {
	m := new(EdiSegment)

	m.Name = "PAC"
	m.Elements = []*EdiElement{
		&EdiElement{
			Name: "NoOfPackages",
			Ref:  "PAC010",
			Tag:  "7224",
		},
		&EdiElement{
			Name: "PackagingDetails",
			Ref:  "PAC020",
			Tag:  "C531",
			SubElements: []*EdiElement{
				&EdiElement{
					Name: "PackagingLevelCoded",
					Ref:  "PAC020-010",
					Tag:  "7075",
				},
				&EdiElement{
					Name: "PackagingRelatedDescriptionCode",
					Ref:  "PAC020-020",
					Tag:  "7233",
				},
			},
		},
		&EdiElement{
			Name: "PackageType",
			Ref:  "PAC030",
			Tag:  "C202",
			SubElements: []*EdiElement{
				&EdiElement{
					Name: "PackagesIdentification",
					Ref:  "PAC030-010",
					Tag:  "7065",
				},
			},
		},
	}

	// m.elements = []string{
	// 	"SegmentName",
	// 	"NoOfPackages",                    //Ref: PAC010 	 Tag: 7224
	// 	"PackagingDetails",                //Ref: PAC020 	 Tag: C531
	// 	"PackagingLevelCoded",             //Ref: PAC020-010 Tag: 7075
	// 	"PackagingRelatedDescriptionCode", //Ref: PAC020-020 Tag: 7233
	// 	"PackageType",                     //Ref: PAC030 Tag: C202
	// 	"PackagesIdentification",          //Ref: PAC030-010 Tag: 7065
	// }

	// m.separators = map[string]string{
	// 	"SegmentName":                     "+",
	// 	"NoOfPackages":                    "+",
	// 	"PackagingDetails":                "+",
	// 	"PackagingLevelCoded":             ":",
	// 	"PackagingRelatedDescriptionCode": "+",
	// 	"PackageType":                     "+",
	// 	"PackagesIdentification":          "",
	// }

	return m
}
