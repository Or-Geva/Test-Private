package main

import "github.com/jfrog/jfrog-desktop-server/pkg/api"

func main() {
	var x api.FilesDownloadRequest
	x.Checksums = append(x.Checksums, "1234")
	println("Hello, World!" + x.Checksums[0])
}
