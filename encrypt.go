package easyencdec

import (
	"crypto/aes"
	"crypto/cipher"
	crand "crypto/rand"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	mrand "math/rand"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Init() {
	var src, dst, key string

	flag.StringVar(&src, "filename", "", "flag for input file being encrypted")
	flag.StringVar(&dst, "outputfile", defaultOut(src), "optional output file, defaults to <<filename>>.crd")
	flag.StringVar(&key, "keyfile", "", "file containing the key used to encrypt")
	flag.Parse()

	if err := EncryptFile(src, dst, key); err != nil {
		log.Panicln("<<file failed to encrypt>>")
	}
}

func defaultOut(fn string) string {
	b := filepath.Base(fn)
	if len(filepath.Ext(b)) > 0 {
		b = strings.ReplaceAll(b, filepath.Ext(b), "")
	}
	return b
}

func GetMacAddress() ([]byte, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	var currentIP, currentNetworkHardwareName string

	for _, address := range addrs {

		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				currentIP = ipnet.IP.String()
			}
		}
	}

	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, interf := range interfaces {

		if addrs, err := interf.Addrs(); err == nil {
			for _, addr := range addrs {

				// only interested in the name with current IP address
				if strings.Contains(addr.String(), currentIP) {
					currentNetworkHardwareName = interf.Name
				}
			}
		}
	}

	// extract the hardware information base on the interface name
	// capture above
	netInterface, err := net.InterfaceByName(currentNetworkHardwareName)
	if err != nil {
		return nil, err
	}

	macAddress := netInterface.HardwareAddr

	hwAddr, err := net.ParseMAC(macAddress.String())

	if err != nil {
		fmt.Println("No able to parse MAC address : ", err)
		return nil, err
	}

	return []byte(hwAddr.String()), nil
}

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = mrand.NewSource(time.Now().UnixNano())

func GenerateKey(n int) []byte {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return b
}

func GenerateKeyFile(output string) error {
	n := 15
	bytes := GenerateKey(n)

	if _, err := os.Stat(filepath.Dir(output)); os.IsNotExist(err) {
		if err = os.Mkdir(filepath.Dir(output), 0700); err != nil {
			return err
		}
	}

	err := os.WriteFile(output, bytes, 0700)
	if err != nil {
		return err
	}
	return nil
}

func CombineKeyAndMac(keyFile string) ([]byte, error) {
	b, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}
	m, err := GetMacAddress()
	if err != nil {
		return nil, err
	}
	comb := append(b, m...)
	return comb, nil
}

func EncryptFile(src string, dst string, key string) error {

	body, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	comb, err := CombineKeyAndMac(key)
	if err != nil {
		return err
	}

	c, err := aes.NewCipher(comb)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(crand.Reader, nonce); err != nil {
		return err
	}

	err = ioutil.WriteFile(dst, gcm.Seal(nonce, nonce, body, nil), 0777)
	if err != nil {
		return err
	}
	return nil
}
