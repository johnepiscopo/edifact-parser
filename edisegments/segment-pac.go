package edisegments

type PACModel struct {
	elements   []string
	separators map[string]string
}

func NewPACModel() *PACModel {
	m := new(PACModel)
	m.elements = []string{
		"SegmentName",
		"NoOfPackages",                    //Ref: PAC010 	 Tag: 7224
		"PackagingDetails",                //Ref: PAC020 	 Tag: C531
		"PackagingLevelCoded",             //Ref: PAC020-010 Tag: 7075
		"PackagingRelatedDescriptionCode", //Ref: PAC020-020 Tag: 7233
		"PackageType",                     //Ref: PAC030 Tag: C202
		"PackagesIdentification",          //Ref: PAC030-010 Tag: 7065
	}

	m.separators = map[string]string{
		"SegmentName":                     "+",
		"NoOfPackages":                    "+",
		"PackagingDetails":                "+",
		"PackagingLevelCoded":             ":",
		"PackagingRelatedDescriptionCode": "+",
		"PackageType":                     "+",
		"PackagesIdentification":          "",
	}

	return m
}

func (u *PACModel) GetElements() []string {
	return u.elements
}

func (u *PACModel) GetSeparator(element string) string {
	return u.separators[element]
}
