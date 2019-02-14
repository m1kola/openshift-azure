// Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// data/master-startup.sh
// data/node-startup.sh
package arm

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

var _masterStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x7c\x7d\x73\xdb\x36\xf2\xf0\xdf\xc7\x4f\xb1\x27\xe7\x9e\xb4\x99\x50\xb2\x1d\x27\x4d\xd4\xb8\x33\xb2\xac\x24\x9e\x38\xb6\x4e\x56\x72\xd7\xa7\xd7\xf1\x40\xe4\x4a\x42\x45\x02\x2c\x00\x2a\x56\xd2\x7c\xf7\xdf\xe0\x85\x14\x28\x51\x8a\x6c\xa7\x77\x75\x67\x1a\x11\xd8\x5d\x2c\xf6\x15\x58\x80\xdc\xfb\x7b\x6b\x44\x59\x6b\x44\xe4\x14\x42\xbc\x09\x82\x3d\x18\x5e\x9e\x5e\xb6\xa1\x85\x2a\x6a\xc5\x4c\xa6\x44\xfe\xde\x8c\x5b\x5c\xd0\x09\x65\x61\x9e\x49\x25\x90\xa4\x61\xcc\x64\x33\xe2\x6c\x0c\x54\x42\x94\x0b\x81\x4c\x25\x0b\x98\x12\x11\x47\x3c\xc6\xf8\x47\xa0\x2a\xd8\x83\x4c\xf0\x11\x19\x25\x0b\x90\x53\x9e\x27\x31\x7b\xa8\x60\x84\x41\x70\xd5\x1b\x7c\x38\xeb\xf6\xae\x87\x3f\xf7\x7b\xc7\x96\x72\x40\xc7\xf0\x0b\x84\x63\x68\x98\x81\xe5\x42\x6a\xea\x74\xd2\x22\x8a\xa7\x34\x0a\x79\x86\x4c\x4e\xe9\x58\x85\x8c\xc7\xd8\x80\x5f\x7f\x04\x35\x45\x16\x00\x00\x54\xc8\xad\xc2\x07\x63\xaa\x27\x25\x30\xe5\x73\x04\x81\x13\x2a\x95\x58\x40\x84\x42\xd1\x31\x8d\x88\x42\x90\x7c\xac\x12\xca\x66\x30\x16\x3c\x85\x98\x47\x33\x14\x41\xce\x4c\x93\x15\x83\x69\x6a\x69\x1c\xd9\x8c\x5b\x05\x91\x26\x89\x22\x94\xb2\x29\x30\x9e\x12\xd5\x8c\x78\xda\xb2\x3f\xc3\x88\x34\x23\xa1\x02\x3d\xa9\xbf\xc3\x44\x60\x06\xad\x39\x11\xad\x84\x8e\x1c\x2d\x4b\x77\x2c\x15\x19\x95\x13\x91\x0b\xa9\x30\x8d\x54\x02\x52\xf1\xcc\xf1\xd1\x94\x28\xe6\x34\xc2\x00\x20\x9d\x8d\x65\xf3\x66\x2c\xb5\x94\x5a\x31\xce\x5b\x31\x95\xb3\x16\xf9\x94\x0b\x6c\x09\x94\x3c\x17\x11\x86\x19\x11\xea\x20\x00\xc0\x68\xca\xe1\xe1\x76\x30\x58\xe3\x0a\x34\x79\x98\x88\xec\xf7\x9c\x2b\x02\xb0\x0f\xfb\x0f\xe1\xa7\x9f\x96\xcc\x6a\x36\x78\xce\xd4\x2a\x66\x00\x20\x50\x2a\x2e\x30\xe2\x0c\xc2\x41\x4d\x7f\x44\x14\xfc\xe4\x8b\x33\x26\x98\x72\xd6\xfc\x4d\x72\x06\x2f\x5f\x3e\xec\x5d\xbe\x7a\x18\x7c\x0e\x00\x1a\x09\x9f\x84\xb1\xa0\x73\x14\x8d\x36\x34\x7e\xe3\xb9\x60\x24\x89\x1b\xc1\x97\xa0\x77\xf9\x6a\x45\x50\x44\xa8\x55\x49\x59\x85\x8f\x69\x82\xce\xec\x60\x84\xd0\x12\x9c\xab\x56\xb3\xd0\xa4\xb1\x2d\x33\xf4\x63\x18\xe5\x0a\x48\xa4\x72\x92\x24\x0b\x60\x88\x31\x50\x05\x94\x05\x7b\xcb\x49\x58\x1b\x05\x35\x25\x6c\x26\x41\x71\x98\x2a\x95\xc9\x76\xab\x35\xa1\x6a\x9a\x8f\x8c\xea\x67\xf9\x08\x05\x43\x85\xd2\xff\x49\xa5\xcc\x51\xb6\x8e\x9e\x1e\x3d\xff\x21\x38\xed\x77\x86\x6f\x8e\x57\xa8\x16\x4c\x05\xe9\x2c\xa6\x02\xc2\x0c\x1e\x18\xb8\x60\x44\x24\x3e\x3b\x82\x30\x86\x97\x2f\x5f\xc2\xe7\xcf\xd0\xec\x5a\xb6\xcf\x52\x32\x41\x69\xff\xe9\xe7\x49\x72\x85\x91\x40\x05\x7f\xc0\x89\xc1\xe8\x31\xed\x82\xf0\xe5\x0b\xfc\xf4\xe0\xb3\x21\xf5\xc5\x9f\x70\x10\x4d\x53\x1e\xc3\xfe\xb3\xfd\x7d\xa8\xed\x2f\x43\x40\xc4\x99\xa4\x31\x0a\x18\x93\x48\xe9\xc9\xab\x35\xbb\x93\x91\xa4\x07\xad\x24\x67\xfb\x9b\x02\x81\x8e\x02\x44\x28\xaa\x28\x67\x5b\xd0\x7f\xf4\x46\xcb\x55\x2e\x10\xa4\x12\x44\xe1\x64\x01\x63\x2e\xb4\x71\xd1\x4f\x28\x81\x8e\x83\x3d\xa3\x23\x8c\xeb\x9c\x0b\x55\x14\xd7\xb9\x56\xe9\x3a\x5b\xf8\xff\xe3\x0f\x50\x22\xc7\x8d\xbe\xe3\x81\xae\x0c\x68\xbd\x26\xc6\x31\xc9\x13\x25\x77\xf2\x1a\x8d\xb7\xd9\x67\x4c\xaf\x36\x63\xc3\x49\xe3\xe4\xf2\x72\x78\x35\x1c\x74\xfa\xd7\xdd\xcb\x8b\x57\x67\xaf\xaf\x2f\x3a\xef\x7a\xc7\x3a\x0e\x86\x56\x6f\x61\x4a\xa4\x42\xd1\x28\x06\x5d\x46\xcf\x07\x9f\xfd\xe0\xf8\xc5\x04\xcf\x20\xf8\xfc\x39\x04\x3a\x86\x07\xcd\xde\x8d\x12\xa4\x39\x44\xa9\x9c\x69\x0d\x72\xc6\x28\x9b\xbc\x67\x31\x0a\xdd\x0c\x5f\xbe\x04\x12\x63\x08\x29\x84\x08\x0d\xb9\x77\xda\x3b\x79\xff\xfa\xfa\xfc\xf2\xf5\x79\xef\x43\xef\xfc\xf8\x70\xb5\xe1\x68\xaf\x01\xbb\xf0\xa0\x59\x40\x16\x6b\xfa\x81\x56\x70\x2c\xb5\xcf\x81\x8a\xb2\xc7\xcf\x8f\x8e\x8e\x7e\x84\x98\x07\x7f\xcb\x04\x57\xfc\xf8\xc1\xe7\x58\xaa\x7f\xfc\xe3\xf1\xa3\x2f\xc1\xdf\x32\x2e\x94\x6d\xd8\xdb\x7b\xf4\xf8\x4b\xf0\x37\x9a\x29\x32\x4a\x50\x42\xd8\x81\xcb\xab\xeb\x57\x67\x83\xde\xbf\x3a\xe7\xe7\xd7\x9d\xf3\xf3\xcb\x7f\x19\x67\x32\x44\x20\x4c\x75\xb0\x50\x08\x61\x68\xff\xbd\xe8\xfd\x4b\x37\x16\xdd\x61\xac\x49\xc3\x03\xf3\xff\xf0\x37\xe8\x74\xbb\xbd\xfe\x30\x88\x39\xc3\x20\x28\x06\x09\x25\x99\x23\xac\xca\xb8\xe8\x0d\x02\x91\x42\x28\xc6\x76\xfe\x5a\x87\xad\x47\xf6\xb7\x73\x74\xab\xa5\xd6\xa3\x60\xdd\xb5\xbb\x28\x54\x47\x9e\x2c\x14\xca\xd2\xcd\xbb\xcb\x14\x25\x9b\x3d\x15\xc5\x5d\x62\xda\xea\x1c\x7d\x39\xa4\x4b\x3d\x77\x1b\xe1\x0a\xc5\x1c\xc5\x0e\xa3\x48\x0b\x58\x3b\x52\x5f\xd0\x39\x51\xf8\x16\x17\xbb\x8e\xf7\x16\x17\x3b\x0d\x37\xc3\xc5\x1d\x27\xd6\xc7\x9d\xa6\x95\xe1\x37\x98\x94\x19\xeb\xab\x53\x32\x43\xe9\x09\xad\x8f\xf5\x33\x49\x93\x77\x44\xc8\x29\x49\xca\x51\x3a\x71\x4a\xd9\xdb\x7c\x84\xd6\xe8\x36\xd2\x76\xa6\xa6\x7d\xcc\xfc\xaf\x39\x2b\x71\xee\x20\xba\xaf\x59\x9c\x3f\x5a\xb1\xe6\x29\xd3\x58\x8d\xe9\x33\x92\x62\xfc\x27\x59\x7f\x75\x24\xfb\x4f\x53\x8b\x3a\xbc\xb3\x4b\xdc\x6e\xc8\x4d\xc3\xec\x68\x3a\x5d\xb2\xd5\x68\xd6\xc6\xba\x9b\x33\x5c\x66\xc8\xae\xf4\xba\xb8\xcb\x99\xe4\x09\xde\x66\x82\x46\x7b\x7a\xa9\x60\x10\xef\x33\xd7\x35\x2e\x76\x9f\x79\x95\x89\xbb\x09\xe1\x95\xe0\x4c\xf5\x05\xbf\x59\xdc\x4e\xc3\x63\x8d\x17\x66\x1a\xf1\xee\x46\x75\x65\x97\xa9\x57\x74\xa2\x33\xed\xed\x18\x70\x4b\xdc\x50\xd2\x09\xbb\x67\xa4\x5a\x63\x63\x77\x15\xac\x70\x71\xf7\xa8\xdc\x4d\x28\x32\x75\x67\xb7\xb6\xd8\xf7\x0d\xd7\x8e\x89\xdd\xa7\x5f\xc3\xc3\xdd\x44\xd0\x99\x4c\x04\x4e\x88\xe2\x62\x69\x90\xb7\x11\x06\x29\xf1\x43\xcf\x32\xef\x25\x90\x5a\x96\x76\x17\xcd\x06\x8e\xee\x26\x9e\x77\x86\xa6\xce\x79\x09\xaa\x3b\x9b\xca\xcc\xe2\x7f\x0b\x6b\xa9\x63\xe8\xd6\x66\xb3\xc2\xcf\x7d\x44\x63\x23\xd8\x5d\x05\xe3\xc2\xd8\xb7\x12\x8b\xcf\xcc\xad\x85\x52\xe1\xe5\x3e\x22\xd9\x65\x01\x5b\xcb\xc1\x37\x58\xd0\x56\x38\xb8\xb5\x08\xb6\x2d\x71\xfb\xf9\x28\xa1\x51\xcd\xf8\x2e\x88\x77\xa2\x48\x6f\x2f\xdf\xe2\xa2\x59\x82\xde\x2e\x96\x13\x4b\x41\x36\x33\x83\xbf\x81\x8d\x8d\x72\x58\xe3\xe3\xae\xa3\xdb\x11\xea\x87\xff\x06\xcb\xe2\x22\x4c\x69\xac\xad\x2b\xe3\xba\xb1\x96\xee\x7f\xab\xc1\x96\x95\x4b\x2f\x02\xb8\x71\x83\x65\x51\xcc\x6c\x0a\xf4\xff\x6c\x59\xf5\xe5\xcb\xde\xe5\xab\xa0\x37\xec\x9e\x5e\x77\x4e\x3f\xf4\x06\xc3\xb3\xab\xde\x75\xf7\xfc\xac\x77\x31\xbc\x7e\x3f\x38\xbf\x3a\x2e\x2a\x50\x0f\xbe\x9b\x72\xa9\xf4\xaa\xe8\xfb\xf6\xe1\x93\x1f\x5e\x58\x9c\x6e\x6f\x30\xbc\x7e\x75\x76\xde\x3b\xae\xdd\xb3\x59\x18\x4b\xcd\x80\x76\xde\x0f\xdf\x1c\x9b\xf2\x87\xe9\x3a\xed\x0c\x3b\xd7\xa7\x67\x83\xe3\x6a\x49\xc2\xf4\xf5\xce\x7b\xdd\xe1\xd9\xe5\xc5\xf5\xf0\xec\x5d\xef\xf2\xfd\xf0\xf8\xf0\xe9\xfe\xbe\xed\x7a\xd3\xeb\x0c\x86\x27\xbd\xce\xf0\xfa\xec\x62\xd8\x1b\x7c\xe8\x9c\x1f\x97\x7d\x67\x17\x67\xc3\xb3\xce\xb9\x37\x9b\x7e\xaf\x37\xd8\x36\x97\xe7\x2b\x98\xdd\xf3\xf7\x57\xc3\xde\xe0\xd8\x8a\x31\xdc\x37\x7f\x25\x6e\xa5\xd5\x60\x3f\xf6\x9b\x0e\x6a\x01\x0f\xd6\x01\x0f\x6b\x01\x0f\x3d\x7e\xde\xf6\x7e\xde\x20\x5a\x6d\xb6\x06\xe4\xfc\xec\x6a\xd8\xbb\xa8\xd5\xd7\x7e\xd3\xfc\xe7\xe9\xca\x01\xaf\x8b\x63\x09\x5a\x0c\x6d\x0a\x3f\x9e\x94\x6c\xab\xc1\xac\xd3\x78\xb9\x9d\xf5\xc0\x36\x2b\xdd\xf4\xd7\x4c\xae\xdc\xa9\x2e\xa1\x86\x03\xad\x8a\xd3\xeb\x6e\x67\x15\xd8\x2d\x8b\x0d\xe8\x3f\xdf\x5f\x0e\x3b\xd7\x27\x9d\xee\xdb\xde\xc5\xe9\xf5\xc9\xcf\xc3\xde\xd5\xf1\xd1\xe1\x8b\xa3\x17\xcf\x7e\x38\x7c\xf1\xcc\xc2\x7c\x9d\xd2\xe5\xab\x20\x88\xaa\xbb\x49\x6f\xbf\x59\xd3\x6e\xb2\x48\xb1\x3e\xff\x1a\x66\x36\xa3\xad\x88\x84\x4a\xe4\x52\xb5\x6c\x71\xbc\x45\x58\x34\xe5\x42\x7a\x9e\xeb\x88\xe5\x59\x4c\x14\x86\x05\xbc\xef\xbe\x75\x31\xdd\x55\xe7\x9a\x0b\x92\x26\xce\xa1\x75\xe4\x91\x92\x72\x66\x63\x4a\x3b\x00\xc8\x92\x7c\x42\xbd\x67\x80\x93\x9c\x26\xf1\xa9\x2b\x27\xda\x26\x00\x4b\x2b\x17\x44\x51\xce\x8a\x46\x00\x92\xd1\x0f\x28\x34\xc5\x36\xcc\x0f\xca\x66\x64\xf3\x36\xfc\xf2\x6b\xf9\x3c\xa3\x2c\x6e\x57\x09\xdb\x11\x4b\x88\xe2\x68\x40\x2e\x69\x03\x24\x34\xa5\x4a\xb6\xe1\xf3\x17\xaf\x51\xe0\xef\x39\x4a\xaf\xd9\x90\xbd\x9c\xa3\x10\x34\xc6\x3b\x32\xec\x31\x58\x52\xf2\x38\x2c\x55\xd1\xa4\xbc\x65\x0b\xdf\x3c\xa1\xd1\xe2\x8e\xd2\xb9\xc1\x28\xd7\x90\x83\x3c\xf1\x27\x1c\x42\x4a\x54\x34\x35\xf4\x3b\x8c\x71\x65\xc8\x55\x24\x12\xc2\x0c\x17\x6d\xa0\xb6\x04\x5f\x61\x2b\x46\xb6\x08\x4b\xd2\x1e\x0e\xc0\x9c\x24\x39\xb6\xe1\xa1\x76\xb6\x87\x5e\x8f\x76\xe2\xf6\x92\x9d\x30\x46\x46\x31\xf6\x00\x38\x1b\xd4\xe9\x25\x2c\xd5\xd5\x86\x8c\xc7\x72\x43\xd7\x48\x8b\x53\x56\x54\xf7\x1b\x46\xaa\x5d\x14\xba\x8b\x3f\x39\xa3\xd9\xa5\x19\x29\x31\x7c\xbc\x22\x34\xc9\x05\xae\xc0\x59\x15\x79\xc2\x77\xfa\x21\x79\x4c\xd5\xd2\x7e\x91\x91\x51\x82\xb1\x43\x5e\xee\x08\x3c\x8b\x5f\x2e\x15\xcf\xd8\x98\xdb\x89\x45\x28\xd4\x2b\x9a\x60\x1b\xb6\x6c\x6b\x0c\x1b\xb8\xd8\x0a\xa7\x23\x15\xc9\xe8\x39\xce\x31\x91\xed\x20\xd4\x8a\x5f\xb1\x03\x92\xab\xe9\x92\x1d\x67\xcf\x6f\x90\xc4\x28\x1c\x33\x86\xb9\x6e\xa7\x0d\x35\x5b\x7e\x0f\x80\xa7\x29\x67\x17\x24\x2d\xb4\x13\x6e\x60\x2a\xb0\x56\xa7\x04\xb1\xa3\xf4\x05\x8e\xe9\xcd\x12\xeb\xdf\xe1\x00\x53\xae\x30\x34\x05\xf9\xd0\xb4\x4e\x04\xcf\x33\x0b\xbe\x0e\xf7\x5a\x77\x9a\xc6\x5c\xa2\xd0\x66\xb4\x09\xf2\xbd\x44\x11\x44\x9c\x29\xc1\x93\x04\x3d\x2d\x60\x82\xd1\xd2\x5b\x12\x1e\xcd\x2e\x8c\x35\xae\xae\x57\xc2\x25\xb2\x36\x25\xb7\x64\x33\x2b\x3e\x36\xd1\x0b\x60\x4b\xc0\x16\x03\x4a\x7f\x2c\xb5\x59\x53\xb2\x70\xe6\x54\xe8\xb1\xa6\x9c\xe0\x0d\xd9\x86\xc6\xa3\x46\x10\x71\x21\x3b\x49\xc2\x3f\x62\x7c\x69\x22\xad\x6c\x07\x31\x93\xcb\xd9\x8c\x28\x8b\x3b\x71\x2c\x50\xca\x36\x14\x89\xf3\xf9\xfe\xd3\x27\xae\xef\x02\xd5\x47\x2e\x66\x6d\x50\x51\x76\x14\x60\xb9\xef\x2f\x0c\x30\x22\x6d\xa8\xa9\x19\xfa\x33\xd9\x50\x7b\xf0\x66\xb2\xa1\x32\x00\x90\x8b\xc4\x68\x26\x84\x8d\xcb\x36\x8d\x74\xa5\xb8\x20\x13\x5c\xce\x6a\x79\x90\xe8\xba\xac\xe1\xb4\xbd\x8e\x26\xe5\x75\x80\xd5\xb0\xc7\x8b\x7a\xdb\x0a\x19\x3f\x7e\xd5\x80\xf9\x44\x4c\xc4\x5b\x72\x36\xe6\x22\x25\xaa\x5d\x73\x2c\xf9\xca\xf4\xc0\x1f\x80\x32\x22\x99\x5e\x16\x5b\x5c\x3f\x66\x68\x0a\x94\x29\x6d\xb9\xc9\xc0\x9d\xa6\xbf\x71\xf2\x68\xbb\xf3\xdc\xb0\x3c\x66\x77\x27\x6b\x4d\x39\x8f\xda\x4f\xf7\xf7\xf7\x03\x1b\x89\xec\x7a\xdc\x05\xa1\x99\xbf\x2b\xf7\x95\xba\x59\x91\x35\x95\x81\x75\x5d\xd6\x6c\xd7\x01\x32\x2e\x54\x1b\x0e\xf6\x0f\x9f\xee\x07\x4b\xc9\xfb\xfc\xe8\xd1\x49\x46\xed\x66\xb0\x23\x26\x79\x8a\xac\xc8\xe5\x51\xc2\xf3\xd8\xad\x0d\x0a\x77\xf5\xd7\x10\xa6\x3f\x13\x7c\x4e\x63\xbd\x57\xf9\x94\x0b\x34\x7b\x02\x0f\xb9\xe8\x2d\xa3\x8e\x06\xb2\x6e\x68\x35\x17\x8e\x48\x34\x43\x16\x17\x00\xda\xba\x9e\x54\x00\x52\x8c\x29\x09\xd5\x22\xc3\x92\x48\x96\x25\x7a\x2b\x4b\x39\x6b\xcd\x59\xdc\xf4\x6c\xcc\x1c\x92\x8d\x72\xcd\xc2\xd2\x35\xff\x9b\xd3\x8a\x92\xdc\xc4\x22\x69\xcb\x95\xa1\x56\x66\x38\xd6\x8a\xaa\x19\xa9\x5a\x8f\xaf\x43\x9f\xe1\x62\x07\x6c\xab\x6c\xfb\x7c\xd6\x6f\xc3\xc1\xe1\x0f\x26\xb0\x1c\x7c\x3d\x87\x6d\xaa\xb1\x54\x02\xdf\xa6\xe2\x07\x80\x8c\xa6\x18\xe7\x65\xb8\xb6\xe0\x75\xfb\xe7\x02\xce\x9e\xdc\x97\xd1\x59\x5e\xf0\x18\xfb\x5c\xa8\x01\x61\x13\xbd\xf0\x78\xe8\xf5\x5d\xe5\x23\x86\xda\x7e\x7f\x38\x6c\x3e\x31\x81\xb2\x75\xf0\x4c\xf7\xeb\xe5\x4e\xa4\x31\x6d\x46\xd3\xeb\x47\x3b\x8e\x9d\xa7\xd1\x33\xde\x58\xaf\x7d\x5b\x1a\x47\xd7\xa5\x42\xc6\x6c\x3e\x59\x59\x08\x92\x28\xc2\x4c\x77\x2b\x64\x6a\xb8\xc8\x34\xe1\x1d\x2c\xed\xb1\x0f\xe3\x26\x07\x30\xca\x85\x54\x6d\x38\xda\xdf\x0f\xdc\xaa\xaf\xa0\xba\x13\x51\x83\xf4\x7b\x26\xdb\x70\x68\x28\xac\x4f\x46\xff\x72\x1e\x6c\x85\x56\x46\xc8\x73\xce\x33\xed\x54\xff\x83\xe9\x3e\xbb\xf7\x74\x9f\x18\x0a\x6b\x73\xf1\x67\xbb\xad\x4c\x61\x5b\x6c\x69\xe9\xfd\xe0\xdc\x44\xfd\x4c\x50\xa6\xa0\x51\xe4\xb3\x06\x7c\xd7\x3c\x45\x41\xe7\x18\xbb\x1a\x54\x11\xcf\x4d\x7e\x50\x84\x32\x5b\x1b\xa3\x11\x7e\x0f\x7f\xc0\xef\x39\x57\x26\x37\x30\x9b\x98\x97\x61\xd3\xf9\xaa\x4b\xd8\x2e\x6d\x46\x34\x16\x3a\xe2\x36\x0f\x0e\x9f\x5b\x8b\x3d\x32\xb3\xd3\x49\xd4\xda\xf3\x39\xb2\x89\x9a\xb6\xe1\x85\xa7\xd7\xb3\xbe\xa3\xd2\x3d\x3b\x1d\x38\x4a\x6e\x6d\xd0\xd2\x02\x71\x63\xf7\xcd\x0e\xcc\xae\x7e\xec\x05\x2e\x6f\xeb\xc7\xe7\x32\x74\x70\x99\xc9\x60\x4b\x4f\xf2\x88\xaf\x7a\x13\xaf\x2e\x31\x89\x94\xa8\xbe\xa1\xfc\x1a\xc5\x01\x58\xab\xe1\xcb\x52\x2f\x1a\x09\x53\xfe\x5e\x32\x45\x35\xe5\x71\x1b\x48\xae\x74\x86\xa7\x31\x32\x45\xd5\xa2\xef\x22\xad\x13\x4a\xc2\x27\x94\x79\xcb\xfd\x94\x64\x19\x65\x93\x77\x0e\x39\x4a\x08\x4d\x83\xe5\x86\xe5\xf3\x67\xf8\x8e\xb2\x18\x6f\xd6\x79\x6b\xf6\x05\xcf\x50\x28\x8a\xb2\xd9\xc9\xd5\xb4\x2f\xb8\x0e\xb4\xcd\xb3\xd5\x81\x61\xff\xfb\xa6\x16\x79\x95\x7f\x13\x58\xbd\x1c\xb0\x61\x0b\x67\x18\xf2\x76\x44\x98\x12\x9a\xf8\xfb\x38\xd3\x50\x3e\xd3\xd8\xef\x93\xf9\x28\xa8\x6c\xc0\xbc\x3e\xfd\x5c\x3e\x66\x02\xc7\x28\x04\xc6\xef\xdd\x1a\xdb\x87\xcc\x19\xfd\x3d\xc7\x6b\x0f\xc1\x46\xf0\xb3\xd3\x6f\x27\x9f\xe2\xa1\xe9\x12\xcd\xe9\xaa\xb0\x8a\x41\xed\x0d\xb0\x3f\x61\xe0\xf2\x6a\xd9\xca\xb0\x76\x05\x76\x99\x21\x3b\x3b\x5d\xa5\xe0\x40\x8a\x35\xaf\x53\x63\xae\xa6\x5c\xd0\x4f\x58\x67\xfc\xc6\xfc\x9a\x29\x8d\x04\x97\x7c\xac\x38\x4b\x28\x43\x73\xa9\xae\xf1\x8d\xe7\x33\x44\x46\x8c\x20\x1b\x2d\xe3\xa2\x87\xad\x92\xb1\xc6\xfa\x2c\x01\x14\x9f\x21\xfb\xab\xb1\x6c\x98\x5a\x61\xd7\x25\xea\x8e\xb7\xe0\xfd\x33\xa3\x76\x41\x7d\x23\xdd\xad\x93\x7f\xf5\xcf\xd3\x8b\x2a\x35\x89\x2b\x55\xb1\xb2\xe9\x1d\xb9\xe9\x4c\xf0\x4a\xa7\xa2\x58\x67\xb2\x22\x19\xba\x6e\x1b\xb5\xa5\x64\x7e\xa3\xb5\x5a\xb9\x79\xd9\x64\xc1\x42\x69\xe1\x4c\x81\x2e\x70\xca\xf6\x59\xb0\xf7\x7b\x87\xba\x79\x85\x8d\xe7\xcf\x8a\x35\x48\x69\x3f\x75\x60\x4f\xf7\xf7\x83\x4c\xf0\xdf\x30\xf2\x82\xb2\xdb\xce\xe8\x85\xd6\x95\xd9\x84\x73\xd1\x06\x73\x9b\x4f\x98\x2b\x0d\xfe\xd6\xae\x15\xf1\x34\xcb\x15\x1e\xbb\xd0\x2c\x31\xca\x05\x55\x0b\xbd\x17\x8e\x88\xc6\xb4\xf1\x3a\x92\x65\x8b\x5b\xf2\xc9\xfd\x76\xeb\xb0\xe8\x3c\x27\x23\x4c\x64\xdf\x9c\xd2\xd9\x0a\xd0\x53\x5b\x3c\xa0\xf1\x2a\xde\xc1\x7e\xf1\x17\x1e\xbc\x28\xfe\x5a\xa6\x35\x10\x3c\x57\x7a\xd3\x5f\x4e\x45\xe6\xa3\x98\xa7\x84\xb2\x9d\x43\xcf\x80\xe7\xf6\xb0\x50\x5b\xbe\x35\x74\x63\x77\x57\x05\x25\xdf\x2e\x64\xe5\x58\x69\x39\x6c\x4a\x18\x99\x60\x5c\x16\x5e\xc2\x42\xa6\xe6\xb7\xa9\x7a\x99\x30\xa4\xdb\xb3\x84\x2f\xcc\x43\x8d\x87\x64\xe5\x81\x56\xa5\x12\x51\x7b\x1c\x05\x90\x15\x47\x6b\x1a\xd8\x8d\xbb\xe5\xf8\x4c\xda\x0a\x49\xb1\x3b\xa8\x2d\x50\x1c\x1d\xd5\xd7\x27\x6a\xb6\x12\xde\xe9\x8d\x5f\x9f\x2a\xe7\xb2\xba\xab\xf0\x8e\x24\xf4\xdc\x6f\x06\xae\x6a\x7b\xc6\x5e\x25\x74\x32\x55\xd6\x38\xcb\xea\xd7\x90\xa6\xc8\x73\xb5\xea\x67\xe6\xb6\x8d\x7f\xca\x69\xec\x2d\xf4\xd8\xdb\xe9\x9e\x50\xb5\xe0\xb3\xd3\xad\x9e\x32\x3d\x97\x49\x24\x34\xc5\x86\x5d\xc3\x95\x6f\x47\x73\x9e\xe4\xa9\x57\xbc\x88\x17\x8c\xa4\x34\x32\x21\x56\x07\x02\xca\x26\xbd\x4a\xd5\xd2\x1e\x3a\x6c\x28\xee\xd7\x45\x8f\xf2\x02\xfb\xca\x72\xc5\xe6\xc9\xab\x4a\x58\x0a\x1c\x62\x3b\x08\x4d\xf4\xd0\xc1\xdf\x2e\xe4\x2b\xe5\x94\x0a\x92\x4e\x17\xab\xe7\x8d\x95\x00\x8a\x2c\x12\x8b\x6c\x3b\x91\x1e\x8b\xb6\xd0\xd8\x3e\xe7\xca\x46\xb3\x7a\x5d\x7f\x39\xe5\x46\x1b\x1a\xf3\x83\xc6\x63\xdd\xaa\x67\xae\x9f\x6d\xd9\xc7\xb6\x65\x02\x63\x6b\x46\x8d\x36\xfc\x62\xd4\xfa\xd9\x29\xb7\xa1\xb5\xa8\xe1\x2f\xf8\x07\xa3\xad\xff\xcf\x99\xd1\x58\x42\x23\xd5\x30\x40\x5f\x1e\xd7\x63\xbc\x23\x37\xbd\x93\xab\x0f\x4e\xc7\x39\xfb\x3a\xf8\xeb\x6e\xaf\x7f\x7a\x1b\x84\xce\xa7\x5c\xe0\x29\x95\xb3\x5b\x20\xa9\x68\x7a\xc6\x74\xac\xe3\x71\x67\x3c\xa6\x8c\xaa\xc5\x76\x94\x0b\xae\x47\xd8\x6d\xce\xaf\x91\xa1\x20\x49\x7f\x29\xd0\xad\xe0\x7d\x1e\x0f\x79\x82\x42\x43\xea\xbc\x33\x24\x94\xa9\xaf\xe0\x74\xa7\x18\xcd\x34\xf0\x3b\x4c\xb9\x58\xf4\x75\xe4\xca\x05\xee\x88\xa4\xa7\x72\x0b\x14\x2b\xd7\x13\xca\x62\xca\x26\xf5\xf0\xc4\x95\x9b\x1a\xed\xb2\x0d\xa0\x51\xa4\x88\x42\xc2\x7e\x27\x40\x23\x31\x89\xaf\xb4\xb6\xb2\x5d\xe0\x44\xdb\xab\xd7\xb8\x3c\x2b\x2b\x56\x7f\x8e\x01\x8f\xd5\x81\x87\xa5\xa1\x7e\x75\x56\x4d\xb9\xa0\x3a\xbb\x6d\xb6\xea\x22\xd1\x5f\x65\x02\x49\xdc\xb7\x18\xd6\x2b\x0c\xdc\x47\xd4\x31\xb9\xd1\x86\x83\xad\xb2\x5a\x35\xa8\x3b\x13\x3a\x47\x22\x95\x4b\x09\x78\x77\x7e\x4e\x48\x42\x58\x84\x71\x71\x48\xe5\x16\x13\x5a\x48\xb7\x25\x65\x2a\x56\x66\xc3\xd5\x99\x73\x1a\xf7\x79\x2c\xb7\xb1\x65\xd6\x23\x5f\xa3\x77\x6f\x31\x19\x3f\x71\x9e\x43\x39\xbb\x35\x9d\xed\x36\xcb\x14\xdd\x6a\xb7\x9a\x83\x4f\x9c\x61\x63\x07\xd3\xd4\xe1\xb2\x86\xad\xc3\xa5\xad\xba\xd7\xa7\xec\x3b\x25\x0f\x4d\x66\x35\x4b\x04\x38\x78\xf6\xbc\xf9\xec\x49\xf3\xe0\xf0\x45\xf3\xe0\xd9\xc3\x9a\x9b\xdd\x02\x25\x4f\xe6\xb6\x58\x5b\x7b\xbb\xbb\x52\xda\xad\x49\x22\x9b\x4a\xbf\xcb\x3c\xe2\x25\xf6\xae\x06\x2e\x76\x3f\x3a\x1a\xd6\xe6\xf6\x2b\x25\x28\x9b\x94\x79\xab\x78\x65\x89\xb2\x39\x4a\x45\x27\x44\x21\xa8\x29\x42\x18\xa6\x84\xd1\x31\x4a\x15\xe6\x22\x01\x77\x77\x0f\x32\x22\x48\x8a\x0a\x05\x10\x16\x83\x44\x04\x3a\x06\xaa\x20\xd5\x62\x0b\xf6\x60\x8a\x49\x06\xb9\x04\xa2\x80\x24\xc9\xfa\x84\x8c\x58\x32\x1e\x4b\x7b\x63\xc7\x3f\xe0\xaf\xcb\xff\x7d\x1e\x07\x29\x2a\x12\x13\x45\x4c\x39\x68\xf5\x6c\x79\x99\x5e\x49\x92\x4d\xc9\xea\x1e\x40\x47\x97\x88\x24\x61\xc6\x63\x57\x96\xb4\x71\xcd\x22\x57\x0e\xa1\xf5\x7e\x81\x33\x64\xaa\x0d\xee\x1d\xa3\x35\x00\x53\xbf\x0f\xb3\x84\x30\xff\x48\xda\x56\x76\xdc\x91\x9f\x43\x35\x46\x92\x91\x08\xed\x81\x53\x68\x5f\xbc\x0b\x64\x86\x51\xdb\x1d\x05\x18\xb5\xb8\xc5\x30\x11\x93\xf2\xfc\xf1\x0f\x67\x8a\x12\x15\x84\xc4\x3d\x34\xa1\xe6\xaa\x93\xeb\xc3\x1b\x8c\x96\x2c\x47\x3c\x4d\xc9\xf2\xf8\xc2\xbc\xa3\x2a\xa7\xee\x29\x8c\xcc\x0f\x73\xaa\x54\x77\x02\x65\x0f\x62\x7a\x2a\x8a\x57\xb7\xf4\xb4\x78\x65\xce\xdd\x1e\x80\x4e\xf2\x91\x2c\xec\x09\x79\x42\xe7\xc8\x50\xca\xbe\xe0\xa3\xb2\xd2\xa3\xb9\x5a\x16\x31\x2a\x5c\x41\x79\xb0\x12\xa9\xc4\x6b\x09\xc3\x88\x98\xc3\x05\xaf\x6d\xed\x56\x4b\x05\xbc\x38\xcb\xa8\x45\x28\xef\xf0\xf8\x28\xc5\xf9\xc5\x66\x8c\x72\x09\xed\x30\x90\xc5\x19\xd7\x49\xdf\x6b\xdd\x78\x16\xb9\x04\x29\x0e\x4e\xa6\x48\x12\x35\x75\x1d\x3a\x62\x51\x92\x9c\x62\x42\x16\xe5\x86\xe1\xc8\xee\x28\xdd\x8d\x86\x42\x8f\xc5\x7e\xd5\x14\xc4\x6f\x54\x21\x39\xbd\xb1\xa2\x09\x4e\xca\x05\xb7\x6e\xb4\xab\xf4\x77\x66\x23\x55\x28\xde\xbc\x36\xd7\x27\x6a\xda\x5e\x4e\xd0\xd1\x58\x8e\xe4\x0e\x9d\x5c\xbb\x4e\xb0\x97\x2c\x59\x78\x94\xab\x74\x56\xde\xc1\x5b\xa1\xa5\x7d\xd4\xb4\xea\xed\x18\x65\x93\x53\x2a\xd6\x71\xb4\xbc\x96\x3b\x36\x3b\x8c\x5b\x06\x2c\xba\x09\x91\xd2\x95\x25\x8c\xbf\x84\xf6\x4d\x3d\xe7\xc4\x41\x31\xd1\xe2\x40\x98\x4b\xcb\x58\x21\x9a\x9a\xc9\xd6\x4d\x75\x23\xe6\xda\xf4\x56\x27\x57\xbf\xd0\x5f\x86\x34\x92\x51\xb7\x79\xfc\xfa\xce\xe6\x7f\x13\xd9\x48\x46\xef\x18\xd8\x2c\xe6\x3d\xe2\x9a\x79\xe3\xb8\x30\x2a\x43\xd3\x3d\x14\x3c\x19\x8f\x36\x4a\x3a\xde\xe9\x62\xd8\x2d\x5f\xc6\x2c\x86\x48\xf8\x24\xc1\x39\x26\xc7\x47\xf6\x55\xca\x44\x62\x6d\xf7\xa1\xff\xa6\x65\x4d\x64\x5d\xbe\x27\xbf\x35\xa4\x76\xad\x70\xfb\x5a\xb6\xf7\x0e\xaa\x3a\xee\xbc\x46\xb5\x0c\xa3\xd6\x72\x6d\x8c\xf9\xb4\x6c\x35\xa7\xec\xb6\x44\x62\xff\xb4\x31\x69\x65\xbe\x19\x0e\xfb\x57\x3b\x04\x23\x00\xb5\x52\xd6\x38\xd8\xf7\x5c\xa2\x50\x99\x8e\x18\xf4\x96\x5c\xb6\x34\xd2\xe2\x5b\xf0\xea\x58\xda\xc8\xeb\xb7\x0e\xa1\x15\x83\xac\xc4\xbf\x8a\x71\xee\x1a\x4d\x37\x2e\x03\xeb\x28\x57\x4e\xf9\x6f\x1b\xaf\xdd\x28\x75\x84\xcb\xa0\xbd\xce\x5b\x36\xa3\x75\x18\xb6\xf9\xbf\x13\xc7\xd7\x25\x5e\x2f\xef\x5d\x88\xac\xcb\x78\xbb\x84\xbf\x9a\x26\xdc\xe7\x37\x36\x89\x74\x2b\x4f\x85\x6c\xd7\x24\xfb\xb5\x14\xe3\xdd\xef\xba\x43\x92\xb1\xc3\x55\xaf\xa5\x6d\x08\xe9\xbb\xe5\x93\x2a\xa9\xed\x79\xa5\xa1\xcd\xa4\x71\xbf\x5c\xb1\x3a\xde\xad\x73\x46\x19\xe6\xa9\x54\xc8\xd6\x2e\x8c\x3f\x3f\x3a\x3a\xfa\x4b\xe5\x15\x64\xf3\xa2\xc3\x2a\xef\xed\xfb\x93\xde\xf5\xbb\xce\xbf\xaf\xfb\xa7\xd7\x1f\x2e\xcf\x8b\x00\xe9\xee\xc9\xfa\xdb\xc1\x77\xe4\xe6\x94\x28\x72\x4a\xe5\x4c\xf6\x51\x7c\x78\xb7\xbd\xd4\xfb\x17\xcd\x62\x46\x23\x9b\x53\x43\x9d\x45\xff\xc9\x91\xbf\x36\x8e\x7a\xa1\x7f\xb7\x28\x5f\x4f\x65\x3d\xcc\xaf\x53\x5b\x3a\x4e\xf1\x3d\x1d\x8a\xb2\x19\x57\x09\x4a\x3a\x61\x44\xe5\x02\x43\x9a\x6a\x41\xd6\xd0\xca\xa5\x89\x63\x7a\xa3\xb6\xfa\x09\x96\x04\x55\x68\x6f\xd7\x17\xbb\x1e\x8b\x28\x78\x46\x26\xae\xe0\xfe\x86\x4b\x35\xe4\xa5\x45\x55\x46\x5f\xa7\xf1\x3f\xc8\x17\xad\xbf\x4a\xc2\xd8\xa2\xb4\x60\xab\xc6\x36\x52\xbc\x85\xea\xea\x15\x62\xd2\x0c\x1d\xc3\x2f\xbf\x40\xc3\xdb\xc4\x36\xe0\xf8\x18\x1a\x95\x17\x82\x1a\xf0\xeb\xf2\x8b\x51\x5b\x12\x93\x5c\xb0\xe8\xce\x19\x49\x23\xdf\x7e\x77\xb1\x31\x5a\x5d\x2d\xcc\x79\xcd\xce\x51\xaa\xc2\xc4\xb7\x8d\x1e\xd7\x7a\xa2\xad\x6b\x9e\xab\xbb\x2f\xeb\xea\x7c\xe7\x96\x1e\x71\x3b\x5b\xd6\xd6\x31\xa6\x41\xf9\xfd\x1a\x61\x3e\x60\xf3\x1d\x3c\xb2\xbb\xdb\x36\x7c\xdf\x7c\xb4\xf7\x9f\x83\x9a\x2a\x48\xf1\x0d\x9b\xdd\x5f\xe4\xbb\xf7\x18\xb5\x1f\xec\xd8\x83\x37\x9d\xee\x5b\x9d\x18\xb2\x05\xac\x74\x82\xe2\x30\xe2\x5c\x49\x25\x48\xe6\xb7\x4b\x6e\x3f\xd5\x54\xb2\xeb\xbe\x92\xa5\xf1\x21\xd7\x8c\x9a\x2f\x59\xa1\xb5\x94\x3d\xb0\x9f\xe1\x92\xa8\xe0\x23\x4d\x12\x60\x5c\xc1\x98\xd0\xc4\x55\x46\x95\x01\xb5\x13\xb6\x24\xec\x82\x0a\x22\x2e\x04\x46\x2a\x59\x34\x6b\xdf\xb1\x5a\xe5\x76\x0d\xa0\x8e\xf7\xb5\x2f\x13\x9d\xc9\x01\x46\x7c\x8e\x62\xa1\x1d\x20\xe1\x93\x09\x0a\x08\x55\xa1\x70\x33\xb1\x3c\x6b\xca\x29\x34\xcc\x6f\xca\x26\x20\x0a\x0c\xce\xc0\x0f\x0a\xc1\x1e\x48\x85\x19\x1c\xc0\xc4\xcd\x6a\x44\xa2\x59\x9e\x99\xaf\xf6\x0c\xc6\xd5\xc2\x49\xeb\x51\xa0\x30\xcd\x4e\x48\x34\x3b\xa5\xa2\xfa\x1a\x64\xcb\xe1\xd9\x52\xf8\x03\x0f\x2e\xf8\x38\xa5\x09\xc2\xdf\xdd\x0d\x76\xc8\xf2\x24\xa9\x71\xec\x9e\x8a\xe2\x13\x43\x03\xbe\x7c\xf9\x11\x62\x0e\x32\x41\xcd\xd9\xbe\x7e\x60\x18\x38\x74\x91\x33\x08\x43\xcd\x5e\x71\x03\xd1\xf8\x08\xfc\x47\xbb\xcb\xbc\xc6\x44\xdb\xeb\x4d\x75\xc0\xd5\x0c\xde\xf6\xbc\xbb\x80\xf6\x27\xd5\x6e\xf1\x5c\xb5\x3f\x99\xae\xaf\xcc\xc5\xa2\x8f\x12\x3e\xd2\x22\x3f\xd6\xd0\x56\x8f\x16\xe2\x24\xe1\x23\x73\xe5\xaf\x80\x8c\x51\x2a\xca\x4c\x26\x3e\xd6\xa3\x38\xc9\x36\xe3\x11\x34\x62\xfe\x91\x25\x9c\xc4\x8d\xed\x6a\xb7\x18\x50\x40\x63\xac\xfd\xa2\x01\x0f\xbe\x4b\x64\x65\x16\xdf\x07\x85\xfe\x0f\x8b\xaf\x6e\x6d\x27\x6c\x81\xb4\x41\x49\x46\x32\x39\xe5\xaa\xb1\xa3\x7a\xbd\xf2\xf3\xbd\xd5\xab\xed\xad\x5d\xfe\x2a\xbb\x7c\x73\x6c\x57\x9f\xac\xa2\x42\x84\xde\xb0\x7b\xda\x1d\x9e\x5f\x77\xfa\x67\xc7\x8d\x27\x8d\x0d\xfa\xab\x30\x6b\x60\x5c\x45\xbb\x9c\x76\x21\xae\x8a\x38\x3d\x4d\x99\xe1\xcc\x76\x31\xd4\xfe\x50\x75\x15\x86\x1f\x1d\x80\xb9\xa5\xe1\xf9\xa3\x6b\x76\x25\x91\xd0\x15\x9a\x57\xd2\xf6\x7f\xf5\xed\xde\x46\x3d\x4b\xa1\xb9\x16\x66\xeb\xa7\x63\x2e\x42\x73\x7c\xb5\x02\x4a\xe2\x39\x0a\x45\x25\x86\x19\xa2\x08\x73\x91\xc8\x0d\xd5\xf5\xe7\xfb\x41\x90\xce\xd7\xa5\xd4\x7a\xb4\xd2\x56\x7e\x4c\x6c\x15\xb2\xe8\xa8\x04\x9e\xca\x87\xe4\x56\xc8\xef\x62\xe7\x68\x2c\xb3\x51\xf9\x2c\xdb\x5e\x11\xfd\x35\xdd\x34\x45\xa6\xbd\xcb\x7b\x27\xd6\xdf\x17\x43\x66\x6f\x74\xe9\xa0\x3b\xa6\x42\x2a\x93\x9b\x02\x95\x33\x8c\x43\x12\xa7\xcb\xfe\x7a\x7c\xed\x9d\x8c\x2b\x6c\x43\xdd\xd7\xe2\x20\x12\x44\x4e\x21\xe1\x3c\x93\x90\x33\x45\x93\x22\x23\x51\x09\x79\x16\x2c\x3f\x08\x69\x5f\x3e\xac\x25\x52\x7e\x1f\x72\xf5\xf3\x91\xdb\x80\xe1\xff\x69\xce\x62\x2a\xc9\x28\x31\xc9\x45\x2e\x64\xc2\x27\x20\x29\x8b\x10\x3e\xa2\xbb\x8a\x06\xa8\x33\x8e\x9a\x6a\x10\x35\x15\x3c\x9f\x4c\xa1\xf8\x6e\xa5\x37\x9e\xa5\x83\x05\x95\x5a\x8e\x78\xb6\xd6\xed\x1d\xba\xda\xaf\x58\xea\x8c\xb9\x9a\x75\x37\x7c\xda\xc0\x47\x71\x5f\x79\x0c\xfe\x2f\x00\x00\xff\xff\x10\x34\x0e\xde\xf4\x55\x00\x00")

