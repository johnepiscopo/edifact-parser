package edisegments

type BGMModel struct {
	elements   []string
	separators map[string]string
}

func NewBGMModel() *BGMModel {
	m := new(BGMModel)
	m.elements = []string{
		"MessageName",
		"MessageNameCoded",
		"MessageNumber",
		"MessageFunctionCoded",
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
