package edisegments

type Segmenter interface {
	GetElements() []string
	GetSeparator(element string) string
}

var Segments = map[string]Segmenter{
	"UNH": NewUNHModel(),
	"BGM": NewBGMModel(),
}
