//go:build windows
// +build windows

package dll

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"
)

const (
	DLL_FILE = "DP4CAPI.dll"
)

var (
	dp4capi, _      = syscall.LoadLibrary(DLL_FILE)
	fAtivate, _     = syscall.GetProcAddress(dp4capi, "DP4C_Activate")
	fValidPWD, _    = syscall.GetProcAddress(dp4capi, "DP4C_validPWD")
	fGenPassword, _ = syscall.GetProcAddress(dp4capi, "DP4C_GenPasswordEx")
)

func stringConvert(s string) uintptr {
	b := append([]byte(s), 0)
	return uintptr(unsafe.Pointer(&b[0]))
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

func ValidPWD(out1, out2 []byte, magicPin string) ([]byte, error) {
	out3 := make([]byte, 100)

	ret, _, callErr := syscall.Syscall6(uintptr(fValidPWD), 4, uintptr(unsafe.Pointer(&out1[0])), uintptr(unsafe.Pointer(&out2[0])), stringConvert(magicPin), uintptr(unsafe.Pointer(&out3[0])), 0, 0)

	if callErr != 0 {
		return nil, callErr
	}
	retInt := int(ret)
	if retInt != 1 {
		return nil, fmt.Errorf("DP4C_validPWD returned %d", retInt)
	}
	return out3, nil
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

func CleanUp() {
	syscall.FreeLibrary(dp4capi)
}
