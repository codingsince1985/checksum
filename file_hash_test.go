package checksum

import (
	m "crypto/md5"
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

func TestTransition(t *testing.T) {
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

	md5sum, errMD5 := MD5sum(file)
	if errMD5 != nil {
		t.Logf("error calculating old hash: %s", err)
		t.FailNow()
	}
	sum, errSum := Sum(m.New(), file)
	if errSum != nil {
		t.Logf("error calculating new hash: %s", err)
		t.FailNow()
	}
	if md5sum != sum {
		t.Logf("hash mismatch: old [%s] != new [%s]", md5sum, sum)
		t.Fail()
	}
}

func TestMd5sumFile(t *testing.T) {
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

	if md5sum, err := MD5sum(file); err != nil || md5sum != "1e50210a0202497fb79bc38b6ade6c34" {
		t.Error("Md5sum(file) failed", md5sum, err)
	}
}

func TestMd5sumDir(t *testing.T) {
	homeDirectory, err := homedir.Dir()
	if err != nil {
		t.Logf("could not get home directory: %s", err)
		t.FailNow()
	}
	file := path.Join(homeDirectory, "Downloads")

	if md5sum, err := MD5sum(file); err != nil || md5sum != "" {
		t.Error("Md5sum(dir) failed", md5sum, err)
	}
}
