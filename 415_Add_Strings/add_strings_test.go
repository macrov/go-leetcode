package Add_Strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddStrings(t *testing.T) {
	assert.Equal(t, "10", addStrings("1", "9"))
}
