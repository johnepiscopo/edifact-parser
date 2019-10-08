package edisegments

type CPSModel struct {
	elements   []string
	separators map[string]string
}

func NewCPSModel() *CPSModel {
	m := new(CPSModel)
	m.elements = []string{
		"SegmentName",
		"HierachicalIDNumber", //Ref: CPS010 	 Tag: 7164
		"HierachicalParentId", //Ref: CPS020 	 Tag: 7166
		"PackagingLevelCoded", //Ref: CPS030 Tag: 7075
	}

	m.separators = map[string]string{
		"SegmentName":         "+",
		"MessageName":         "+",
		"HierachicalIDNumber": "+",
		"HierachicalParentId": "+",
		"PackagingLevelCoded": "",
	}

	return m
}

func (u *CPSModel) GetElements() []string {
	return u.elements
}

func (u *CPSModel) GetSeparator(element string) string {
	return u.separators[element]
}
