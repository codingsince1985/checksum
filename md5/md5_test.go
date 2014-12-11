package md5_test

import (
	"github.com/codingsince1985/checksum/md5"
	"testing"
)

func TestMd5sumFile(t *testing.T) {
	file := "/home/jerry/Downloads/ubuntu-gnome-14.10-desktop-i386.iso"

	md5sum, error := md5.MD5sum(file)
	if error != nil || md5sum != "8041c76ef092626f4564baa0ffb595b6" {
		t.Error("Md5sum(file) failed", md5sum, error)
	}
}

func TestMd5sumDir(t *testing.T) {
	file := "/home/jerry/Downloads"

	md5sum, error := md5.MD5sum(file)
	if error != nil || md5sum != "" {
		t.Error("Md5sum(dir) failed", md5sum, error)
	}
}
