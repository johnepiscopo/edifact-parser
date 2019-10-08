package edisegments

type BGMModel struct {
	elements   []string
	separators map[string]string
}

func NewBGMModel() *BGMModel {
	m := new(BGMModel)
	m.elements = []string{
		"MessageName",          //Ref: BGM010 	 Tag: C002
		"MessageNameCoded",     //Ref: BGM010-010 	 Tag: 1001
		"MessageNumber",        //Ref: BGM020 Tag: 1004
		"MessageFunctionCoded", //Ref: BGM030 Tag: 1225
	}

	m.separators = map[string]string{
		"MessageName":          "+",
		"MessageNameCoded":     "+",
		"MessageNumber":        "+",
		"MessageFunctionCoded": "",
	}

	return m
}

func (u *BGMModel) GetElements() []string {
	return u.elements
}

func (u *BGMModel) GetSeparator(element string) string {
	return u.separators[element]
}
