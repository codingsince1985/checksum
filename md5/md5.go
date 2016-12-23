// Package md5 computes MD5 checksum for large files
package md5

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

const bufferSize = 65536

// MD5sum returns MD5 checksum of filename
func MD5sum(filename string) (string, error) {
	if info, err := os.Stat(filename); err != nil || info.IsDir() {
		return "", err
	}

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	buf := make([]byte, bufferSize)
	reader := bufio.NewReader(file)
	for {
		if n, err := reader.Read(buf); err == nil {
			hash.Write(buf[:n])
		} else if err == io.EOF {
			break
		} else {
			return "", err
		}
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
