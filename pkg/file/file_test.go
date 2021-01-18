package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLatest(t *testing.T) {
	output, err := Latest()

	if assert.NoError(t, err) {
		assert.Contains(t, output, "8086")
	}
}

func TestLatestTesting(t *testing.T) {
	Testing = true
	output, err := Latest()

	if assert.NoError(t, err) {
		assert.Contains(t, output, "121a")
	}
}
