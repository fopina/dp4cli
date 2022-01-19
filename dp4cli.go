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
	"unsafe"

	"github.com/fopina/dp4cli/dll"
)

const (
	XML_URL   = "http://sc.vasco.com/update/dp4windows/50/digipass.xml"
	MAGIC_PIN = "111111"
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

	out1, out2, err := dll.Activate(vector, serial, code, MAGIC_PIN)

	if err != nil {
		return err
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

	out3, err := dll.ValidPWD(out1, out2, MAGIC_PIN)

	if err != nil {
		return "", err
	}

	fmt.Println(out1)
	fmt.Println(out2)
	fmt.Println(string(out3[:]))
	//fmt.Println(capi.ValidPWD(out1, out2, MAGIC_PIN))

	pin, err := dll.GenPassword(out1, out2, out3)

	if err != nil {
		return "", err
	}

	return pin, nil
}

func main() {
	//defer syscall.FreeLibrary(dp4capi)

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
