package dll

import (
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
