package first_missing_positive

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirstMissingPositive(t *testing.T) {
	assert.Equal(t,3, firstMissingPositive([]int{1,2,0}))
	assert.Equal(t,3, firstMissingPositive([]int{1,2}))
	assert.Equal(t,3, firstMissingPositive([]int{1,2,0,4}))
}