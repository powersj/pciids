package query

import (
	"fmt"
	"testing"

	"github.com/powersj/pciids/pkg/file"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	file.Testing = true
	ids, err := All()

	if assert.NoError(t, err) {
		assert.NotEmpty(t, ids)
	}
}

func TestDevice(t *testing.T) {
	file.Testing = true

	tests := []struct {
		vendorID string
		deviceID string
		matches  int
	}{
		{"121a", "0009", 2},
		{"121a", "0003", 3},
		{"9999", "9999", 0},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("query=%s %s", tc.vendorID, tc.deviceID), func(t *testing.T) {
			ids, _ := Device(tc.vendorID, tc.deviceID)
			if tc.matches != len(ids) {
				t.Fatalf("want %d, got %d", tc.matches, len(ids))
			}
		})
	}
}

func TestSubDevice(t *testing.T) {
	file.Testing = true

	tests := []struct {
		vendorID    string
		deviceID    string
		subVendorID string
		subDeviceID string
		matches     int
	}{
		{"121a", "0003", "121a", "0001", 1},
		{"121a", "0003", "121a", "0003", 1},
		{"121a", "0009", "121a", "0003", 1},
		{"121a", "0009", "121a", "0009", 1},
		{"9999", "9999", "9999", "9999", 0},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("subQuery=%s %s", tc.vendorID, tc.deviceID), func(t *testing.T) {
			ids, _ := SubDevice(tc.vendorID, tc.deviceID, tc.subVendorID, tc.subDeviceID)
			if tc.matches != len(ids) {
				t.Fatalf("want %d, got %d", tc.matches, len(ids))
			}
		})
	}
}
