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

var _staticMainJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x57\x4b\x8f\xdb\x36\x10\xbe\xeb\x57\x4c\x90\x00\x92\x0c\xaf\xbc\x97\x5e\x0c\x38\x40\x9b\xa0\x68\x0e\x45\xda\x6e\x7b\x0a\x82\x05\x2d\x8d\x2d\x35\x32\xa9\x90\xa3\xf5\xba\x59\xff\xf7\x82\x4f\x91\xb2\xdd\xc7\x65\xe1\x25\xe7\xf5\x7d\x9c\x97\x56\x8b\x45\x06\x0b\xf8\x81\x29\x84\xba\x67\x4a\xc1\x4e\x48\x10\xdb\x3f\xb1\x26\x05\xc7\xb6\xab\x5b\x60\x12\x01\x9f\x90\x13\x28\x31\xca\x1a\x55\x95\xc1\x62\x95\x59\xf1\x8f\x5b\x85\xf2\x89\x6d\x7b\x84\x6f\x19\x80\xe0\x85\x11\x5d\x42\xcb\x78\xd3\xa3\x2c\xcd\x31\xc0\x16\x19\xaf\x04\x2f\xa8\xed\xd4\x12\x66\x32\x19\xc0\x39\xcb\x00\x76\x9d\x44\xaf\x5f\x55\x55\xc3\x88\x25\xfa\xe6\x3e\xb1\xe0\xa5\x8c\x85\x73\x96\x39\x3c\x3f\x23\xc9\xae\x7e\x40\x02\xe4\x35\x1b\xd4\xd8\x33\x42\x05\xd4\x22\x28\x24\x10\x3b\x50\xd8\x63\x4d\xd8\xc0\xc1\x88\x2a\x20\x01\x5b\x84\xbd\x64\x43\x8b\x8d\xb6\xc1\x78\x63\x14\x6a\xd1\x3b\x51\xed\xc9\xf0\x43\x2d\x23\x2b\x1a\x53\x11\x39\x7d\x26\xe4\xcd\x05\x39\xab\xc5\x02\x06\x26\x35\x93\x9d\x02\xe6\x14\x14\x74\x5c\x11\xe3\x35\x3a\xbe\x35\xc0\xc8\x98\x11\x55\xe3\xd6\x05\xbe\x58\x65\x00\xb5\xe0\x8a\xe4\x58\x93\x90\x85\xb5\xe8\x79\x52\xe3\x80\xb2\x28\xcd\x6f\x6d\xa8\x72\x0e\x37\xce\xf3\x74\xe1\x81\x6f\x80\xe3\x11\x1e\x90\x62\x2d\x03\x75\x03\xdf\xcc\xb3\x24\x96\xf4\x23\xe6\xe3\xd0\x30\xc2\x7c\x09\x45\x09\x9b\xb7\xce\x35\xc0\x6a\x05\xef\xd8\x40\xa3\x44\xc3\xd5\x20\x3a\x4e\x36\xa3\x18\x3f\x5d\x50\xee\x94\xf4\x75\xd1\x23\x01\x67\x07\xd4\x08\xe3\xe8\xca\x60\xdb\x05\xa1\x48\x48\x7c\x34\x96\x0b\xad\x50\xba\x6b\x17\xa8\x93\x32\x89\x32\x45\xa9\xcf\xac\xe0\xb9\x8c\x00\x5d\x43\x92\x40\x95\xc8\x1b\xcd\x26\x9c\x21\xe4\x28\x6b\x9a\xc2\x06\xe7\x63\xeb\x76\x50\xbc\x32\x6a\x2d\x53\xfe\x6e\x0a\x3c\x86\x53\x45\xda\xd9\x0d\x54\x57\xae\x2d\x1c\xd6\x34\xf9\x12\x6e\x5e\x3b\x24\x0e\xa7\x0f\x57\xe2\x41\x3c\xe1\x95\x88\xff\x63\xc0\x0d\xf6\x48\x38\x0b\xca\x1e\x4e\x99\xf2\xc9\x5e\x7f\xbe\x0c\xca\xba\xff\xff\x61\xd7\x3d\x32\x4d\xfc\xb7\x8b\x7c\xad\xdc\xd5\xb5\x5c\x9d\x99\x36\x92\x79\x39\x3f\x8e\x3c\x1a\x5f\x11\x09\xce\x9f\x44\x1a\x25\x4f\xdd\x46\x52\x5e\xf1\xca\xab\x39\x03\x3a\x9d\x5d\xf6\x6f\x2e\x69\x82\x97\x17\xf8\x64\xc9\xb2\x42\xd5\x30\xaa\xb6\x08\xfc\x77\x07\x54\xc4\x0e\xc3\x3a\x49\xc6\x70\xbc\x74\x72\x4f\xac\x1f\x31\x95\x71\xc1\x26\xef\x71\x9e\x51\x15\xa2\xd8\x38\xef\x1e\xce\xd0\x0b\xea\x4f\x8f\x5a\xa6\x98\x31\xf1\xd1\x0c\x84\xea\x0b\x9e\x54\x11\x0c\x95\xd5\x81\x0d\xa6\x08\xe3\x1e\xa0\xa1\x4b\x76\x4c\x70\x6b\x19\x9f\x1c\xce\xe2\x54\xd6\xcf\x6b\x2d\x6f\x6c\x99\x78\xb4\x31\xf3\x63\x42\x5c\x2e\x83\xf4\xe9\xa6\xb4\xa1\x23\x92\x3c\x88\x06\xd7\x90\xf7\x1d\x47\x95\x4f\xc7\x3a\x96\xb5\xf9\x1b\x3a\x47\x60\xe9\x72\x84\xa8\x9b\x03\x84\xf5\x7d\x98\x1d\x6e\x54\x74\x12\x0e\x42\x11\x48\xac\x4d\xa7\x5d\xd8\x27\x52\x4b\xd8\x21\xd5\x2d\x36\xb0\x93\xe2\x60\x8c\xac\x1a\xdc\x8e\xfb\xd5\x21\x38\x69\x2c\x88\x8b\x71\xa2\x6e\x0d\x93\x78\x04\xdc\x6c\xfe\x53\x8f\x8f\xab\x43\x83\xd7\x67\x2e\x07\xcd\x59\xe0\x5a\xcf\x83\xb1\xef\xa3\xb6\x8f\xb2\x13\x0d\x6c\xe0\xbb\xfb\xfb\xfb\xb8\x96\x7a\x42\x39\x49\x9b\x14\x32\x78\xff\xbd\x98\xa2\xfc\x34\x6a\xb6\x22\x03\x0c\xf6\xdc\x09\x55\xed\x91\x8a\x3c\xe5\x29\xf7\xfd\xa3\xa2\x16\x79\x51\x48\x35\x24\xf3\xe7\x02\xb5\x54\x83\xc9\xc0\x2a\x1d\x39\x33\x1e\x2e\xb2\xdb\x0f\x9f\x54\x3c\xa1\x08\x8f\xf0\xde\xc4\xec\xb3\x68\x16\xda\xb5\xb8\xfc\x38\x49\x4f\x6f\xce\xaa\x90\x94\xa6\xd5\x30\x49\xd3\x3b\x93\x18\xe2\x67\xee\x38\xe9\xd4\xe8\x61\x03\xc7\x8e\x37\xe2\x58\x29\xa4\x0f\xee\xd0\xc5\x62\x04\x3d\xd1\xcb\xf8\x69\xe3\x7e\x36\x14\x17\x33\xc2\x1b\x9f\x06\x84\xf3\x61\xda\x6b\xf0\x92\xca\xa6\xdd\x5c\x21\xfd\x68\xd2\xa5\xb0\x59\x93\xf4\xf5\x90\x48\xf6\x47\x76\x8d\x2e\x37\xca\xdc\x34\x9e\x7a\xec\xf4\xd2\xd3\x83\x5e\xed\x4a\x57\xfa\x4e\xdc\x0a\x96\xd1\x69\xdc\x56\x6d\x3e\x9b\x3d\x23\x16\xf1\x9b\xcc\x5a\xe7\x00\xab\xfc\xbf\x66\x46\xc4\x4b\x89\x6f\x2e\xfa\x25\x1d\x4e\x57\x1c\x3a\xb6\x57\x31\xfc\x97\x17\x07\xc6\x80\xa8\x3a\x5e\xf7\x63\x83\x2e\x1f\x1d\x6b\x6e\x77\x79\x53\xe4\xaf\x1d\xf0\xbb\x5a\x8c\x9c\xf2\xb2\x22\x7c\xa6\x22\xe2\xa0\x47\xbe\xa7\xb6\x0c\xf2\xd6\x02\x36\xa9\x82\x9f\x6d\x5e\x7a\x66\xfe\x8e\x74\xb7\xc9\xcb\xaa\xa5\x43\x5f\x10\x1e\x06\xd3\x05\x5d\x79\x3c\x9a\xdb\x47\x7f\x1c\xc6\x97\xae\xb7\x75\xb2\xe3\x9d\xcb\xd0\x5c\x77\x23\xaf\xa9\x13\x1c\x24\x36\x92\x1d\xc3\x63\x1a\x1e\xf5\x00\xf2\xd3\x3c\x25\x36\x19\x4d\x46\xe1\x17\x73\x52\x39\x33\x41\xbd\x4c\x9c\x28\xa4\xc7\x51\xf6\xce\x8b\x4e\x98\xd1\xd5\xee\x1f\xbf\x7d\x30\x86\xc6\xca\x6e\x28\xbf\x8e\x28\x4f\x45\x6e\xa3\x36\x6d\x46\xe7\x7f\x1a\x84\x67\x4b\x75\x7f\x21\xbc\x85\x7b\x1f\xfb\xa8\xb7\xba\xd4\xc0\x12\xbe\x97\x92\x9d\x2a\xdd\xef\x6f\x58\xb1\xa3\xa1\x28\x1d\x37\x00\x6d\xa7\xf7\x89\x93\xd9\x05\x1e\xc8\x30\x7a\x5e\x42\x9e\x2f\x61\xac\x48\x3c\x90\xec\xf8\xbe\x28\x53\x80\x03\x93\x0a\x67\x10\xbf\xc6\x10\xab\xaf\x26\x2e\x92\x23\x7a\x50\x5f\xab\x2b\xcb\xe0\x69\x40\xb1\x8b\xae\x36\x1b\xc8\x95\x71\x99\x4f\x75\x9f\x02\xd1\xab\x6c\x50\xb0\xef\x0c\xd8\x2b\x0c\xe2\x61\xb1\x77\x29\x2f\x76\x30\xf7\x7d\xdd\x6a\xba\x2b\x9e\xa3\x5e\x12\x63\xaf\xc5\x61\xe8\xe2\x04\xb4\x36\x84\xeb\x2e\xae\xde\x7f\x32\x9f\x98\x5b\x26\x55\xe5\x14\x8a\x37\x93\xa4\xcd\xec\x19\xa9\x24\xf6\xfb\x1e\x1f\x6d\x14\xb6\x9e\x03\xb9\x7a\xf8\xcf\x92\xd3\xd1\xaa\x90\xa6\xfa\x0f\x8d\x1a\xc9\x25\x58\xe8\x0b\x09\x49\xfa\x5e\x23\x0e\x97\x49\x20\x1d\xef\x7c\xcf\x0f\xc5\xa7\x47\xb9\x51\xbd\x5a\x84\xeb\x4b\x56\xd2\x62\xbe\xf3\xe7\xb9\x4f\xbb\xf7\xd1\xf6\xfc\x3e\x1a\x95\x2e\x8b\xdc\x16\x62\x8a\x25\x81\x9d\xdc\xeb\xaf\xc7\x58\xb9\x9c\x19\x0b\x83\x47\xb7\x18\x57\xba\x1c\x8f\xfa\x57\x91\xd7\x2d\x93\x94\x2f\xff\xa9\xe4\x83\x3d\x7d\x0a\x1b\xd3\xa3\xac\x5a\xf9\xe9\xfe\xf3\x3c\xb4\xf4\xc3\xce\xb6\x88\x8b\xf8\x53\x21\xd7\x29\x2e\xe2\x76\x73\x57\x3f\xcb\x9b\xa2\x11\xf5\x78\xd0\xdf\xda\x95\x44\xd6\x9c\xa2\x21\x6f\x5f\x4a\x83\x8b\xa5\xb4\x87\xba\xef\xea\x2f\xf9\x12\x72\x92\xce\xe8\x9d\x14\x47\xa0\xc6\x34\x6a\xfd\xdd\x89\x7d\xb4\x2b\xa4\xb9\x87\x7d\x45\x4c\xee\x91\x4c\xbb\x7e\x27\x38\x69\xbb\x99\x5b\x0c\x6e\xba\x7a\x6d\x3f\x7c\xe6\xb6\x53\xf8\xd3\x67\xd4\x55\x63\x1d\x1f\x46\x32\xc6\xec\xe8\xb8\x6e\x2d\xd0\x14\x06\x3c\xf6\x55\x3d\x4a\xfd\x31\xf2\xbb\x8d\xdc\xae\xe5\x99\xff\xf4\x8e\x3a\x56\x76\x2e\xb3\xbf\x03\x00\x00\xff\xff\x3f\x56\x19\x32\x80\x12\x00\x00")

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

	info := bindataFileInfo{name: "static/main.js", size: 4736, mode: os.FileMode(420), modTime: time.Unix(1480957223, 0)}
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

