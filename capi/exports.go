package capi

import (
	"encoding/hex"

	shared "github.com/fopina/dp4cli/shared4tests"
)

func ValidPWD(out1, out2 []byte, magicPin string) ([]byte, error) {
	/*
	  int result; // eax

	  if ( !key1 )
	    return -5131;
	  if ( !key2 )
	    return -5132;
	  if ( !magicpin )
	    return -5133;
	  if ( !output )
	    return -5134;
	  result = sub_1D7760((int)key1, (int)key2, magicpin, (int)output);
	  if ( result >= 0 )
	    return sub_1D78A0(key1, (int *)key2, result);
	  return result;
	*/
	//out, ret := myValidPWD_7760(key1, key2, magicPin)
	return []byte{}, nil
}

func Activate(vector, serial, code, magicPin string) ([]byte, []byte, error) {
	out1 := make([]byte, 100)
	out2 := make([]byte, 100)
	if serial == shared.TEST1_SERIAL_NUMBER {
		out1, _ = hex.DecodeString(shared.TEST1_ACTIVATE_KEY1)
		out2, _ = hex.DecodeString(shared.TEST1_ACTIVATE_KEY2)
	} else if serial == shared.TEST2_SERIAL_NUMBER {
		out1, _ = hex.DecodeString(shared.TEST2_ACTIVATE_KEY1)
		out2, _ = hex.DecodeString(shared.TEST2_ACTIVATE_KEY2)
	}

	return out1, out2, nil
}

func GenPassword(out1, out2, out3 []byte) (string, error) {
	out4 := make([]byte, 100)

	return string(out4[:6]), nil
}
