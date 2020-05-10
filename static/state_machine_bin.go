// Code generated by go-bindata.
// sources:
// schemas/state-machine.json
// schemas/state-machine-strict-arn.json
// DO NOT EDIT!

package static

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

var _schemasStateMachineJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x5f\x6f\xdb\x36\x10\x7f\xf7\xa7\x20\xd4\x3c\x3a\x75\xb3\x97\x0d\x79\x4b\xb3\x6e\x2b\xb0\x34\x41\x13\x74\x0f\x81\x0b\x5c\xe4\x73\xcc\x46\xa2\x54\xf2\xb4\x25\x28\xfc\xdd\x07\xfd\xb1\x22\x4a\xa4\x2c\x53\x6a\x6c\x04\x2a\x50\xa0\x20\x79\x7f\x78\xbc\xdf\xef\x8e\x16\xfb\x63\xc2\x98\x77\xc4\x17\xde\x29\xf3\xde\x28\x02\xc2\xe3\x10\xfc\x15\x17\xe8\x4d\xb3\x29\xe5\xaf\x30\x84\x74\x7a\x45\x14\x9f\xce\x66\xdf\x54\x24\x8e\xf3\xd1\xb7\x91\xbc\x9f\x2d\x24\x2c\xe9\xf8\xdd\xaf\xb3\x7c\xec\x4d\x2e\x47\x4f\x31\xa6\x42\xd1\xdd\x37\xf4\x29\x1f\x8b\x65\x14\xa3\x24\x8e\xca\x3b\x65\xa9\x61\xc6\xbc\xf3\x28\x0c\x51\x50\x39\x50\x11\x55\x24\xb9\xb8\xf7\xb2\xe1\xf5\x34\x5f\x7e\x4d\x20\xe9\x6c\x97\xe5\x54\xb1\x66\xf6\x2b\x1b\x8f\x81\x08\xa5\xb8\x6a\xba\x98\x4d\x7f\x7d\xfb\xe3\x64\x7a\xf2\xcb\x6f\xeb\x23\x6d\x3c\x0d\x90\xc4\x65\x16\xbc\xd9\x02\x97\x5c\x70\xe2\x91\x50\xb3\x2c\x92\x5e\xb9\x6e\x5d\xfc\x6b\x5d\xda\x83\xc5\x22\x5b\x0a\x81\x66\x72\x09\x81\x42\x6d\x0b\x5f\x50\x2a\x1e\x89\xae\x3b\xbe\xe1\x21\x46\x09\x5d\xa3\x1f\x89\x85\x71\xe7\x5c\x10\xde\xa3\x7c\xde\x7a\xc8\x05\x0f\x93\xd0\x3b\x65\xef\x26\x1b\x67\x33\x75\xad\x4e\x66\x0b\x24\x7e\x4f\xb8\xc4\x34\x7b\x6e\xf5\x03\xd2\x0f\x60\xc2\xd8\x3c\x13\xa8\xc4\xe8\x39\x07\xf2\x60\x75\x38\xa4\x48\xe0\xe5\xb2\xb4\x95\xfe\xe9\x70\x14\xfe\x2a\xe2\x7e\xf5\x2c\xa6\xbb\x48\x2f\x81\x07\xae\xb2\x31\x48\x08\x02\xec\x21\xaf\x94\xab\xac\x4a\x7c\x1f\x71\xe1\x2a\x4e\xa0\x1e\x5c\x65\xff\x03\x4e\xae\xb2\x21\xc4\x4d\xd4\xcc\xb5\x0c\x2f\xce\xf3\x74\x3b\xa4\x2d\x58\xbe\xc9\xd7\x6b\xde\xd4\x20\x35\xad\xce\x15\xd4\x90\x4e\x7f\x3d\xcf\x8c\x1f\x19\xb7\xe7\x7d\xc2\x47\xda\xa2\xd8\x28\xf7\x41\x2c\xea\x62\x28\x32\x40\xde\x92\x4c\x70\x6e\x14\x6a\x12\x67\x67\x7b\x97\x09\xc5\x09\x5d\x01\xad\x2c\xd2\xb7\x65\x1c\x98\x27\x92\x20\xf0\xcc\x2e\x7c\x14\x83\xa8\xc9\x63\xaa\x6c\x3b\x01\x29\xe1\x49\x3f\x11\x4e\x18\xd6\xd7\x5b\x53\xea\x32\x46\x09\x14\x49\xaf\xb2\x78\x6d\x74\xe4\x77\x5c\x42\x12\x74\x0f\x69\x83\xd4\xab\x74\x98\xa7\xd9\xf4\x79\x7b\xf3\xdd\xb9\x3f\x23\x9f\xfd\x24\xfa\x1f\xc0\x03\x4b\x9a\xbf\xa2\xcc\x83\x44\x6d\x8b\x90\x19\xb1\x52\x46\x72\xc8\x3c\x71\x48\x8e\x94\x2a\xf7\x93\x1b\x17\x10\x8f\x0c\xc8\x06\xcc\xc3\xcf\xa8\x92\x60\x08\x77\x52\x5e\xec\xaf\xe6\x02\x1e\xcf\x23\xe1\x27\x52\xa2\xf0\x9f\x6c\xf1\x15\x49\x78\x57\x69\x23\x99\xa9\x95\x34\x38\x98\x73\xb1\xad\x89\xd6\x6f\x20\x46\x1d\x57\x20\x21\x44\x42\x69\xad\x18\x05\x06\x2c\xa1\x26\x69\xdd\x52\xf7\x52\x93\x4a\x28\x03\xde\x8a\x69\x0b\xea\x8a\xd9\x8c\x3d\x3e\x7c\x4f\x20\x30\x4d\xb7\xfa\xd3\xea\x55\xdd\xb7\x3a\x02\xca\x78\xd4\x46\xd6\x75\x0b\xde\x47\x41\x28\xff\x85\xa0\x79\x8d\x30\x78\x69\x48\x84\x62\x85\x29\x1d\x6c\x26\x2f\xe0\xf1\x8c\x08\xc3\x98\x5e\xc4\xdc\x7b\xf0\x1f\xa2\xe5\xf2\xb3\x7e\xf3\x18\xca\xdc\xa4\xc5\xb8\xce\xfe\xd5\x5c\x98\x6f\xed\x50\xce\x81\x7c\x1b\xb8\xc7\xec\xdd\xa8\x30\x54\xa1\x9a\x6b\x66\xdd\x8e\xc7\x36\x2d\x2c\x5a\x8e\xaf\x51\xfd\xcb\x5b\x6c\xa5\x1c\x6b\xca\x35\x6d\xeb\x29\xb3\xad\x4b\xab\x66\xb9\x6c\xde\xde\x85\x96\xcc\xeb\xd0\x69\x94\xf7\xd8\xfd\xb4\x1b\x57\x85\xf9\xb1\xe7\x60\x07\xd8\x73\xbc\x97\x20\xfc\xd5\x10\xd7\xb7\x0d\xf2\x1b\x8c\xb3\xad\x3b\x60\x35\xec\x9a\xb9\x73\xac\xfc\x63\xe5\x1f\x2b\xff\xeb\xce\xde\xb1\xf2\x37\x2b\x7f\x49\xd0\x4e\x95\x5f\xa9\xbd\x55\x7d\xa5\xc6\x8a\xcf\x5e\xb2\xe2\xb7\x6e\xa6\xef\xb5\x5b\x15\xbf\xac\x1e\x02\x2a\x1c\x90\xb0\xf9\x9e\xb2\x1f\x30\x5c\xe7\xd6\x07\xfa\x41\xb6\x71\x02\x83\x44\x28\xfb\x64\xb4\x9f\xf0\xdc\x80\x7a\x18\xb9\x82\x0d\xcb\x15\x51\x22\xfd\xc6\x89\x18\xb0\xda\xba\xc5\x1a\x5e\x59\x1b\x5f\xa4\xa0\x75\x23\xae\xce\x9b\x1a\x6f\x00\xe3\x0d\x60\xbc\x01\xbc\xe6\xec\x3d\xac\x1b\x40\xe5\xf8\xac\x0f\x73\xd8\x96\x8c\xab\xe4\xda\x89\x51\xf3\x5f\x08\x92\xee\x10\x7e\x86\x6e\x87\xae\xef\x00\xfa\xbb\x69\xa5\x80\x39\x74\x32\xd9\x03\x96\xfd\x74\x32\xff\x00\xa7\xb1\x93\x61\x03\x76\x32\xbd\x40\x61\xfe\x8a\x99\x42\x59\x11\x84\xb1\x4b\x78\x0a\x7f\xfa\x6f\xac\xf4\xc2\x51\xd5\x01\xe0\xd4\x01\x9b\xe5\x6b\x9e\x2e\x6f\x3a\x2d\x00\xfd\x02\x92\xc3\x5d\xe0\xf4\xf2\xc3\x15\x85\x67\x4d\x14\xbe\xfc\x03\xa7\x4b\xeb\x9b\x95\x97\xf3\xe1\x53\xd4\x88\x5f\x57\x95\xda\x67\x87\x28\x0a\x10\x84\xb1\x17\x29\x77\x75\x97\x2f\xb2\x9c\x64\x12\xa2\xe4\x7e\xbb\x86\x82\x1e\xda\x14\xfc\x29\x11\x08\xe5\xcd\x0a\xc4\x30\x5a\x7a\x7b\xf4\x37\x2a\xd5\xd3\x9d\x8d\x0a\x77\x5f\xae\x33\x30\xb4\xcb\xb7\xf2\x64\x36\xd7\x21\xb8\x3b\x29\xe9\xeb\xcf\xb6\xd0\x76\xd7\xe0\xee\x49\x49\xfc\x03\xa8\xe8\x19\x5f\x93\x9e\x01\xbc\xea\x13\xe5\x86\x92\x1d\xfd\xd9\xbd\x2c\x9e\x55\xcb\xdd\xb3\x27\x9a\x39\x4d\x40\xe7\x2e\x63\x75\xb7\x0b\xa7\xfc\xb9\xab\x88\xc6\x74\x6e\xc2\xd5\x44\xe9\xad\xa1\x97\x27\x65\x72\xf4\x13\x77\xf3\xe1\x52\xee\x2a\xa1\xf1\xa0\x93\x6c\x8f\xd0\xdb\xf8\xcf\x49\x8d\x6b\xe0\x8d\xbc\xb7\xa3\x8e\x3a\xe7\xb9\x8a\xf7\x88\x65\x0b\xd7\xb9\x6a\x72\x8d\xa8\x8d\xe3\x9a\xdd\x7d\xf1\x5f\x27\x26\xe9\xdf\xf5\xe4\xff\x00\x00\x00\xff\xff\x67\x30\xbe\x62\xed\x35\x00\x00")

