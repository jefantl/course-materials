package hscan

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
)

//==========================================================================\\

var shalookup sync.Map
var md5lookup sync.Map

func GuessSingle(sourceHash string, filename string) (string, error) {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		password := scanner.Text()

		if len(sourceHash) == 32 {
			hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
			if hash == sourceHash {
				fmt.Printf("[+] Password found (MD5): %s\n", password)
				return password, nil
			}
		} else if len(sourceHash) == 64 {
			hash := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
			if hash == sourceHash {
				fmt.Printf("[+] Password found (SHA-256): %s\n", password)
				return password, nil
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
		return "", err
	}

	return "", fmt.Errorf("no passwords found")
}

func GenHashMaps(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		scannedText := scanner.Text()

		go func(password string) {
			hashMD5 := fmt.Sprintf("%x", md5.Sum([]byte(password)))
			md5lookup.Store(hashMD5, password)
		}(scannedText)
		go func(password string) {
			hashSHA := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
			shalookup.Store(hashSHA, password)
		}(scannedText)
	}
}

func GetSHA(hash string) (string, error) {
	password, ok := shalookup.Load(hash)
	if ok {
		return password.(string), nil

	} else {

		return "", errors.New("password does not exist")

	}
}

func GetMD5(hash string) (string, error) {
	password, ok := md5lookup.Load(hash)
	if ok {
		return password.(string), nil

	} else {

		return "", errors.New("password does not exist")

	}
}
