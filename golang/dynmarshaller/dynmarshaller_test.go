package dynmarshaller

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type X interface {
	H() int
}

type C struct {
	Type string
	B    string
	C    int
}

func (c *C) H() int {
	return c.C
}

type A struct {
	Type string
	B    string
	C    int
}

func (a *A) Initialize() error {
	a.C = a.C + 1
	return nil
}

func (a *A) H() int {
	return a.C
}

func TestDynamicUnmarshaller_Normal(t *testing.T) {
	registry := map[string]X{
		"A": (*A)(nil),
		"C": (*C)(nil),
	}
	d := NewDynamicUnmarshaller[X](registry)
	data := []byte(`{"Type":"C","B":"B","C":1}`)
	x, err := d.Unmarshal(data)
	assert.Nil(t, err)
	assert.Equal(t, 1, x.H())
}

func TestDynamicUnmarshaller_Initializable(t *testing.T) {
	registry := map[string]X{
		"A": (*A)(nil),
	}
	d := NewDynamicUnmarshaller[X](registry)
	data := []byte(`{"Type":"A","B":"B","C":1}`)
	x, err := d.Unmarshal(data)
	assert.Nil(t, err)
	assert.Equal(t, 2, x.H())
	data = []byte(`{"Type":"B","B":"B","C":1}`)
	x, err = d.Unmarshal(data)
	assert.NotNil(t, err)
}
