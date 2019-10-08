package edisegments

type DTMModel struct {
	elements   []string
	separators map[string]string
}

func NewDTMModel() *DTMModel {
	m := new(DTMModel)
	m.elements = []string{
		"DateTimePeriod",                //Ref: DTM010 	 Tag: C507
		"DateTimePeriodQualifier",       //Ref: DTM010-010 	 Tag: 2005
		"DateTimePeriodValue",           //Ref: DTM010-020 Tag: 2380
		"DateTimePeriodFormatQualifier", //Ref: DTM010-030 Tag: 2379
	}

	m.separators = map[string]string{
		"DateTimePeriod":                "+",
		"DateTimePeriodQualifier":       ":",
		"DateTimePeriodValue":           ":",
		"DateTimePeriodFormatQualifier": "",
	}

	return m
}

func (u *DTMModel) GetElements() []string {
	return u.elements
}

func (u *DTMModel) GetSeparator(element string) string {
	return u.separators[element]
}
