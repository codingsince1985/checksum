package checksum_test

import (
	"github.com/codingsince1985/checksum"
	"strings"
	"testing"
)

func TestSHA1sumReader(t *testing.T) {
	if sha1sum, err := checksum.SHA1sumReader(strings.NewReader("some data")); err != nil || sha1sum != "baf34551fecb48acc3da868eb85e1b6dac9de356" {
		t.Error("SHA1sum(reader) failed", sha1sum, err)
	}
}

func TestSHA256sumReader(t *testing.T) {
	if sha256sum, err := checksum.SHA256sumReader(strings.NewReader("some data")); err != nil || sha256sum != "1307990e6ba5ca145eb35e99182a9bec46531bc54ddf656a602c780fa0240dee" {
		t.Error("SHA256sum(reader) failed", sha256sum, err)
	}
}

func TestMd5sumReader(t *testing.T) {
	if md5sum, err := checksum.MD5sumReader(strings.NewReader("some data")); err != nil || md5sum != "1e50210a0202497fb79bc38b6ade6c34" {
		t.Error("Md5sum(reader) failed", md5sum, err)
	}
}
