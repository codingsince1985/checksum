package md5

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func MD5sum(filename string) (string, error) {
	const bufferSize = 65536

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	defer file.Close()

	hash := md5.New()
	for buf, reader := make([]byte, bufferSize), bufio.NewReader(file); ; {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}

			return "", err
		}

		io.WriteString(hash, string(buf[:n]))
	}

	checksum := fmt.Sprintf("%x", hash.Sum(nil))
	return checksum, nil
}
