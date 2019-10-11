package edisegments

// type LOCModel struct {
// 	elements   []string
// 	separators map[string]string
// }

func NewLOCModel() *EdiSegment {
	m := new(EdiSegment)

	m.Name = "LOC"
	m.Elements = []*EdiElement{
		&EdiElement{
			Name: "LocationQualifier",
			Ref:  "LOC010",
			Tag:  "3227",
		},
		&EdiElement{
			Name: "LocationIdentifier",
			Ref:  "LOC020",
			Tag:  "C517",
			SubElements: []*EdiElement{
				&EdiElement{
					Name: "PlaceLocationIdentifier",
					Ref:  "LOC020-010",
					Tag:  "3225",
				},
			},
		},
	}

	// m.elements = []string{
	// 	"LocationQualifier",       //Ref: LOC010 	 Tag: 3227
	// 	"LocationIdentifier",      //Ref: LOC020 	 Tag: C517
	// 	"PlaceLocationIdentifier", //Ref: LOC020-010 Tag: 3225
	// }

	// m.separators = map[string]string{
	// 	"LocationQualifier":       "+",
	// 	"LocationIdentifier":      "+",
	// 	"PlaceLocationIdentifier": ""}

	return m
}

// func (u *LOCModel) GetElements() []string {
// 	return u.elements
// }

// func (u *LOCModel) GetSeparator(element string) string {
// 	return u.separators[element]
// }
