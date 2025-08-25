package AsciiArt

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"os"
)

func CheckFileHashing(fileName string) bool {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	buf := make([]byte, 30*1024)
	sha256 := sha256.New()
	for {
		n, err := file.Read(buf)
		if n > 0 {
			_, err := sha256.Write(buf[:n])
			if err != nil {
				log.Fatal(err)
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Read %d bytes: %v", n, err)
			break
		}
	}
	sum1 := sha256.Sum(nil)
	sum := hex.EncodeToString(sum1)
	standard := "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	if fileName == "standard.txt" && string(sum) == standard {
		return true
	}
	return false
}
