package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"unsafe"

	"github.com/fopina/dp4cli/dll"
)

const (
	XML_URL = "http://sc.vasco.com/update/dp4windows/50/digipass.xml"
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

func crack(vector, serial string) string {
	code := uint64(10000000000000000000)

	for {
		codeStr := strconv.FormatUint(code, 10)
		_, _, err := dll.Activate(vector, serial, codeStr)
		if err == nil {
			return codeStr
		}
		code++
	}
	return ""
}

func main() {
	//defer syscall.FreeLibrary(dp4capi)

	err := downloadXML()
	if err != nil {
		panic(err)
	}

	vector, err := loadStaticVector()
	if err != nil {
		panic(err)
	}

	code := crack(vector, "1234567")
	fmt.Println(code)
	code = crack(vector, "7654321")
	fmt.Println(code)
}
