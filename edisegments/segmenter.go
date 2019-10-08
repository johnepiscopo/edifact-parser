package edisegments

//Segmenter is the interface defining the contract
//for each of the EDI Segments
type Segmenter interface {
	GetElements() []string
	GetSeparator(element string) string
}

//Segments is the map that maps the Segment identifier to the EDI Segment model
var Segments = map[string]Segmenter{
	"UNH": NewUNHModel(),
	"BGM": NewBGMModel(),
}
