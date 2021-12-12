package wordnumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestVersion is a cynical test case for coverage
func TestVersion(t *testing.T) {
	v := Version()
	assert.Equal(t, "unknown-snapshot", v)
}
