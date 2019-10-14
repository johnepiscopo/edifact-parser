package ediparser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SingleSegment(t *testing.T) {
	s := "UNH+123456+DESADV:D:96A:UN"

	e := new(EdiParser)

	m := e.ParseEdifact(s, []string{})

	assert.Equal(t, 1, len(m))
	assert.Equal(t, "UNH", m[0].Name)
	assert.Equal(t, "123456", m[0].Elements[0].Value)
	assert.Equal(t, "DESADV", m[0].Elements[1].SubElements[0].Value,
		fmt.Sprintf("Got %s Expected DESADV", m[0].Elements[1].SubElements[0].Value))
	assert.Equal(t, "D", m[0].Elements[1].SubElements[1].Value,
		fmt.Sprintf("Got %s Expected D", m[0].Elements[1].SubElements[1].Value))
	assert.Equal(t, "96A", m[0].Elements[1].SubElements[2].Value,
		fmt.Sprintf("Got %s Expected 96A", m[0].Elements[1].SubElements[2].Value))
	assert.Equal(t, "UN", m[0].Elements[1].SubElements[3].Value,
		fmt.Sprintf("Got %s Expected UN", m[0].Elements[1].SubElements[3].Value))
	assert.Equal(t, "", m[0].Elements[1].SubElements[4].Value,
		fmt.Sprintf("Got %s Expected ", m[0].Elements[1].SubElements[4].Value))
}

func Test_MultipleSegments(t *testing.T) {
	s := "UNH+123456+DESADV:D:96A:UN'BGM+351+012345+9"

	e := new(EdiParser)

	m := e.ParseEdifact(s, []string{})

	assert.Equal(t, 2, len(m))
	//m[0]
	assert.Equal(t, "UNH", m[0].Name)
	assert.Equal(t, "123456", m[0].Elements[0].Value)
	assert.Equal(t, "DESADV", m[0].Elements[1].SubElements[0].Value,
		fmt.Sprintf("Got %s Expected DESADV", m[0].Elements[1].SubElements[0].Value))
	assert.Equal(t, "D", m[0].Elements[1].SubElements[1].Value,
		fmt.Sprintf("Got %s Expected D", m[0].Elements[1].SubElements[1].Value))
	assert.Equal(t, "96A", m[0].Elements[1].SubElements[2].Value,
		fmt.Sprintf("Got %s Expected 96A", m[0].Elements[1].SubElements[2].Value))
	assert.Equal(t, "UN", m[0].Elements[1].SubElements[3].Value,
		fmt.Sprintf("Got %s Expected UN", m[0].Elements[1].SubElements[3].Value))
	assert.Equal(t, "", m[0].Elements[1].SubElements[4].Value,
		fmt.Sprintf("Got %s Expected ", m[0].Elements[1].SubElements[4].Value))

	//m[1]
	assert.Equal(t, "BGM", m[1].Name)
	assert.Equal(t, "351", m[1].Elements[0].SubElements[0].Value)
	assert.Equal(t, "012345", m[1].Elements[1].Value)
	assert.Equal(t, "9", m[1].Elements[2].Value)

}

func Test_IgnoreElements(t *testing.T) {
	s := "BGM+++9"

	e := new(EdiParser)

	m := e.ParseEdifact(s, []string{"MessageNameCoded", "MessageNumber"})

	assert.Equal(t, 1, len(m))
	//m[0]
	assert.Equal(t, "BGM", m[0].Name)
	assert.Equal(t, "", m[0].Elements[0].SubElements[0].Value, fmt.Sprintf("Expected '' Got %s", m[0].Elements[0].SubElements[0].Value))
	assert.Equal(t, "", m[0].Elements[1].Value, fmt.Sprintf("Expected '' Got %s", m[0].Elements[1].Value))
	assert.Equal(t, "9", m[0].Elements[2].Value)
}
