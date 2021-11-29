package checksum_test

import (
	"github.com/codingsince1985/checksum"
	"strings"
	"testing"
)

func TestSHA1sumReader(t *testing.T) {
	if result, err := checksum.SHA1sumReader(strings.NewReader("some data")); err != nil || result != "baf34551fecb48acc3da868eb85e1b6dac9de356" {
		t.Error(result, err)
	}
}

func TestSHA256sumReader(t *testing.T) {
	if result, err := checksum.SHA256sumReader(strings.NewReader("some data")); err != nil || result != "1307990e6ba5ca145eb35e99182a9bec46531bc54ddf656a602c780fa0240dee" {
		t.Error(result, err)
	}
}

func TestMd5sumReader(t *testing.T) {
	if result, err := checksum.MD5sumReader(strings.NewReader("some data")); err != nil || result != "1e50210a0202497fb79bc38b6ade6c34" {
		t.Error(result, err)
	}
}

func TestCrcReader(t *testing.T) {
	if result, err := checksum.CRCReader(strings.NewReader("some data")); err != nil || result != "d9c2e91e" {
		t.Error(result, err)
	}
}

func TestBlake2s256Reader(t *testing.T) {
	if result, err := checksum.Blake2s256Reader(strings.NewReader("some data")); err != nil || result != "54fc4fe89148c8f82479348f56168f71c4165eedda67961daec1d46015db3884" {
		t.Error(result, err)
	}
}
