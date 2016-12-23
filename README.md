checksum
==
[![GoDoc](https://godoc.org/github.com/codingsince1985/checksum?status.svg)](https://godoc.org/github.com/codingsince1985/checksum)

Computing message digest in golang for potentially large files.

Install
--
`go get -d gopkg.in/codingsince1985/checksum.v2`

Usage
--
```go
package main

import (
	"fmt"
	"gopkg.in/codingsince1985/checksum.v2/md5"
)

func main() {
	file := "/home/jerry/Downloads/ubuntu-gnome-16.04-desktop-amd64.iso"
	md5sum, _ := md5.MD5sum(file)
	fmt.Println(md5sum)
}
```
License
==
checksum is distributed under the terms of the MIT license. See LICENSE for details.