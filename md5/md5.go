// Package md5 computes MD5 checksum for large files
package md5

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"hash"
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
	defer func() { _ = file.Close() }()

	hasher := md5.New()
	buf := make([]byte, bufferSize)
	reader := bufio.NewReader(file)
here:
	for {
		switch n, err := reader.Read(buf); err {
		case nil:
			hasher.Write(buf[:n])
		case io.EOF:
			break here
		default:
			return "", err
		}
	}
	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}

// Sum calculates the hash based on a provided hash provider
func Sum(hashAlgorithm hash.Hash, filename string) (string, error) {
	if info, err := os.Stat(filename); err != nil || info.IsDir() {
		return "", err
	}

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer func() { _ = file.Close() }()

	buf := make([]byte, bufferSize)
	reader := bufio.NewReader(file)
here:
	for {
		switch n, err := reader.Read(buf); err {
		case nil:
			hashAlgorithm.Write(buf[:n])
		case io.EOF:
			break here
		default:
			return "", err
		}
	}
	return fmt.Sprintf("%x", hashAlgorithm.Sum(nil)), nil
}
