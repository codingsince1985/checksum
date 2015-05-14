checksum
==
[![GoDoc](https://godoc.org/github.com/codingsince1985/checksum?status.svg)](https://godoc.org/github.com/codingsince1985/checksum)

Computing message digest in golang for potentially large files.

Install
--
`go get -d gopkg.in/codingsince1985/checksum.v1`

Usage
--
```go
package main

import (
	"fmt"
	"gopkg.in/codingsince1985/checksum.v1/md5"
)

func main() {
	file := "/home/jerry/Downloads/ubuntu-gnome-15.04-desktop-amd64.iso"
	md5sum, _ := md5.MD5sum(file)
	fmt.Println(md5sum)
}
```
License
==
geo-goang is distributed under the terms of the MIT license. See LICENSE for details.