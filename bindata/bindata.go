// Code generated by go-bindata.
// sources:
// static/index.html
// static/main.js
// DO NOT EDIT!

package bindata

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _staticIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x57\xcd\x8e\xdb\x36\x10\xbe\xef\x53\x30\x0c\x7a\x14\x69\x77\x93\x22\xd8\xd2\x02\x8a\x6d\x0a\x14\x68\xd1\xa2\x48\x8b\xe6\x48\x8b\xa3\x88\x2e\x45\xaa\xe4\x68\xd7\x86\xa1\x77\x2f\x28\x59\xb6\x2c\x4b\xde\x6c\xb3\x87\xf8\x22\x0d\xf9\xcd\x37\x9c\x5f\xd1\xe2\xd5\x8f\xbf\xdd\x7f\xf8\xf8\xfb\x7b\x52\x60\x69\xd2\x1b\xd1\x3d\x08\x21\x44\x14\x20\x55\xf7\xda\x8a\xa8\xd1\x40\x2a\x78\xf7\x3c\xad\x97\x80\x92\x64\x85\xf4\x01\x70\x45\x6b\xcc\x93\x77\x94\xf0\x01\xc0\x68\xfb\x0f\x29\x3c\xe4\x2b\x5a\x20\x56\xe1\x8e\xf3\x4c\xd9\x4d\x60\x99\x71\xb5\xca\x8d\xf4\xc0\x32\x57\x72\xb9\x91\x5b\x6e\xf4\x3a\xf0\x75\x6d\x4a\xc9\x17\xec\x5b\x76\xcb\xb3\x70\x90\x59\xa9\x2d\xcb\x42\xa0\xc4\x83\x59\xd1\x80\x3b\x03\xa1\x00\x40\x9a\x0a\x1e\x6d\x0c\x4c\x86\xcc\xeb\x0a\x49\xf0\xd9\x67\xdb\xdc\xfc\x5b\x83\xdf\xf1\x5b\xb6\x64\xcb\x83\xd0\x5a\xdc\x84\xc8\xdf\x11\x7e\x99\x85\x42\x5a\x65\x60\x2d\x7d\x60\x9b\xc0\xdf\xb0\x05\xfb\x6e\xb8\xf6\xb2\xc6\xe4\x56\xbb\xc0\x17\x6c\xf9\x96\xdd\x76\xc2\x0b\x1b\x58\x83\xb4\x7c\xc9\xa2\x85\xf6\xfd\x85\xe9\xff\xfc\xe3\xe7\x18\xa5\x25\x5b\xbe\x63\x6f\x5a\xe9\xf9\xfc\xac\x32\x0e\x99\xd9\xf1\xf8\x34\xbb\xc4\x48\x84\x80\x33\x44\x82\x77\xf5\x7e\xd3\x49\x6b\xa7\x76\x03\x0b\x4a\x3f\x90\xcc\xc8\x10\x56\xb4\x00\xef\x88\x0e\x89\xd1\x9f\x0a\xa4\x27\xcc\x14\x2e\x89\x34\x23\xcc\x18\x97\x39\x8b\x52\x5b\xf0\x91\x34\x37\xb5\x56\x13\x0a\xad\xd2\xab\x24\x21\xf7\x85\xf4\x48\x7e\xf0\x20\x49\x92\xcc\xe0\x06\xe4\x6b\xb7\x9d\x61\x3b\x22\xb5\x5a\xd1\xd8\xbe\x63\x5f\xce\x80\x5c\xe9\x87\x19\x6b\xd3\x5b\x13\xcb\xa3\xa5\xb1\x38\x15\x12\x7a\x33\x13\x80\x7b\x67\xd1\x3b\x13\x2e\x82\x20\xac\x3c\xd2\x18\x78\x00\xf3\x44\xf4\x5b\x4c\x62\x20\x9f\x73\xff\x12\xac\x11\xca\x6b\xb1\xaa\x86\x6e\x78\x67\x48\x21\x43\x22\x95\x72\x36\x5c\x51\x6b\x55\xb5\xad\x6a\x6c\x33\x92\x6b\x83\xe0\x69\x4f\xd5\x6e\x50\x82\xbb\x0a\x56\x14\x61\x8b\x94\x54\x46\x66\x50\x38\xa3\xc0\xaf\xe8\x4f\x2d\x9c\x94\x80\x5e\x67\x81\x31\x76\x35\x99\xd5\xb3\x52\xf9\xff\xa2\x10\x2a\x69\x07\x9e\x80\x4a\x32\x57\x5b\xa4\xe9\xdf\x82\xc7\xbd\x94\xb8\x7c\x80\x3a\x9c\xbc\x07\x7d\x3c\x80\xbe\xb0\xe4\xa6\x8f\xee\x27\x3a\xf7\x88\xae\xa6\xdc\x14\xb2\x3d\x24\x6c\x2b\xe7\xf1\x98\x94\x75\x8d\xe8\x2c\x4d\xdf\xb7\xcb\x82\xcb\x74\x3e\xb4\xd7\x68\x33\x03\xd2\x8f\x58\xe3\x30\x50\xd2\x7e\x02\x4f\xd3\xfb\xb8\x3f\x4b\x3f\xd9\x69\x56\x8e\x97\x62\xef\xfc\xda\x05\x99\xfc\xa2\x03\x5e\x76\x4e\x3f\x0c\x0e\xa9\x78\x72\x6c\x99\xba\x9c\xad\xe8\x4b\xe4\x53\x63\xe8\xb3\x06\xd6\xc4\x39\x13\x94\x6b\x03\x71\x9a\xcf\x16\xef\x4c\x94\x9e\xd8\x7a\xfe\x04\x1b\x7e\x86\x4a\x79\xe5\x73\x15\xaf\x2e\x03\xe6\xd7\xed\xf8\x25\xfb\xf3\x23\x3c\x6a\x85\xc5\x1d\x59\x2e\x16\xdf\x7c\x7f\xbe\x53\x40\x2c\xe0\x3b\xf2\x76\xb1\xa8\xb6\x83\xbd\xe6\xf4\x8a\x9e\x1d\xc2\xe3\xdd\x23\xd3\x21\x09\x60\x20\x43\x50\x04\xd5\xd8\x50\xee\x2c\x26\x8f\x07\xce\xb5\x33\x6a\x92\x52\xf0\xc1\xb1\x05\xef\xbe\x8f\x9d\x10\x4b\xeb\x03\x94\x55\xfb\x7d\x3d\xd6\x55\x1f\x8e\x71\xaa\x12\x3c\x40\x87\xe3\x8c\x6f\x93\xd3\x55\xe8\x84\x18\x5e\x3f\xa3\x6e\x5f\x25\x9d\x10\xdd\x42\xaf\x2b\x18\x7f\x31\x05\x9e\xdf\x5d\x4f\xeb\x7e\xa6\x04\xb0\x48\xbb\xe6\x10\x1c\x8b\x79\xcc\x7d\xed\x3d\x58\x24\x7f\x49\x53\xc3\x34\x54\xf0\xb1\x91\x88\xbb\x38\x8e\xc0\xf3\x1b\x46\xff\xdb\xef\x5f\x83\xcc\x0a\xa2\x24\xca\xa6\x99\xf2\xa0\x8f\xc1\x29\xbf\x51\x49\xe7\xa4\x4f\x71\xd3\x0c\xf2\xbd\xdf\x73\x9d\x37\xcd\x5c\x9b\xa2\xea\xe9\xac\x2c\x81\xa6\xfb\x7d\x7c\x36\x8d\xe0\x38\x11\xbf\x91\xca\x43\x8c\x42\xd4\x69\x5f\xe6\x94\x2e\x03\xd2\xf9\xc9\xa3\x9f\x23\x17\x05\x9f\x08\x8b\xc0\xdc\x39\xfc\x8a\xb2\x79\x7e\x1c\xc1\xdb\x6a\xec\x1b\xa3\xef\xf6\x1b\xc1\xbb\xff\x52\xff\x05\x00\x00\xff\xff\x7b\x8d\x49\xb1\x63\x0d\x00\x00")

func staticIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_staticIndexHtml,
		"static/index.html",
	)
}

func staticIndexHtml() (*asset, error) {
	bytes, err := staticIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/index.html", size: 3427, mode: os.FileMode(420), modTime: time.Unix(1480955498, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticMainJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x57\x4d\x8f\xdb\xbc\x11\xbe\xeb\x57\x4c\x90\x00\x92\x0c\xaf\xbc\x97\x5e\x16\x70\x80\x36\x41\xd1\x1c\x8a\xb4\xdd\xf6\x14\x04\x0b\x5a\x1a\x5b\x6a\x64\x52\x4b\x52\xeb\xf5\x9b\xf5\x7f\x7f\x41\x72\x48\x91\xb2\xfc\x7e\x5c\x0c\x99\x1c\xce\xc7\x33\x0f\x67\x86\x9b\xd5\x2a\x83\x15\xfc\x8d\x29\x84\xba\x67\x4a\xc1\x5e\x48\x10\xbb\xff\x63\xad\x15\x9c\xda\xae\x6e\x81\x49\x04\x7c\x41\xae\x41\x89\x51\xd6\xa8\xaa\x0c\x56\x9b\xcc\x89\x7f\xdd\x29\x94\x2f\x6c\xd7\x23\xfc\xcc\x00\x04\x2f\xac\xe8\x1a\x5a\xc6\x9b\x1e\x65\x69\x97\x01\x76\xc8\x78\x25\x78\xa1\xdb\x4e\xad\x61\x26\x93\x01\x5c\xb2\x0c\x60\xdf\x49\xf4\xe7\xab\xaa\x6a\x98\x66\xc9\x79\xbb\x9f\x68\xf0\x52\x56\xc3\x25\xcb\x28\x9e\x7f\xa2\x96\x5d\xfd\x88\x1a\x90\xd7\x6c\x50\x63\xcf\x34\x2a\xd0\x2d\x82\x42\x0d\x62\x0f\x0a\x7b\xac\x35\x36\x70\xb4\xa2\x0a\xb4\x80\x1d\xc2\x41\xb2\xa1\xc5\xc6\xe8\x60\xbc\xb1\x07\x6a\xd1\x93\xa8\xb1\x64\xf1\xd1\x2d\xd3\x4e\x34\x86\x22\x32\xfa\xaa\x91\x37\x57\xe0\x6c\x56\x2b\x18\x98\x34\x48\x76\x0a\x18\x1d\x50\xd0\x71\xa5\x19\xaf\x91\xf0\x36\x01\x46\xca\xac\xa8\x1a\x77\xe4\xf8\x6a\x93\x01\xd4\x82\x2b\x2d\xc7\x5a\x0b\x59\x38\x8d\x1e\x27\x35\x0e\x28\x8b\xd2\x7e\x1b\x45\x15\x19\xdc\x92\xe5\x69\xc3\x07\xbe\x05\x8e\x27\x78\x44\x1d\x9f\xb2\xa1\x6e\xe1\xa7\x4d\x4b\xa2\xc9\x24\x31\x1f\x87\x86\x69\xcc\xd7\x50\x94\xb0\xfd\x48\xa6\x01\x36\x1b\xf8\xc4\x06\x3d\x4a\xb4\x58\x0d\xa2\xe3\xda\x31\x8a\xf1\xf3\x15\xe4\x74\xc8\x6c\x17\x3d\x6a\xe0\xec\x88\x26\xc2\xd8\xbb\x32\xe8\x26\x27\x94\x16\x12\x9f\xac\xe6\xc2\x1c\x28\x69\x9b\x1c\x25\x29\x4b\x94\xc9\x4b\xb3\xe6\x04\x2f\x65\x14\xd0\x52\x24\x49\xa8\x12\x79\x63\xd0\x84\x0b\x04\x8e\xb2\xa6\x29\x9c\x73\xde\xb7\x6e\x0f\xc5\x3b\x7b\xac\x65\xca\xef\x4d\x8e\xc7\xe1\x54\xd1\xe9\xec\x46\x54\x0b\xdb\x2e\x1c\xd6\x34\xf9\x1a\x6e\x6e\x53\x24\x14\xa7\x77\x57\xe2\x51\xbc\xe0\x82\xc7\x7f\xd0\xe1\x06\x7b\xd4\x38\x73\xca\x2d\x4e\x4c\xf9\xe6\xb6\xbf\x5f\x3b\xe5\xcc\xff\x79\xb7\xeb\x1e\x99\x01\xfe\xe7\x15\x5f\x2b\xda\x5a\xe2\xea\x4c\xb5\x95\xcc\xcb\xf9\x72\x64\xd1\xda\x8a\x40\x20\x7b\x12\xf5\x28\x79\x6a\x36\x92\xf2\x07\x17\xb2\x46\x0a\x0c\x9d\x89\xfd\xdb\x6b\x98\xe0\xed\x0d\xbe\x39\xb0\x9c\x50\x35\x8c\xaa\x2d\x02\xfe\xdd\x11\x95\x66\xc7\xe1\x21\x21\x63\x58\x5e\x93\xdc\x0b\xeb\x47\x4c\x65\xc8\xd9\x24\x1f\x97\x19\x54\xc1\x8b\x2d\x59\xf7\xe1\x0c\xbd\xd0\xfd\xf9\xc9\xc8\x14\x33\x24\xbe\xda\x86\x50\xfd\xc0\xb3\x2a\x82\xa2\xb2\x3a\xb2\xc1\x5e\xc2\xb8\x06\x98\xd0\x25\x3b\x25\x71\x1b\x19\x4f\x0e\xd2\x38\x5d\xeb\xd7\x07\x23\x6f\x75\x59\x7f\x8c\x32\xfb\x31\x45\x5c\xae\x83\xf4\xf9\xa6\xb4\x85\x23\x92\x3c\x8a\x06\x1f\x20\xef\x3b\x8e\x2a\x9f\x96\x8d\x2f\x0f\xf6\x37\x54\x8e\x80\xd2\x75\x0b\x51\x37\x1b\x08\xeb\xfb\xd0\x3b\xa8\x55\x74\x12\x8e\x42\x69\x90\x58\xdb\x4a\xbb\x72\x29\x52\x6b\xd8\xa3\xae\x5b\x6c\x60\x2f\xc5\xd1\x2a\xd9\x34\xb8\x1b\x0f\x9b\x63\x30\xd2\xb8\x20\xae\xda\x89\xba\xd5\x4c\xe2\x16\x70\xb3\xf8\x4f\x35\x3e\xbe\x1d\x26\x78\xb3\x46\x1c\xb4\x6b\x01\x6b\xd3\x0f\xc6\xbe\x8f\xca\x3e\xca\x4e\x34\xb0\x85\xbf\xdc\xdf\xdf\xc7\x77\xa9\xd7\x28\x27\x69\x4b\x21\x1b\xef\xef\x5f\xa6\x88\x9f\xf6\x98\xbb\x91\x21\x0c\xf6\xda\x09\x55\x1d\x50\x17\x79\x8a\x53\xee\xeb\x47\xa5\x5b\xe4\x45\x21\xd5\x90\xf4\x9f\xab\xa8\xa5\x1a\x2c\x03\xab\xb4\xe5\xcc\x70\xb8\x62\xb7\x6f\x3e\xa9\x78\x02\x11\x9e\xe0\xb3\xf5\xd9\xb3\x68\xe6\xda\x92\x5f\xbe\x9d\xa4\xab\x37\x7b\x55\x20\xa5\x2d\x35\x4c\xea\x29\xcf\x5a\x0c\x71\x9a\x3b\xae\x0d\x35\x7a\xd8\xc2\xa9\xe3\x8d\x38\x55\x0a\xf5\x17\x5a\x24\x5f\xac\xa0\x07\x7a\x1d\xa7\x36\xae\x67\x43\x71\xd5\x23\xbc\xf2\xa9\x41\x90\x0d\x5b\x5e\x83\x95\x54\x36\xad\xe6\x0a\xf5\xdf\x2d\x5d\x0a\xc7\x9a\xa4\xae\x07\x22\xb9\x8f\x6c\x09\x2e\x6a\x65\xd4\x8d\xa7\x1a\x3b\x65\x7a\x4a\xe8\x62\x55\x5a\xa8\x3b\x71\x29\x58\x47\xab\x71\x59\x75\x7c\xb6\x73\x46\x2c\xe2\x27\x99\x07\xc3\x01\x56\xf9\xbf\xb6\x47\xc4\x43\x89\x2f\x2e\x26\x93\x14\x27\x5d\x0e\xe3\xdb\xbb\x38\xfc\xb7\x37\x0a\xc6\x06\x51\x75\xbc\xee\xc7\x06\x89\x8f\x84\x1a\xcd\x2e\x1f\x8a\xfc\x3d\x05\x7e\x57\x8b\x91\xeb\xbc\xac\x34\xbe\xea\x22\xc2\xa0\x47\x7e\xd0\x6d\x19\xe4\x9d\x06\x6c\xd2\x03\xbe\xb7\x79\xe9\x99\xfa\x3b\x6d\xaa\x4d\x5e\x56\xad\x3e\xf6\x85\xc6\xe3\x60\xab\x20\x5d\x8f\x27\xbb\xfb\xe4\x97\x43\xfb\x32\xf7\xed\x21\x99\xf1\x2e\xde\x73\x93\xb2\x8e\x0f\xa3\x19\x47\x27\xaf\xe8\x56\x1b\xc2\xd9\x4d\xe7\x5b\x09\xef\xb6\x31\x41\x26\xfa\x39\xa1\xc0\x39\xda\x8e\x18\x77\xc9\xb2\xfd\xc8\x6b\xdd\x09\x0e\x12\x1b\xc9\x4e\x44\x1a\x9b\x2d\xd3\xe6\xfc\xcc\x90\xa6\x2f\x69\x80\x19\xc0\xbf\xec\xff\x8a\x54\x84\xc3\x65\x62\xc0\xac\x90\x7a\x92\xe7\x78\x32\x5f\x45\x5e\xb7\x4c\xea\x7c\xfd\x5b\x56\xca\xd8\x2b\x42\xc5\x1d\x2b\xbf\xdd\x7f\x4f\x0c\x29\xd4\x4f\xa3\xec\xc9\x96\x81\x72\xa4\x42\xf4\xbf\xff\x7c\xb1\xfe\x8e\x95\x1b\xb7\xfe\x3d\xa2\x3c\x17\xb9\x4b\x41\xbe\xb0\x13\xe1\x6e\x50\x4f\xfd\xf3\xa4\x50\xdd\x2f\x08\x1f\xe1\xde\x03\x3f\x9a\xe1\x35\x55\xbd\x86\xbf\x4a\xc9\xce\x95\x69\x6b\x37\xb4\xb8\x0e\x58\x94\xa5\xbb\xc6\x91\x41\x2f\x91\xa6\x37\xb6\x42\x6e\x12\x80\x33\x79\x52\xd7\x76\x66\x0a\x3b\xdb\x09\xea\x51\x5b\x1e\x5e\xd6\x90\xe7\x6b\x18\x2b\x2d\x1e\xb5\xec\xf8\xc1\xc0\x9c\xa4\x8c\x49\x85\x33\x2c\x9f\x63\x2c\xab\x67\xeb\x80\x96\x23\x7a\x8c\x9e\xab\x85\x11\xfa\x3c\xa0\xd8\x47\x5b\xdb\x2d\xe4\xca\x9a\xcc\x27\xba\xa6\xb8\x98\x07\x40\x38\xe0\x38\x0b\xd8\x2b\x0c\xe2\xe1\x39\x44\x85\x42\xec\x61\x6e\x7b\x59\x6b\x3a\x61\x5f\xa2\xfb\xe0\x23\x48\x91\x4e\x40\x9d\x2a\xf4\x73\x02\x70\x0c\x5b\x2d\x8e\x43\x17\xdf\x78\x67\x5e\x90\x46\x2a\xb0\xff\xb0\x6f\xfa\x1d\x93\xaa\xa2\x03\xc5\x87\x49\xd2\x95\x92\x59\x3e\xb4\x38\x1c\x7a\x7c\x72\xbe\xb8\x02\x1a\xf2\x62\xa6\xad\xd9\x3d\xa5\x78\x14\xea\xa9\xe0\x86\xce\x88\x9a\xa8\x1e\x0a\x71\x82\xaf\xd9\x37\x60\x85\xcd\xc4\x91\x8e\x77\xfe\x2e\x87\x6a\x67\x66\x27\x7b\x74\xb1\xea\x3d\x5c\xa3\x92\x56\xcf\x3b\xbf\x9e\x7b\xc6\x7e\x8e\x9e\x2b\x71\x0e\x88\x80\x34\xf6\x15\xa1\x32\x84\x67\x73\xbc\x6f\x9e\xeb\xf1\xe1\x72\xa6\x2c\x74\x7a\x3f\xd8\xbb\xaf\x94\x35\xc9\xfb\xd7\x55\xb9\x2b\xab\xa9\x10\xd5\xa0\x2b\x6b\x34\x9e\x18\x30\x3f\x14\x8d\xa8\xc7\x23\x72\x5d\x56\x12\x59\x73\x8e\x66\x21\x87\xaf\x71\x24\x96\x32\x16\xea\xbe\xab\x7f\xe4\x6b\xc8\xb5\x24\xa5\x77\x52\x9c\x40\x37\xb6\x9f\x99\xe7\x39\xf6\xd1\x48\x95\x32\x06\xfb\x4a\x33\x79\x40\xd7\x39\x3e\x09\xae\x8d\xde\x8c\xe6\xa7\x9b\xa6\xde\xbb\xf7\xe1\x5c\x77\x1a\xfe\xf4\xda\x5c\x54\x66\x7b\x91\x55\x16\x8a\xd5\x82\xb6\xeb\x5b\x86\x7d\x55\x8f\xd2\xbc\xd9\xfe\xeb\x3c\x77\xaf\x97\xc9\x8c\x9b\xb1\x0c\x84\xa6\x0e\x17\xae\xef\x98\x5c\x4e\xb5\x2b\xbb\x94\xd9\xaf\x01\x00\x00\xff\xff\x86\x4e\x8e\x72\xc0\x13\x00\x00")

func staticMainJsBytes() ([]byte, error) {
	return bindataRead(
		_staticMainJs,
		"static/main.js",
	)
}

func staticMainJs() (*asset, error) {
	bytes, err := staticMainJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/main.js", size: 5056, mode: os.FileMode(420), modTime: time.Unix(1480958409, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"static/index.html": staticIndexHtml,
	"static/main.js": staticMainJs,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"static": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{staticIndexHtml, map[string]*bintree{}},
		"main.js": &bintree{staticMainJs, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

