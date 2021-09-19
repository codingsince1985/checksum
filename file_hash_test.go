package checksum_test

import (
	"fmt"
	"github.com/codingsince1985/checksum"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func prepareFile() (string, error) {
	file, err := ioutil.TempFile("", "prefix")
	if err != nil {
		return "", err
	}
	if err := ioutil.WriteFile(file.Name(), []byte("some data"), 0600); err != nil {
		return "", err
	}
	return file.Name(), nil
}

func testFile(t *testing.T, checksumFunc func(string) error) {
	file, err := prepareFile()
	if err != nil {
		t.Logf("could not create test file: %s", err)
		t.FailNow()
	}
	defer func() {
		err := os.Remove(file)
		if err != nil {
			t.Logf("could not remove test file: %s", err)
		}
	}()
	if err := checksumFunc(file); err != nil {
		t.Error(err)
	}
}

func TestSHA1sumFile(t *testing.T) {
	testFile(t, func(filename string) error {
		if result, err := checksum.SHA1sum(filename); err != nil || result != "baf34551fecb48acc3da868eb85e1b6dac9de356" {
			return fmt.Errorf(result, err)
		}
		return nil
	})
}

func TestSHA256sumFile(t *testing.T) {
	testFile(t, func(filename string) error {
		if result, err := checksum.SHA256sum(filename); err != nil || result != "1307990e6ba5ca145eb35e99182a9bec46531bc54ddf656a602c780fa0240dee" {
			return fmt.Errorf(result, err)
		}
		return nil
	})
}

func TestMd5sumFile(t *testing.T) {
	testFile(t, func(filename string) error {
		if result, err := checksum.MD5sum(filename); err != nil || result != "1e50210a0202497fb79bc38b6ade6c34" {
			return fmt.Errorf(result, err)
		}
		return nil
	})
}

func TestCrc32File(t *testing.T) {
	testFile(t, func(filename string) error {
		if result, err := checksum.CRC32(filename); err != nil || result != "d9c2e91e" {
			return fmt.Errorf(result, err)
		}
		return nil
	})
}

func TestMd5sumDir(t *testing.T) {
	homeDirectory, err := homedir.Dir()
	if err != nil {
		t.Logf("could not get home directory: %s", err)
		t.FailNow()
	}
	file := path.Join(homeDirectory, "Downloads")

	if md5sum, err := checksum.MD5sum(file); err != nil || md5sum != "" {
		t.Error("Md5sum(dir) failed", md5sum, err)
	}
}
