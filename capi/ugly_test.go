package capi

import (
	"bytes"
	"testing"
)

// test serial + bruteforced code using DLL calls
const (
	XML_VECTOR            = "3806564553AEF4E69858FB1FA8165D51F7EA40293792D20F0600010F6301000000C802010600010200000000000000000000000000000000"
	TEST1_SERIAL_NUMBER   = "1234567"
	TEST1_ACTIVATION_CODE = "10000000000000000053"
	TEST2_SERIAL_NUMBER   = "7654321"
	TEST2_ACTIVATION_CODE = "10000000000000000326"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestMyValidPWD_E980(t *testing.T) {
	v11 := []byte{0x56, 0x45, 0x53, 0x30, 0x30, 0x39, 0x38, 0x30}
	vv := []byte{0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11}
	want := []byte{0x47, 0x54, 0x42, 0x21, 0x21, 0x28, 0x29, 0x21}
	actual := myValidPWD_E980(v11, vv, 8)
	if bytes.Compare(actual, want) != 0 {
		t.Fatalf(`E980() = %q, want match for %#q, nil`, actual, want)
	}
}

func TestMyValidPWD_18B0(t *testing.T) {
	v1 := []byte{0x47, 0x54, 0x42, 0x21, 0x21, 0x28, 0x29, 0x21}
	v2 := []byte{}
	want := []byte{0x53, 0x29, 0xb4, 0x78, 0x94, 0x04, 0x61, 0xfb}
	myValidPWD_18B0(v1, 0, v2)
	if bytes.Compare(want, want) != 0 {
		t.Fatalf(`18B0() = %q, want match for %#q, nil`, v2, want)
	}
}
