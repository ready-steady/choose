package choose

import (
	"testing"

	"github.com/ready-steady/assert"
)

func TestUniformUint(t *testing.T) {
	cases := []struct {
		sequence []uint
		choose   uint
		chosen   []uint
	}{
		{
			sequence: []uint{11, 21, 31, 41, 51},
			choose:   0,
			chosen:   []uint{},
		},
		{
			sequence: []uint{11, 21, 31, 41, 51},
			choose:   1,
			chosen:   []uint{11},
		},
		{
			sequence: []uint{11, 21, 31, 41, 51},
			choose:   2,
			chosen:   []uint{11, 51},
		},
		{
			sequence: []uint{11, 21, 31, 41, 51},
			choose:   10,
			chosen:   []uint{11, 21, 31, 41, 51},
		},
		{
			sequence: []uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 21, 31, 35, 41},
			choose:   5,
			chosen:   []uint{1, 11, 21, 31, 41},
		},
		{
			sequence: []uint{0, 2, 3, 40, 100},
			choose:   4,
			chosen:   []uint{0, 3, 40, 100},
		},
	}

	for _, c := range cases {
		indices := UniformUint(c.sequence, c.choose)
		chosen := make([]uint, len(indices))
		for i, j := range indices {
			chosen[i] = c.sequence[j]
		}
		assert.Equal(chosen, c.chosen, t)
	}
}
