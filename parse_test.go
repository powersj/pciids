package pciids

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	t.Parallel()

	input := "# test input\n" +
		"121a  3Dfx Interactive, Inc.\n" +
		"\t0001  Voodoo\n" +
		"\t0009  Voodoo 4 / Voodoo 5\n" +
		"\t\t121a 0003  Voodoo5 PCI 5500\n" +
		"\t\t121a 0009  Voodoo5 AGP 5500/6000\n"

	ids := parse(input)

	assert.Equal(t, 4, len(ids))
}

func TestParseDevice(t *testing.T) {
	pciid := parseDevice("\t0009  Voodoo 4 / Voodoo 5")

	assert.Equal(t, "", pciid.VendorID)
	assert.Equal(t, "", pciid.VendorName)
	assert.Equal(t, "0009", pciid.DeviceID)
	assert.Equal(t, "Voodoo 4 / Voodoo 5", pciid.DeviceName)
}

func TestParseSubDevice(t *testing.T) {
	pciid := parseSubDevice("\t\t121a 0009  Voodoo5 AGP 5500/6000")

	assert.Equal(t, "", pciid.VendorID)
	assert.Equal(t, "", pciid.VendorName)
	assert.Equal(t, "", pciid.DeviceID)
	assert.Equal(t, "", pciid.DeviceName)

	assert.Equal(t, "121a", pciid.SubVendorID)
	assert.Equal(t, "", pciid.SubVendorName)
	assert.Equal(t, "0009", pciid.SubDeviceID)
	assert.Equal(t, "Voodoo5 AGP 5500/6000", pciid.SubDeviceName)
}

func TestParseVendors(t *testing.T) {
	t.Parallel()

	input := "# comment\n" +
		"2077 Araska\n" +
		" 0001 Militech\n" +
		"\n" +
		"\t0009\n" +
		"\t\t121a 0009\n" +
		"C 09"

	vendors := parseVendors(input)

	assert.Equal(t, 2, len(vendors))
}

func TestValid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input  string
		result bool
	}{
		{"121a  3Dfx Interactive, Inc.", true},
		{"\t0009  Voodoo 4 / Voodoo 5", true},
		{"\t\t121a 0009  Voodoo5 AGP 5500/6000", true},
		{"", false},
		{"# comment", false},
		{"C 09  Input device controller", false},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("valid=%s", tc.input), func(t *testing.T) {
			t.Parallel()

			result := valid(tc.input)
			if result != tc.result {
				t.Fatalf("got %v; want %v", result, tc.result)
			}
		})
	}
}

func TestValidVendor(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input  string
		result bool
	}{
		{"121a  3Dfx Interactive, Inc.", true},
		{"\t0009  Voodoo 4 / Voodoo 5", false},
		{"\t\t121a 0009  Voodoo5 AGP 5500/6000", false},
		{"", false},
		{"# comment", false},
		{"C 09  Input device controller", false},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("valid=%s", tc.input), func(t *testing.T) {
			t.Parallel()

			result := validVendor(tc.input)
			if result != tc.result {
				t.Fatalf("got %v; want %v", result, tc.result)
			}
		})
	}
}
