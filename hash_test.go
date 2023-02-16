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
	if result, err := checksum.SHA256sumReader(strings.NewReader("dot")); err != nil || result != "e392dad8b08599f74d4819cd291feef81ab4389e0a6fae2b1286f99411b0c7ca" {
		t.Error(result, err)
	}
}

func TestMd5sumReader(t *testing.T) {
	if result, err := checksum.MD5sumReader(strings.NewReader("dot")); err != nil || result != "69eb76c88557a8211cbfc9beda5fc062" {
		t.Error(result, err)
	}
}

func TestCrcReader(t *testing.T) {
	if result, err := checksum.CRCReader(strings.NewReader("dot")); err != nil || result != "059278a3" {
		t.Error(result, err)
	}
}

func TestBlake2s256Reader(t *testing.T) {
	if result, err := checksum.Blake2s256Reader(strings.NewReader("dot")); err != nil || result != "e364d604d2573afa9f0882f8a50458c4c9c16ca185ab97c5ec1fc71bfc7063bf" {
		t.Error(result, err)
	}
}
