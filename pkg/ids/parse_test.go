package ids

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	var input string = "# test input\n" +
		"121a  3Dfx Interactive, Inc.\n" +
		"\t0001  Voodoo\n" +
		"\t0009  Voodoo 4 / Voodoo 5\n" +
		"\t\t121a 0003  Voodoo5 PCI 5500\n" +
		"\t\t121a 0009  Voodoo5 AGP 5500/6000\n"

	ids, err := Parse(input)

	if assert.NoError(t, err) {
		assert.Equal(t, 4, len(ids))
	}
}

func TestParseSubDevice(t *testing.T) {
	var output PCIID = parseSubDevice("\t\t121a 0009  Voodoo5 AGP 5500/6000")

	fmt.Println(output)

	assert.Equal(t, "121a", output.VendorID)
	assert.Equal(t, "3Dfx Interactive, Inc.", output.VendorName)
	assert.Equal(t, "0009", output.DeviceID)
	assert.Equal(t, "Voodoo 4 / Voodoo 5", output.DeviceName)

	assert.Equal(t, "121a", output.SubVendorID)
	assert.Equal(t, "3Dfx Interactive, Inc.", output.SubVendorName)
	assert.Equal(t, "0009", output.SubDeviceID)
	assert.Equal(t, "Voodoo5 AGP 5500/6000", output.SubDeviceName)
}

func TestParseDevice(t *testing.T) {
	var output PCIID = parseDevice("\t0009  Voodoo 4 / Voodoo 5")

	assert.Equal(t, "121a", output.VendorID)
	assert.Equal(t, "3Dfx Interactive, Inc.", output.VendorName)
	assert.Equal(t, "0009", output.DeviceID)
	assert.Equal(t, "Voodoo 4 / Voodoo 5", output.DeviceName)
}

func TestParseVendors(t *testing.T) {
	var input string = "# comment\n" +
		"2077 Araska\n" +
		" 0001 Militech\n" +
		"\n" +
		"\t0009\n" +
		"\t\t121a 0009\n" +
		"C 09"

	var vendors map[string]string = parseVendors(input)

	assert.Equal(t, 2, len(vendors))
}

func TestValid(t *testing.T) {
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
			var result bool = valid(tc.input)
			if result != tc.result {
				t.Fatalf("got %v; want %v", result, tc.result)
			}
		})
	}
}

func TestValidVendor(t *testing.T) {
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
			var result bool = validVendor(tc.input)
			if result != tc.result {
				t.Fatalf("got %v; want %v", result, tc.result)
			}
		})
	}
}
