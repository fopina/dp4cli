package capi

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"
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
	out, ret := myValidPWD_7760(key1, key2, magicPin)
	return out, ret
}

func Activate(vector, serial, code, magicPin string) ([]byte, []byte, error) {
	out1 := make([]byte, 100)
	out2 := make([]byte, 100)

	ret, _, callErr := syscall.Syscall9(uintptr(fAtivate), 8, stringConvert(vector), stringConvert(strings.TrimSpace(serial)), stringConvert(strings.TrimSpace(code)), 0, stringConvert(magicPin), 0, uintptr(unsafe.Pointer(&out1[0])), uintptr(unsafe.Pointer(&out2[0])), 0)

	if callErr != 0 {
		return nil, nil, callErr
	}

	retInt := int(ret)
	if retInt != 0 {
		return nil, nil, fmt.Errorf("DP4C_Activate returned %d", retInt)
	}

	return out1, out2, nil
}

func GenPassword(out1, out2, out3 []byte) (string, error) {
	out4 := make([]byte, 100)
	out5 := make([]byte, 100)

	ret, _, callErr := syscall.Syscall9(uintptr(fGenPassword), 7, uintptr(unsafe.Pointer(&out1[0])), uintptr(unsafe.Pointer(&out2[0])), 0, uintptr(unsafe.Pointer(&out3[0])), 0, uintptr(unsafe.Pointer(&out4[0])), uintptr(unsafe.Pointer(&out5[0])), 0, 0)
	if callErr != 0 {
		return "", callErr
	}
	retInt := int(ret)
	if retInt != 0 {
		return "", fmt.Errorf("DP4C_GenPasswordEx returned %d", retInt)
	}

	return string(out4[:6]), nil
}
