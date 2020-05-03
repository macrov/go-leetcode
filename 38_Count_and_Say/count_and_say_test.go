package Count_and_Say

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountAndSay(t *testing.T) {
	assert.Equal(t, countAndSay(1), "1")
	assert.Equal(t, countAndSay(2), "11", "should be the same")
	assert.Equal(t, countAndSay(3), "21", "should be the same")
	assert.Equal(t, countAndSay(4), "1211", "should be the same")
	assert.Equal(t, countAndSay(5), "111221", "should be the same")
}
