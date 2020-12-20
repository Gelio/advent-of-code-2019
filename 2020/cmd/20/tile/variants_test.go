package tile

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAllVariants(t *testing.T) {
	cases := []struct {
		input         string
		variantsCount int
	}{
		{input: "Tile 1:\n123\n456\n789", variantsCount: 4 * 3},
		{input: "Tile 2:\n111\n111\n111", variantsCount: 1},
		{input: "Tile 2:\n111\n222\n111", variantsCount: 2},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			tile, err := Parse(strings.Split(tt.input, "\n"))

			require.NoError(t, err, "parsing tile")

			variants := tile.GetAllVariants()

			assert.Len(t, variants, tt.variantsCount, "invalid number of variants")
		})
	}
}
