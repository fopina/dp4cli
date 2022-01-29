package capi

/*
All these tests should use constants and match expected values from dll_test.go
*/

import (
	"bytes"
	"encoding/hex"
	"testing"

	shared "github.com/fopina/dp4cli/shared4tests"
	"github.com/stretchr/testify/assert"
)

func TestActivate1(t *testing.T) {
	out1, out2, err := Activate(shared.XML_VECTOR, shared.TEST1_SERIAL_NUMBER, shared.TEST1_ACTIVATION_CODE, shared.MAGIC_PIN)

	assert.Nil(t, err)
	assert.Equal(t, shared.TEST1_ACTIVATE_KEY1, hex.EncodeToString(out1))
	assert.Equal(t, shared.TEST1_ACTIVATE_KEY2, hex.EncodeToString(out2))
}

func TestActivate2(t *testing.T) {
	out1, out2, err := Activate(shared.XML_VECTOR, shared.TEST2_SERIAL_NUMBER, shared.TEST2_ACTIVATION_CODE, shared.MAGIC_PIN)

	assert.Nil(t, err)
	assert.Equal(t, shared.TEST2_ACTIVATE_KEY1, hex.EncodeToString(out1))
	assert.Equal(t, shared.TEST2_ACTIVATE_KEY2, hex.EncodeToString(out2))
}

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
