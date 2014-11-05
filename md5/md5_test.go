package md5

import (
	"testing"
)

func TestMd5(t *testing.T) {
	file := "/home/jerry/Downloads/ubuntu-gnome-14.10-desktop-i386.iso"

	md5sum, error := MD5sum(file)
	if error != nil || md5sum != "8041c76ef092626f4564baa0ffb595b6" {
		t.Error("Md5sum() failed", md5sum, error)
	}
}