func schemasStateMachineJsonBytes() ([]byte, error) {
	return bindataRead(
		_schemasStateMachineJson,
		"schemas/state-machine.json",
	)
}

func schemasStateMachineJson() (*asset, error) {
	bytes, err := schemasStateMachineJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schemas/state-machine.json", size: 13805, mode: os.FileMode(493), modTime: time.Unix(1589122196, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemasStateMachineStrictArnJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\xdd\x6f\xdb\x36\x10\x7f\xf7\x5f\x21\xa8\x79\x48\x1a\x3b\x4e\xf6\xb2\xd5\x2f\x83\x9b\x65\x5b\x81\xa6\x09\x1a\xa3\x03\x96\xba\xc3\x45\x3e\xc7\x6c\x24\x4a\x25\xa9\x35\x59\xeb\xff\x7d\xd0\x87\x15\x51\x22\x65\x99\x52\x63\x23\x50\x81\x02\x02\x79\x5f\xbc\x8f\xdf\x1d\x2d\xe5\x5b\xcf\xb2\xec\x3d\x32\xb3\x47\x96\xfd\x82\x0b\x10\x38\xf0\xc0\x59\x10\x8a\x76\x3f\xde\xe2\xce\x02\x3d\x88\xb6\x17\x42\x04\xa3\xe1\xf0\x33\xf7\xe9\x20\x59\x3d\xf2\xd9\xed\x70\xc6\x60\x2e\x06\xc7\x3f\x0f\x93\xb5\x17\x09\x9f\x78\x08\x30\x62\xf2\x6f\x3e\xa3\x23\x92\xb5\x80\xf9\x01\x32\x41\x90\xdb\x23\x2b\x52\x6c\x59\xf6\xa9\xef\x79\x48\x45\xb6\x90\x63\xe5\x82\x11\x7a\x6b\xc7\xcb\xcb\x7e\x42\x7e\x25\x80\x89\xf1\x26\xe4\x22\xa7\x4d\x6d\x57\xbc\x1e\x80\x10\xc8\xe8\x65\xd9\xc4\x78\xfb\xd3\xd1\xb7\x93\xfe\xc9\x4f\xbf\x2c\xf7\xa4\xf5\xc8\x41\x0c\xe7\xb1\xf3\x86\x33\x9c\x13\x4a\x04\xf1\x29\x1f\xc6\x9e\xb4\x33\xba\x65\xfa\xb4\xcc\xf4\xc1\x6c\x16\x93\x82\x2b\xa9\x9c\x83\xcb\x51\x3a\xc2\x07\x64\x9c\xf8\xb4\xee\x89\x27\xc4\x43\x3f\x14\x57\xe8\xf8\x74\xa6\x3c\x39\xa1\x02\x6f\x91\x3d\x1e\xdd\x23\x94\x78\xa1\x67\x8f\xac\xe3\xde\xca\xd8\x58\x5c\xa5\x91\x31\x01\xc3\x2f\x21\x61\x18\x65\xcf\xb5\x1c\x20\x39\x00\x3d\xcb\x9a\xc6\x0c\x39\x1f\x3d\xe6\x40\xe2\xac\x1a\x41\xf2\x29\x5e\xcc\x33\x5d\xd1\xbf\x1a\xa1\x70\x16\x3e\x71\xf2\xb1\xe8\x6f\xc2\x3d\x07\xe2\x9a\xf2\x06\xc0\xc0\x75\xb1\x01\x3f\xe7\xa6\xbc\x3c\x74\x1c\xc4\x99\x29\xbb\x00\x7e\x67\xca\xfb\x15\x88\x30\xe5\xf5\x20\x28\x57\xcd\x54\xca\xf0\x34\x9e\xa3\xf5\x25\xad\xa9\xe5\x49\x42\x2f\x59\x53\x28\xa9\x7e\x7e\x2f\x85\x86\x68\xfb\xd3\x69\xac\x7c\x4f\x79\x3c\xfb\x1d\xde\x8b\x35\x82\x95\x7c\x67\x74\x56\x64\x43\x1a\x17\xe4\xb5\x60\x21\x4e\x95\x4c\x65\xe0\xac\xad\xef\x22\x14\x41\x28\x2e\x41\x2c\x34\xdc\xd7\x99\x1f\x2c\x9b\x86\xae\x6b\xab\x4d\x78\x43\x5b\x11\x93\xf8\x94\xeb\x4e\x02\x8c\xc1\x83\x1c\x11\x22\xd0\x2b\xd2\x6b\x53\xea\x22\x40\x06\xc2\x67\x76\x8e\x78\xa9\x34\xe4\x37\x9c\x43\xe8\xd6\x77\x69\x09\xd4\xf3\x70\x98\xa4\x59\xff\xf1\x78\xd3\xcd\xb1\x3f\x06\x9f\xed\x24\xfa\xef\x40\x5c\x4d\x9a\x3f\xa3\xcc\x83\x90\xaf\xf3\x90\xba\x62\x19\xf3\x59\x9b\x79\x62\x90\x1c\x11\x54\x6e\x27\x37\xce\x21\xe8\x10\xd0\x6a\x31\x0f\xdf\x23\x0f\xdd\x36\xcc\x89\x70\xb1\xb9\x98\x73\xb8\x3f\xf5\xa9\x13\x32\x86\xd4\x79\xd0\xf9\x97\x86\xde\x4d\x6e\x8c\xb4\x54\xa3\xa4\xc2\xc0\x04\x8b\x75\x43\xb4\x7c\x03\x51\xca\xb8\x04\x06\x1e\x0a\x64\xda\x8e\x91\xd6\x80\xc6\xd5\x82\x69\x8f\x54\xbf\xd5\x44\x1c\x5c\x51\x6f\xe9\xb6\xa6\xea\xd2\xdd\x18\x3d\xce\xbe\x84\xe0\xaa\xb6\x2b\xed\xa9\xb4\xaa\x68\x5b\xb1\x02\x32\x7f\x14\x56\x96\x45\x0d\xf6\x1b\x2a\x90\xfd\x0b\x6e\xf9\x1a\xa1\xb0\x52\x91\x08\x29\x85\x2a\x1d\x74\x2a\xcf\xe1\x7e\x2c\x04\x7a\x81\x78\x12\x75\xaf\xc1\xb9\xf3\xe7\xf3\xf7\xf2\xcd\xa3\x2d\x75\xbd\x0a\xe5\x32\xfa\xe7\x73\x61\xba\x76\x42\x39\x05\xe1\xe8\x8a\xbb\xcb\xde\x95\x08\x45\x17\x2a\x98\xa6\x96\x6d\x18\xb6\x7e\xaa\x51\x13\xbe\x52\xf7\xcf\x6e\xb1\xb9\x76\x2c\x09\x97\xa4\x2d\xfb\x96\x8e\x2e\xea\x9a\x19\xd9\xb4\x7a\x0a\xcd\x90\xd7\x60\xd2\xc8\xee\xb1\xdb\x19\x37\x2e\x53\xf5\xdd\xcc\x61\xed\xe0\xcc\xf1\x9a\x01\x75\x16\x6d\x5c\xdf\x56\x95\x5f\x42\x9c\x75\xd3\x81\x55\xa8\x5d\x35\x76\x76\x9d\xbf\xeb\xfc\x5d\xe7\x7f\xde\xd9\xdb\x75\xfe\x72\xe7\xcf\x00\xda\xa8\xf3\x73\xbe\xb5\xae\xcf\x79\xd7\xf1\xad\xa7\xec\xf8\x95\x87\x69\x7a\xed\xe6\xe9\x2f\xab\xbb\x50\x15\x06\x95\xb0\x7a\x9f\xb2\x9d\x62\xb8\x4a\xb4\xb7\xf4\x83\x6c\x29\x02\xad\x78\x28\x7e\x65\xb4\x1d\xf7\x4c\x80\xdf\x75\x58\x61\xb5\x8b\x15\x7e\xc8\x9c\x52\x44\x14\xb5\xba\x2e\x52\x85\x58\x01\xa3\x23\xf8\xca\x47\xfb\xd7\x30\xf8\x6f\xfa\x7d\x70\x70\xb8\x7a\xbc\x3e\x1e\xbc\x8a\x16\x5e\x8e\xe2\xa7\x97\x79\x92\xe8\x69\x3c\xf8\xfb\x78\xf0\x6a\xf0\xcf\xd1\xf4\x70\x7f\xb4\xff\xf1\xe3\xde\xdb\xf1\xe4\xec\x6a\xf2\x3d\xbf\x37\x3d\x3c\x38\xf8\x75\x4f\x7a\xd5\xd3\x57\xcd\x40\x2a\xbc\x8a\x40\xc3\x0c\x38\x6b\x3b\xb5\xbb\x81\x74\x37\x90\xee\x06\xf2\x9c\xb3\x77\xb7\x6e\x20\xb9\xf0\x69\x3f\x0c\xb2\xd6\x64\x5c\x2e\xd7\x4e\x94\x92\xff\x44\x60\xe2\x06\xe1\x47\xc8\x36\x98\x3a\x77\x60\xbe\xec\xe7\x1a\xa8\xc1\x24\x15\x7f\x40\xb3\x9d\x49\xea\x2f\x20\xa2\x9b\xa4\xac\x16\x27\xa9\x46\x45\xa1\x7e\x8b\x1a\x95\x32\x17\xe0\x05\x26\xee\x49\xed\x69\x7e\xb0\xcc\x0a\x43\x51\x3b\x50\xa7\x06\xb5\x99\x7d\x4d\x54\xe7\x9b\x52\x4d\x81\x7e\x00\x46\xe0\xc6\x35\xfa\xf2\xc4\xb4\x0a\xc7\xe5\x2a\x7c\xfa\x0f\xac\x2e\xb4\xdf\xcc\x3c\x9d\x0d\xef\xfc\x92\xff\xea\x8a\x94\x5e\x7b\xf8\xbe\x8b\x40\x95\xb3\x48\x76\xaa\x9b\x84\x48\x13\xc9\xd0\x43\x46\x9c\x6a\x09\x29\x3c\x54\x09\xf8\x83\x21\x08\x64\x93\x05\xd0\x76\xa4\x34\xb6\xe8\x2d\x72\xde\xd0\x9c\x95\x08\x73\x5b\xae\xe2\x62\xa8\xe6\xaf\xc4\xc9\x78\xaf\x86\x73\x37\x12\xd2\xd4\x9e\x75\xae\xad\x2f\xc1\xdc\x92\x0c\xf8\x5b\x10\xd1\xd0\xbf\x2a\x39\x2d\x58\xd5\xc4\xcb\x25\x21\x1b\xda\xb3\x79\x5b\x1c\xe7\xdb\xdd\xa3\x25\x92\x3a\x89\x41\xc6\x2e\x65\x77\xd7\x33\x47\xf8\xb9\x29\x8b\x84\x74\x66\xcc\xf9\x44\x69\x2c\xa1\x91\x25\x59\x72\x34\x63\x37\xb3\xe1\x82\x6d\xca\x21\xe1\xa0\x11\x6f\x03\xd7\xeb\xf0\xcf\x48\x8c\xa9\xe3\x95\xb8\xb7\xa1\x8c\x22\xe6\x99\xb2\x37\xf0\x65\x05\xd6\x99\x4a\x32\xf5\xa8\x0e\xe3\xca\xd3\x7d\xfa\xa7\x1b\xbd\xe8\xff\xb2\xf7\x7f\x00\x00\x00\xff\xff\x4e\x80\xd7\x4e\x6d\x36\x00\x00")

func schemasStateMachineStrictArnJsonBytes() ([]byte, error) {
	return bindataRead(
		_schemasStateMachineStrictArnJson,
		"schemas/state-machine-strict-arn.json",
	)
}

func schemasStateMachineStrictArnJson() (*asset, error) {
	bytes, err := schemasStateMachineStrictArnJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schemas/state-machine-strict-arn.json", size: 13933, mode: os.FileMode(493), modTime: time.Unix(1587363797, 0)}
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
	"schemas/state-machine.json": schemasStateMachineJson,
	"schemas/state-machine-strict-arn.json": schemasStateMachineStrictArnJson,
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
	"schemas": &bintree{nil, map[string]*bintree{
		"state-machine-strict-arn.json": &bintree{schemasStateMachineStrictArnJson, map[string]*bintree{}},
		"state-machine.json": &bintree{schemasStateMachineJson, map[string]*bintree{}},
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
