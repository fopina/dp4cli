//go:build !windows
// +build !windows

package wrapper

import (
	"github.com/fopina/dp4cli/capi"
)

func Activate(vector, serial, code, magicPin string) ([]byte, []byte, error) {
	return capi.Activate(vector, serial, code, magicPin)
}

func ValidPWD(out1, out2 []byte, magicPin string) ([]byte, error) {
	return capi.ValidPWD(out1, out2, magicPin)
}

func GenPassword(out1, out2, out3 []byte) (string, error) {
	return capi.GenPassword(out1, out2, out3)
}

func CleanUp() {
}
