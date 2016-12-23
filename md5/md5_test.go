package md5_test

import (
	"github.com/codingsince1985/checksum/md5"
	"testing"
)

func TestMd5sumFile(t *testing.T) {
	file := "/home/jerry/Downloads/ubuntu-gnome-16.04-desktop-amd64.iso"

	if md5sum, error := md5.MD5sum(file); error != nil || md5sum != "d49a40366d6319501ff5b2d11b3bbf0b" {
		t.Error("Md5sum(file) failed", md5sum, error)
	}
}

func TestMd5sumDir(t *testing.T) {
	file := "/home/jerry/Downloads"

	if md5sum, error := md5.MD5sum(file); error != nil || md5sum != "" {
		t.Error("Md5sum(dir) failed", md5sum, error)
	}
}
