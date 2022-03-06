//go:build windows
// +build windows

package wrapper

import (
	"github.com/fopina/dp4cli/dll"
)

func Activate(vector, serial, code, magicPin string) ([]byte, []byte, error) {
	return dll.Activate(vector, serial, code, magicPin)
}

func ValidPWD(out1, out2 []byte, magicPin string) ([]byte, error) {
	return dll.ValidPWD(out1, out2, magicPin)
}

func GenPassword(out1, out2, out3 []byte) (string, error) {
	return dll.GenPassword(out1, out2, out3)
}

func CleanUp() {
	dll.CleanUp()
}
