package edisegments

import (
	"fmt"
)

//Segments is the map that maps the Segment identifier to the EDI Segment model
var Segments = map[string]*EdiSegment{
	"UNH": NewUNHModel(),
	"BGM": NewBGMModel(),
	"DTM": NewDTMModel(),
	"RFF": NewRFFModel(),
	"NAD": NewNADModel(),
	"TDT": NewTDTModel(),
	"LOC": NewLOCModel(),
	"CPS": NewCPSModel(),
	"PAC": NewPACModel(),
	"MEA": NewMEAModel(),
	"UNB": NewUNBModel(),
}

//SegmentError is a placeholder comment
type SegmentError struct {
	segmentCode string
}

func (s *SegmentError) Error() string {
	return fmt.Sprintf("%s is not a valid segment", s.segmentCode)
}

//NewSegmentError returns a new segment error object
func NewSegmentError(segmentCode string) error {
	return &SegmentError{segmentCode: segmentCode}
}
