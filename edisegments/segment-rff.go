package edisegments

type RFFModel struct {
	elements   []string
	separators map[string]string
}

func NewRFFModel() *RFFModel {
	m := new(RFFModel)
	m.elements = []string{
		"Reference",          //Ref: RFF010 	 Tag: C506
		"ReferenceQualifier", //Ref: RFF010-010 	 Tag: 1153
		"ReferenceNumber",    //Ref: RFF010-020 Tag: 1154
	}

	m.separators = map[string]string{
		"Reference":          "+",
		"ReferenceQualifier": ":",
		"ReferenceNumber":    "",
	}

	return m
}

func (u *RFFModel) GetElements() []string {
	return u.elements
}

func (u *RFFModel) GetSeparator(element string) string {
	return u.separators[element]
}
