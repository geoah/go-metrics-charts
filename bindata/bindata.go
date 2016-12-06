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

var _staticIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x59\xdf\x6f\xdb\xb6\x13\x7f\xef\x5f\xc1\xaa\xf8\xbe\x55\xa2\xfd\x4d\xba\x15\x9d\x6c\xa0\x48\xdb\x61\xc0\xba\x15\x6d\x37\xac\x4f\xc3\x59\x3c\x45\x4c\x29\x52\x23\x4f\x89\x8d\xc0\xff\xfb\x40\xfd\xb0\x15\x59\xb2\x93\xd4\x68\xb2\xbc\x48\xa4\x8e\x77\xbc\x0f\x3f\x77\x3e\x5e\xe2\xa7\x6f\x7e\x3f\xfb\xfc\xe5\xc3\x5b\x96\x51\xae\xe6\x4f\xe2\xfa\xc1\x18\x63\x71\x86\x20\xea\xd7\x6a\x48\x92\x14\xce\x63\x5e\x3f\xb7\xf3\x39\x12\xb0\x24\x03\xeb\x90\x66\x41\x49\x69\xf8\x32\x60\xbc\x23\xa0\xa4\xfe\xca\x32\x8b\xe9\x2c\xc8\x88\x0a\xf7\x8a\xf3\x44\xe8\x0b\x17\x25\xca\x94\x22\x55\x60\x31\x4a\x4c\xce\xe1\x02\x96\x5c\xc9\x85\xe3\x8b\x52\xe5\xc0\x27\xd1\xff\xa3\x13\x9e\xb8\x66\x1c\xe5\x52\x47\x89\x73\x01\xb3\xa8\x66\x81\xa3\x95\x42\x97\x21\x52\x30\x8f\xb9\xb7\xf1\x4d\x26\x53\xa3\x29\x84\x2b\x74\x26\x47\x7e\x1a\xfd\x18\x4d\x2a\xcb\xdd\xe9\xbb\x6c\xc0\x25\x56\x16\xc4\x9c\x4d\x6e\xbd\x83\x8b\x7f\x4a\xb4\x2b\x7e\x12\x4d\xa3\x69\x33\xa8\x2c\x5e\x38\xaf\xbf\x56\xf8\x6d\x16\x32\xd0\x42\xe1\x02\xac\x8b\x2e\x1c\x3f\x8d\x26\xd1\x0f\xdd\xb9\xe3\x1a\x83\xa5\x34\x8e\x4f\xa2\xe9\x8b\xe8\xa4\x1e\x1c\xd9\xc0\x02\x41\xf3\x69\xe4\x2d\x54\xef\x47\x56\xff\xc7\xc7\x5f\x3c\x4a\xd3\x68\xfa\x32\x3a\xad\x46\xc7\xd5\xff\x4e\x2a\xfc\x04\x97\x68\x6b\x2b\x27\xd1\x49\x67\xea\xee\xa6\xa2\x42\x19\x8a\xd4\x8a\xfb\xa7\x5a\x85\x0a\x08\x1d\x8d\x28\x8a\x79\x1d\xdb\xf5\x60\x61\xc4\xaa\x79\xaf\xc6\x42\x5e\xb2\x44\x81\x73\xb3\x20\x43\x6b\x98\x74\xa1\x92\xe7\x19\x05\xdb\x5d\x0c\xc9\x85\x5e\x4f\x4f\xa6\x2f\x97\x18\x4d\x20\x35\x5a\xaf\x34\x55\xa5\x14\x03\x0b\xaa\x45\x4f\xc3\x90\x9d\x65\x60\x89\xbd\xb6\x08\x2c\x0c\x47\xe4\x3a\xca\x17\x66\x39\xa2\x6d\x23\x29\xc5\x2c\xf0\xa9\xaa\xef\xcb\x0d\x41\x2e\xe4\xe5\x88\xb5\xe1\x4f\x03\xd3\xbd\xa9\x66\x38\x88\xf1\x06\x93\xe0\xc9\x08\x02\x67\x46\x93\x35\xca\xed\xa0\x10\x6b\xd8\xa8\x51\x78\x89\xea\x00\xfc\x95\x4c\xa8\x30\x1d\xf3\x7f\x57\x58\x12\xe6\xfb\xc0\x2a\xba\x6e\x58\xa3\x58\x06\x2e\x04\x21\x8c\x76\x7b\x96\x55\x4b\xa5\x2e\x4a\xaa\x8e\xa4\x7a\x0b\x53\xa9\x08\x6d\xc0\xaa\x1f\x98\x59\xf0\xae\x1a\xb2\xc5\x8a\xe5\x48\x56\x26\x4c\x43\x8e\xcf\x19\x24\x09\x16\xe4\x98\xc5\xf3\x52\x81\x65\xb8\x2c\x2c\x3a\x27\xbd\xc1\xbd\xf6\x9a\xbf\x66\xbb\x95\xc9\x80\xd1\xaa\xc0\x59\x40\xb8\xa4\x80\x15\x0a\x12\xcc\x8c\x12\x68\x37\xd6\x6b\xd3\x2e\x8a\xa2\x43\xde\xc0\x86\x86\x25\x91\xd1\x41\xe5\xd8\x82\x74\x98\x28\x04\xdb\x3a\xb7\x5f\x49\xa5\xc8\x15\xa0\x37\xbb\x4c\x8c\xf6\xa1\xe2\x72\x50\x43\x87\x3b\x8c\x6b\xbb\x3a\x05\x96\x42\x98\x28\xe3\xd0\x67\x00\x79\x0b\xe3\xdc\x5b\x3f\xe0\x29\x87\x7d\xc1\x53\xdc\x29\x74\xd8\xbd\x48\x57\x41\xe4\xf1\xad\x51\x45\x11\x26\xa6\xd4\x14\xcc\xff\x6a\x1c\x60\x26\xed\x48\x35\x87\xd8\x0a\x7d\xd9\xe7\xe5\xed\x43\x7c\x78\xeb\xb6\xc9\x94\x87\xb2\x9a\x27\x1e\xab\xa9\xb2\x1b\xd7\xf7\x47\xe6\x1b\xc2\xb1\x47\xe0\xaa\xa4\xa3\xb0\x19\x48\x17\x42\x42\xf2\x12\xb7\xbc\x56\x52\x63\x78\x28\x99\x6e\xb4\x1f\x9b\xd5\x5d\xeb\xdf\x87\xda\xec\x10\x48\x5b\x68\xc0\x22\x3c\x1c\x34\x5d\xeb\x8f\x0e\x9a\x05\xd8\x87\x43\xa6\x63\xfc\x71\xa4\xc3\xf1\x2c\xf1\x89\xc0\x12\x77\x64\x0a\x06\x5a\x30\x8b\xa9\x45\x97\xb1\x02\xad\x34\xe2\x51\x26\x8c\xed\x19\x17\x50\xfa\x5f\x9c\xef\x7f\xbe\x8d\xe1\xef\x48\xfa\x7e\x0d\x53\x9f\xcf\xa6\x86\xf9\x78\xe3\xd8\x9e\x33\xa9\x99\xc3\xc4\x68\x71\xff\x5a\x45\x97\xf9\xc2\x57\x49\xb9\xd4\xb3\x60\x1a\xb0\x4b\x50\x25\xce\x82\x17\x7b\x2b\xda\xfb\xb2\xf0\xbd\x74\x09\xf3\x89\xff\x11\xfd\x4a\xb5\x2c\xc3\x65\x61\x2c\x85\x85\x3e\x0f\x7a\x44\x7c\x00\xe6\x09\x73\xa5\x95\x01\x71\x44\xf2\x6d\xf6\x39\xff\xf0\xdb\xcf\xc7\x4a\xd1\x2d\x78\xa6\x40\x1d\x5a\xb8\xea\x43\xd7\x34\x4d\xb8\xc0\x45\x79\xce\x9b\xc2\x29\x60\x04\xf6\x1c\x69\x16\xfc\xbd\x50\xa0\xbf\x3e\x48\x15\x6b\xc4\x31\x23\x7b\xb3\xc7\xf9\x47\xb8\xfa\x8f\x54\xbd\x47\x89\x9a\xfa\x36\x22\x80\x60\x27\x6a\xce\xfc\x27\xf6\x06\x08\xee\xc2\xa4\x4a\x61\x4f\x97\x3f\x6c\x01\xfa\xbc\xbe\xe7\x9b\x92\x7c\xb5\x26\x5a\x0b\xaf\x95\x3a\x2a\x9a\x83\xd7\x6f\x0d\x97\xf3\x3a\x87\xbd\xd5\x62\xe8\x1e\xbd\x7b\xd9\x7e\x5f\xb3\x9d\xfd\x2a\x1d\xed\x5e\xb5\xdb\xf6\x41\x1b\x13\x87\x1a\x1d\xaa\xcc\x47\xcf\x65\x57\xf2\x50\xe3\xe2\x56\x2d\x8e\x81\x7d\x86\x04\x0b\x55\x05\xce\x28\x11\x47\x20\x3c\xf0\x69\xb4\xe7\xb1\x05\x7d\x07\xcf\x7e\x3b\x64\x33\xec\xf6\xb4\x72\xd8\xd3\xfb\xa2\x55\xb7\xf5\xcc\x9e\x55\xc5\x1c\xbb\xbe\xb9\xbb\x2b\x29\x28\x7b\xc5\xa6\x93\xc9\xff\x7e\xba\xf9\x25\x43\x7f\x3b\x7b\xc5\x5e\x4c\x26\xc5\xb2\xf3\x6d\xbd\x7d\x25\x1b\x35\xc8\x59\x73\x15\xf9\xa4\x85\x0a\x13\x42\xc1\x48\xf4\x0d\x55\xcd\xe1\xab\x46\xe7\xc2\x28\x31\xac\xf2\x59\xb7\x34\x18\xd9\xec\x29\xe6\x83\x8b\x63\xde\xf1\x39\xe6\x9d\x56\x5d\x05\xf3\x67\xcc\x8b\xaa\xd3\xb7\xe1\x6b\x8b\x65\x9f\x02\x21\x35\xa2\xdd\x9e\x07\x5f\x86\xdb\xfe\xef\x56\xa2\xdb\xf4\xf7\x6b\x5b\xf6\xd5\x03\x8f\x09\x59\x59\x60\xbf\x77\x17\xd3\xcd\xff\x18\x6c\xe7\xed\x08\xb5\x28\x9b\xd7\x24\x89\x39\x65\xe3\x32\x67\xa5\xb5\xa8\x89\xfd\xe9\x4b\x9d\x61\xd1\x98\xf7\x8d\x78\xb9\x9d\xed\xc4\x54\x23\xd8\x5f\x7e\x7d\xfd\x0c\x21\xc9\x98\xcf\x8b\xeb\xf5\x90\x07\x2d\x06\x5b\x72\xf8\x45\x32\x65\x2d\x3f\xd6\xeb\x0e\x59\xae\xaf\xb9\x4c\xd7\xeb\xb1\xf0\x27\xd1\xaa\xd3\x90\x63\x30\xbf\xbe\xf6\xcf\xf5\x3a\xe6\x34\x80\x5f\x6f\x49\x55\xf0\xf9\x35\xd5\xcb\xd8\xa2\x5d\x40\x6a\x3f\xb9\xf7\xb3\xe7\x62\xcc\x07\x60\x89\x29\x35\x86\x1e\xd1\x69\xde\xdc\x4e\xcc\x2b\x36\xb6\x81\xd1\xa6\x8a\x27\x31\xaf\xff\x83\xf5\x6f\x00\x00\x00\xff\xff\xdf\x71\xbc\xa3\xd9\x1a\x00\x00")

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

	info := bindataFileInfo{name: "static/index.html", size: 6873, mode: os.FileMode(420), modTime: time.Unix(1481054189, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticMainJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x59\x5d\x8f\xdb\xb6\xd2\xbe\xae\x7f\xc5\x14\x2d\x22\x6a\x63\xcb\x0e\xf0\xbe\x07\x07\xde\x3a\x40\x92\xa6\x68\x81\x16\xed\x49\xd2\xab\xc5\x62\x41\x4b\x63\x8b\x1b\x99\x54\x48\x6a\xbd\x6e\xe2\xff\x7e\xc0\x0f\x49\xa4\x2c\x67\xb7\x07\xdb\x5c\x24\x0a\x39\x9c\x79\x66\x38\xc3\xf9\xf0\xfc\xe2\x62\x02\x17\xf0\x9a\x2a\x84\xbc\xa2\x4a\xc1\x46\x48\x10\xeb\x5b\xcc\xb5\x82\x7d\xc9\xf2\x12\xa8\x44\xc0\x3b\xe4\x1a\x94\x68\x64\x8e\x2a\x9b\xc0\xc5\x7c\xe2\xc8\x7f\x5f\x2b\x94\x77\x74\x5d\x21\x7c\x9e\x00\x08\x4e\x2c\xe9\x14\x4a\xca\x8b\x0a\x65\x6a\x97\x01\xd6\x48\x79\x26\x38\xd1\x25\x53\x53\x18\xd0\x4c\x00\x8e\x93\x09\xc0\x86\x49\x6c\xcf\x67\x59\x56\x50\x4d\xa3\xf3\x76\x3f\xe2\xd0\x52\x59\x0e\xc7\xc9\xc4\xeb\xf3\x1b\x6a\xc9\xf2\xf7\xa8\x01\x79\x4e\x6b\xd5\x54\x54\xa3\x02\x5d\x22\x28\xd4\x20\x36\xa0\xb0\xc2\x5c\x63\x01\x3b\x4b\xaa\x40\x0b\x58\x23\x6c\x25\xad\x4b\x2c\x0c\x0f\xca\x0b\x7b\x20\x17\x95\x27\x35\x92\xac\x7d\x74\x49\xb5\x23\x0d\x4d\x11\x08\xbd\xd7\xc8\x8b\x13\xe3\xcc\x2f\x2e\xa0\xa6\xd2\x58\x92\x29\xa0\xfe\x80\x02\xc6\x95\xa6\x3c\x47\x6f\x6f\xa3\x60\xc0\xcc\x92\xaa\x66\xed\x81\x5f\xcc\x27\x00\xb9\xe0\x4a\xcb\x26\xd7\x42\x12\xc7\xb1\xb5\x93\x6a\x6a\x94\x24\xb5\xdf\x86\x51\xe6\x05\xae\xbc\xe4\x7e\xa3\x55\x7c\x05\x1c\xf7\xf0\x1e\x75\x78\xca\xaa\xba\x82\xcf\xf6\x5a\x22\x4e\xe6\x12\x93\xa6\x2e\xa8\xc6\x64\x0a\x24\x85\xd5\x4b\x2f\x1a\x60\x3e\x87\x37\xb4\xd6\x8d\x44\x6b\xab\x5a\x30\xae\x9d\x47\x51\x7e\x38\x31\xb9\x3f\x64\xb6\x49\x85\x1a\x38\xdd\xa1\xd1\x30\x44\x97\x76\xbc\x3d\x08\xa5\x85\xc4\x1b\xcb\x99\x98\x03\xa9\xdf\xf6\x40\x3d\x95\x75\x94\x1e\xa5\x59\x73\x84\xc7\x34\x50\x68\x4c\x93\x48\x55\x89\xbc\x30\xd6\x84\x23\x74\x3e\x4a\x8b\x82\x38\x70\x2d\x36\xb6\x01\xf2\xad\x3d\x56\x52\xd5\xee\xf5\xc0\x43\x75\xb2\xe0\xf4\xe4\x8c\x56\x23\xdb\x4e\x1d\x5a\x14\xc9\x14\xce\x6e\x7b\x4d\xbc\x9e\x2d\x5c\x89\x3b\x71\x87\x23\x88\x1f\x09\xb8\xc0\x0a\x35\x0e\x40\xb9\xc5\xde\x53\xae\xdc\xf6\xf5\x29\x28\x27\xfe\xef\xc3\xd6\x62\xbb\xad\xfe\x07\xd8\xb1\xba\x8e\x27\x60\xa5\x30\x26\x1b\x5e\x43\x27\x37\xaf\x90\x9a\x0b\xff\x7c\x12\x27\x99\xdf\x1a\x8b\x91\x81\x4a\x96\x32\x49\x87\xcb\x81\xa6\xbd\xac\x1b\xc3\xa5\x13\xd8\xc5\x82\x13\x7a\x3e\x1a\x4e\x4c\x0f\x2b\xb8\xba\xee\x74\xf9\xaa\xdc\xc0\x7a\x9e\x9f\x44\xdd\x48\x1e\xab\x1b\x50\xb5\x07\x47\xbc\xd4\x33\x30\x90\x7d\xb4\xaf\x46\xb0\x7d\xf9\xd2\x82\x73\x44\x59\xdd\xa8\x92\x74\xba\xb0\x1d\x2a\x4d\x77\xf5\x32\x0a\xbe\x6e\x79\xea\xe9\xee\x68\xd5\x60\x4c\xe3\xc1\x46\xfe\x77\x1c\x5c\x51\x60\x21\x27\xbd\x55\xa7\xae\x84\xae\x0e\xf1\x05\x78\x4b\xfc\x6e\x13\x60\xf6\x11\x0f\x8a\x74\x8c\xd2\x6c\x47\x6b\xfb\xe8\x84\x6f\x9e\x51\x5d\xd2\x7d\xa4\xb7\xa1\x69\x83\xc1\x73\xec\x9f\xb1\xfb\xa5\xa1\xb7\xbc\x2c\x1e\xc3\xcc\x7e\xf4\x1a\xa7\xd3\x8e\xfa\x70\x96\xda\x9a\x23\xa0\xd4\x87\x1a\x97\x90\xa8\x9c\x6a\x8d\x32\xe9\x37\x0c\x9a\xa5\xfd\xbb\x7b\x2b\x3b\x3b\x9d\x26\x4d\x75\x36\x65\xd2\xaa\xea\xb2\xa5\x4f\x8e\x4c\xc2\x4e\x28\x0d\x12\x73\x9b\x5b\x2e\xdc\x25\xa9\x29\x6c\x50\xe7\x25\x16\xb0\x91\x62\x67\x99\xcc\x0b\x5c\x37\xdb\xf9\xae\x13\x52\x38\x35\x4e\x12\xa8\x3a\x97\x3e\xc3\xa4\x77\x36\xdd\xf5\x59\x2d\x8c\x4b\xa3\xbc\xea\x43\xc4\xae\x75\xd6\x36\x19\xb0\xa9\xaa\x20\xd1\xa1\x64\xa2\x80\x15\xfc\xff\x62\xb1\x08\xa3\xa9\xd2\x28\x7b\x6a\xeb\x44\x56\xdf\x87\xc3\x29\xf0\x50\x7b\xcc\xc5\x64\xa7\x06\xbd\x67\x42\x65\x5b\xd4\x24\x89\xed\x94\xb4\x2f\x66\xa6\x4b\xe4\x84\x48\x55\x47\x19\x17\xe0\x47\xaa\x69\x26\x6a\xcd\x04\x57\x99\x77\xe9\x8a\x1e\x44\x63\xfc\x49\x57\x08\x2b\x90\xaa\xb6\x7e\x99\xe5\xbb\xa2\x62\x1c\xb3\x5b\xc1\x38\x49\xa0\x63\x7e\x62\xbb\xee\x44\x9c\xaa\xc7\x2c\x87\x7b\x03\x01\x49\xcf\xab\x7b\xc4\x94\xa6\xd6\x71\xc2\x68\x0a\x38\xef\xcc\xbe\x4a\xc3\x14\xef\xc2\xc9\xda\x34\x46\xe1\x68\xaf\xcc\xdf\xd7\x01\xb5\xcd\x08\x87\x1a\xc5\x86\xb8\x70\x80\xd5\x0a\x12\xde\xec\xd6\x28\x93\x98\xb1\x63\xed\xa2\x17\x92\xad\xe8\x98\x66\x09\x3c\xb7\x50\x23\xe2\xe8\xf2\x6c\x3c\xc3\xca\x01\x0b\xc8\x8e\x93\xd3\xaf\xc8\xdf\x4e\xde\x91\xf6\x21\x6f\x23\x71\x70\xbd\x83\xbb\xf5\x29\xcd\x15\x21\xf1\xea\xd9\x0a\xa7\x0b\x6c\xfb\x60\x53\xa9\xfb\x58\xd1\xa2\x0e\x43\x85\x71\x6d\xc2\xab\x82\x15\xec\x19\x2f\xc4\x3e\x53\xa8\x7f\xf1\x8b\x1e\x8b\x25\x6c\x9d\x75\x1a\x86\x47\x98\x15\x6a\x72\x92\xa2\x5b\xe6\xfd\x25\x78\x19\x36\xeb\x75\x52\x62\xda\x30\xb9\x05\xe8\xba\xe8\xec\xd2\xb4\x42\x7d\xe3\x60\x10\x8f\x26\xcc\xd8\x5d\xfc\xba\x8f\x01\x30\x75\x23\x1b\xce\x19\xdf\x92\x61\xf1\x10\x18\xa8\x5b\xb1\x06\x8c\x85\x87\x1c\xe2\x88\x7f\x2d\x44\x85\x94\x9f\x6a\xd5\x82\xfe\xc9\x3e\x21\xc4\xbd\x24\x11\xe6\xee\x71\x71\x1f\x93\xf1\xeb\x37\x7a\x37\xb2\x22\x69\x5f\xde\xf9\x0a\xb5\xcf\xc3\x12\x7d\x58\xbe\xc3\xed\xdb\xfb\x9a\x84\xec\xbf\x7c\x81\x24\xbb\xf0\x71\xdf\xd7\x19\x5d\xd6\xb6\x7e\x3b\x9a\xe6\x46\x12\x59\x98\x59\xa6\xc1\x6a\x98\xa7\xdd\xf3\x68\x0b\xf5\x90\xa4\x6d\x05\x96\xee\xf9\x6a\xff\x6b\x8b\x8e\xb0\xaa\xef\x03\xeb\x98\x7a\x1d\xfc\x5b\x6b\xb0\x49\xcc\xf0\x1e\x73\xbf\x62\xc1\xa7\xbe\xbe\xff\x9e\x24\xdf\x79\xdd\x66\xb9\x68\xb8\x4e\xd2\x4c\xe3\xbd\x26\x81\x9a\x15\xf2\xad\x2e\xd3\x8e\xde\xf1\xc7\x22\x3e\xd0\xd6\x43\x2d\xf5\x80\xfd\x4c\x9b\xfc\x94\xa4\x59\xa9\x77\x15\xd1\xb8\xab\x6d\xde\xf4\x81\x7e\x63\x77\x6f\xda\xe5\xae\xe4\x31\xaf\xda\x32\xea\x83\x8e\x2d\x72\x73\x2b\x8c\xd7\x8d\x69\xd9\x8c\x18\xfb\x3d\x73\xd8\xfc\xc5\x19\x5f\xb6\xcb\xc6\xbc\x24\x85\x6f\x57\xa1\x0f\xf5\x4e\xdd\xd3\x84\xdb\x81\x33\xdb\xdc\x0f\xb3\x27\xf9\x63\xd8\xfe\x51\x09\xad\x19\xdf\x9a\xef\xa7\xe1\x6a\x8a\x83\xc9\xa6\xe1\xb9\xc9\x6f\x70\x63\xf2\x1b\xe9\xc7\x02\xde\x25\xff\xb0\x59\x2f\xe3\xb8\x37\x5f\x24\xc9\x4b\x2a\x75\x32\xb5\x56\x9e\x7e\x25\x41\x8e\xee\xa5\xc6\x2a\x9d\x44\xb3\x76\x63\x72\xe6\x8d\x65\x4a\x22\xb9\x0e\x4e\xec\xc1\x51\x51\x39\xc2\x8b\x4a\xa4\x11\x2f\x9f\x28\xf3\x8f\x58\xbc\x92\x68\x1a\x8a\xf6\x00\xd1\x92\xe6\xa8\xa2\xa7\xd5\xae\xc0\xb3\x67\xe0\xbe\x5a\xaf\xec\xdf\x31\xbb\x7c\xb5\xb8\x36\xb7\x6d\xde\xce\x44\x8b\xbf\x50\x8a\x43\x12\x11\x64\x1b\x21\xdf\xd2\xbc\x24\x0a\x25\x33\x99\xea\x25\x7c\x76\x9f\xd9\x21\x93\x78\x87\x52\x21\x49\x2f\xc1\xaf\xdd\xf7\x6b\x26\xbb\x7c\x63\x73\x3c\xb9\xa3\x12\xd8\xea\xc5\x25\xb0\x1f\x22\x34\x97\xc0\x9e\x3f\x8f\x1a\x77\x07\x8a\x85\xa0\x38\xde\xeb\x43\x62\x58\xf5\xcc\x6e\x57\x8b\x4b\xb8\xfd\x81\xfc\x46\x75\x99\xed\x58\x6b\x80\x2b\x76\x7d\x95\x1c\x92\x6b\xcf\x7e\xda\xf1\x9b\xbd\x88\x36\xd2\xf4\x12\x6e\x9d\xe4\x6f\xbe\x09\xa4\x5a\x9a\xab\xdb\x6b\x78\xbe\x3a\x39\x7a\x75\x7b\x7d\xe9\x40\x1c\xcd\x3f\xc7\x27\xb4\x92\x67\xda\x16\x83\x96\xe3\xe5\xc4\xc9\x88\x1c\x28\xb8\xfd\xaf\x3a\xd3\x88\x37\xad\xa9\x3c\x71\x26\xdf\x96\x7e\x85\xd3\x24\xac\xd1\x9c\x6a\x62\x03\xe1\xb4\xcd\x6b\x64\x4a\x2b\x73\x5d\x6b\x2a\x93\x31\xe4\x6e\xf4\x76\x9c\x4c\xcc\xed\x59\x40\x1b\x6e\xd2\x6e\x1c\x34\x03\xd0\x71\x0c\xf9\x53\x24\x7d\xf2\xc7\xe8\xcf\x77\xbf\xc2\xcf\x94\x9b\x8a\xf7\x1f\x7b\x90\xba\x7c\xdc\x59\xbf\xf1\xc9\xf7\xcf\x77\xbf\x58\x43\x37\x7e\x24\xf1\x9f\x06\xe5\x81\x24\xee\xc9\x4f\x46\x76\x82\x17\xde\x44\x7a\x7c\x7d\x6d\x12\x52\xec\x2f\x84\x97\xb0\x68\xef\xa9\xc9\x68\x51\xc4\xac\xa7\xf0\x4a\x4a\x7a\xc8\x4c\xe3\x75\x86\x8b\xeb\xd1\xac\x47\xb9\x4b\xed\x04\xb6\x14\x71\x26\x09\xa5\x78\x98\xfe\xe9\x1c\xd0\x7b\x76\x25\x53\x5a\xc8\x83\xed\xf2\xdf\x6b\x9b\xf7\x8e\x53\x48\x92\x29\x34\x99\x16\xef\xb5\x74\xc5\x57\xec\xce\x54\x2a\x1c\xd8\xf2\x53\x68\xcb\xec\x93\x05\xa0\x65\x83\xad\x8d\x3e\x65\x23\xf3\x21\xd7\x0d\xf4\x5b\xa6\x21\x50\x56\x64\xd0\x10\xc4\x76\xa1\x45\xd1\x1f\x18\x9b\x19\x8d\x8c\x65\x86\xb2\xc7\xb9\xc6\x53\xaf\x63\x90\x7a\x5b\x0d\x62\x4b\x47\x46\xed\xeb\xc5\x4f\x91\x81\x9f\x3a\x50\xde\xa3\x6e\x6a\x78\x06\x6f\xed\x8c\xff\x67\x37\x93\x57\xff\x54\xc8\x30\xce\xda\x27\xa0\xab\x96\x4c\xb7\x6e\xf5\x1f\xad\x9a\x96\x1e\xd3\x9a\x4a\x95\xe5\x62\x57\xb3\x0a\xc9\xb0\x02\x9b\xb5\xd4\x6d\x29\xd6\x39\xf7\x8f\x7e\x4c\x67\x05\xf8\x5c\xbf\xec\x6e\xcd\x3d\x8c\xcb\xe0\x16\x55\x29\xf6\xbf\x32\xfe\x71\x09\x1b\x5a\xa9\xa0\xb2\xdd\x89\x02\x5f\x53\xf9\xba\xd1\x5a\x70\xf5\x41\xbc\xb3\xe1\xbb\x84\xab\x44\x8b\x5f\x76\x74\x6b\xfa\xb0\x44\x21\x2f\x8c\xc4\x0f\xe2\x4d\x25\x9a\x22\xb9\xee\xcf\x17\x4c\xd5\x15\x3d\x54\x62\x2b\x3c\xef\xd6\x2f\xa6\x11\x18\x5f\x9d\xf4\x98\xbc\xdf\x4c\x03\x13\xa9\xa5\x8d\x0c\x3f\x31\x21\xa1\xae\x9d\x07\xfa\xe8\xe9\x7e\x49\x88\x42\x3c\x1d\x52\xc7\xd3\x70\x83\xe4\x01\x12\xff\xf6\x75\x54\x9d\xdb\x76\x3d\x53\xb4\xdc\x36\x91\x13\x08\x63\xdd\x38\xf3\xf7\xa4\x10\x79\xb3\x43\xae\xd3\x4c\x22\x2d\x0e\x41\x1f\xec\xbc\xc5\x54\xc6\xf3\x39\x7c\xb0\x43\xe1\x36\x08\x1d\x2a\xe3\x52\xfb\x12\x39\xe4\x15\xcb\x3f\x32\xbe\x05\xc1\x41\x97\xb8\x03\x66\xff\x0d\x8a\xec\xf9\x1c\xac\xab\x4c\x4c\xfd\xde\x8b\x34\x4a\xd9\xc3\xe6\xfa\xb4\xf4\x80\x67\x52\xec\x41\x17\xb6\x55\x48\xa6\x40\x30\x68\xcd\x63\xa3\xf8\x51\x35\x66\x9a\xca\x2d\x6a\xdb\x35\xbc\x11\x5c\x1b\xe6\x93\xf6\x17\x87\xf9\xdc\x84\x59\x08\x08\x2a\xa6\x74\xdf\xed\x0d\x11\xd9\xda\xdd\x20\x8a\x1b\x80\x31\x24\xa7\xef\x05\x66\x79\x23\x25\x72\xfd\xc1\x41\x72\xb3\x91\x31\x30\x12\x37\x12\x55\xd9\x77\xca\x0f\xe1\x70\x84\x0f\xe1\x68\x9b\xf3\x51\x20\x70\x01\x2f\x16\x8b\x45\x84\xe7\x4d\x85\xa6\x00\x6c\x94\x83\x65\xeb\x06\x5f\xbf\x57\x48\xef\xcc\xbd\xf6\xb6\x0b\xae\x9e\x71\x4d\x73\xfd\xb5\x1b\xfd\x6e\xad\xf9\xcc\x0e\x1d\x66\x86\xdf\xd7\xef\x32\x1c\xc9\x9f\xc2\xeb\x91\x51\x5e\x9c\x60\x79\x1c\x88\x47\xc8\x8f\x45\x7b\xa7\xb7\x55\x9d\xff\x29\xf2\x31\xc2\x6a\xda\xa8\x13\xc7\xb5\xbd\x65\x2e\xb8\x6f\x2d\x3b\x3a\x60\x41\x67\x19\xdd\xe5\xf8\xb4\x64\x10\xef\xc1\xd4\xc4\x70\xf7\x75\xcd\x9b\x8a\x2a\x45\x92\x0d\xf5\x58\x22\x12\x5a\x14\xc1\x7e\x45\x0f\xc9\x68\xe2\x3d\xf3\xb0\x9c\x97\xd4\x73\x1a\x15\x14\x00\x39\x9e\xde\xef\x4f\xe7\x82\x71\xcc\x99\xc6\x03\xf2\xb4\x63\xb7\xfd\x77\xe2\x85\x9e\x89\x57\xde\x54\x55\x74\xe9\x6f\xef\x6b\x61\xdc\xcc\xce\x25\x59\x0e\xcc\xa4\x97\x07\x81\xa1\x3d\x35\xab\xf9\x76\x08\xcb\xd5\xdd\xd1\xac\x71\x5b\xc0\xea\x65\x34\xf2\xf4\x6d\xb4\x4f\x66\x64\x5b\x4c\x07\x03\x54\x5b\x08\xed\xa8\x5e\x26\x56\xc4\x60\xaf\x44\xb6\x2d\xf5\x12\xfe\xb5\x58\x0c\xb7\xf6\xac\xd0\xe5\x12\x5e\xfc\x9f\x1f\xa7\xb7\x7f\x8e\xe9\x60\xfe\x69\xbc\xdc\x26\x95\x78\xf4\x14\xb4\x35\x7e\x8a\x48\xb5\x58\x77\xd4\x99\xaa\x2b\xa6\x49\x32\x4d\xd2\xab\x17\xd7\x69\x74\x6c\xdd\x6c\xc0\x27\x42\x5b\x17\xbf\x6e\x36\x1b\x94\xf6\x68\x34\x00\x6a\xe9\xef\x18\xee\xdb\xb2\x93\x71\xfd\x6f\x7b\x88\xac\x9b\xcd\x60\xbe\xed\x62\x09\x56\xb0\xb8\x04\x06\x3f\x40\xc0\xd0\xf6\xbf\x10\x4f\x9f\x0d\xdb\x2b\x76\x0d\x2b\x47\x68\x5e\x91\x37\xa2\xc0\x57\x9a\xb0\x14\x9e\xc1\xe2\x7e\xb3\x19\x99\x25\x5b\x05\x2a\xb1\xf6\x80\x5e\x57\x62\x4d\xae\x0c\xab\xeb\x29\x7c\xb6\xbf\xf7\x24\xd6\x33\xe6\xe6\x3e\x8e\x3d\xc2\x76\x9c\x4b\xef\xf0\x95\x22\x86\xc3\x14\x92\xad\x98\xb5\x9e\x67\xc8\xe3\x69\x71\x97\x17\xf6\x4c\xe7\x25\x68\x01\xa6\x7f\x73\xaf\xdd\x83\x7e\x67\x48\x67\xed\xd0\x65\x18\x0e\x56\x57\x3d\x5b\xdb\xaa\x29\x49\xe3\x90\x65\x6a\x46\x73\xcd\xee\xda\xa0\x6c\x5f\xa5\x80\x63\x1a\x84\xf0\x90\xfc\x7c\xcb\x19\x39\xfd\x88\x76\x54\x22\x7d\xa4\x76\x86\xf4\x69\xb5\x0b\x38\xfe\x0d\xed\xfa\xc9\xd1\x83\xda\xad\xa9\x7c\xa4\x72\x6b\x2a\x9f\x56\xb7\x9e\xe1\xdf\x50\xad\x1b\x63\x9c\xd1\xec\x1d\x16\x92\xee\x6d\xfe\x35\x9b\xa6\xc4\x73\x2e\x0e\x12\x4d\x4b\x6c\xd5\x74\x2b\x06\xa6\x59\x22\xae\x82\x3d\xa6\x93\xff\x06\x00\x00\xff\xff\xf6\x69\xeb\x34\xe3\x24\x00\x00")

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

	info := bindataFileInfo{name: "static/main.js", size: 9443, mode: os.FileMode(420), modTime: time.Unix(1481054562, 0)}
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

