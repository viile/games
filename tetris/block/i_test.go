package block

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestBlockI_Rotate(t *testing.T) {
	i := BlockI{Te{[]Block{
		{7, 5, true},
		{8, 5, false},
		{9, 5, false},
		{10, 5, false}}}}

	assert.Equal(t, i.iss(), false, "iss")

	t.Log(i.bs())

	i = BlockI{Te{[]Block{
		{7, 5, true},
		{7, 4, false},
		{7, 3, false},
		{7, 2, false}}}}

	assert.Equal(t, i.iss(), true, "iss")
	t.Log(i.bh())
}
