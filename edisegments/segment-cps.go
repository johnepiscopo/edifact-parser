package edisegments

func NewCPSModel() *EdiSegment {
	m := new(EdiSegment)

	m.Name = "CPS"
	m.Elements = []*EdiElement{
		&EdiElement{
			Name: "HierachicalIDNumber",
			Ref:  "CPS010",
			Tag:  "7164",
		},
		&EdiElement{
			Name: "HierachicalParentId",
			Ref:  "CPS020",
			Tag:  "7166",
		},
		&EdiElement{
			Name: "PackagingLevelCoded",
			Ref:  "CPS030",
			Tag:  "7075",
		},
	}

	// m.elements = []string{
	// 	"SegmentName",
	// 	"HierachicalIDNumber", //Ref: CPS010 	 Tag: 7164
	// 	"HierachicalParentId", //Ref: CPS020 	 Tag: 7166
	// 	"PackagingLevelCoded", //Ref: CPS030 Tag: 7075
	// }

	// m.separators = map[string]string{
	// 	"SegmentName":         "+",
	// 	"MessageName":         "+",
	// 	"HierachicalIDNumber": "+",
	// 	"HierachicalParentId": "+",
	// 	"PackagingLevelCoded": "",
	// }

	return m
}
