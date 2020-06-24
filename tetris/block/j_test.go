package block

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestBlockJ_Rotate(t *testing.T) {
	j := NewBlockJ(10,10)

	t.Log(j.Blocks)

	o := j.Origin()
	assert.Equal(t,o.X ,10,"Origin")
	assert.Equal(t,o.Y ,10,"Origin")
	assert.Equal(t,j.status() ,BlockJUp,"status()")
}