func masterStartupShBytes() ([]byte, error) {
	return bindataRead(
		_masterStartupSh,
		"master-startup.sh",
	)
}

func masterStartupSh() (*asset, error) {
	bytes, err := masterStartupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master-startup.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nodeStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\x6b\x6f\x1b\xb7\x12\xfd\xce\x5f\x31\x91\x8c\xeb\x7b\x83\xbb\xbb\x76\xea\xba\xa9\x63\x1b\xf0\x43\x49\x83\x3c\x64\xd8\x4e\x80\xa0\x28\x0c\x2e\x39\xab\x65\xc4\x25\x37\xc3\x59\x25\xaa\xa2\xff\x5e\x70\x57\xaf\xd8\x6e\x9c\xa2\x9f\xb4\x22\x0f\xcf\x1c\xce\x93\xfd\x47\x59\x6e\x5c\x96\xcb\x50\x42\x82\x5f\x84\xe8\xc3\xf5\xf0\x7c\x78\x00\x19\xb2\xca\xb4\x0b\x95\x0c\x9f\x52\x9d\x79\x32\x23\xe3\x92\xa6\x0e\x4c\x28\xab\x44\xbb\x90\x2a\xef\x0a\x30\x01\x54\x43\x84\x8e\xed\x14\x4a\x49\x5a\x79\x8d\xfa\x19\x18\x16\x7d\xa8\xc9\xe7\x32\xb7\x53\x08\xa5\x6f\xac\x76\xdb\x0c\x39\x0a\x71\x35\xb8\x7c\xff\xf2\x6c\x70\x73\xfd\xe1\x62\x70\xd4\x31\x0b\x53\xc0\xef\x90\x14\xd0\x6b\x0d\x87\x69\x88\xec\x66\x94\x49\xf6\x95\x51\x89\xaf\xd1\x85\xd2\x14\x9c\x38\xaf\xb1\x07\x7f\x3c\x03\x2e\xd1\x09\x00\x80\x6f\xe8\x6e\xe3\x45\x61\xe2\xa5\x08\x2b\x3f\x41\x20\x1c\x99\xc0\x34\x05\x85\xc4\xa6\x30\x4a\x32\x42\xf0\x05\x5b\xe3\xc6\x50\x90\xaf\x40\x7b\x35\x46\x12\x8d\x6b\x97\x3a\x37\xb4\x4b\x59\x3c\x13\x52\x9d\x2d\x49\x52\xa9\x14\x86\x90\x12\xea\x52\x72\xaa\x7c\x95\x75\x9f\x89\x92\xa9\x22\x06\x11\x6f\xf5\x08\x46\x84\x35\x64\x13\x49\x99\x35\xf9\x82\xac\x23\x2e\x02\xcb\x7c\x75\x93\x30\x0d\x8c\x95\x62\x0b\x81\x7d\xbd\x10\x92\x06\xa4\x89\x51\x28\x00\xaa\x71\x11\xd2\x2f\x45\x88\x6e\xca\x34\x4e\x32\x6d\xc2\x38\x93\x7f\x36\x84\x19\x61\xf0\x0d\x29\x4c\x6a\x49\xbc\x2b\x00\x50\x95\x1e\xb6\xbf\x0f\x83\x3b\xaa\x20\xd2\xc3\x88\xea\x4f\x8d\x67\x09\xb0\x03\x3b\xdb\x70\x7c\xbc\x16\x1b\x65\xf8\xc6\xf1\xed\x93\x02\x80\x30\xb0\x27\x54\xde\x41\x72\x79\x67\x7f\x36\x4b\xc0\x14\x80\x9f\x20\x1d\x7c\x61\x92\xe9\xa5\xb7\x08\x3d\xe3\x0a\x92\x3d\x98\xcf\x05\x80\x92\x0c\xc7\x9b\x0e\xd7\x12\x2b\xef\xd2\x8f\xc1\x3b\x38\x3c\xdc\x1e\x0c\x9f\x6f\x8b\x99\x00\xe8\x59\x3f\x4a\x34\x99\x09\x52\xef\x00\x7a\x1f\x7d\x43\x4e\x5a\xdd\x13\x73\x31\x18\x3e\x6f\x4d\xa1\xd3\x1d\xe9\xa6\x53\x25\xf1\x6d\xaf\x76\xd9\x51\x18\x8b\x8b\x1c\x85\x1c\x21\x23\xef\x39\x4b\x97\x61\x6f\x13\xb1\x55\xf1\x7f\xc8\x1b\x06\xa9\xb8\x91\xd6\x4e\xc1\x21\x6a\x30\x0c\xc6\x89\xfe\xfa\xc2\x5d\x42\x03\x97\xd2\x8d\x03\xb0\x87\x92\xb9\x0e\x07\x59\x36\x32\x5c\x36\x79\x9b\x27\xe3\x26\x47\x72\xc8\x18\x36\x3f\x4d\x08\x0d\x86\x6c\xef\xe7\xbd\xa7\xbf\x88\xf3\x8b\x93\xeb\xdf\x8e\x6e\xb1\x2e\x45\x89\x6a\xac\x0d\x41\x52\xc3\x56\x8b\x13\xb9\x0c\xb8\xbf\x07\x89\x86\xc3\xc3\x43\x98\xcd\x20\x3d\xeb\x64\xbf\xac\xe4\x08\x43\xf7\x73\xd1\x58\x7b\x85\x8a\x90\xe1\x2b\x9c\xb6\x27\x06\x2e\xd6\x2b\xcc\xe7\x70\xbc\x35\x6b\xa9\xe6\x9b\x17\x16\xaa\xac\xbc\x86\x9d\xfd\x9d\x1d\xb8\x77\x5f\x74\x79\x76\x3a\x1c\x5e\x5f\x5d\x5f\x9e\x5c\xdc\x9c\x0d\xdf\x3e\x7f\xf9\xe2\xe6\xed\xc9\x9b\xc1\x51\x2c\xd5\xa4\x43\x27\x51\xd2\x46\xe0\xe7\xf3\x55\x5e\xad\x6b\x7d\x6b\xb6\x59\xca\xf3\xb6\xd4\xc5\x32\x73\xb6\x16\xc7\xaf\x31\xf0\xe2\x6e\x97\x8d\x73\xc6\x8d\xde\x39\x8d\x14\x97\x63\xc4\x03\x6a\x48\x0c\x24\x08\xbd\xd0\x3f\x1f\x9c\xbe\x7b\x71\xf3\x7a\xf8\xe2\xf5\xe0\xfd\xe0\xf5\xd1\x93\xdb\x0b\x7b\xfd\x1e\xfc\x88\x86\x8d\x8c\x12\x54\x41\x42\x45\x77\x0c\x59\xe9\xec\x71\xf7\xbd\x08\x50\x25\x03\x23\x65\x8f\xc5\xdd\x90\x7c\xe3\xf1\xff\x7e\x90\x95\x7d\x23\x29\x94\xd2\xae\x82\xf5\xd6\x6b\x3c\xf5\x9e\x03\x93\xac\x5f\x35\x39\x76\xa2\xfe\xd7\xc6\x67\xd3\x4a\x54\x95\xe5\x4b\x64\x3a\x5e\x41\x1f\xb2\x7a\x86\xc4\x27\xe1\x74\xca\x18\x56\x56\xcf\xd6\xbd\x30\x7c\x2b\xa1\xdd\xfa\x1b\xeb\x6d\x70\x57\x12\x6a\xa4\xd8\xf2\x1e\x32\x7f\x41\x66\x22\x19\x5f\xe1\xf4\x1f\x88\x78\x85\xd3\x1f\xd6\x30\xc6\xe9\x66\xce\xfe\xd8\x89\xbb\xb0\x7b\x5d\xfb\xaf\x7d\x7b\x26\xbf\xe7\xd0\x6e\x68\x08\x55\xdf\x95\xb3\x98\x27\xed\x7a\x3d\x36\x99\x92\x09\x53\x13\x38\xeb\xba\x79\x26\x9d\x2a\x3d\x85\x6c\x3d\x25\x17\x64\x4d\xad\x25\x63\xb2\xc4\x2f\xcb\xd5\xc9\x0a\x63\x07\x44\x82\xdd\xfd\xa7\xe9\xfe\x4f\xe9\xee\x93\x5f\xd3\xdd\xfd\xed\x7b\x64\xc5\x91\x61\x27\xed\xb0\x5f\x77\x9d\x4d\x94\xb2\xbe\xd1\x35\xf9\x89\xd1\x48\x42\xac\xbb\xf8\x7d\xfb\xdd\x1c\xea\x9e\x0e\xab\x96\x3e\x83\xf4\x1c\x63\x3b\xd7\xe9\x59\x04\x5f\x2c\xc0\xd1\x87\xad\x27\x59\x1a\x87\x74\xd5\xf5\x6c\xf8\x0a\x57\x4c\xc6\x8d\x62\x39\xc6\x7e\x2f\xfa\xe0\x3c\xe3\x01\xdc\x57\xbb\xa0\x28\xbe\x6d\xac\xf7\x75\x80\xc6\xb1\xb1\xd0\x95\x68\x7c\xb9\x34\xb5\x58\xcf\x07\x74\x32\xb7\x78\x2f\xc9\x6a\x5c\xdc\x9e\x26\xdf\x03\xc3\x7f\xa2\x32\x6d\x82\xcc\x6d\x54\x4b\x61\x1a\xac\x1f\x41\x30\x4e\x21\x7c\x46\xa8\xa4\x93\x23\x04\x9c\x20\x4d\xb9\x8c\x10\x2e\xc9\x37\xa3\x12\x96\x13\x6d\xc3\x5e\xc7\x83\x4b\x96\x7b\x15\xf9\xfa\xce\xb6\xe8\x5b\x2f\x75\x7c\x61\x00\x37\x0e\x75\x7c\x92\xb5\xd3\xee\x73\x69\x54\x19\x7d\x10\x27\x76\x55\xa1\xd3\xd8\xc1\x4a\x1f\x18\x3e\x7a\xe3\x42\xfb\x57\xd9\x26\x7a\xeb\xe1\xf9\xdd\xd2\x27\x52\x57\x2b\x13\x1b\xd9\xe8\x1d\x93\xb7\x49\x6d\xa5\x5b\x74\x53\x1b\xf0\xa1\x53\xb7\x5b\xef\x5f\x01\x00\x00\xff\xff\xa5\x4f\xd2\x0f\xac\x0a\x00\x00")

func nodeStartupShBytes() ([]byte, error) {
	return bindataRead(
		_nodeStartupSh,
		"node-startup.sh",
	)
}

func nodeStartupSh() (*asset, error) {
	bytes, err := nodeStartupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "node-startup.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"master-startup.sh": masterStartupSh,
	"node-startup.sh":   nodeStartupSh,
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
	"master-startup.sh": {masterStartupSh, map[string]*bintree{}},
	"node-startup.sh":   {nodeStartupSh, map[string]*bintree{}},
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
