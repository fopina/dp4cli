package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"
	"unsafe"
)

const (
	XML_URL   = "http://sc.vasco.com/update/dp4windows/50/digipass.xml"
	DLL_FILE  = "DP4CAPI.dll"
	MAGIC_PIN = "111111"
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

func configDir() string {
	dirname, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}
	d := filepath.Join(dirname, "dp4cli")
	if _, err := os.Stat(d); os.IsNotExist(err) {
		os.MkdirAll(d, 0700)
	}
	return d
}

func downloadXML() error {
	// Get the data
	resp, err := http.Get(XML_URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	cd := configDir()
	out, err := os.Create(filepath.Join(cd, "digipass.xml"))
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func loadStaticVector() (string, error) {
	cd := configDir()
	content, err := ioutil.ReadFile(filepath.Join(cd, "digipass.xml"))
	if err != nil {
		return "", nil
	}
	re := regexp.MustCompile(`<StaticVector>(.*?)<\/StaticVector>`)
	m := re.FindSubmatch(content)
	if len(m) < 2 {
		return "", fmt.Errorf("StaticVector not found in XML")
	}
	return string(m[1]), nil
}

func activate() error {
	err := downloadXML()
	if err != nil {
		return err
	}

	vector, err := loadStaticVector()
	if err != nil {
		return err
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Serial: ")
	serial, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	fmt.Print("Activation Code: ")
	code, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	var out1, out2 [100]byte

	ret, _, callErr := syscall.Syscall9(uintptr(fAtivate), 8, stringConvert(vector), stringConvert(strings.TrimSpace(serial)), stringConvert(strings.TrimSpace(code)), 0, stringConvert(MAGIC_PIN), 0, uintptr(unsafe.Pointer(&out1[0])), uintptr(unsafe.Pointer(&out2[0])), 0)

	if callErr != 0 {
		return callErr
	}

	retInt := int(ret)
	if retInt != 0 {
		return fmt.Errorf("DP4C_Activate returned %d", retInt)
	}

	cd := configDir()
	err = ioutil.WriteFile(filepath.Join(cd, "key1.dat"), out1[:], 0600)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.Join(cd, "key2.dat"), out2[:], 0600)
	if err != nil {
		return err
	}
	return nil
}

func generatePIN() (string, error) {
	cd := configDir()
	out1, err := ioutil.ReadFile(filepath.Join(cd, "key1.dat"))
	if err != nil {
		return "", err
	}
	out2, err := ioutil.ReadFile(filepath.Join(cd, "key2.dat"))
	if err != nil {
		return "", err
	}

	var out3 [100]byte
	var out4, out5 [100]byte

	ret, _, callErr := syscall.Syscall6(uintptr(fValidPWD), 4, uintptr(unsafe.Pointer(&out1[0])), uintptr(unsafe.Pointer(&out2[0])), stringConvert(MAGIC_PIN), uintptr(unsafe.Pointer(&out3[0])), 0, 0)

	if callErr != 0 {
		return "", callErr
	}
	retInt := int(ret)
	if retInt != 1 {
		return "", fmt.Errorf("DP4C_validPWD returned %d", retInt)
	}

	ret, _, callErr = syscall.Syscall9(uintptr(fGenPassword), 7, uintptr(unsafe.Pointer(&out1[0])), uintptr(unsafe.Pointer(&out2[0])), 0, uintptr(unsafe.Pointer(&out3[0])), 0, uintptr(unsafe.Pointer(&out4[0])), uintptr(unsafe.Pointer(&out5[0])), 0, 0)
	if callErr != 0 {
		return "", callErr
	}
	retInt = int(ret)
	if retInt != 0 {
		return "", fmt.Errorf("DP4C_GenPasswordEx returned %d", retInt)
	}

	return string(out4[:6]), nil
}

func main() {
	defer syscall.FreeLibrary(dp4capi)

	setup := flag.Bool("setup", false, "Activate")
	flag.Parse()

	if *setup {
		err := activate()
		if err != nil {
			log.Fatal(err)
		}
	}

	pin, err := generatePIN()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pin)
}
